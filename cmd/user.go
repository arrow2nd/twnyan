package cmd

import (
	"github.com/arrow2nd/ishell"
	"github.com/arrow2nd/twnyan/api"
	"github.com/arrow2nd/twnyan/util"
)

func (cmd *Cmd) newUserCmd() {
	uc := &ishell.Cmd{
		Name:    "user",
		Aliases: []string{"ur"},
		Func:    cmd.userCmd,
		Help:    "get a user timeline",
		LongHelp: createLongHelp(
			"Get a user timeline.\nIf you omit the counts, the default value in the configuration file (25 by default) will be specified.",
			"ur",
			"user [<username/tweetnumber>] [counts]",
			"user github 25\n  user 2",
		),
	}

	uc.AddCmd(&ishell.Cmd{
		Name: "own",
		Help: "get your own timeline",
		LongHelp: createLongHelp(
			"Get your own timeline.\nIf you omit the counts, the default value in the configuration file (25 by default) will be specified.",
			"",
			"user own [counts]",
			"user own 25",
		),
		Func: func(c *ishell.Context) {
			counts := cmd.getCountFromCmdArg(c.Args)
			cmd.loadUserTimeline("", counts)
		},
	})

	cmd.shell.AddCmd(uc)
}

func (cmd *Cmd) userCmd(c *ishell.Context) {
	// 引数をパース
	value, counts, err := cmd.parseTLCmdArgs(c.Args)
	if err != nil {
		cmd.drawWrongArgMessage(c.Cmd.Name)
		return
	}
	// ツイート番号ならスクリーンネームに置換
	if util.IsNumber(value) {
		value, err = cmd.view.GetDataFromTweetNum(value, "screenname")
		if err != nil {
			cmd.drawErrorMessage(err.Error())
			return
		}
	}
	// ユーザータイムラインを取得
	cmd.loadUserTimeline(value, counts)
}

func (cmd *Cmd) loadUserTimeline(screenName, counts string) {
	// ユーザータイムラインを取得
	v := api.CreateURLValues(counts)
	v.Add("screen_name", screenName)
	t, err := cmd.api.GetTimeline("user", v)
	if err != nil {
		cmd.drawErrorMessage(err.Error())
		return
	}
	// ユーザーとの関係を取得
	u := (*t)[0].User
	fs, err := cmd.api.GetFriendships(u.IdStr)
	if err != nil {
		cmd.drawErrorMessage(err.Error())
		return
	}
	// 描画
	cmd.view.RegisterTweets(t)
	cmd.view.DrawTweets()
	cmd.view.DrawUser(&u, fs)
}
