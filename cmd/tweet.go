package cmd

import (
	"github.com/arrow2nd/twnyan/twitter"
	"github.com/gookit/color"
	"gopkg.in/abiosoft/ishell.v2"
)

func init() {
	tweetCmd := &ishell.Cmd{
		Name:    "tweet",
		Aliases: []string{"tw"},
		Help:    "post a tweet",
		LongHelp: createLongHelp(
			"Post a tweet.\nIf there is no tweet text, \"にゃーん\" will be posted.\nIf you are submitting an image, please add the file name separated by a space.",
			"tw",
			"tweet [text] [image]...",
			"tweet nyaan! cat.png dog.png",
		),
		Func: func(c *ishell.Context) {
			status, media := parsingTweetCommandArgument(c.Args)
			twitter.PostTweet(status, "", media)
		},
	}

	tweetCmd.AddCmd(&ishell.Cmd{
		Name:    "remove",
		Aliases: []string{"rm"},
		Help:    "delete the tweet",
		LongHelp: createLongHelp(
			"Delete the specified tweet.\nIf there is more than one, please separate them with a space.",
			"rm",
			"tweet remove [<tweet number>]",
			"tweet remove 0 1",
		),
		Func: deleteTweet,
	})

	shell.AddCmd(tweetCmd)
}

func deleteTweet(c *ishell.Context) {
	if len(c.Args) <= 0 {
		showWrongMsg("tweet " + c.Cmd.Name)
		return
	}

	for _, v := range c.Args {
		id, err := tweets.GetDataFromTweetNum(v, "TweetID")
		if err != nil {
			color.Error.Tips(err.Error())
			return
		}
		twitter.DeleteTweet(id)
	}
}
