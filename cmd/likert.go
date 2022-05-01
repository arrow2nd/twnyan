package cmd

import (
	"github.com/arrow2nd/ishell"
)

func (cmd *Cmd) newLikertCmd() *ishell.Cmd {
	return &ishell.Cmd{
		Name:    "likert",
		Aliases: []string{"lr", "fr"},
		Func: func(c *ishell.Context) {
			cmd.actionOnTweet("LIKED&RT", c.Cmd.Name, cmd.cfg.Color.Favorite, c.Args, cmd.execLikert)
		},
		Help: "like and retweet a tweet",
		LongHelp: createLongHelp(
			`Like and Retweet a tweet.
If there is more than one, please separate them with a space.`,
			"lr, fr",
			"likert [<tweetnumber>]...",
			"likert 0 1",
		),
	}
}

func (cmd *Cmd) execLikert(tweetID string) (string, error) {
	if _, err := cmd.api.Favorite(tweetID); err != nil {
		return "", err
	}

	return cmd.api.Retweet(tweetID)
}
