package main

import (
	"bufio"
	"flag"
	"os"

	"github.com/inovarka/lab4/engine"
)

func main() {
	inputFile := flag.String("f", "", "Run file with EventLoop")
	flag.Parse()

	eventLoop := new(engine.EventLoop)
	eventLoop.Start()
	if input, err := os.Open(inputFile); err == nil {
		defer input.Close()
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			commandLine := scanner.Text()
			cmd := Parse(commandLine)
			eventLoop.Post(cmd)
		}
	}

	eventLoop.AwaitFinish()

}
