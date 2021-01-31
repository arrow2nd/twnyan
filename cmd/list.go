package cmd

import (
	"fmt"

	"github.com/arrow2nd/twnyan/util"
	"github.com/gookit/color"
	"gopkg.in/abiosoft/ishell.v2"
)

func (cmd *Cmd) newListCmd() {
	cmd.shell.AddCmd(&ishell.Cmd{
		Name:    "list",
		Aliases: []string{"ls"},
		Func: func(c *ishell.Context) {
			name, counts, err := cmd.parseTLCmdArgs(c.Args)
			if err != nil {
				showWrongMsg(c.Cmd.Name)
				return
			}
			i := util.IndexOf(cmd.api.ListNames, name)
			if i == -1 {
				color.Error.Prompt("No list exists!")
				return
			}
			t, err := cmd.api.GetListTimeline(cmd.api.ListIDs[i], counts)
			if err != nil {
				return
			}
			cmd.view.RegisterTweets(t)
			cmd.view.DrawTweets()
		},
		Help: "get the list timeline",
		LongHelp: createLongHelp(
			"Get the list timeline.\nYou can use the tab key to complete the list name.\nIf you omit the counts, the default value in the configuration file (25 by default) will be specified.",
			"ls",
			"list [<listname>] [counts]",
			"list cats 50",
		),
		Completer: func([]string) []string {
			if cmd.api.ListNames == nil {
				return nil
			}
			cmp := make([]string, len(cmd.api.ListNames))
			for i, v := range cmd.api.ListNames {
				if util.ChkRegexp("\\s", v) {
					cmp[i] = fmt.Sprintf("\"%s\"", v)
				} else {
					cmp[i] = v
				}
			}
			return cmp
		},
	})
}
