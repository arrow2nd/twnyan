package cmd

import (
	"gopkg.in/abiosoft/ishell.v2"
)

func (cmd *Cmd) newRetweetCmd() {
	rc := &ishell.Cmd{
		Name:    "retweet",
		Aliases: []string{"rt"},
		Func: func(c *ishell.Context) {
			cmd.actionOnTweet("RETWEETED", c.Cmd.Name, cmd.cfg.Color.Retweet, c.Args, cmd.api.Retweet)
		},
		Help: "retweet tweet",
		LongHelp: createLongHelp("Retweet tweet.\nIf there is more than one, please separate them with a space.",
			"rt",
			"retweet [<tweetnumber>]...",
			"retweet 0 1",
		),
	}

	rc.AddCmd(&ishell.Cmd{
		Name:    "quote",
		Aliases: []string{"qt"},
		Func: func(c *ishell.Context) {
			// 引数をチェック
			if len(c.Args) < 1 {
				cmd.drawWrongArgMessage(c.Cmd.Name)
				return
			}
			// 引用するツイートのURLを取得
			uri, err := cmd.view.GetTweetURL(c.Args[0])
			if err != nil {
				cmd.drawErrorMessage(err.Error())
				return
			}
			// 引数をパース
			status, media := cmd.parseTweetCmdArgs(c.Args[1:])
			// URLを追加してツイート
			status += " " + uri
			tweetStr, err := cmd.api.PostTweet(status, "", media)
			if err != nil {
				cmd.drawErrorMessage(err.Error())
				return
			}
			cmd.drawMessage("QUOTED", tweetStr, cmd.cfg.Color.Retweet)
		},
		Help: "quote tweet",
		LongHelp: createLongHelp(
			"Quote tweet.\nIf there is no tweet text, \"にゃーん\" will be posted.\nIf you are submitting an image, please add the file name separated by a space.",
			"qt",
			"retweet quote [<tweetnumber>] [text] [image]...",
			"retweet quote 0 cute!! cat.png",
		),
	})

	rc.AddCmd(&ishell.Cmd{
		Name:    "remove",
		Aliases: []string{"rm"},
		Func: func(c *ishell.Context) {
			cmd.actionOnTweet("UN-RETWEETED", "retweet "+c.Cmd.Name, cmd.cfg.Color.Retweet, c.Args, cmd.api.UnRetweet)
		},
		Help: "un-retweet tweet",
		LongHelp: createLongHelp(
			"UnRetweet tweet.\nIf there is more than one, please separate them with a space.",
			"rm",
			"retweet remove [<tweetnumber>]...",
			"retweet remove 0 1",
		),
	})

	cmd.shell.AddCmd(rc)
}
