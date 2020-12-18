package cmd

import (
	"github.com/arrow2nd/twnyan/util"
	"github.com/gookit/color"
	"gopkg.in/abiosoft/ishell.v2"
)

func init() {
	showUserTL := func(screenName, counts string) {
		err := tweets.LoadUserTL(screenName, counts)
		if err != nil {
			return
		}
		tweets.DrawTweets()
		tweets.DrawUserInfo(0)
	}

	userCmd := &ishell.Cmd{
		Name:    "user",
		Aliases: []string{"ur"},
		Help:    "get a user timeline",
		LongHelp: createLongHelp(
			"Get a user timeline.\nIf you omit the counts, the default value in the configuration file (25 by default) will be specified.",
			"ur",
			"user [<username/tweetnumber>] [counts]",
			"user github 25\n  user 2",
		),
		Func: func(c *ishell.Context) {
			// 引数をパース
			user, counts, err := parseTLCmdArgs(c.Args)
			if err != nil {
				showWrongMsg(c.Cmd.Name)
				return
			}

			// スクリーンネームを取得
			if util.IsNumber(user) {
				user, err = tweets.GetDataFromTweetNum(user, "ScreenName")
				if err != nil {
					color.Error.Prompt(err.Error())
					return
				}
			}

			// 表示
			showUserTL(user, counts)
		},
	}

	userCmd.AddCmd(&ishell.Cmd{
		Name: "myuser",
		Help: "get your own timeline",
		LongHelp: createLongHelp(
			"Get your own timeline.\nIf you omit the counts, the default value in the configuration file (25 by default) will be specified.",
			"",
			"user myuser [counts]",
			"user myuser 25",
		),
		Func: func(c *ishell.Context) {
			// 取得件数
			counts := getCountsFromCmdArg(c.Args)

			// 表示
			showUserTL("", counts)
		},
	})

	shell.AddCmd(userCmd)
}
