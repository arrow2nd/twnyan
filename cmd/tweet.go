package cmd

import (
	"net/url"

	"github.com/arrow2nd/ishell"
)

func (cmd *Cmd) addTweetCmd() {
	// tweet
	tweetCmd := &ishell.Cmd{
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

	// tweet multi
	tweetCmd.AddCmd(&ishell.Cmd{
		Name:    "multi",
		Aliases: []string{"ml"},
		Func:    cmd.tweetMultiCmd,
		Help:    "post a multi-line tweet",
		LongHelp: createLongHelp(
			"Post a multi-line tweet.\nEnter a semicolon to end the input.\nAnd if you want to cancel, press Ctrl+c on an empty line.",
			"ml",
			"tweet multi",
			"",
		),
	})

	// tweet remove
	tweetCmd.AddCmd(&ishell.Cmd{
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

	cmd.shell.AddCmd(tweetCmd)
}

func (cmd *Cmd) tweetCmd(c *ishell.Context) {
	text, images := cmd.parseTweetCmdArgs(c.Args)
	cmd.tweet(text, images)
}

func (cmd *Cmd) tweetMultiCmd(c *ishell.Context) {
	text, images := cmd.inputMultiLine()
	if text == "" {
		return
	}

	cmd.tweet(text, images)
}

func (cmd *Cmd) tweetRemoveCmd(c *ishell.Context) {
	if len(c.Args) <= 0 {
		cmd.showWrongArgMessage("tweet " + c.Cmd.Name)
		return
	}

	// 引数の数だけ削除処理
	for _, tweetNumStr := range c.Args {
		tweetID, err := cmd.view.GetDataFromTweetNum(tweetNumStr, "tweetID")
		if err != nil {
			cmd.showErrorMessage(err.Error())
			return
		}

		tweetText, err := cmd.api.DeleteTweet(tweetID)
		if err != nil {
			cmd.showErrorMessage(err.Error())
			return
		}

		cmd.showMessage("DELETED", tweetText, cmd.cfg.Color.Accent2)
	}
}

func (cmd *Cmd) tweet(text string, images []string) {
	query := url.Values{}

	// 画像をアップロード
	err := cmd.upload(images, &query)
	if err != nil {
		cmd.showErrorMessage(err.Error())
		return
	}

	// ツイートを投稿
	tweetText, err := cmd.api.PostTweet(query, text)
	if err != nil {
		cmd.showErrorMessage(err.Error())
		return
	}

	cmd.showMessage("TWEETED", tweetText, cmd.cfg.Color.Accent2)
}
