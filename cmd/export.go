package cmd

import (
	"fmt"

	"github.com/arrow2nd/twnyan/util"
	"github.com/gookit/color"
	"gopkg.in/abiosoft/ishell.v2"
)

func init() {
	shell.AddCmd(&ishell.Cmd{
		Name:    "export",
		Aliases: []string{"ep"},
		Help:    "export tweets to a file",
		LongHelp: createLongHelp(
			"Export the tweets you are currently viewing to a file.",
			"ep",
			"export [<format>] [<filename>]",
			"export json tweets",
		),
		Func: func(c *ishell.Context) {
			// 引数取得
			if len(c.Args) != 2 {
				showWrongMsg(c.Cmd.Name)
				return
			}

			// フォーマット・ファイル名
			format, filename := c.Args[0], c.Args[1]
			if c.Args[0] == "" || c.Args[1] == "" {
				showWrongMsg(c.Cmd.Name)
				return
			}

			// 出力
			err := tweets.OutPut(format, filename)
			if err != nil {
				color.Error.Prompt(err.Error())
				return
			}

			util.ShowSuccessMsg("Exported", fmt.Sprintf("Export completed! (%s.%s)", filename, format), cfg.Color.BoxFg, cfg.Color.Accent3)
		},
		Completer: func([]string) []string {
			return []string{"json", "yaml"}
		},
	})
}
