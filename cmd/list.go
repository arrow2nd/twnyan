package cmd

import (
	"fmt"

	"github.com/arrow2nd/ishell"
	"github.com/arrow2nd/twnyan/util"
)

func (cmd *Cmd) addListCmd() {
	cmd.shell.AddCmd(&ishell.Cmd{
		Name:    "list",
		Aliases: []string{"ls"},
		Func:    cmd.listCmd,
		Help:    "get the list timeline",
		LongHelp: createLongHelp(
			"Get the list timeline.\nYou can use the tab key to complete the list name.\nIf you omit the counts, the default value in the configuration file (25 by default) will be specified.",
			"ls",
			"list [<listname>] [counts]",
			"list cats 50",
		),
		Completer: cmd.listCmdCompleter,
	})
}

func (cmd *Cmd) listCmd(c *ishell.Context) {
	name, counts, err := cmd.parseTimelineCmdArgs(c.Args)
	if err != nil {
		cmd.showWrongArgMessage(c.Cmd.Name)
		return
	}

	// 指定されたリスト名が存在するかチェック
	listIndex := util.IndexOf(cmd.api.ListNames, name)
	if listIndex == -1 {
		cmd.showErrorMessage("No list exists!")
		return
	}

	// リストのツイートを取得
	tweets, err := cmd.api.FetchListTweets(cmd.api.ListIDs[listIndex], counts)
	if err != nil {
		cmd.showErrorMessage(err.Error())
		return
	}

	// 登録して表示
	cmd.view.RegisterTweets(tweets)
	cmd.view.ShowRegisteredTweets()
}

func (cmd *Cmd) listCmdCompleter([]string) []string {
	// リストが無いならreturn
	if cmd.api.ListNames == nil {
		return nil
	}

	// 入力補完用のスライスを作成
	cmp := make([]string, len(cmd.api.ListNames))
	for i, name := range cmd.api.ListNames {
		// リスト名が空白を含んでいるならダブルクオートで囲む
		if util.MatchesRegexp("\\s", name) {
			cmp[i] = fmt.Sprintf("\"%s\"", name)
		} else {
			cmp[i] = name
		}
	}

	return cmp
}
