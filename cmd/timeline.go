package cmd

import "gopkg.in/abiosoft/ishell.v2"

func init() {
	shell.AddCmd(&ishell.Cmd{
		Name:    "timeline",
		Aliases: []string{"tl"},
		Help:    "get a home timeline",
		LongHelp: createLongHelp(
			"Get a home timeline.\nIf you omit the counts, the default value in the configuration file (25 by default) will be specified.",
			"tl",
			"timeline [counts]",
			"timeline 50",
		),
		Func: func(c *ishell.Context) {
			// 取得件数
			counts := getCountsFromCmdArg(c.Args)

			// ホームTL取得
			err := tweets.LoadHomeTL(counts)
			if err != nil {
				return
			}

			// 表示
			tweets.DrawTweets()
		},
	})
}
