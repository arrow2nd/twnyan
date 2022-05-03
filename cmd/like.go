package cmd

import (
	"github.com/arrow2nd/ishell/v2"
)

func (cmd *Cmd) newLikeCmd() *ishell.Cmd {
	// like
	likeCmd := &ishell.Cmd{
		Name:    "like",
		Aliases: []string{"lk", "fv"},
		Func: func(c *ishell.Context) {
			cmd.actionOnTweet("LIKED", c.Cmd.Name, cmd.config.Color.Favorite, c.Args, cmd.twitter.Favorite)
		},
		Help: "like a tweet",
		LongHelp: createLongHelp(
			`Like a tweet.
If there is more than one, please separate them with a space.`,
			"lk, fv",
			"like [<tweet-number>]...",
			"like 0 1",
		),
	}

	// like remove
	likeCmd.AddCmd(&ishell.Cmd{
		Name:    "remove",
		Aliases: []string{"rm"},
		Func: func(c *ishell.Context) {
			cmd.actionOnTweet("UN-LIKED", "like "+c.Cmd.Name, cmd.config.Color.Favorite, c.Args, cmd.twitter.Unfavorite)
		},
		Help: "un-like a tweet",
		LongHelp: createLongHelp(
			`UnLike a tweet.
If there is more than one, please separate them with a space.`,
			"rm",
			"like remove [<tweet-number>]...",
			"like remove 0 1",
		),
	})

	return likeCmd
}
