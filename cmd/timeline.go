package cmd

import (
	"github.com/arrow2nd/ishell/v2"
	"github.com/arrow2nd/twnyan/twitter"
)

func (cmd *Cmd) newTimelineCmd() *ishell.Cmd {
	return &ishell.Cmd{
		Name:    "timeline",
		Aliases: []string{"tl"},
		Func:    cmd.execTimelineCmd,
		Help:    "displays the home timeline",
		LongHelp: createLongHelp(
			`Displays the home timeline.
If number of acquisitions is omitted, the default value in the configuration file is specified.`,
			"tl",
			"timeline [number]",
			"timeline 50",
		),
	}
}

func (cmd *Cmd) execTimelineCmd(c *ishell.Context) {
	count := cmd.getCountFromCmdArg(c.Args)
	query := twitter.CreateQuery(count)

	// タイムラインのツイートを取得
	tweets, err := cmd.twitter.FetchTimelineTweets(twitter.Home, query)
	if err != nil {
		cmd.showErrorMessage(err.Error())
		return
	}

	cmd.twitter.RegisterTweets(tweets)
	cmd.showTweets()
}
