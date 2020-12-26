package engine

import (
	"sync"
)

// messageQueue struct for main commands stack
type messageQueue struct {
	sync.Mutex

	data          []Command
	receiveSignal chan struct{}
}

// push - add command to stack
func (mq *messageQueue) push(command Command) {
	mq.Lock()
	defer mq.Unlock()

	mq.data = append(mq.data, command)

	go func() {
		mq.receiveSignal <- struct{}{}
	}()
}

// pull - take command from stack
func (mq *messageQueue) pull() Command {
	mq.Lock()
	defer mq.Unlock()

	if len(mq.receiveSignal) < 1 {
		mq.Unlock()
		<-mq.receiveSignal
		mq.Lock()
	} else {
		<-mq.receiveSignal
	}

	res := mq.data[0]
	mq.data[0] = nil
	mq.data = mq.data[1:]

	return res
}

func (mq *messageQueue) empty() bool {
	return len(mq.data) == 0
}

type EventLoop struct {
	mq          *messageQueue
	stopSignal  chan struct{}
	stopRequest bool
}

func (l *EventLoop) mustLoop() bool {
	return !l.stopRequest || !l.mq.empty()
}

// Start runs main loop
func (l *EventLoop) Start() {
	l.mq = &messageQueue{receiveSignal: make(chan struct{}, 1)}
	l.stopSignal = make(chan struct{})
	go l.loop()
}

func (l *EventLoop) loop() {
	for l.mustLoop() {
		cmd := l.mq.pull()
		cmd.Execute(HandlerFunc(func(cmd Command) {
			l.postSystem(cmd)
		}))
	}
	l.stopSignal <- struct{}{}
}

type CommandFunc func(h Handler)

type HandlerFunc func(cmd Command)

func (hf HandlerFunc) Post(cmd Command) {
	hf(cmd)
}

func (cf CommandFunc) Execute(h Handler) {
	cf(h)
}

func (l *EventLoop) AwaitFinish() {
	l.Post(CommandFunc(func(h Handler) {
		l.stopRequest = true
	}))
	<-l.stopSignal
}

func (l *EventLoop) Resume() {
	if l.mustLoop() {
		l.stopRequest = false
	} else {
		l.stopRequest = false
		loopLaunched := make(chan struct{})
		l.mq.push(CommandFunc(func(h Handler) {
			loopLaunched <- struct{}{}
		}))
		go l.loop()
		<-loopLaunched
	}
}

func (l *EventLoop) Post(cmd Command) {
	if l.mustLoop() {
		l.mq.push(cmd)
	}
}

func (l *EventLoop) postSystem(cmd Command) {
	if !l.mustLoop() {
		l.Resume()
		l.Post(cmd)
		l.AwaitFinish()
	} else {
		l.Post(cmd)
	}
}
