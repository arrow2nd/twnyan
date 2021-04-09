package cmd

import (
	"github.com/arrow2nd/ishell"
)

func (cmd *Cmd) addStreamCmd() {
	cmd.shell.AddCmd(&ishell.Cmd{
		Name:    "stream",
		Aliases: []string{"st"},
		Func:    cmd.streamCmd,
		Help:    "",
		LongHelp: createLongHelp(
			"",
			"st",
			"stream",
			"stream",
		),
	})
}

func (cmd *Cmd) streamCmd(c *ishell.Context) {
	// 初めにタイムラインを取得
	// v := api.CreateQuery("25")
	// tweets, err := cmd.api.FetchTimelineTweets("home", v)
	// if err != nil {
	// 	cmd.showErrorMessage(err.Error())
	// 	return
	// }
	// cmd.view.ShowTweets(*tweets)
	// tweetSinceId := (*tweets)[0].IdStr
}
