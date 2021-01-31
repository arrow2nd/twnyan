package cmd

import (
	"github.com/gookit/color"
	"github.com/pkg/browser"
	"gopkg.in/abiosoft/ishell.v2"
)

func (cmd *Cmd) newOpenCmd() {
	cmd.shell.AddCmd(&ishell.Cmd{
		Name:    "open",
		Aliases: []string{"op"},
		Func: func(c *ishell.Context) {
			if len(c.Args) != 1 {
				showWrongMsg(c.Cmd.Name)
				return
			}
			uri, err := cmd.view.GetTweetURL(c.Args[0])
			if err != nil {
				color.Error.Prompt(err.Error())
				return
			}
			// util.ShowSuccessMsg("Open", uri, cfg.Color.BoxFg, cfg.Color.Accent3)
			browser.OpenURL(uri)
		},
		Help: "view the tweet in your browser",
		LongHelp: createLongHelp(
			"View the tweet in your browser.",
			"op",
			"open [<tweetnumber>]",
			"open 2",
		),
	})
}
