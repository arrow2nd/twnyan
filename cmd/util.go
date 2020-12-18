package cmd

import (
	"errors"
	"fmt"

	"github.com/arrow2nd/twnyan/util"
	"github.com/gookit/color"
)

// parseTweetCmdArgs ツイート系コマンドの引数をパースする
func parseTweetCmdArgs(args []string) (string, []string) {
	status, media := "にゃーん", []string{}

	if len(args) > 0 {
		if util.ChkRegexp("\\.\\w{3,4}$", args[0]) {
			status = ""
			media = args[0:]
		} else {
			status = args[0]
			media = args[1:]
		}
	}

	return status, media
}

// parseTLCmdArgs タイムライン系コマンドの引数をパースする
func parseTLCmdArgs(args []string) (string, string, error) {
	// 引数エラー
	if len(args) <= 0 {
		return "", "", errors.New("No arguments")
	}

	// 文字列、デフォルトの取得件数
	str, counts := args[0], cfg.Default.Counts

	// 取得件数
	if len(args) >= 2 {
		counts = args[1]
	}

	return str, counts, nil
}

// getCountsFromCmdArg 引数から取得件数を取得
func getCountsFromCmdArg(args []string) string {
	// 引数無し・数値以外ならデフォルト値を返す
	if len(args) <= 0 || !util.IsNumber(args[0]) {
		return cfg.Default.Counts
	}

	return args[0]
}

// reactToTweet ツイートにリアクションする
func reactToTweet(args []string, cmdName string, function func(string)) {
	// 引数エラー
	if len(args) <= 0 {
		showWrongMsg(cmdName)
		return
	}

	for _, v := range args {
		id, err := tweets.GetDataFromTweetNum(v, "TweetID")
		if err != nil {
			color.Error.Prompt(err.Error())
			return
		}
		function(id)
	}
}

// reactToUser ユーザーにリアクションする
func reactToUser(args []string, cmdName string, function func(string)) {
	// 引数エラー
	if len(args) <= 0 {
		showWrongMsg(cmdName)
		return
	}

	screenName := args[0]

	// ツイート番号ならスクリーンネームに置換
	if util.IsNumber(args[0]) {
		var err error
		screenName, err = tweets.GetDataFromTweetNum(args[0], "ScreenName")
		if err != nil {
			color.Error.Prompt(err.Error())
			return
		}
	}

	function(screenName)
}

// createCompleter 入力補完用リストを作成
func createCompleter(list []string) []string {
	if list == nil {
		return nil
	}
	r := make([]string, len(list))
	for i, v := range list {
		if util.ChkRegexp("\\s", v) {
			r[i] = fmt.Sprintf("\"%s\"", v)
		} else {
			r[i] = v
		}
	}
	return r
}

// CreateLongHelp 詳しいヘルプ文を作成
func createLongHelp(help, alias, use, exp string) string {
	return fmt.Sprintf("%s\n\nAlias:\n  %s\n\nUse:\n  %s\n\nExample:\n  %s", help, alias, use, exp)
}

// showWrongMsg 引数ミスのメッセージを表示
func showWrongMsg(cmdName string) {
	color.Error.Prompt("Wrong argument, try '%s help'", cmdName)
}
