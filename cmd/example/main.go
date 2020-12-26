package main

import (
	"bufio"
	"flag"
	"os"
	"strings"

	"github.com/inovarka/lab4/commands"
	"github.com/inovarka/lab4/engine"
)

func parse(cmdLine string) engine.Command {
	cmdFields := strings.Fields(cmdLine)
	cmdName := cmdFields[0]
	args := cmdFields[1:]
	res := commands.Construct(cmdName, args)
	return res
}

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
 			cmd := parse(commandLine) 
 			eventLoop.Post(cmd)
 		}
	}
		
	eventLoop.AwaitFinish()

}