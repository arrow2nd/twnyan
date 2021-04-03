package cmd

import (
	"errors"
	"fmt"
	"html"
	"net/url"
	"strings"

	"github.com/arrow2nd/twnyan/util"
	"github.com/gookit/color"
)

// setDefaultPrompt デフォルトのプロンプトを設定
func (cmd *Cmd) setDefaultPrompt() {
	prompt := fmt.Sprintf("@%s : ", cmd.api.OwnUser.ScreenName)
	cmd.shell.SetPrompt(prompt)
}

// parseTweetCmdArgs ツイート系コマンドの引数をパース
func (cmd *Cmd) parseTweetCmdArgs(args []string) (string, []string) {
	status, images := "にゃーん", []string{}

	if len(args) > 0 {
		if util.ContainsStr("\\.\\w{3,4}$", args[0]) {
			status = ""
			images = args[0:]
		} else {
			status = args[0]
			images = args[1:]
		}
	}

	return status, images
}

// parseTLCmdArgs タイムライン系コマンドの引数をパース
func (cmd *Cmd) parseTLCmdArgs(args []string) (string, string, error) {
	// 引数をチェック
	if len(args) <= 0 {
		return "", "", errors.New("no arguments")
	}
	str, counts := args[0], cmd.cfg.Option.Counts

	// 取得件数の指定があれば置換
	if len(args) >= 2 {
		counts = args[1]
	}

	return str, counts, nil
}

// getCountFromCmdArg 引数から取得件数を取得
func (cmd *Cmd) getCountFromCmdArg(args []string) string {
	// 引数無し、数値以外ならデフォルト値を返す
	if len(args) <= 0 || !util.IsNumber(args[0]) {
		return cmd.cfg.Option.Counts
	}

	return args[0]
}

// inputMultiLine マルチラインツイート入力
func (cmd *Cmd) inputMultiLine() (string, []string) {
	// プロンプトを変更
	cmd.shell.SetPrompt("... ")
	defer cmd.setDefaultPrompt()

	// ツイート文入力
	cmd.drawMessage("INPUT", "End typing with a semicolon (cancel with Ctrl+c on an empty line)", cmd.cfg.Color.Accent3)
	text := cmd.shell.ReadMultiLines(";")
	if util.IsEndLFCode(text) {
		cmd.drawMessage("CANCELED", "Canceled input", cmd.cfg.Color.Accent2)
		return "", nil
	}

	// 添付画像ファイル名入力
	cmd.drawMessage("IMAGE", "Enter the file name of the attached image (separated by a space)", cmd.cfg.Color.Accent3)
	img := cmd.shell.ReadLine()
	if util.IsEndLFCode(img) {
		cmd.drawMessage("CANCELED", "Canceled input", cmd.cfg.Color.Accent2)
		return "", nil
	}

	// 戻り値を作成
	tweet := strings.TrimRight(text, ";")
	images := strings.Fields(img)

	return tweet, images
}

// upload 画像をアップロード
func (cmd *Cmd) upload(files []string, val *url.Values) error {
	// ファイルが無ければ処理しない
	if len(files) <= 0 {
		return nil
	}

	// プログレスバー開始
	fmt.Print("Uploading...🐾 ")
	cmd.shell.ProgressBar().Indeterminate(true)
	cmd.shell.ProgressBar().Start()

	// アップロード
	mediaIDs, err := cmd.api.UploadImage(files)
	cmd.shell.ProgressBar().Stop()
	if err != nil {
		return err
	}

	// media_idsを追加
	val.Add("media_ids", mediaIDs)

	return nil
}

// actionOnTweet ツイートに対しての操作
func (cmd *Cmd) actionOnTweet(actionName, cmdName, bgColor string, args []string, actionFunc func(string) (string, error)) {
	// 引数をチェック
	if len(args) <= 0 {
		cmd.drawWrongArgMessage(cmdName)
		return
	}

	// 引数の数だけ処理
	for _, v := range args {
		id, err := cmd.view.GetDataFromTweetNum(v, "tweetID")
		if err != nil {
			cmd.drawErrorMessage(err.Error())
			return
		}

		tweetStr, err := actionFunc(id)
		if err != nil {
			cmd.drawErrorMessage(err.Error())
			return
		}

		cmd.drawMessage(actionName, tweetStr, bgColor)
	}
}

// actionOnUser ユーザーに対しての操作
func (cmd *Cmd) actionOnUser(actionName, cmdName, bgColor string, args []string, actionFunc func(string) (string, error)) {
	var err error

	// 引数をチェック
	if len(args) <= 0 {
		cmd.drawWrongArgMessage(cmdName)
		return
	}

	// ツイート番号ならスクリーンネームに置換
	screenName := args[0]
	if util.IsNumber(args[0]) {
		screenName, err = cmd.view.GetDataFromTweetNum(args[0], "screenname")
		if err != nil {
			cmd.drawErrorMessage(err.Error())
			return
		}
	}

	// 処理を実行
	userStr, err := actionFunc(screenName)
	if err != nil {
		cmd.drawErrorMessage(err.Error())
		return
	}

	cmd.drawMessage(actionName, userStr, bgColor)
}

// drawMessage メッセージを表示
func (cmd *Cmd) drawMessage(tips, text, bgColor string) {
	width := util.GetWindowWidth()
	util.AllReplace(&text, "[\t\n\r]", " ")
	text = html.UnescapeString(text)
	text = util.TruncateStr(text, width-len(tips)-3)
	tips = color.HEXStyle(cmd.cfg.Color.BoxForground, bgColor).Sprintf(" %s ", tips)
	fmt.Printf("%s %s\n", tips, text)
}

// drawErrorMsg エラーメッセージを表示
func (cmd *Cmd) drawErrorMessage(text string) {
	width := util.GetWindowWidth()
	text = util.TruncateStr(text, width-9)
	errMsg := color.HEXStyle(cmd.cfg.Color.BoxForground, cmd.cfg.Color.Error).Sprintf(" ERROR: %s ", text)
	fmt.Printf("%s\n", errMsg)
}

// drawWrongArgError 引数ミスのメッセージを表示
func (cmd *Cmd) drawWrongArgMessage(cmdName string) {
	text := fmt.Sprintf("Wrong argument, try '%s help'", cmdName)
	cmd.drawErrorMessage(text)
}

// createLongHelp 詳細なヘルプ文を作成
func createLongHelp(help, alias, use, exp string) string {
	return fmt.Sprintf("%s\n\nAlias:\n  %s\n\nUse:\n  %s\n\nExample:\n  %s", help, alias, use, exp)
}
