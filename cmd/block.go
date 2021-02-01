package cmd

import (
	"gopkg.in/abiosoft/ishell.v2"
)

func (cmd *Cmd) newBlockCmd() {
	bc := &ishell.Cmd{
		Name:    "block",
		Aliases: []string{"bk"},
		Func: func(c *ishell.Context) {
			cmd.actionOnUser("BLOCKED", c.Cmd.Name, cmd.cfg.Color.Block, c.Args, cmd.api.Block)
		},
		Help: "block user",
		LongHelp: createLongHelp(
			"Block user.\nIf you specify a tweetnumber, the person posting the tweet will be selected.",
			"bk",
			"block [<username/tweetnumber>]",
			"block arrow_2nd\n  block 2",
		),
	}

	bc.AddCmd(&ishell.Cmd{
		Name:    "remove",
		Aliases: []string{"rm"},
		Func: func(c *ishell.Context) {
			cmd.actionOnUser("UNBLOCKED", "block "+c.Cmd.Name, cmd.cfg.Color.Block, c.Args, cmd.api.Unblock)
		},
		Help: "unblock user",
		LongHelp: createLongHelp(
			"Unblock user.\nIf you specify a tweetnumber, the person posting the tweet will be selected.",
			"rm",
			"block remove [<userName/tweetnumber>]",
			"block remove arrow_2nd\n  block rm 2",
		),
	})

	cmd.shell.AddCmd(bc)
}
