package cmd

import (
	"github.com/arrow2nd/twnyan/util"
	"github.com/gookit/color"
	"github.com/pkg/browser"
	"gopkg.in/abiosoft/ishell.v2"
)

func init() {
	shell.AddCmd(&ishell.Cmd{
		Name:    "open",
		Aliases: []string{"op"},
		Help:    "view the tweet in your browser",
		LongHelp: createLongHelp(
			"View the tweet in your browser.",
			"op",
			"open [<tweetnumber>]",
			"open 2",
		),
		Func: func(c *ishell.Context) {
			// 引数エラー
			if len(c.Args) != 1 {
				showWrongMsg(c.Cmd.Name)
				return
			}

			// URL取得
			uri, err := tweets.GetTweetURL(c.Args[0])
			if err != nil {
				color.Error.Prompt(err.Error())
				return
			}

			// ブラウザで開く
			util.ShowSuccessMsg("Open", uri, cfg.Color.BoxFg, cfg.Color.Accent3)
			browser.OpenURL(uri)
		},
	})
}
