package cmd

import (
	"io/ioutil"
	"net/url"
	"os"
	"syscall"

	"github.com/arrow2nd/ishell/v2"
	"github.com/arrow2nd/twnyan/twitter"
	"golang.org/x/term"
)

func (cmd *Cmd) newTweetCmd() *ishell.Cmd {
	// tweet
	tweetCmd := &ishell.Cmd{
		Name:    "tweet",
		Aliases: []string{"tw"},
		Func:    cmd.execTweetCmd,
		Help:    "post a tweet",
		LongHelp: createLongHelp(
			`Post a tweet.
If there is no tweet text, 'にゃーん' will be posted.
If you are submitting an image, please add the file name separated by a space.`,
			"tw",
			"tweet [text] [images]...",
			"tweet meow🐱 cat.png supercat.jpg",
		),
	}

	// tweet multi
	tweetCmd.AddCmd(&ishell.Cmd{
		Name:    "multi",
		Aliases: []string{"ml"},
		Func:    cmd.execTweetMultiCmd,
		Help:    "post a multi-line tweet",
		LongHelp: createLongHelp(
			`Post a multi-line tweet.
Enter a semicolon to end the input.
And if you want to cancel, input ":exit".`,
			"ml",
			"tweet multi [images]...",
			"",
		),
	})

	// tweet remove
	tweetCmd.AddCmd(&ishell.Cmd{
		Name:    "remove",
		Aliases: []string{"rm"},
		Func:    cmd.execTweetRemoveCmd,
		Help:    "delete a tweet",
		LongHelp: createLongHelp(
			`Delete a tweet.
If there is more than one, please separate them with a space.`,
			"rm",
			"tweet remove [<tweet-number>]",
			"tweet remove 0 1",
		),
	})

	return tweetCmd
}

func (cmd *Cmd) execTweetCmd(c *ishell.Context) {
	// パイプからの入力を処理
	if len(c.Args) == 0 && !term.IsTerminal(syscall.Stdin) {
		stdin, _ := ioutil.ReadAll(os.Stdin)
		cmd.tweet(string(stdin), nil)
		return
	}

	cmd.tweet(cmd.parseTweetCmdArgs(c.Args))
}

func (cmd *Cmd) execTweetMultiCmd(c *ishell.Context) {
	// 添付画像を取得
	_, images := cmd.parseTweetCmdArgs(c.Args)

	text := cmd.inputMultiLine()
	if text == "" {
		return
	}

	cmd.tweet(text, images)
}

func (cmd *Cmd) execTweetRemoveCmd(c *ishell.Context) {
	if len(c.Args) <= 0 {
		cmd.showWrongArgMessage("tweet " + c.Cmd.Name)
		return
	}

	// 引数の数だけ削除処理
	for _, tweetNumStr := range c.Args {
		tweetId, err := cmd.twitter.GetDataFromTweetNum(tweetNumStr, twitter.TweetId)
		if err != nil {
			cmd.showErrorMessage(err.Error())
			return
		}

		tweetText, err := cmd.twitter.DeleteTweet(tweetId)
		if err != nil {
			cmd.showErrorMessage(err.Error())
			return
		}

		cmd.showMessage("DELETED", tweetText, cmd.config.Color.Accent3)
	}
}

func (cmd *Cmd) tweet(text string, images []string) {
	query := url.Values{}

	// 画像をアップロード
	if err := cmd.upload(images, &query); err != nil {
		cmd.showErrorMessage(err.Error())
		return
	}

	// ツイートを投稿
	tweetText, err := cmd.twitter.PostTweet(query, text)
	if err != nil {
		cmd.showErrorMessage(err.Error())
		return
	}

	cmd.showMessage("TWEETED", tweetText, cmd.config.Color.Accent3)
}
