package cmd

import (
	"net/url"

	"github.com/arrow2nd/ishell"
	"github.com/arrow2nd/twnyan/twitter"
)

func (cmd *Cmd) newReplyCmd() *ishell.Cmd {
	// reply
	replyCmd := &ishell.Cmd{
		Name:    "reply",
		Aliases: []string{"rp"},
		Func:    cmd.execReplyCmd,
		Help:    "post a reply",
		LongHelp: createLongHelp(
			`Post a reply.
If there is no tweet text, 'にゃーん' will be posted.
If you are submitting an image, please add the file name separated by a space.`,
			"rp",
			"reply [<tweet-number>] [text] [images]...",
			"reply 2 meow cat.jpg",
		),
	}

	// reply multi
	replyCmd.AddCmd(&ishell.Cmd{
		Name:    "multi",
		Aliases: []string{"ml"},
		Func:    cmd.execReplyMultiCmd,
		Help:    "post a multi-line reply",
		LongHelp: createLongHelp(
			`Post a multi-line reply.
Enter a semicolon to end the input.
And if you want to cancel, input ":exit".`,
			"ml",
			"reply multi [<tweet-number>]",
			"reply multi 2",
		),
	})

	return replyCmd
}

func (cmd *Cmd) execReplyCmd(c *ishell.Context) {
	if len(c.Args) < 1 {
		cmd.showWrongArgMessage(c.Cmd.Name)
		return
	}

	status, files := cmd.parseTweetCmdArgs(c.Args[1:])
	cmd.execReply(c.Args[0], status, files)
}

func (cmd *Cmd) execReplyMultiCmd(c *ishell.Context) {
	if len(c.Args) < 1 {
		cmd.showWrongArgMessage("reply " + c.Cmd.Name)
		return
	}

	// 添付画像を取得
	_, images := cmd.parseTweetCmdArgs(c.Args[1:])

	text := cmd.inputMultiLine()
	if text == "" {
		return
	}

	cmd.execReply(c.Args[0], text, images)
}

func (cmd *Cmd) execReply(tweetNumStr, status string, files []string) {
	// リプライ先のツイートIDを取得
	tweetId, err := cmd.twitter.GetDataFromTweetNum(tweetNumStr, twitter.TweetId)
	if err != nil {
		cmd.showErrorMessage(err.Error())
		return
	}

	// リプライ先を設定
	query := url.Values{}
	query.Add("in_reply_to_status_id", tweetId)
	query.Add("auto_populate_reply_metadata", "true")

	// 画像をアップロード
	if err := cmd.upload(files, &query); err != nil {
		cmd.showErrorMessage(err.Error())
		return
	}

	// リプライを投稿
	tweetText, err := cmd.twitter.PostTweet(query, status)
	if err != nil {
		cmd.showErrorMessage(err.Error())
		return
	}

	cmd.showMessage("REPLYED", tweetText, cmd.config.Color.Reply)
}
