package cmd

import (
	"github.com/arrow2nd/ishell"
)

func (cmd *Cmd) newSearchCmd() {
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
	// 引数をパース
	keyword, counts, err := cmd.parseTLCmdArgs(c.Args)
	if err != nil {
		cmd.drawWrongArgMessage(c.Cmd.Name)
		return
	}

	// 検索結果を取得
	t, err := cmd.api.GetSearchResult(keyword, counts)
	if err != nil {
		cmd.drawErrorMessage(err.Error())
		return
	}

	// 描画
	cmd.view.RegisterTweets(t)
	cmd.view.DrawTweets()
}
