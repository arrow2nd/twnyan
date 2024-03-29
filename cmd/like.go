package cmd

import (
	"github.com/arrow2nd/ishell"
)

func (cmd *Cmd) addLikeCmd() {
	// like
	likeCmd := &ishell.Cmd{
		Name:    "like",
		Aliases: []string{"lk", "fv"},
		Func: func(c *ishell.Context) {
			cmd.actionOnTweet("LIKED", c.Cmd.Name, cmd.cfg.Color.Favorite, c.Args, cmd.api.Favorite)
		},
		Help: "like a tweet",
		LongHelp: createLongHelp(
			"Like a tweet.\nIf there is more than one, please separate them with a space.",
			"lk, fv",
			"like [<tweetnumber>]...",
			"like 0 1",
		),
	}

	// like remove
	likeCmd.AddCmd(&ishell.Cmd{
		Name:    "remove",
		Aliases: []string{"rm"},
		Func: func(c *ishell.Context) {
			cmd.actionOnTweet("UN-LIKED", "like "+c.Cmd.Name, cmd.cfg.Color.Favorite, c.Args, cmd.api.Unfavorite)
		},
		Help: "un-like a tweet",
		LongHelp: createLongHelp(
			"UnLike a tweet.\nIf there is more than one, please separate them with a space.",
			"rm",
			"like remove [<tweetnumber>]...",
			"like remove 0 1",
		),
	})

	cmd.shell.AddCmd(likeCmd)
}
