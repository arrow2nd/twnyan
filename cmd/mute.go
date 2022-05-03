package cmd

import (
	"github.com/arrow2nd/ishell/v2"
)

func (cmd *Cmd) newMuteCmd() *ishell.Cmd {
	// mute
	muteCmd := &ishell.Cmd{
		Name:    "mute",
		Aliases: []string{"mu"},
		Func: func(c *ishell.Context) {
			cmd.actionOnUser("MUTED", c.Cmd.Name, cmd.config.Color.Mute, c.Args, cmd.twitter.Mute)
		},
		Help: "mute a user",
		LongHelp: createLongHelp(
			`Mute a user.
If you specify a tweet-number, the person posting the tweet will be selected.`,
			"mu",
			"mute [<username / tweet-number>]",
			"mute arrow_2nd\n  mute 2",
		),
	}

	// mute remove
	muteCmd.AddCmd(&ishell.Cmd{
		Name:    "remove",
		Aliases: []string{"rm"},
		Func: func(c *ishell.Context) {
			cmd.actionOnUser("UNMUTED", "mute "+c.Cmd.Name, cmd.config.Color.Mute, c.Args, cmd.twitter.Unmute)
		},
		Help: "unmute a user",
		LongHelp: createLongHelp(
			`Unmute a user.
If you specify a tweet-number, the person posting the tweet will be selected.`,
			"rm",
			"mute remove [<username / tweet-number>]",
			"mute remove arrow_2nd\n  mute rm 2",
		),
	})

	return muteCmd
}
