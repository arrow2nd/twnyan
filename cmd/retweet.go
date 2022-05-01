package cmd

import (
	"github.com/arrow2nd/ishell"
)

func (cmd *Cmd) newRetweetCmd() *ishell.Cmd {
	// retweet
	retweetCmd := &ishell.Cmd{
		Name:    "retweet",
		Aliases: []string{"rt"},
		Func: func(c *ishell.Context) {
			cmd.actionOnTweet("RETWEETED", c.Cmd.Name, cmd.cfg.Color.Retweet, c.Args, cmd.api.Retweet)
		},
		Help: "retweet a tweet",
		LongHelp: createLongHelp(
			`Retweet a tweet.
If there is more than one, please separate them with a space.`,
			"rt",
			"retweet [<tweetnumber>]...",
			"retweet 0 1",
		),
	}

	// retweet remove
	retweetCmd.AddCmd(&ishell.Cmd{
		Name:    "remove",
		Aliases: []string{"rm"},
		Func: func(c *ishell.Context) {
			cmd.actionOnTweet("UN-RETWEETED", "retweet "+c.Cmd.Name, cmd.cfg.Color.Retweet, c.Args, cmd.api.UnRetweet)
		},
		Help: "un-retweet a tweet",
		LongHelp: createLongHelp(
			`UnRetweet a tweet.
If there is more than one, please separate them with a space.`,
			"rm",
			"retweet remove [<tweetnumber>]...",
			"retweet remove 0 1",
		),
	})

	return retweetCmd
}
