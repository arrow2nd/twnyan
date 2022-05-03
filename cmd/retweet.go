package cmd

import (
	"github.com/arrow2nd/ishell/v2"
)

func (cmd *Cmd) newRetweetCmd() *ishell.Cmd {
	// retweet
	retweetCmd := &ishell.Cmd{
		Name:    "retweet",
		Aliases: []string{"rt"},
		Func: func(c *ishell.Context) {
			cmd.actionOnTweet("RETWEETED", c.Cmd.Name, cmd.config.Color.Retweet, c.Args, cmd.twitter.Retweet)
		},
		Help: "retweet a tweet",
		LongHelp: createLongHelp(
			`Retweet a tweet.
If there is more than one, please separate them with a space.`,
			"rt",
			"retweet [<tweet-number>]...",
			"retweet 0 1",
		),
	}

	// retweet remove
	retweetCmd.AddCmd(&ishell.Cmd{
		Name:    "remove",
		Aliases: []string{"rm"},
		Func: func(c *ishell.Context) {
			cmd.actionOnTweet("UNRETWEETED", "retweet "+c.Cmd.Name, cmd.config.Color.Retweet, c.Args, cmd.twitter.UnRetweet)
		},
		Help: "unretweet a tweet",
		LongHelp: createLongHelp(
			`UnRetweet a tweet.
If there is more than one, please separate them with a space.`,
			"rm",
			"retweet remove [<tweet-number>]...",
			"retweet remove 0 1",
		),
	})

	return retweetCmd
}
