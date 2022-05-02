package cmd

import (
	"github.com/arrow2nd/ishell"
	"github.com/arrow2nd/twnyan/api"
)

func (cmd *Cmd) newTimelineCmd() *ishell.Cmd {
	return &ishell.Cmd{
		Name:    "timeline",
		Aliases: []string{"tl"},
		Func:    cmd.execTimelineCmd,
		Help:    "get a home timeline",
		LongHelp: createLongHelp(
			`Get a home timeline.
If you omit the counts, the default value in the configuration file (25 by default) will be specified.`,
			"tl",
			"timeline [counts]",
			"timeline 50",
		),
	}
}

func (cmd *Cmd) execTimelineCmd(c *ishell.Context) {
	count := cmd.getCountFromCmdArg(c.Args)
	query := api.CreateQuery(count)

	// タイムラインのツイートを取得
	tweets, err := cmd.api.FetchTimelineTweets(api.Home, query)
	if err != nil {
		cmd.showErrorMessage(err.Error())
		return
	}

	cmd.view.RegisterTweets(tweets)
	cmd.view.ShowRegisteredTweets()
}
