package cmd

import (
	"net/url"

	"github.com/arrow2nd/ishell"
)

func (cmd *Cmd) newReplyCmd() {
	rc := &ishell.Cmd{
		Name:    "reply",
		Aliases: []string{"rp"},
		Func:    cmd.replyCmd,
		Help:    "post a reply",
		LongHelp: createLongHelp(
			"Post a reply.\nIf there is no tweet text, 'にゃーん' will be posted.\nIf you are submitting an image, please add the file name separated by a space.",
			"rp",
			"reply [<tweetnumber>] [text] [image]...",
			"reply 2 meow cat.jpg",
		),
	}

	rc.AddCmd(&ishell.Cmd{
		Name:    "multi",
		Aliases: []string{"ml"},
		Func:    cmd.replyMultiCmd,
		Help:    "post a multi-line reply",
		LongHelp: createLongHelp(
			"Post a multi-line reply.\nEnter a semicolon to end the input.\nAlso, if it is blank, the tweet will be canceled.",
			"ml",
			"reply multi [<tweetnumber>]",
			"reply multi 2",
		),
	})

	cmd.shell.AddCmd(rc)
}

func (cmd *Cmd) replyCmd(c *ishell.Context) {
	// 引数をチェック
	if len(c.Args) < 1 {
		cmd.drawWrongArgMessage(c.Cmd.Name)
		return
	}

	// 引数をパース
	status, files := cmd.parseTweetCmdArgs(c.Args[1:])

	// リプライ
	cmd.reply(c, status, files)
}

func (cmd *Cmd) replyMultiCmd(c *ishell.Context) {
	// 引数をチェック
	if len(c.Args) < 1 {
		cmd.drawWrongArgMessage("reply " + c.Cmd.Name)
		return
	}

	// 入力
	status, files := cmd.inputMultiLine()
	if status == "" {
		return
	}

	// リプライ
	cmd.reply(c, status, files)
}

func (cmd *Cmd) reply(c *ishell.Context, status string, files []string) {
	// リプライ先のツイートIDを取得
	tweetID, err := cmd.view.GetDataFromTweetNum(c.Args[0], "tweetID")
	if err != nil {
		cmd.drawErrorMessage(err.Error())
		return
	}

	// リプライ先を設定
	val := url.Values{}
	val.Add("in_reply_to_status_id", tweetID)
	val.Add("auto_populate_reply_metadata", "true")

	// 画像をアップロード
	err = cmd.upload(files, &val)
	if err != nil {
		cmd.drawErrorMessage(err.Error())
		return
	}

	// リプライ
	tweetStr, err := cmd.api.PostTweet(val, status)
	if err != nil {
		cmd.drawErrorMessage(err.Error())
		return
	}

	cmd.drawMessage("REPLYED", tweetStr, cmd.cfg.Color.Reply)
}
