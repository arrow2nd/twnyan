package cmd

import (
	"gopkg.in/abiosoft/ishell.v2"
)

func (cmd *Cmd) newFollowCmd() {
	fc := &ishell.Cmd{
		Name:    "follow",
		Aliases: []string{"fw"},
		Func: func(c *ishell.Context) {
			cmd.reactToUser(c.Args, c.Cmd.Name, cmd.api.Follow)
		},
		Help: "follow user",
		LongHelp: createLongHelp(
			"Follow user.\nIf you specify a tweetnumber, the person posting the tweet will be selected.",
			"fw",
			"follow [<username/tweetnumber>]",
			"follow arrow_2nd\n  follow 2",
		),
	}

	fc.AddCmd(&ishell.Cmd{
		Name:    "remove",
		Aliases: []string{"rm"},
		Func: func(c *ishell.Context) {
			cmd.reactToUser(c.Args, c.Cmd.Name, cmd.api.Unfollow)
		},
		Help: "unfollow user",
		LongHelp: createLongHelp(
			"Unfollow user.\nIf you specify a tweetnumber, the person posting the tweet will be selected.",
			"rm",
			"follow remove [<username/tweetnumber>]",
			"follow remove arrow_2nd\n  follow rm 2",
		),
	})

	cmd.shell.AddCmd(fc)
}
