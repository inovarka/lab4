package main

import (
	"bufio"
	"errors"
	"flag"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/inovarka/lab4/commands"
	"github.com/inovarka/lab4/engine"
)

var (
	inputFile = flag.String("f", "", "example")
)

func scanFile(inputFile *string, loop *engine.EventLoop) error {
	var source io.Reader

	if *inputFile != "" {
		data, err := ioutil.ReadFile(*inputFile)
		if err != nil {
			return err
		}
		source = strings.NewReader(string(data))
	} else {
		return errors.New("No file")
	}

	scanner := bufio.NewScanner(source)
	for scanner.Scan() {
		commandLine := scanner.Text()
		cmd := commands.Parse(commandLine)
		loop.Post(cmd)
	}

	return nil
}

func main() {

	filename := "example.txt"
	eventLoop := new(engine.EventLoop)
	eventLoop.Start()
	if input, err := os.Open(filename); err == nil {
		defer input.Close()
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			commandLine := scanner.Text()
			cmd := commands.Parse(commandLine)
			eventLoop.Post(cmd)
		}
	} else {
		panic(err)
	}
	eventLoop.AwaitFinish()
	// Doesn't execute after finish
	eventLoop.Post(commands.Parse("print never_run"))

	//flag.Parse()

	//eventLoop := new(engine.EventLoop)
	//eventLoop.Start()

	//if err := scanFile(inputFile, eventLoop); err != nil {
	//	_, _ = os.Stderr.WriteString(err.Error() + "\n")
	//}

	//eventLoop.AwaitFinish()
}
