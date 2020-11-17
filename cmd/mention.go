package cmd

import (
	"github.com/arrow2nd/twnyan/util"
	"gopkg.in/abiosoft/ishell.v2"
)

func init() {
	shell.AddCmd(&ishell.Cmd{
		Name:    "mention",
		Aliases: []string{"mt"},
		Help:    "displays the Mentions to you",
		LongHelp: createLongHelp(
			"Displays the Mentions to you.\nIf you omit the counts, the default value in the configuration file (25 by default) will be specified.",
			"mt",
			"mention [counts]",
			"mention 50",
		),
		Func: func(c *ishell.Context) {
			counts := cfg.Default.Counts

			if len(c.Args) > 0 && util.IsNumber(c.Args[0]) {
				counts = c.Args[0]
			}

			err := tweets.LoadMentionTL(counts)
			if err == nil {
				tweets.DrawTweets()
			}
		},
	})
}
