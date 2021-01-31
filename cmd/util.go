package cmd

import (
	"errors"
	"fmt"

	"github.com/arrow2nd/twnyan/util"
	"github.com/gookit/color"
)

// parseTweetCmdArgs ツイート系コマンドの引数をパースする
func (cmd *Cmd) parseTweetCmdArgs(args []string) (string, []string) {
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
func (cmd *Cmd) parseTLCmdArgs(args []string) (string, string, error) {
	// 引数をチェック
	if len(args) <= 0 {
		return "", "", errors.New("No arguments")
	}
	str, counts := args[0], cmd.cfg.Option.Counts
	// 取得件数の指定があれば置換
	if len(args) >= 2 {
		counts = args[1]
	}
	return str, counts, nil
}

// getCountsFromCmdArg 引数から取得件数を取得
func (cmd *Cmd) getCountsFromCmdArg(args []string) string {
	// 引数無し、数値以外ならデフォルト値を返す
	if len(args) <= 0 || !util.IsNumber(args[0]) {
		return cmd.cfg.Option.Counts
	}
	return args[0]
}

// reactToTweet ツイートにリアクションする
func (cmd *Cmd) reactToTweet(args []string, cmdName string, function func(string) error) {
	// 引数をチェック
	if len(args) <= 0 {
		showWrongMsg(cmdName)
		return
	}
	// 引数の数だけ処理
	for _, v := range args {
		id, err := cmd.view.GetDataFromTweetNum(v, "tweetID")
		if err != nil {
			color.Error.Prompt(err.Error())
			return
		}
		err = function(id)
		if err != nil {
			color.Error.Prompt(err.Error())
			return
		}
	}
}

// reactToUser ユーザーにリアクションする
func (cmd *Cmd) reactToUser(args []string, cmdName string, function func(string) error) {
	var err error
	// 引数をチェック
	if len(args) <= 0 {
		showWrongMsg(cmdName)
		return
	}

	screenName := args[0]

	// ツイート番号ならスクリーンネームに置換
	if util.IsNumber(args[0]) {
		screenName, err = cmd.view.GetDataFromTweetNum(args[0], "screenname")
		if err != nil {
			color.Error.Prompt(err.Error())
			return
		}
	}

	err = function(screenName)
	if err != nil {
		color.Error.Prompt(err.Error())
	}
}

// CreateLongHelp 詳しいヘルプ文を作成
func createLongHelp(help, alias, use, exp string) string {
	return fmt.Sprintf("%s\n\nAlias:\n  %s\n\nUse:\n  %s\n\nExample:\n  %s", help, alias, use, exp)
}

// showWrongMsg 引数ミスのメッセージを表示
func showWrongMsg(cmdName string) {
	color.Error.Prompt("Wrong argument, try '%s help'", cmdName)
}
