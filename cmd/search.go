package cmd

import "gopkg.in/abiosoft/ishell.v2"

func (cmd *Cmd) newSearchCmd() {
	cmd.shell.AddCmd(&ishell.Cmd{
		Name:    "search",
		Aliases: []string{"sh"},
		Func: func(c *ishell.Context) {
			keyword, counts, err := cmd.parseTLCmdArgs(c.Args)
			if err != nil {
				showWrongMsg(c.Cmd.Name)
				return
			}
			t, err := cmd.api.GetSearchResult(keyword, counts)
			if err != nil {
				return
			}
			cmd.view.RegisterTweets(t)
			cmd.view.DrawTweets()
		},
		Help: "search for tweets from the past 7 days",
		LongHelp: createLongHelp(
			"Search for tweets from the past 7 days.\nIf you omit the counts, the default value in the configuration file (25 by default) will be specified.",
			"sh",
			"search [<keyword>] [counts]",
			"search cats 50",
		),
	})
}
