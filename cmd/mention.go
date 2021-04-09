package cmd

import (
	"github.com/arrow2nd/ishell"
	"github.com/arrow2nd/twnyan/api"
)

func (cmd *Cmd) addMentionCmd() {
	cmd.shell.AddCmd(&ishell.Cmd{
		Name:    "mention",
		Aliases: []string{"mt"},
		Func:    cmd.mentionCmd,
		Help:    "get a Mentions to you",
		LongHelp: createLongHelp(
			"Get a Mentions to you.\nIf you omit the counts, the default value in the configuration file (25 by default) will be specified.",
			"mt",
			"mention [counts]",
			"mention 50",
		),
	})
}

func (cmd *Cmd) mentionCmd(c *ishell.Context) {
	count := cmd.getCountFromCmdArg(c.Args)

	// メンションタイムラインを取得
	query := api.CreateQuery(count)
	tweets, err := cmd.api.FetchTimelineTweets("mention", query)
	if err != nil {
		cmd.showErrorMessage(err.Error())
		return
	}

	cmd.view.RegisterTweets(tweets)
	cmd.view.ShowRegisteredTweets()
}
