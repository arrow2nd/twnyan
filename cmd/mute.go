package cmd

import (
	"github.com/arrow2nd/twnyan/twitter"
	"gopkg.in/abiosoft/ishell.v2"
)

func init() {
	muteCmd := &ishell.Cmd{
		Name:    "mute",
		Aliases: []string{"mu"},
		Help:    "mute user",
		LongHelp: createLongHelp(
			"Mute user.\nIf you specify a tweet number, the person posting the tweet will be selected.",
			"mu",
			"mute [<userID / tweet number>]",
			"mute arrow_2nd",
		),
		Func: func(c *ishell.Context) {
			reactToUser(c.Args, c.Cmd.Name, twitter.Mute)
		},
	}

	muteCmd.AddCmd(&ishell.Cmd{
		Name:    "remove",
		Aliases: []string{"rm"},
		Help:    "unmute user",
		LongHelp: createLongHelp(
			"Unmute user.\nIf you specify a tweet number, the person posting the tweet will be selected.",
			"rm",
			"mute remove [<userID / tweet number>]",
			"mute remove arrow_2nd",
		),
		Func: func(c *ishell.Context) {
			reactToUser(c.Args, c.Cmd.Name, twitter.Unmute)
		},
	})

	shell.AddCmd(muteCmd)
}
