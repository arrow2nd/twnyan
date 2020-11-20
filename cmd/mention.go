package cmd

import (
	"github.com/arrow2nd/twnyan/util"
	"gopkg.in/abiosoft/ishell.v2"
)

func init() {
	shell.AddCmd(&ishell.Cmd{
		Name:    "mention",
		Aliases: []string{"mt"},
		Help:    "get a Mentions to you",
		LongHelp: createLongHelp(
			"Get a Mentions to you.\nIf you omit the counts, the default value in the configuration file (25 by default) will be specified.",
			"mt",
			"mention [counts] [data format(json|yaml)]",
			"mention 50",
		),
		Func: func(c *ishell.Context) {
			counts, dataFmt := cfg.Default.Counts, ""

			args, _ := util.FetchStringSpecifiedType(c.Args, "num", "str")
			if args != nil {
				counts, dataFmt = args[0], args[1]
				if counts == "" {
					counts = cfg.Default.Counts
				}
			}

			err := tweets.LoadMentionTL(counts)
			if err != nil {
				return
			}

			if dataFmt == "" {
				tweets.DrawTweets()
			} else {
				tweets.OutData(dataFmt)
			}
		},
	})
}
