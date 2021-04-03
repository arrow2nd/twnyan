package cmd

import (
	"github.com/arrow2nd/ishell"
	"github.com/arrow2nd/twnyan/api"
)

func (cmd *Cmd) newMentionCmd() {
	cmd.shell.AddCmd(&ishell.Cmd{
		Name:    "mention",
		Aliases: []string{"mt"},
		Func:    cmd.mentionCmd,
		Help:    "get a Mentions to you",
		LongHelp: createLongHelp(
			"Get a Mentions to you.\nIf you omit the counts, the default value in the configuration file (25 by default) will be specified.",
			"mt",
			"mention [counts]",
			"mention 50",
		),
	})
}

func (cmd *Cmd) mentionCmd(c *ishell.Context) {
	// 引数をパース
	counts := cmd.getCountFromCmdArg(c.Args)

	// メンションタイムラインを取得
	v := api.CreateURLValues(counts)
	t, err := cmd.api.GetTimeline("mention", v)
	if err != nil {
		cmd.drawErrorMessage(err.Error())
		return
	}

	// 描画
	cmd.view.RegisterTweets(t)
	cmd.view.DrawTweets()
}
