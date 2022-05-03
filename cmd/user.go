package cmd

import (
	"github.com/arrow2nd/ishell/v2"
	"github.com/arrow2nd/twnyan/twitter"
	"github.com/arrow2nd/twnyan/util"
)

func (cmd *Cmd) newUserCmd() *ishell.Cmd {
	// user
	userCmd := &ishell.Cmd{
		Name:    "user",
		Aliases: []string{"ur"},
		Func:    cmd.execUserCmd,
		Help:    "get a user timeline",
		LongHelp: createLongHelp(
			`Get a user timeline.
If you omit the counts, the default value in the configuration file (25 by default) will be specified.`,
			"ur",
			"user [<username / tweet-number>] [counts]",
			"user github 25\n  user 2",
		),
	}

	// use own
	userCmd.AddCmd(&ishell.Cmd{
		Name: "own",
		Help: "get your own timeline",
		LongHelp: createLongHelp(
			`Get your own timeline.
If you omit the counts, the default value in the configuration file (25 by default) will be specified.`,
			"",
			"user own [counts]",
			"user own 25",
		),
		Func: func(c *ishell.Context) {
			count := cmd.getCountFromCmdArg(c.Args)
			cmd.showUserTimeline("", count)
		},
	})

	return userCmd
}

func (cmd *Cmd) execUserCmd(c *ishell.Context) {
	screenName, count, err := cmd.parseTimelineCmdArgs(c.Args)
	if err != nil {
		cmd.showWrongArgMessage(c.Cmd.Name)
		return
	}

	// ツイート番号ならスクリーンネームに置換
	if util.IsThreeDigitsNumber(screenName) {
		screenName, err = cmd.twitter.GetDataFromTweetNum(screenName, twitter.ScreenName)
		if err != nil {
			cmd.showErrorMessage(err.Error())
			return
		}
	}

	cmd.showUserTimeline(screenName, count)
}

// showUserTimeline ユーザタイムラインを表示
func (cmd *Cmd) showUserTimeline(screenName, count string) {
	query := twitter.CreateQuery(count)
	query.Add("screen_name", screenName)

	// ユーザのツイートを取得
	tweets, err := cmd.twitter.FetchTimelineTweets(twitter.User, query)
	if err != nil {
		cmd.showErrorMessage(err.Error())
		return
	}

	user := cmd.twitter.OwnUser
	relationships := []string{}

	if screenName != "" {
		// ユーザ情報を取得
		user, err = cmd.twitter.FetchUserInfo(screenName)
		if err != nil {
			cmd.showErrorMessage(err.Error())
			return
		}

		// ユーザとの関係を取得
		relationships, err = cmd.twitter.FetchRelationships(user.IdStr)
		if err != nil {
			cmd.showErrorMessage(err.Error())
			return
		}
	}

	cmd.twitter.RegisterTweets(tweets)
	cmd.showTweets()

	cmd.view.ShowUserInfo(user, relationships)
}
