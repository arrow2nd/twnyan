package cmd

import (
	"github.com/arrow2nd/twnyan/util"
	"github.com/gookit/color"
	"gopkg.in/abiosoft/ishell.v2"
)

func init() {
	userCmd := &ishell.Cmd{
		Name:    "user",
		Aliases: []string{"ur"},
		Help:    "get a user timeline",
		LongHelp: createLongHelp(
			"Get a user timeline.\nIf you don't specify a user name, you will see your own timeline.\nIf you omit the counts, the default value in the configuration file (25 by default) will be specified.",
			"ur",
			"user [user name] [counts] [data format(json|yaml)]",
			"user github 25",
		),
		Func: func(c *ishell.Context) {
			user, counts, dataFmt := "", "", ""

			// 引数をパース
			args, _ := util.FetchStringSpecifiedType(c.Args, "str", "num", "str")
			if args != nil {
				user, counts, dataFmt = args[0], args[1], args[2]
			}
			if counts == "" {
				counts = cfg.Default.Counts
			}

			showUserTL(user, counts, dataFmt)
		},
	}

	userCmd.AddCmd(&ishell.Cmd{
		Name:    "number",
		Aliases: []string{"num", "no"},
		Help:    "get the user timeline from the tweet number",
		LongHelp: createLongHelp(
			"Get the user timeline from the tweet number.\nIf you omit the counts, the default value in the configuration file (25 by default) will be specified.",
			"no",
			"user number [tweet number] [counts]",
			"user number 1 25",
		),
		Func: func(c *ishell.Context) {
			tweetNum, counts := "", ""

			// 引数をパース
			args, _ := util.FetchStringSpecifiedType(c.Args, "num", "num")
			if args != nil {
				tweetNum, counts = args[0], args[1]
			}
			if counts == "" {
				counts = cfg.Default.Counts
			}

			// ツイート番号からスクリーンネームを取得
			screenName, err := tweets.GetDataFromTweetNum(tweetNum, "ScreenName")
			if err != nil {
				color.Error.Tips(err.Error())
				return
			}

			showUserTL(screenName, counts, "")
		},
	})

	shell.AddCmd(userCmd)
}

func showUserTL(screenName, counts, dataFmt string) {
	err := tweets.LoadUserTL(screenName, counts)
	if err != nil {
		return
	}

	if dataFmt == "" {
		tweets.DrawTweets()
		tweets.DrawUserInfo(0)
	} else {
		tweets.OutData(dataFmt)
	}
}
