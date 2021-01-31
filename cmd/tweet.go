package cmd

import (
	"github.com/gookit/color"
	"gopkg.in/abiosoft/ishell.v2"
)

func (cmd *Cmd) newTweetCmd() {
	tc := &ishell.Cmd{
		Name:    "tweet",
		Aliases: []string{"tw"},
		Func: func(c *ishell.Context) {
			status, media := cmd.parseTweetCmdArgs(c.Args)
			cmd.api.PostTweet(status, "", media)
		},
		Help: "post tweet",
		LongHelp: createLongHelp(
			"Post tweet.\nIf there is no tweet text, \"にゃーん\" will be posted.\nIf you are submitting an image, please add the file name separated by a space.",
			"tw",
			"tweet [text] [image]...",
			"tweet nyaan! cat.png dog.png",
		),
	}

	tc.AddCmd(&ishell.Cmd{
		Name:    "remove",
		Aliases: []string{"rm"},
		Func: func(c *ishell.Context) {
			if len(c.Args) <= 0 {
				showWrongMsg("tweet " + c.Cmd.Name)
				return
			}
			for _, v := range c.Args {
				id, err := cmd.view.GetDataFromTweetNum(v, "tweetID")
				if err != nil {
					color.Error.Prompt(err.Error())
					return
				}
				cmd.api.DeleteTweet(id)
			}
		},
		Help: "delete tweet",
		LongHelp: createLongHelp(
			"Delete tweet.\nIf there is more than one, please separate them with a space.",
			"rm",
			"tweet remove [<tweetnumber>]",
			"tweet remove 0 1",
		),
	})

	cmd.shell.AddCmd(tc)
}
