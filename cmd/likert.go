package cmd

import (
	"github.com/arrow2nd/ishell"
)

func (cmd *Cmd) newLikertCmd() *ishell.Cmd {
	return &ishell.Cmd{
		Name:    "likert",
		Aliases: []string{"lr", "fr"},
		Func: func(c *ishell.Context) {
			cmd.actionOnTweet("LIKE&RT", c.Cmd.Name, cmd.cfg.Color.Favorite, c.Args, cmd.execLikert)
		},
		Help: "like and retweet a tweet",
		LongHelp: createLongHelp(
			`Like and Retweet a tweet.
If there is more than one, please separate them with a space.`,
			"lr, fr",
			"likert [<tweet-number>]...",
			"likert 0 1",
		),
	}
}

func (cmd *Cmd) execLikert(tweetId string) (string, error) {
	if _, err := cmd.api.Favorite(tweetId); err != nil {
		return "", err
	}

	return cmd.api.Retweet(tweetId)
}
