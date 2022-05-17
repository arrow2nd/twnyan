package cmd

import (
	"net/url"

	"github.com/arrow2nd/ishell/v2"
)

func (cmd *Cmd) newQuoteCmd() *ishell.Cmd {
	// quote
	quoteCmd := &ishell.Cmd{
		Name:    "quote",
		Aliases: []string{"qt"},
		Func:    cmd.execQuoteCmd,
		Help:    "quote a tweet",
		LongHelp: createLongHelp(
			`Quote a tweet.
If there is no tweet text, 'にゃーん' will be posted.
If you are submitting an image, please add the file name separated by a space.
(Only available in interactive mode)`,
			"qt",
			"quote <tweet-number> [text] [image]...",
			"quote 0 cute!! cat.png",
		),
	}

	// quote multi
	quoteCmd.AddCmd(&ishell.Cmd{
		Name:    "multi",
		Aliases: []string{"ml"},
		Func:    cmd.execQuoteMultiCmd,
		Help:    "post a multi-line quote tweet",
		LongHelp: createLongHelp(
			`Post a multi-line quote tweet.
Enter a semicolon to end the input.
And if you want to cancel, input ":exit".`,
			"ml",
			"quote multi <tweet-number> [image]...",
			"quote multi 0 apple.png",
		),
	})

	return quoteCmd
}

func (cmd *Cmd) execQuoteCmd(c *ishell.Context) {
	if len(c.Args) < 1 {
		cmd.showWrongArgMessage(c.Cmd.Name)
		return
	}

	status, files := cmd.parseTweetCmdArgs(c.Args[1:])
	cmd.execQuote(c.Args[0], status, files)
}

func (cmd *Cmd) execQuoteMultiCmd(c *ishell.Context) {
	if len(c.Args) < 1 {
		cmd.showWrongArgMessage("quote " + c.Cmd.Name)
		return
	}

	// 添付画像を取得
	_, images := cmd.parseTweetCmdArgs(c.Args[1:])

	text := cmd.inputMultiLine()
	if text == "" {
		return
	}

	cmd.execQuote(c.Args[0], text, images)
}

func (cmd *Cmd) execQuote(tweetNumStr, text string, images []string) {
	if cmd.checkCommandLineMode() {
		return
	}

	query := url.Values{}

	// 引用するツイートのURLを取得
	tweetUrl, err := cmd.twitter.GetTweetURL(tweetNumStr)
	if err != nil {
		cmd.showErrorMessage(err.Error())
		return
	}

	// 画像をアップロード
	if err := cmd.upload(images, &query); err != nil {
		cmd.showErrorMessage(err.Error())
		return
	}

	// URLを追加してツイート
	text += " " + tweetUrl
	tweetText, err := cmd.twitter.PostTweet(query, text)
	if err != nil {
		cmd.showErrorMessage(err.Error())
		return
	}

	cmd.showMessage("QUOTED", tweetText, cmd.config.Color.Retweet)
}
