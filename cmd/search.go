package cmd

import (
	"github.com/arrow2nd/ishell"
)

func (cmd *Cmd) addSearchCmd() {
	cmd.shell.AddCmd(&ishell.Cmd{
		Name:    "search",
		Aliases: []string{"sh"},
		Func:    cmd.searchCmd,
		Help:    "search for tweets from the past 7 days",
		LongHelp: createLongHelp(
			"Search for tweets from the past 7 days.\nIf you omit the counts, the default value in the configuration file (25 by default) will be specified.",
			"sh",
			"search [<keyword>] [counts]",
			"search cats 50",
		),
	})
}

func (cmd *Cmd) searchCmd(c *ishell.Context) {
	keyword, count, err := cmd.parseTimelineCmdArgs(c.Args)
	if err != nil {
		cmd.showWrongArgMessage(c.Cmd.Name)
		return
	}

	// 検索結果を取得
	tweets, err := cmd.api.FetchSearchResult(keyword, count)
	if err != nil {
		cmd.showErrorMessage(err.Error())
		return
	}

	cmd.view.RegisterTweets(tweets)
	cmd.view.ShowRegisteredTweets()
}
