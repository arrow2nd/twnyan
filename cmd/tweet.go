package cmd

import (
	"net/url"

	"gopkg.in/abiosoft/ishell.v2"
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
		Func:    cmd.multiCmd,
		Help:    "post a tweet that contains a new line",
		LongHelp: createLongHelp(
			"Post a tweet that contains a new line.\nEnter a semicolon to end the input.\nAlso, if it is blank, the tweet will be canceled.",
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
	val := url.Values{}
	// 引数をパース
	status, files := cmd.parseTweetCmdArgs(c.Args)
	// 画像をアップロード
	if len(files) != 0 {
		mediaIDs, err := cmd.upload(files)
		if err != nil {
			cmd.drawErrorMessage(err.Error())
			return
		}
		val.Add("media_ids", mediaIDs)
	}
	// ツイート
	tweetStr, err := cmd.api.PostTweet(val, status)
	if err != nil {
		cmd.drawErrorMessage(err.Error())
		return
	}
	cmd.drawMessage("TWEETED", tweetStr, cmd.cfg.Color.Accent2)
}

func (cmd *Cmd) multiCmd(c *ishell.Context) {
	// マルチライン入力
	cmd.drawMessage("MULTILINE MODE", "Enter a semicolon to end the input", cmd.cfg.Color.Accent3)
	status := c.ReadMultiLines(";")
	// セミコロンのみならキャンセル
	if status == ";" {
		cmd.drawMessage("CANCELED", "Tweet post cancelled", cmd.cfg.Color.Accent3)
		return
	}
	// セミコロンを除去
	status = status[:len(status)-1]
	// ツイート
	tweetStr, err := cmd.api.PostTweet(url.Values{}, status)
	if err != nil {
		cmd.drawErrorMessage(err.Error())
		return
	}
	cmd.drawMessage("TWEETED", tweetStr, cmd.cfg.Color.Accent2)
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

// upload 画像をアップロード
func (cmd *Cmd) upload(medias []string) (string, error) {
	// プログレスバー開始
	cmd.shell.Print("Uploading...🐾 ")
	cmd.shell.ProgressBar().Indeterminate(true)
	cmd.shell.ProgressBar().Start()
	// アップロード
	mediaIDs, err := cmd.api.UploadImage(medias)
	cmd.shell.ProgressBar().Stop()
	if err != nil {
		return "", err
	}
	return mediaIDs, nil
}
