package cmd

import "gopkg.in/abiosoft/ishell.v2"

func init() {
	shell.AddCmd(&ishell.Cmd{
		Name:    "search",
		Aliases: []string{"sh"},
		Help:    "search for tweets from the past 7 days",
		LongHelp: createLongHelp(
			"Search for tweets from the past 7 days.\nIf you omit the counts, the default value in the configuration file (25 by default) will be specified.",
			"sh",
			"search [<keyword>] [counts]",
			"search cats 50",
		),
		Func: func(c *ishell.Context) {
			// 引数をパース
			keyword, counts, err := parseTLCmdArgs(c.Args)
			if err != nil {
				showWrongMsg(c.Cmd.Name)
				return
			}

			// 検索結果取得
			err = tweets.LoadSearchResult(keyword, counts)
			if err != nil {
				return
			}

			// 表示
			tweets.DrawTweets()
		},
	})
}
