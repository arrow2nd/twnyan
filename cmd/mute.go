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
			"Mute user.\nIf you specify a tweetnumber, the person posting the tweet will be selected.",
			"mu",
			"mute [<username/tweetnumber>]",
			"mute arrow_2nd\n  mute 2",
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
			"Unmute user.\nIf you specify a tweetnumber, the person posting the tweet will be selected.",
			"rm",
			"mute remove [<username/tweetnumber>]",
			"mute remove arrow_2nd\n  mute rm 2",
		),
		Func: func(c *ishell.Context) {
			reactToUser(c.Args, c.Cmd.Name, twitter.Unmute)
		},
	})

	shell.AddCmd(muteCmd)
}
