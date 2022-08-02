package cmd

import (
	"fmt"

	"github.com/arrow2nd/ishell/v2"
)

const version = "1.9.4"

func (cmd *Cmd) newVersionCmd() *ishell.Cmd {
	return &ishell.Cmd{
		Name:    "version",
		Aliases: []string{"ver"},
		Func: func(c *ishell.Context) {
			fmt.Printf("🐈 twnyan ver.%s\n", version)
		},
		Help:     "display version",
		LongHelp: "Display twnyan version.",
	}
}
