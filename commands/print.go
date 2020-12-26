package commands

import (
	"fmt"

	"github.com/inovarka/lab4/engine"
)

type printCommand struct {
	Arg string
}

func (prnt *printCommand) Execute(loop engine.Handler) {
	fmt.Println(prnt.Arg)
}