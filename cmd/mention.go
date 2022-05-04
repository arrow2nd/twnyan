package cmd

import (
	"github.com/arrow2nd/ishell/v2"
	"github.com/arrow2nd/twnyan/twitter"
)

func (cmd *Cmd) newMentionCmd() *ishell.Cmd {
	return &ishell.Cmd{
		Name:    "mention",
		Aliases: []string{"mt"},
		Func:    cmd.execMentionCmd,
		Help:    "displays mention tweets",
		LongHelp: createLongHelp(
			`Displays Mention Tweets addressed to you.
If you omit the counts, the default value in the configuration file (25 by default) will be specified.`,
			"mt",
			"mention [counts]",
			"mention 50",
		),
	}
}

func (cmd *Cmd) execMentionCmd(c *ishell.Context) {
	count := cmd.getCountFromCmdArg(c.Args)
	query := twitter.CreateQuery(count)

	// メンションタイムラインを取得
	tweets, err := cmd.twitter.FetchTimelineTweets(twitter.Mention, query)
	if err != nil {
		cmd.showErrorMessage(err.Error())
		return
	}

	cmd.twitter.RegisterTweets(tweets)
	cmd.showTweets()
}
