package cmd

import (
	"fmt"

	"github.com/arrow2nd/ishell"
	"github.com/arrow2nd/twnyan/util"
)

func (cmd *Cmd) newListCmd() *ishell.Cmd {
	return &ishell.Cmd{
		Name:    "list",
		Aliases: []string{"ls"},
		Func:    cmd.execListCmd,
		Help:    "get the list timeline",
		LongHelp: createLongHelp(
			"Get the list timeline.\nYou can use the tab key to complete the list name.\nIf you omit the counts, the default value in the configuration file (25 by default) will be specified.",
			"ls",
			"list [<listname>] [counts]",
			"list cats 50",
		),
		Completer: cmd.listCmdCompleter,
	}
}

func (cmd *Cmd) execListCmd(c *ishell.Context) {
	name, counts, err := cmd.parseTimelineCmdArgs(c.Args)
	if err != nil {
		cmd.showWrongArgMessage(c.Cmd.Name)
		return
	}

	// リスト名からリストIDを取得
	listId, ok := cmd.api.List[name]
	if !ok {
		cmd.showErrorMessage("No list exists!")
		return
	}

	// リストのツイートを取得
	tweets, err := cmd.api.FetchListTweets(listId, counts)
	if err != nil {
		cmd.showErrorMessage(err.Error())
		return
	}

	// 登録して表示
	cmd.view.RegisterTweets(tweets)
	cmd.view.ShowRegisteredTweets()
}

func (cmd *Cmd) listCmdCompleter([]string) []string {
	// リストが無いなら処理しない
	if len(cmd.api.List) == 0 {
		return nil
	}

	// 入力補完用のスライスを作成
	items := []string{}

	for name := range cmd.api.List {
		// リスト名が空白を含んでいるならダブルクオートで囲む
		if util.MatchesRegexp("\\s", name) {
			items = append(items, fmt.Sprintf("\"%s\"", name))
		} else {
			items = append(items, name)
		}
	}

	return items
}
