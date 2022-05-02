package cmd

import (
	"github.com/arrow2nd/ishell"
	"github.com/gookit/color"
	"github.com/pkg/browser"
)

func (cmd *Cmd) newOpenCmd() *ishell.Cmd {
	return &ishell.Cmd{
		Name:    "open",
		Aliases: []string{"op"},
		Func:    cmd.execOpenCmd,
		Help:    "view the tweet in your browser",
		LongHelp: createLongHelp(
			"View the tweet in your browser.",
			"op",
			"open [<tweet-number>]",
			"open 2",
		),
	}
}

func (cmd *Cmd) execOpenCmd(c *ishell.Context) {
	if len(c.Args) != 1 {
		cmd.showWrongArgMessage(c.Cmd.Name)
		return
	}

	// 該当ツイートのURLを取得
	url, err := cmd.view.GetTweetURL(c.Args[0])
	if err != nil {
		color.Error.Prompt(err.Error())
		return
	}

	cmd.showMessage("OPENED", url, cmd.cfg.Color.Accent2)
	browser.OpenURL(url)
}
