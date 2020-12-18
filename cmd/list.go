package cmd

import (
	"github.com/arrow2nd/twnyan/util"
	"github.com/gookit/color"
	"gopkg.in/abiosoft/ishell.v2"
)

func init() {
	shell.AddCmd(&ishell.Cmd{
		Name:    "list",
		Aliases: []string{"ls"},
		Help:    "get the list timeline",
		LongHelp: createLongHelp(
			"Get the list timeline.\nYou can use the tab key to complete the list name.\nIf you omit the counts, the default value in the configuration file (25 by default) will be specified.",
			"ls",
			"list [<listname>] [counts]",
			"list cats 50",
		),
		Func: func(c *ishell.Context) {
			// 引数をパース
			name, counts, err := parseTLCmdArgs(c.Args)
			if err != nil {
				showWrongMsg(c.Cmd.Name)
				return
			}

			// リストの存在チェック
			idx := util.IndexOf(listName, name)
			if idx == -1 {
				color.Error.Prompt("No list exists!")
				return
			}

			// リストTL読み込み
			err = tweets.LoadListTL(listID[idx], counts)
			if err != nil {
				return
			}

			// 表示
			tweets.DrawTweets()
		},
		Completer: func([]string) []string {
			return createCompleter(listName)
		},
	})
}
