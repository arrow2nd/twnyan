package cmd

import (
	"github.com/arrow2nd/twnyan/util"
	"gopkg.in/abiosoft/ishell.v2"
)

func init() {
	shell.AddCmd(&ishell.Cmd{
		Name:    "search",
		Aliases: []string{"sh"},
		Help:    "search for tweets from the past 7 days",
		LongHelp: createLongHelp(
			"Search for tweets from the past 7 days.\nIf you omit the counts, the default value in the configuration file (25 by default) will be specified.",
			"sh",
			"search [<keyword>] [counts]",
			"search catto 50",
		),
		Func: getSearch,
	})
}

func getSearch(c *ishell.Context) {
	args, err := util.FetchStringSpecifiedType(c.Args, "str", "num")
	if err != nil || args[0] == "" {
		showWrongMsg(c.Cmd.Name)
		return
	}

	keyword, counts := args[0], args[1]
	if counts == "" {
		counts = cfg.Default.Counts
	}

	err = tweets.LoadSearchResult(keyword, counts)
	if err == nil {
		tweets.DrawTweets()
	}
}
