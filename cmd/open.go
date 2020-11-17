package cmd

import (
	"fmt"

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
			"open [tweet number>]",
			"open 2",
		),
		Func: func(c *ishell.Context) {
			if len(c.Args) != 1 {
				showWrongMsg(c.Cmd.Name)
				return
			}
			uri, err := tweets.GetTweetURL(c.Args[0])
			if err != nil {
				color.Error.Tips(err.Error())
				return
			}
			green := color.LightGreen.Render
			fmt.Printf("%s %s\n", green("Open:"), uri)
			browser.OpenURL(uri)
		},
	})
}
