package cmd

import (
	"github.com/arrow2nd/twnyan/twitter"
	"github.com/gookit/color"
	"gopkg.in/abiosoft/ishell.v2"
)

func init() {
	shell.AddCmd(&ishell.Cmd{
		Name:    "reply",
		Aliases: []string{"rp"},
		Help:    "post a reply",
		LongHelp: createLongHelp(
			"Post a reply.\nIf there is no tweet text, \"にゃーん\" will be posted.\nIf you are submitting an image, please add the file name separated by a space.",
			"rp",
			"reply [<tweetnumber>] [text] [image]...",
			"reply 2 meow cat.jpg",
		),
		Func: func(c *ishell.Context) {
			// 引数エラー
			if len(c.Args) < 1 {
				showWrongMsg(c.Cmd.Name)
				return
			}

			// 返信先のツイートIDを取得
			tweetID, err := tweets.GetDataFromTweetNum(c.Args[0], "TweetID")
			if err != nil {
				color.Error.Prompt(err.Error())
				return
			}

			// ツイート
			status, media := parseTweetCmdArgs(c.Args[1:])
			twitter.PostTweet(status, tweetID, media)
		},
	})
}
