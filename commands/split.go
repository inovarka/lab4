package commands

import (
	"strings"

	"github.com/inovarka/lab4/engine"
)

type SplitCmd struct {
	Str string
	Sep string
}

func (sCmd *SplitCmd) Execute(h engine.Handler) {
	splitted := strings.Split(sCmd.Str, sCmd.Sep)
	joined := strings.Join(splitted, "\n")
	h.Post(&PrintCmd{joined})
}
