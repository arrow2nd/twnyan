package cmd

import (
	"net/url"

	"github.com/arrow2nd/ishell"
)

func (cmd *Cmd) newTweetCmd() {
	tc := &ishell.Cmd{
		Name:    "tweet",
		Aliases: []string{"tw"},
		Func:    cmd.tweetCmd,
		Help:    "post a tweet",
		LongHelp: createLongHelp(
			"Post a tweet.\nIf there is no tweet text, 'にゃーん' will be posted.\nIf you are submitting an image, please add the file name separated by a space.",
			"tw",
			"tweet [text] [image]...",
			"tweet meow🐱 cat.png supercat.jpg",
		),
	}

	tc.AddCmd(&ishell.Cmd{
		Name:    "multi",
		Aliases: []string{"ml"},
		Func:    cmd.tweetMultiCmd,
		Help:    "post a multi-line tweet",
		LongHelp: createLongHelp(
			"Post a multi-line tweet.\nEnter a semicolon to end the input.\nAnd if you want to cancel, press Ctrl+c on an empty line.",
			"ml",
			"tweet multi",
			"tweet multi",
		),
	})

	tc.AddCmd(&ishell.Cmd{
		Name:    "remove",
		Aliases: []string{"rm"},
		Func:    cmd.tweetRemoveCmd,
		Help:    "delete a tweet",
		LongHelp: createLongHelp(
			"Delete a tweet.\nIf there is more than one, please separate them with a space.",
			"rm",
			"tweet remove [<tweetnumber>]",
			"tweet remove 0 1",
		),
	})

	cmd.shell.AddCmd(tc)
}

func (cmd *Cmd) tweetCmd(c *ishell.Context) {
	// 引数をパース
	status, files := cmd.parseTweetCmdArgs(c.Args)
	// ツイート
	cmd.tweet(status, files)
}

func (cmd *Cmd) tweetMultiCmd(c *ishell.Context) {
	// 入力
	status, files := cmd.inputMultiLine()
	if status == "" {
		return
	}
	// ツイート
	cmd.tweet(status, files)
}

func (cmd *Cmd) tweetRemoveCmd(c *ishell.Context) {
	// 引数をチェック
	if len(c.Args) <= 0 {
		cmd.drawWrongArgMessage("tweet " + c.Cmd.Name)
		return
	}

	// 引数の数だけ削除処理
	for _, v := range c.Args {
		id, err := cmd.view.GetDataFromTweetNum(v, "tweetID")
		if err != nil {
			cmd.drawErrorMessage(err.Error())
			return
		}

		tweetStr, err := cmd.api.DeleteTweet(id)
		if err != nil {
			cmd.drawErrorMessage(err.Error())
			return
		}

		cmd.drawMessage("DELETED", tweetStr, cmd.cfg.Color.Accent2)
	}
}

func (cmd *Cmd) tweet(status string, files []string) {
	val := url.Values{}

	// 画像をアップロード
	err := cmd.upload(files, &val)
	if err != nil {
		cmd.drawErrorMessage(err.Error())
		return
	}

	// ツイート
	tweetStr, err := cmd.api.PostTweet(val, status)
	if err != nil {
		cmd.drawErrorMessage(err.Error())
		return
	}

	cmd.drawMessage("TWEETED", tweetStr, cmd.cfg.Color.Accent2)
}
