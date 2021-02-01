package cmd

import (
	"github.com/arrow2nd/twnyan/api"
	"gopkg.in/abiosoft/ishell.v2"
)

func (cmd *Cmd) newTimelineCmd() {
	cmd.shell.AddCmd(&ishell.Cmd{
		Name:    "timeline",
		Aliases: []string{"tl"},
		Func: func(c *ishell.Context) {
			// 引数をパース
			counts := cmd.getCountFromCmdArg(c.Args)
			// タイムラインを取得
			v := api.CreateURLValues(counts)
			t, err := cmd.api.GetTimeline("home", v)
			if err != nil {
				cmd.drawErrorMessage(err.Error())
				return
			}
			// 描画
			cmd.view.RegisterTweets(t)
			cmd.view.DrawTweets()
		},
		Help: "get a home timeline",
		LongHelp: createLongHelp(
			"Get a home timeline.\nIf you omit the counts, the default value in the configuration file (25 by default) will be specified.",
			"tl",
			"timeline [counts]",
			"timeline 50",
		),
	})
}
