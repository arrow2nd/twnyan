package cmd

import (
	"github.com/arrow2nd/ishell/v2"
)

func (cmd *Cmd) newLikertCmd() *ishell.Cmd {
	return &ishell.Cmd{
		Name:    "likert",
		Aliases: []string{"lr", "fr"},
		Func: func(c *ishell.Context) {
			cmd.actionOnTweet("LIKE&RT", c.Cmd.Name, cmd.config.Color.Favorite, c.Args, cmd.execLikertCmd)
		},
		Help: "like & retweet a tweet",
		LongHelp: createLongHelp(
			`Like & Retweet a tweet.
If there is more than one, please separate them with a space.
(Only available in interactive mode)`,
			"lr, fr",
			"likert <tweet-number>...",
			"likert 0 1",
		),
	}
}

func (cmd *Cmd) execLikertCmd(tweetId string) (string, error) {
	if _, err := cmd.twitter.Favorite(tweetId); err != nil {
		return "", err
	}

	return cmd.twitter.Retweet(tweetId)
}
