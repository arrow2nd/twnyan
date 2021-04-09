package cmd

import (
	"net/url"

	"github.com/arrow2nd/ishell"
)

func (cmd *Cmd) addQuoteCmd() {
	// quote
	quoteCmd := &ishell.Cmd{
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

	// quote multi
	quoteCmd.AddCmd(&ishell.Cmd{
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

	cmd.shell.AddCmd(quoteCmd)
}

func (cmd *Cmd) quoteCmd(c *ishell.Context) {
	if len(c.Args) < 1 {
		cmd.showWrongArgMessage(c.Cmd.Name)
		return
	}

	status, files := cmd.parseTweetCmdArgs(c.Args[1:])
	cmd.quote(c, status, files)
}

func (cmd *Cmd) quoteMultiCmd(c *ishell.Context) {
	if len(c.Args) < 1 {
		cmd.showWrongArgMessage("quote " + c.Cmd.Name)
		return
	}

	status, files := cmd.inputMultiLine()
	cmd.quote(c, status, files)
}

func (cmd *Cmd) quote(c *ishell.Context, text string, images []string) {
	query := url.Values{}

	// 引用するツイートのURLを取得
	tweetUrl, err := cmd.view.GetTweetURL(c.Args[0])
	if err != nil {
		cmd.showErrorMessage(err.Error())
		return
	}

	// 画像をアップロード
	err = cmd.upload(images, &query)
	if err != nil {
		cmd.showErrorMessage(err.Error())
		return
	}

	// URLを追加してツイート
	text += " " + tweetUrl
	tweetText, err := cmd.api.PostTweet(query, text)
	if err != nil {
		cmd.showErrorMessage(err.Error())
		return
	}

	cmd.showMessage("QUOTED", tweetText, cmd.cfg.Color.Retweet)
}
