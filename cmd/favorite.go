package cmd

import (
	"gopkg.in/abiosoft/ishell.v2"
)

func (cmd *Cmd) newFaoriteCmd() {
	fc := &ishell.Cmd{
		Name:    "favorite",
		Aliases: []string{"like", "fv"},
		Func: func(c *ishell.Context) {
			cmd.reactToTweet(c.Args, c.Cmd.Name, cmd.api.Favorite)
		},
		Help: "like tweet",
		LongHelp: createLongHelp(
			"Like tweet.\nIf there is more than one, please separate them with a space.",
			"like, fv",
			"favorite [<tweetnumber>]...",
			"favorite 0 1",
		),
	}

	fc.AddCmd(&ishell.Cmd{
		Name:    "remove",
		Aliases: []string{"rm"},
		Func: func(c *ishell.Context) {
			cmd.reactToTweet(c.Args, "favorite "+c.Cmd.Name, cmd.api.Unfavorite)
		},
		Help: "un-like tweet",
		LongHelp: createLongHelp(
			"UnLike tweet.\nIf there is more than one, please separate them with a space.",
			"rm",
			"favorite remove [<tweetnumber>]...",
			"favorite remove 0 1",
		),
	})

	cmd.shell.AddCmd(fc)
}
