package cmd

import (
	"fmt"

	"github.com/arrow2nd/ishell/v2"
	"github.com/arrow2nd/twnyan/util"
)

func (cmd *Cmd) newListCmd() *ishell.Cmd {
	return &ishell.Cmd{
		Name:      "list",
		Aliases:   []string{"ls"},
		Completer: cmd.listCmdCompleter,
		Func:      cmd.execListCmd,
		Help:      "get the list timeline",
		LongHelp: createLongHelp(
			`Get the list timeline.
You can use the tab key to complete the list name.
If you omit the counts, the default value in the configuration file (25 by default) will be specified.`,
			"ls",
			"list [<listname>] [counts]",
			"list cats 50",
		),
	}
}

func (cmd *Cmd) execListCmd(c *ishell.Context) {
	name, counts, err := cmd.parseTimelineCmdArgs(c.Args)
	if err != nil {
		cmd.showWrongArgMessage(c.Cmd.Name)
		return
	}

	// リスト名からリストIDを取得
	listId, ok := cmd.twitter.List[name]
	if !ok {
		cmd.showErrorMessage("Not found in list")
		return
	}

	// リストのツイートを取得
	tweets, err := cmd.twitter.FetchListTweets(listId, counts)
	if err != nil {
		cmd.showErrorMessage(err.Error())
		return
	}

	cmd.twitter.RegisterTweets(tweets)
	cmd.showTweets()
}

func (cmd *Cmd) listCmdCompleter([]string) []string {
	// リストが無いなら処理しない
	if len(cmd.twitter.List) == 0 {
		return nil
	}

	// 入力補完用のスライスを作成
	items := []string{}

	for name := range cmd.twitter.List {
		// リスト名が空白を含んでいるならダブルクオートで囲む
		if util.MatchesRegexp("\\s", name) {
			name = fmt.Sprintf("\"%s\"", name)
		}

		items = append(items, name)
	}

	return items
}
