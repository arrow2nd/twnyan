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
		Help:    "displays the user's timeline",
		LongHelp: createLongHelp(
			"Displays the user's timeline.\nIf you don't specify a user name, you will see your own timeline.\nIf you omit the counts, the default value in the configuration file (25 by default) will be specified.",
			"ur",
			"user [user name] [counts]",
			"user github 25",
		),
		Func: getUserTimeline,
	}

	userCmd.AddCmd(&ishell.Cmd{
		Name:    "number",
		Aliases: []string{"num", "no"},
		Help:    "Displays the user timeline of the person who posted the tweet",
		LongHelp: createLongHelp(
			"Displays the user timeline of the person who posted the tweet.\nIf you omit the counts, the default value in the configuration file (25 by default) will be specified.",
			"no",
			"user number [tweet number] [counts]",
			"user number 1 25",
		),
		Func: getUserTimelineFromTweetNum,
	})

	shell.AddCmd(userCmd)
}

func getUserTimeline(c *ishell.Context) {
	user, counts := "", ""

	// 引数をパース
	args, _ := util.FetchStringSpecifiedType(c.Args, "str", "num")
	if args != nil {
		user, counts = args[0], args[1]
	}
	if counts == "" {
		counts = cfg.Default.Counts
	}

	err := tweets.LoadUserTL(user, counts)
	if err == nil {
		tweets.DrawTweets()
		tweets.DrawUserInfo(0)
	}
}

func getUserTimelineFromTweetNum(c *ishell.Context) {
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

	err = tweets.LoadUserTL(screenName, counts)
	if err == nil {
		tweets.DrawTweets()
		tweets.DrawUserInfo(0)
	}
}
