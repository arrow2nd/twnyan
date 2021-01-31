package cmd

import (
	"gopkg.in/abiosoft/ishell.v2"
)

func (cmd *Cmd) newMuteCmd() {
	mc := &ishell.Cmd{
		Name:    "mute",
		Aliases: []string{"mu"},
		Func: func(c *ishell.Context) {
			cmd.reactToUser(c.Args, c.Cmd.Name, cmd.api.Mute)
		},
		Help: "mute user",
		LongHelp: createLongHelp(
			"Mute user.\nIf you specify a tweetnumber, the person posting the tweet will be selected.",
			"mu",
			"mute [<username/tweetnumber>]",
			"mute arrow_2nd\n  mute 2",
		),
	}

	mc.AddCmd(&ishell.Cmd{
		Name:    "remove",
		Aliases: []string{"rm"},
		Func: func(c *ishell.Context) {
			cmd.reactToUser(c.Args, c.Cmd.Name, cmd.api.Unmute)
		},
		Help: "unmute user",
		LongHelp: createLongHelp(
			"Unmute user.\nIf you specify a tweetnumber, the person posting the tweet will be selected.",
			"rm",
			"mute remove [<username/tweetnumber>]",
			"mute remove arrow_2nd\n  mute rm 2",
		),
	})

	cmd.shell.AddCmd(mc)
}
