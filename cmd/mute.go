package cmd

import (
	"github.com/arrow2nd/ishell"
)

func (cmd *Cmd) newMuteCmd() {
	mc := &ishell.Cmd{
		Name:    "mute",
		Aliases: []string{"mu"},
		Func: func(c *ishell.Context) {
			cmd.actionOnUser("MUTED", c.Cmd.Name, cmd.cfg.Color.Mute, c.Args, cmd.api.Mute)
		},
		Help: "mute a user",
		LongHelp: createLongHelp(
			"Mute a user.\nIf you specify a tweetnumber, the person posting the tweet will be selected.",
			"mu",
			"mute [<username/tweetnumber>]",
			"mute arrow_2nd\n  mute 2",
		),
	}

	mc.AddCmd(&ishell.Cmd{
		Name:    "remove",
		Aliases: []string{"rm"},
		Func: func(c *ishell.Context) {
			cmd.actionOnUser("UNMUTED", "mute "+c.Cmd.Name, cmd.cfg.Color.Mute, c.Args, cmd.api.Unmute)
		},
		Help: "unmute a user",
		LongHelp: createLongHelp(
			"Unmute a user.\nIf you specify a tweetnumber, the person posting the tweet will be selected.",
			"rm",
			"mute remove [<username/tweetnumber>]",
			"mute remove arrow_2nd\n  mute rm 2",
		),
	})

	cmd.shell.AddCmd(mc)
}
