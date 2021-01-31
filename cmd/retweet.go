package cmd

import (
	"github.com/gookit/color"
	"gopkg.in/abiosoft/ishell.v2"
)

func (cmd *Cmd) newRetweetCmd() {
	rc := &ishell.Cmd{
		Name:    "retweet",
		Aliases: []string{"rt"},
		Func: func(c *ishell.Context) {
			cmd.reactToTweet(c.Args, c.Cmd.Name, cmd.api.Retweet)
		},
		Help: "retweet tweet",
		LongHelp: createLongHelp("Retweet tweet.\nIf there is more than one, please separate them with a space.",
			"rt",
			"retweet [<tweetnumber>]...",
			"retweet 0 1",
		),
	}

	rc.AddCmd(&ishell.Cmd{
		Name:    "quote",
		Aliases: []string{"qt"},
		Func: func(c *ishell.Context) {
			if len(c.Args) < 1 {
				showWrongMsg(c.Cmd.Name)
				return
			}
			uri, err := cmd.view.GetTweetURL(c.Args[0])
			if err != nil {
				color.Error.Prompt(err.Error())
				return
			}
			status, media := cmd.parseTweetCmdArgs(c.Args[1:])
			status += " " + uri
			cmd.api.PostTweet(status, "", media)
		},
		Help: "quote tweet",
		LongHelp: createLongHelp(
			"Quote tweet.\nIf there is no tweet text, \"にゃーん\" will be posted.\nIf you are submitting an image, please add the file name separated by a space.",
			"qt",
			"retweet quote [<tweetnumber>] [text] [image]...",
			"retweet quote 0 cute!! cat.png",
		),
	})

	rc.AddCmd(&ishell.Cmd{
		Name:    "remove",
		Aliases: []string{"rm"},
		Func: func(c *ishell.Context) {
			cmd.reactToTweet(c.Args, "retweet "+c.Cmd.Name, cmd.api.UnRetweet)
		},
		Help: "un-retweet tweet",
		LongHelp: createLongHelp(
			"UnRetweet tweet.\nIf there is more than one, please separate them with a space.",
			"rm",
			"retweet remove [<tweetnumber>]...",
			"retweet remove 0 1",
		),
	})

	cmd.shell.AddCmd(rc)
}
