package cmd

import (
	"github.com/arrow2nd/twnyan/util"
	"github.com/gookit/color"
	"gopkg.in/abiosoft/ishell.v2"
)

func init() {
	shell.AddCmd(&ishell.Cmd{
		Name:    "list",
		Aliases: []string{"ls"},
		Help:    "get the list timeline",
		LongHelp: createLongHelp(
			"Get the list timeline.\nYou can use the tab key to complete the list name.\nIf you omit the counts, the default value in the configuration file (25 by default) will be specified.",
			"ls",
			"list [<list name>] [counts] [data format(json|yaml)]",
			"list cats 50",
		),
		Func: getListTimeline,
		Completer: func([]string) []string {
			return createCompleter(listName)
		},
	})
}

func getListTimeline(c *ishell.Context) {
	args, err := util.FetchStringSpecifiedType(c.Args, "str", "num", "str")
	if err != nil {
		showWrongMsg(c.Cmd.Name)
		return
	}

	name, counts, dataFmt := args[0], args[1], args[2]
	if counts == "" {
		counts = cfg.Default.Counts
	}

	// リストが存在するかチェック
	idx := util.IndexOf(listName, name)
	if idx == -1 {
		color.Error.Tips("No list exists!")
		return
	}

	err = tweets.LoadListTL(listID[idx], counts)
	if err != nil {
		return
	}

	if dataFmt == "" {
		tweets.DrawTweets()
	} else {
		tweets.OutData(dataFmt)
	}
}
