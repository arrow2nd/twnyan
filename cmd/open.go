package cmd

import (
	"github.com/arrow2nd/ishell/v2"
	"github.com/gookit/color"
	"github.com/pkg/browser"
)

func (cmd *Cmd) newOpenCmd() *ishell.Cmd {
	return &ishell.Cmd{
		Name:    "open",
		Aliases: []string{"op"},
		Func:    cmd.execOpenCmd,
		Help:    "open a tweet in browser",
		LongHelp: createLongHelp(
			`Open a tweet in your browser.
      (Only available in interactive mode)`,
			"op",
			"open <tweet-number>",
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
	url, err := cmd.twitter.GetTweetURL(c.Args[0])
	if err != nil {
		color.Error.Prompt(err.Error())
		return
	}

	cmd.showMessage("OPENED", url, cmd.config.Color.Accent3)
	browser.OpenURL(url)
}
