package cmd

import (
	"github.com/gookit/color"
	"gopkg.in/abiosoft/ishell.v2"
)

func (cmd *Cmd) newReplyCmd() {
	cmd.shell.AddCmd(&ishell.Cmd{
		Name:    "reply",
		Aliases: []string{"rp"},
		Func: func(c *ishell.Context) {
			if len(c.Args) < 1 {
				showWrongMsg(c.Cmd.Name)
				return
			}
			tweetID, err := cmd.view.GetDataFromTweetNum(c.Args[0], "tweetID")
			if err != nil {
				color.Error.Prompt(err.Error())
				return
			}
			status, media := cmd.parseTweetCmdArgs(c.Args[1:])
			cmd.api.PostTweet(status, tweetID, media)
		},
		Help: "post a reply",
		LongHelp: createLongHelp(
			"Post a reply.\nIf there is no tweet text, \"にゃーん\" will be posted.\nIf you are submitting an image, please add the file name separated by a space.",
			"rp",
			"reply [<tweetnumber>] [text] [image]...",
			"reply 2 meow cat.jpg",
		),
	})
}
