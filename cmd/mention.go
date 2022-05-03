package cmd

import (
	"github.com/arrow2nd/ishell"
	"github.com/arrow2nd/twnyan/api"
)

func (cmd *Cmd) newMentionCmd() *ishell.Cmd {
	return &ishell.Cmd{
		Name:    "mention",
		Aliases: []string{"mt"},
		Func:    cmd.execMentionCmd,
		Help:    "get a Mentions to you",
		LongHelp: createLongHelp(
			`Get a Mentions to you.
If you omit the counts, the default value in the configuration file (25 by default) will be specified.`,
			"mt",
			"mention [counts]",
			"mention 50",
		),
	}
}

func (cmd *Cmd) execMentionCmd(c *ishell.Context) {
	count := cmd.getCountFromCmdArg(c.Args)
	query := api.CreateQuery(count)

	// メンションタイムラインを取得
	tweets, err := cmd.api.FetchTimelineTweets(api.Mention, query)
	if err != nil {
		cmd.showErrorMessage(err.Error())
		return
	}

	cmd.view.RegisterTweets(tweets)
	cmd.view.ShowRegisteredTweets()
}
