package cmd

import (
	"github.com/arrow2nd/ishell/v2"
)

func (cmd *Cmd) newBlockCmd() *ishell.Cmd {
	// block
	blockCmd := &ishell.Cmd{
		Name:    "block",
		Aliases: []string{"bk"},
		Func: func(c *ishell.Context) {
			cmd.actionOnUser("BLOCKED", c.Cmd.Name, cmd.config.Color.Block, c.Args, cmd.twitter.Block)
		},
		Help: "block a user",
		LongHelp: createLongHelp(
			`Block a user.
If you specify a tweet-number, the person posting the tweet will be selected.`,
			"bk",
			"block <username / tweet-number>",
			"block arrow_2nd\n  block 2",
		),
	}

	// block remove
	blockCmd.AddCmd(&ishell.Cmd{
		Name:    "remove",
		Aliases: []string{"rm"},
		Func: func(c *ishell.Context) {
			cmd.actionOnUser("UNBLOCKED", "block "+c.Cmd.Name, cmd.config.Color.Block, c.Args, cmd.twitter.Unblock)
		},
		Help: "unblock a user",
		LongHelp: createLongHelp(
			`Unblock a user.
If you specify a tweet-number, the person posting the tweet will be selected.`,
			"rm",
			"block remove <username / tweet-number>",
			"block remove arrow_2nd\n  block rm 2",
		),
	})

	return blockCmd
}
