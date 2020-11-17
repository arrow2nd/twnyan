package cmd

import (
	"github.com/arrow2nd/twnyan/twitter"
	"github.com/gookit/color"
	"gopkg.in/abiosoft/ishell.v2"
)

func init() {
	rtCmd := &ishell.Cmd{
		Name:    "retweet",
		Aliases: []string{"rt"},
		Help:    "retweet tweet",
		LongHelp: createLongHelp(
			"Retweet tweet.\nIf there is more than one, please separate them with a space.",
			"rt",
			"retweet [<tweet number>]...",
			"retweet 0 1",
		),
		Func: func(c *ishell.Context) {
			reactToTweet(c.Args, c.Cmd.Name, twitter.Retweet)
		},
	}

	rtCmd.AddCmd(&ishell.Cmd{
		Name:    "quote",
		Aliases: []string{"qt"},
		Help:    "quote tweet",
		LongHelp: createLongHelp(
			"Quote tweet.\nIf there is no tweet text, \"にゃーん\" will be posted.\nIf you are submitting an image, please add the file name separated by a space.",
			"qt",
			"retweet quote [<tweet number>] [text] [image]...",
			"retweet quote 0 cute!! cat.png",
		),
		Func: quote,
	})

	rtCmd.AddCmd(&ishell.Cmd{
		Name:    "remove",
		Aliases: []string{"rm"},
		Help:    "un-retweet tweet",
		LongHelp: createLongHelp(
			"UnRetweet tweet.\nIf there is more than one, please separate them with a space.",
			"rm",
			"retweet remove [<tweet number>]...",
			"retweet remove 0 1",
		),
		Func: func(c *ishell.Context) {
			reactToTweet(c.Args, "retweet "+c.Cmd.Name, twitter.UnRetweet)
		},
	})

	shell.AddCmd(rtCmd)
}

func quote(c *ishell.Context) {
	if len(c.Args) < 1 {
		showWrongMsg(c.Cmd.Name)
		return
	}

	// URL取得
	uri, err := tweets.GetTweetURL(c.Args[0])
	if err != nil {
		color.Error.Tips(err.Error())
		return
	}

	status, media := parsingTweetCommandArgument(c.Args[1:])
	status += " " + uri
	twitter.PostTweet(status, "", media)
}
