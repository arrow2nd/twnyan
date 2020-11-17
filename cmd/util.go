package cmd

import (
	"fmt"

	"github.com/arrow2nd/twnyan/util"
	"github.com/gookit/color"
)

// parsingTweetCommandArgument ツイートのコマンドをパースする
func parsingTweetCommandArgument(args []string) (string, []string) {
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

// reactToTweet ツイートにリアクションする
func reactToTweet(args []string, cmdName string, function func(string)) {
	if len(args) <= 0 {
		showWrongMsg(cmdName)
		return
	}

	for _, v := range args {
		id, err := tweets.GetDataFromTweetNum(v, "TweetID")
		if err != nil {
			color.Error.Tips(err.Error())
			return
		}
		function(id)
	}
}

// reactToUser ユーザーにリアクションする
func reactToUser(args []string, cmdName string, function func(string)) {
	if len(args) <= 0 {
		showWrongMsg(cmdName)
		return
	}

	screenName := args[0]

	// ツイート番号が指定された場合
	if util.IsNumber(args[0]) {
		var err error
		screenName, err = tweets.GetDataFromTweetNum(args[0], "ScreenName")
		if err != nil {
			color.Error.Tips(err.Error())
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
	color.Error.Tips("Wrong argument, try '%s help'", cmdName)
}
