package cmd

import (
	"github.com/arrow2nd/ishell/v2"
)

func (cmd *Cmd) newSearchCmd() *ishell.Cmd {
	return &ishell.Cmd{
		Name:    "search",
		Aliases: []string{"sh"},
		Func:    cmd.execSearchCmd,
		Help:    "search for tweets in the last 7 days",
		LongHelp: createLongHelp(
			`Search for tweets in the last 7 days.
If number of acquisitions is omitted, the default value in the configuration file is specified.`,
			"sh",
			"search <keyword> [number]",
			"search cats 50",
		),
	}
}

func (cmd *Cmd) execSearchCmd(c *ishell.Context) {
	keyword, count, err := cmd.parseTimelineCmdArgs(c.Args)
	if err != nil {
		cmd.showWrongArgMessage(c.Cmd.Name)
		return
	}

	// 検索結果を取得
	tweets, err := cmd.twitter.FetchSearchResult(keyword, count)
	if err != nil {
		cmd.showErrorMessage(err.Error())
		return
	}

	cmd.twitter.RegisterTweets(tweets)
	cmd.showTweets()
}
