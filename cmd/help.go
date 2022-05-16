package cmd

import (
	"fmt"

	"github.com/arrow2nd/ishell/v2"
)

func (cmd *Cmd) newHelpCmd() *ishell.Cmd {
	return &ishell.Cmd{
		Name: "help",
		Func: func(c *ishell.Context) {
			fmt.Print(cmd.shell.HelpText())
			cmd.flagSet.Usage()
		},
		Help:     "display help",
		LongHelp: "Display twnyan help.",
	}
}
