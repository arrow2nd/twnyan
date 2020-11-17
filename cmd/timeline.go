package cmd

import (
	"github.com/arrow2nd/twnyan/util"
	"gopkg.in/abiosoft/ishell.v2"
)

func init() {
	shell.AddCmd(&ishell.Cmd{
		Name:    "timeline",
		Aliases: []string{"tl"},
		Help:    "displays the home timeline",
		LongHelp: createLongHelp(
			"Displays the home timeline.\nIf you omit the counts, the default value in the configuration file (25 by default) will be specified.",
			"tl",
			"timeline [counts]",
			"timeline 50",
		),
		Func: func(c *ishell.Context) {
			counts := cfg.Default.Counts

			if len(c.Args) > 0 && util.IsNumber(c.Args[0]) {
				counts = c.Args[0]
			}

			err := tweets.LoadHomeTL(counts)
			if err == nil {
				tweets.DrawTweets()
			}
		},
	})
}
