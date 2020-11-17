package cmd

import (
	"github.com/arrow2nd/twnyan/twitter"
	"gopkg.in/abiosoft/ishell.v2"
)

func init() {
	favCmd := &ishell.Cmd{
		Name:    "favorite",
		Aliases: []string{"like", "fv"},
		Help:    "like tweet",
		LongHelp: createLongHelp(
			"Like tweet.\nIf there is more than one, please separate them with a space.",
			"like, fv",
			"favorite [<tweet number>]...",
			"favorite 0 1",
		),
		Func: func(c *ishell.Context) {
			reactToTweet(c.Args, c.Cmd.Name, twitter.Favorite)
		},
	}

	favCmd.AddCmd(&ishell.Cmd{
		Name:    "remove",
		Aliases: []string{"rm"},
		Help:    "un-like tweet",
		LongHelp: createLongHelp(
			"UnLike tweet.\nIf there is more than one, please separate them with a space.",
			"rm",
			"favorite remove [<tweet number>]...",
			"favorite remove 0 1",
		),
		Func: func(c *ishell.Context) {
			reactToTweet(c.Args, "favorite "+c.Cmd.Name, twitter.Unfavorite)
		},
	})

	shell.AddCmd(favCmd)
}
