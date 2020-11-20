package cmd

import (
	"github.com/arrow2nd/twnyan/util"
	"gopkg.in/abiosoft/ishell.v2"
)

func init() {
	shell.AddCmd(&ishell.Cmd{
		Name:    "timeline",
		Aliases: []string{"tl"},
		Help:    "get a home timeline",
		LongHelp: createLongHelp(
			"Get a home timeline.\nIf you omit the counts, the default value in the configuration file (25 by default) will be specified.",
			"tl",
			"timeline [counts] [data format(json|yaml)]",
			"timeline 50",
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

			err := tweets.LoadHomeTL(counts)
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
