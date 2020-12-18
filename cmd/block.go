package cmd

import (
	"github.com/arrow2nd/twnyan/twitter"
	"gopkg.in/abiosoft/ishell.v2"
)

func init() {
	blockCmd := &ishell.Cmd{
		Name:    "block",
		Aliases: []string{"bk"},
		Help:    "block user",
		LongHelp: createLongHelp(
			"Block user.\nIf you specify a tweetnumber, the person posting the tweet will be selected.",
			"bk",
			"block [<username/tweetnumber>]",
			"block arrow_2nd\n  block 2",
		),
		Func: func(c *ishell.Context) {
			reactToUser(c.Args, c.Cmd.Name, twitter.Block)
		},
	}

	blockCmd.AddCmd(&ishell.Cmd{
		Name:    "remove",
		Aliases: []string{"rm"},
		Help:    "unblock user",
		LongHelp: createLongHelp(
			"Unblock user.\nIf you specify a tweetnumber, the person posting the tweet will be selected.",
			"rm",
			"block remove [<username/tweetnumber>]",
			"block remove arrow_2nd\n  block rm 2",
		),
		Func: func(c *ishell.Context) {
			reactToUser(c.Args, c.Cmd.Name, twitter.Unblock)
		},
	})

	shell.AddCmd(blockCmd)
}
