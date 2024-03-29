package cmd

import (
	"errors"
	"fmt"
	"html"
	"net/url"
	"os"
	"strings"

	"github.com/arrow2nd/twnyan/twitter"
	"github.com/arrow2nd/twnyan/util"
	"github.com/gookit/color"
)

// setDefaultPrompt デフォルトのプロンプトを設定
func (cmd *Cmd) setDefaultPrompt() {
	prompt := fmt.Sprintf("@%s : ", cmd.twitter.OwnUser.ScreenName)
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
	argNum := len(args)

	if argNum <= 0 {
		return "", "", errors.New("no arguments")
	}

	str, count := args[0], cmd.config.Option.Counts

	// 2つ目の引数があればcountに代入
	if argNum >= 2 {
		count = args[1]
	}

	return str, count, nil
}

// parseAccountCmdArgs アカウント系のコマンド引数をパース
func (cmd *Cmd) parseAccountCmdArgs(args []string) (string, error) {
	// 対象のスクリーン名が指定されていない
	if len(args) == 0 {
		return "", errors.New("specify the screen name of the target account")
	}

	screenName := strings.Replace(args[0], "@", "", 1)

	// メインアカウントを示す "main" が許可されているなら通す
	if screenName == "main" {
		return "main", nil
	}

	// アカウントの存在チェック
	if _, ok := cmd.config.Cred.Sub[screenName]; !ok {
		return "", errors.New("account does not exist")
	}

	return screenName, nil
}

// getCountFromCmdArg 引数からツイート取得件数を取得
func (cmd *Cmd) getCountFromCmdArg(args []string) string {
	// 引数が無い、または数値以外ならデフォルト値を返す
	if len(args) <= 0 || !util.IsThreeDigitsNumber(args[0]) {
		return cmd.config.Option.Counts
	}

	return args[0]
}

// inputMultiLine マルチラインツイート入力
func (cmd *Cmd) inputMultiLine() string {
	// プロンプトを変更
	cmd.shell.SetPrompt("... ")
	defer cmd.setDefaultPrompt()

	fmt.Println("End typing with a semicolon. (If you want to cancel, input ':exit')")

	input := cmd.shell.ReadMultiLinesFunc(func(f string) bool {
		return f != ":exit" && !strings.HasSuffix(f, ";")
	})

	// 文字列内に:exitがあればキャンセル
	if strings.Contains(input, ":exit") {
		cmd.showMessage("CANCELED", "Input interrupted", cmd.config.Color.Accent2)
		return ""
	}

	return strings.TrimRight(input, ";")
}

// showExecutionConf 実行確認を表示
func (cmd *Cmd) showExecutionConf(msg string) bool {
	cmd.shell.SetPrompt("y/N: ")
	defer cmd.setDefaultPrompt()

	fmt.Println(msg)
	result := cmd.shell.ReadLine()

	return strings.ToLower(result) == "y"
}

// upload 画像をアップロード
func (cmd *Cmd) upload(images []string, query *url.Values) error {
	if len(images) <= 0 {
		return nil
	}

	// プログレスバー開始
	fmt.Print("Uploading... 🐾 ")
	cmd.shell.ProgressBar().Indeterminate(true)
	cmd.shell.ProgressBar().Start()

	mediaIDs, err := cmd.twitter.UploadImage(images)
	cmd.shell.ProgressBar().Stop()
	if err != nil {
		return err
	}

	query.Add("media_ids", mediaIDs)
	return nil
}

// actionOnTweet ツイートに対しての操作
func (cmd *Cmd) actionOnTweet(actionName, cmdName, bgColor string, args []string, actionFunc func(string) (string, error)) {
	if cmd.checkCommandLineMode() {
		return
	}

	if len(args) <= 0 {
		cmd.showWrongArgMessage(cmdName)
		return
	}

	// 引数の数だけ処理
	for _, v := range args {
		tweetId, err := cmd.twitter.GetDataFromTweetNum(v, twitter.TweetId)
		if err != nil {
			cmd.showErrorMessage(err.Error())
			return
		}

		tweetText, err := actionFunc(tweetId)
		if err != nil {
			cmd.showErrorMessage(err.Error())
			return
		}

		cmd.showMessage(actionName, tweetText, bgColor)
	}
}

// actionOnUser ユーザに対しての操作
func (cmd *Cmd) actionOnUser(actionName, cmdName, bgColor string, args []string, actionFunc func(string) (string, error)) {
	var err error

	if len(args) <= 0 {
		cmd.showWrongArgMessage(cmdName)
		return
	}

	screenName := args[0]

	// ツイート番号ならスクリーンネームに置換
	if util.IsThreeDigitsNumber(args[0]) {
		screenName, err = cmd.twitter.GetDataFromTweetNum(args[0], twitter.ScreenName)
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

// showTweets 登録されたツイートを一覧表示
func (cmd *Cmd) showTweets() {
	cmd.view.ShowTweets(cmd.twitter.Tweets, true)
}

// showMessage メッセージを表示
func (cmd *Cmd) showMessage(title, text, bgColor string) {
	width := util.GetWindowWidth()

	// 不要な文字を削除
	text = util.AllReplace(text, "[\t\n\r]", " ")
	text = html.UnescapeString(text)

	// 画面内に収まるよう丸める
	text = util.TruncateString(text, width-len(title)-3)

	tips := color.HEXStyle(cmd.config.Color.BoxForground, bgColor).Sprintf(" %s ", title)
	fmt.Printf("%s %s\n", tips, text)
}

// showErrorMessage エラーメッセージを表示
func (cmd *Cmd) showErrorMessage(msg string) {
	tips := color.HEXStyle(cmd.config.Color.BoxForground, cmd.config.Color.Error).Sprint(" ERROR ")
	fmt.Fprintf(os.Stderr, "%s %s\n", tips, msg)
}

// drawWrongArgError 引数ミスのメッセージを表示
func (cmd *Cmd) showWrongArgMessage(cmdName string) {
	msg := fmt.Sprintf("Wrong argument, try '%s help'", cmdName)
	cmd.showErrorMessage(msg)
}

// checkCommandLineMode コマンドラインモードかどうかチェック
func (cmd *Cmd) checkCommandLineMode() bool {
	if cmd.isCommandLineMode {
		cmd.showErrorMessage("Only available in interactive mode")
	}
	return cmd.isCommandLineMode
}

// createLongHelp 詳細なヘルプ文を作成
func createLongHelp(help, alias, use, exp string) string {
	longHelp := help

	if alias != "" {
		longHelp += fmt.Sprintf("\n\nAlias:\n  %s", alias)
	}

	if use != "" {
		longHelp += fmt.Sprintf("\n\nUse:\n  %s", use)
	}

	if exp != "" {
		longHelp += fmt.Sprintf("\n\nExample:\n  %s", exp)
	}

	return longHelp
}
