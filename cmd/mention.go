package cmd

import (
	"github.com/arrow2nd/twnyan/api"
	"gopkg.in/abiosoft/ishell.v2"
)

func (cmd *Cmd) newMentionCmd() {
	cmd.shell.AddCmd(&ishell.Cmd{
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
			counts := cmd.getCountsFromCmdArg(c.Args)
			v := api.CreateURLValues(counts)
			t, err := cmd.api.GetTimeline("mention", v)
			if err != nil {
				return
			}
			cmd.view.RegisterTweets(t)
			cmd.view.DrawTweets()
		},
	})
}
