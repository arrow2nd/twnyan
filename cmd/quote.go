package cmd

import (
	"net/url"

	"gopkg.in/abiosoft/ishell.v2"
)

func (cmd *Cmd) newQuoteCmd() {
	qc := &ishell.Cmd{
		Name:    "quote",
		Aliases: []string{"qt"},
		Func:    cmd.quoteCmd,
		Help:    "quote a tweet",
		LongHelp: createLongHelp(
			"Quote a tweet.\nIf there is no tweet text, 'にゃーん' will be posted.\nIf you are submitting an image, please add the file name separated by a space.",
			"qt",
			"quote [<tweetnumber>] [text] [image]...",
			"quote 0 cute!! cat.png",
		),
	}

	qc.AddCmd(&ishell.Cmd{
		Name:    "multi",
		Aliases: []string{"ml"},
		Func:    cmd.quoteMultiCmd,
		Help:    "post a multi-line quote tweet",
		LongHelp: createLongHelp(
			"Post a multi-line quote tweet.\nEnter a semicolon to end the input.\nAlso, if it is blank, the tweet will be canceled.",
			"ml",
			"quote multi [<tweetnumber>]",
			"quote multi 0",
		),
	})

	cmd.shell.AddCmd(qc)
}

func (cmd *Cmd) quoteCmd(c *ishell.Context) {
	// 引数をチェック
	if len(c.Args) < 1 {
		cmd.drawWrongArgMessage(c.Cmd.Name)
		return
	}
	// 引数をパース
	status, files := cmd.parseTweetCmdArgs(c.Args[1:])
	// 引用ツイート
	cmd.quote(c, status, files)
}

func (cmd *Cmd) quoteMultiCmd(c *ishell.Context) {
	// 引数をチェック
	if len(c.Args) < 1 {
		cmd.drawWrongArgMessage("quote " + c.Cmd.Name)
		return
	}
	// 入力
	status, files := cmd.inputMultiLine()
	// 引用ツイート
	cmd.quote(c, status, files)
}

func (cmd *Cmd) quote(c *ishell.Context, status string, files []string) {
	val := url.Values{}
	// 引用するツイートのURLを取得
	uri, err := cmd.view.GetTweetURL(c.Args[0])
	if err != nil {
		cmd.drawErrorMessage(err.Error())
		return
	}
	// 画像をアップロード
	err = cmd.upload(files, &val)
	if err != nil {
		cmd.drawErrorMessage(err.Error())
		return
	}
	// URLを追加してツイート
	status += " " + uri
	tweetStr, err := cmd.api.PostTweet(val, status)
	if err != nil {
		cmd.drawErrorMessage(err.Error())
		return
	}
	cmd.drawMessage("QUOTED", tweetStr, cmd.cfg.Color.Retweet)
}
