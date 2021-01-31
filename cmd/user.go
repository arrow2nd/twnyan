package cmd

import (
	"github.com/arrow2nd/twnyan/api"
	"github.com/arrow2nd/twnyan/util"
	"github.com/gookit/color"
	"gopkg.in/abiosoft/ishell.v2"
)

func (cmd *Cmd) newUserCmd() {
	uc := &ishell.Cmd{
		Name:    "user",
		Aliases: []string{"ur"},
		Func: func(c *ishell.Context) {
			value, counts, err := cmd.parseTLCmdArgs(c.Args)
			if err != nil {
				showWrongMsg(c.Cmd.Name)
				return
			}
			if util.IsNumber(value) {
				value, err = cmd.view.GetDataFromTweetNum(value, "screenname")
				if err != nil {
					color.Error.Prompt(err.Error())
					return
				}
			}
			cmd.loadUserTL(value, counts)
		},
		Help: "get a user timeline",
		LongHelp: createLongHelp(
			"Get a user timeline.\nIf you omit the counts, the default value in the configuration file (25 by default) will be specified.",
			"ur",
			"user [<username/tweetnumber>] [counts]",
			"user github 25\n  user 2",
		),
	}

	uc.AddCmd(&ishell.Cmd{
		Name: "own",
		Help: "get your own timeline",
		LongHelp: createLongHelp(
			"Get your own timeline.\nIf you omit the counts, the default value in the configuration file (25 by default) will be specified.",
			"",
			"user own [counts]",
			"user own 25",
		),
		Func: func(c *ishell.Context) {
			counts := cmd.getCountsFromCmdArg(c.Args)
			cmd.loadUserTL("", counts)
		},
	})

	cmd.shell.AddCmd(uc)
}

func (cmd *Cmd) loadUserTL(screenName, counts string) {
	v := api.CreateURLValues(counts)
	v.Add("screen_name", screenName)
	t, err := cmd.api.GetTimeline("user", v)
	if err != nil {
		return
	}
	cmd.view.RegisterTweets(t)
	cmd.view.DrawTweets()
}
