package cmd

import "gopkg.in/abiosoft/ishell.v2"

func init() {
	shell.AddCmd(&ishell.Cmd{
		Name:    "mention",
		Aliases: []string{"mt"},
		Help:    "get a Mentions to you",
		LongHelp: createLongHelp(
			"Get a Mentions to you.\nIf you omit the counts, the default value in the configuration file (25 by default) will be specified.",
			"mt",
			"mention [counts]",
			"mention 50",
		),
		Func: func(c *ishell.Context) {
			// 取得件数
			counts := getCountsFromCmdArg(c.Args)

			// メンションTL読み込み
			err := tweets.LoadMentionTL(counts)
			if err != nil {
				return
			}

			// 表示
			tweets.DrawTweets()
		},
	})
}
