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

// parseTweetCmdArgs ツイート系のコマンドの引数をパース
func (cmd *Cmd) parseTweetCmdArgs(args []string) (string, []string) {
	// 引数がないならにゃーん
	if len(args) == 0 {
		return "にゃーん", []string{}
	}

	// 1つ目の引数に拡張子が含まれているなら画像パスのみを返す
	if util.MatchesRegexp("\\.\\w{3,4}$", args[0]) {
		return "", args[0:]
	}

	// ツイート文と画像パスを返す
	return args[0], args[1:]
}

// parseTimelineCmdArgs タイムライン取得系のコマンドの引数をパース
func (cmd *Cmd) parseTimelineCmdArgs(args []string) (string, string, error) {
	if len(args) <= 0 {
		return "", "", errors.New("no arguments")
	}

	str, count := args[0], cmd.cfg.Option.Counts

	// 2つ目の引数があればcountに代入
	if len(args) >= 2 {
		count = args[1]
	}

	return str, count, nil
}

// getCountFromCmdArg 引数からツイート取得件数を取得
func (cmd *Cmd) getCountFromCmdArg(args []string) string {
	// 引数が無い、または数値以外ならデフォルト値を返す
	if len(args) <= 0 || !util.IsThreeDigitsNumber(args[0]) {
		return cmd.cfg.Option.Counts
	}

	return args[0]
}

// inputMultiLine マルチラインツイート入力
func (cmd *Cmd) inputMultiLine() string {
	// プロンプトを変更
	cmd.shell.SetPrompt("... ")
	defer cmd.setDefaultPrompt()

	cmd.showMessage(
		"INPUT",
		"End typing with a semicolon (If you want to cancel, input \":exit\").",
		cmd.cfg.Color.Accent3,
	)

	input := cmd.shell.ReadMultiLinesFunc(func(f string) bool {
		if f == ":exit" || strings.HasSuffix(f, ";") {
			return false
		}
		return true
	})

	// 文字列内に:exitがあればキャンセル
	if strings.Contains(input, ":exit") {
		cmd.showMessage("CANCELED", "Input interrupted.", cmd.cfg.Color.Accent2)
		return ""
	}

	return strings.TrimRight(input, ";")
}

// upload 画像をアップロード
func (cmd *Cmd) upload(images []string, query *url.Values) error {
	if len(images) <= 0 {
		return nil
	}

	// プログレスバー開始
	fmt.Print("Uploading...🐾 ")
	cmd.shell.ProgressBar().Indeterminate(true)
	cmd.shell.ProgressBar().Start()

	mediaIDs, err := cmd.api.UploadImage(images)
	cmd.shell.ProgressBar().Stop()
	if err != nil {
		return err
	}

	query.Add("media_ids", mediaIDs)

	return nil
}

// actionOnTweet ツイートに対しての操作
func (cmd *Cmd) actionOnTweet(actionName, cmdName, bgColor string, args []string, actionFunc func(string) (string, error)) {
	if len(args) <= 0 {
		cmd.showWrongArgMessage(cmdName)
		return
	}

	// 引数の数だけ処理
	for _, v := range args {
		tweetID, err := cmd.view.GetDataFromTweetNum(v, "tweetID")
		if err != nil {
			cmd.showErrorMessage(err.Error())
			return
		}

		tweetText, err := actionFunc(tweetID)
		if err != nil {
			cmd.showErrorMessage(err.Error())
			return
		}

		cmd.showMessage(actionName, tweetText, bgColor)
	}
}

// actionOnUser ユーザーに対しての操作
func (cmd *Cmd) actionOnUser(actionName, cmdName, bgColor string, args []string, actionFunc func(string) (string, error)) {
	var err error

	if len(args) <= 0 {
		cmd.showWrongArgMessage(cmdName)
		return
	}

	screenName := args[0]

	// ツイート番号ならスクリーンネームに置換
	if util.IsThreeDigitsNumber(args[0]) {
		screenName, err = cmd.view.GetDataFromTweetNum(args[0], "screenname")
		if err != nil {
			cmd.showErrorMessage(err.Error())
			return
		}
	}

	// 受け取った関数を実行
	userName, err := actionFunc(screenName)
	if err != nil {
		cmd.showErrorMessage(err.Error())
		return
	}

	cmd.showMessage(actionName, userName, bgColor)
}

// showMessage メッセージを表示
func (cmd *Cmd) showMessage(tips, msg, bgColor string) {
	width := util.GetWindowWidth()

	// 不要な文字を削除
	msg = util.AllReplace(msg, "[\t\n\r]", " ")
	msg = html.UnescapeString(msg)

	// 画面内に収まるよう丸める
	msg = util.TruncateString(msg, width-len(tips)-3)

	color.HEXStyle(cmd.cfg.Color.BoxForground, bgColor).Printf(" %s ", tips)
	fmt.Printf(" %s\n", msg)
}

// showErrorMessage エラーメッセージを表示
func (cmd *Cmd) showErrorMessage(msg string) {
	width := util.GetWindowWidth()

	msg = util.TruncateString(msg, width-9)
	errMsg := color.HEXStyle(cmd.cfg.Color.BoxForground, cmd.cfg.Color.Error).Sprintf(" ERROR: %s ", msg)

	fmt.Printf("%s\n", errMsg)
}

// drawWrongArgError 引数ミスのメッセージを表示
func (cmd *Cmd) showWrongArgMessage(cmdName string) {
	msg := fmt.Sprintf("Wrong argument, try '%s help'", cmdName)
	cmd.showErrorMessage(msg)
}

// createLongHelp 詳細なヘルプ文を作成
func createLongHelp(help, alias, use, exp string) string {
	longHelp := fmt.Sprintf("%s\n", help)

	if alias != "" {
		longHelp += fmt.Sprintf("\nAlias:\n  %s\n", alias)
	}

	if use != "" {
		longHelp += fmt.Sprintf("\nUse:\n  %s\n", use)
	}

	if exp != "" {
		longHelp += fmt.Sprintf("\nExample:\n  %s\n", exp)
	}

	return longHelp
}
