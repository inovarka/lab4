package commands

import (
	"fmt"
	"strings"

	"github.com/inovarka/lab4/engine"
)

func Parse(command string) engine.Command {
	splitted := strings.FieldsFunc(command, func(r rune) bool {
		return r == ' '
	})

	var parsed engine.Command = &printCommand{
		fmt.Sprintf("PARSING ERROR: Invalid command: %s", command),
	}

	l := len(splitted)
	switch command := splitted[0]; command {
	case "print":
		if l < 2 {
			parsed = &printCommand{
				fmt.Sprintf("SYNTAX ERROR: Trying to print an empty line"),
			}
		} else {
			parsed = &printCommand{strings.Join(splitted[1:], " ")}
		}
	case "split":
		if l != 3 {
			parsed = &printCommand{
				fmt.Sprintf("SYNTAX ERROR: Invalid count of arguments for split: %d", l),
			}
		} else {
			parsed = &splitCommand{
				Str: splitted[1],
				Sep: splitted[2],
			}
		}
	}

	return parsed
}
