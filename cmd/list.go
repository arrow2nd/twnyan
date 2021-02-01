package cmd

import (
	"fmt"

	"github.com/arrow2nd/twnyan/util"
	"gopkg.in/abiosoft/ishell.v2"
)

func (cmd *Cmd) newListCmd() {
	cmd.shell.AddCmd(&ishell.Cmd{
		Name:    "list",
		Aliases: []string{"ls"},
		Func: func(c *ishell.Context) {
			// 引数をパース
			name, counts, err := cmd.parseTLCmdArgs(c.Args)
			if err != nil {
				cmd.drawWrongArgMessage(c.Cmd.Name)
				return
			}
			// リスト名がリスト内にあるかチェック
			i := util.IndexOf(cmd.api.ListNames, name)
			if i == -1 {
				cmd.drawErrorMessage("No list exists!")
				return
			}
			// リストタイムラインを取得
			t, err := cmd.api.GetListTimeline(cmd.api.ListIDs[i], counts)
			if err != nil {
				cmd.drawErrorMessage(err.Error())
				return
			}
			// 描画
			cmd.view.RegisterTweets(t)
			cmd.view.DrawTweets()
		},
		Help: "get the list timeline",
		LongHelp: createLongHelp(
			"Get the list timeline.\nYou can use the tab key to complete the list name.\nIf you omit the counts, the default value in the configuration file (25 by default) will be specified.",
			"ls",
			"list [<listname>] [counts]",
			"list cats 50",
		),
		Completer: func([]string) []string {
			// リストの存在チェック
			if cmd.api.ListNames == nil {
				return nil
			}
			// 補完用スライス作成
			cmp := make([]string, len(cmd.api.ListNames))
			for i, v := range cmd.api.ListNames {
				if util.ContainsStr("\\s", v) {
					// リスト名に空白があればダブルクオートで囲む
					cmp[i] = fmt.Sprintf("\"%s\"", v)
				} else {
					cmp[i] = v
				}
			}
			return cmp
		},
	})
}
