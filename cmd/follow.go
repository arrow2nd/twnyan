package cmd

import (
	"github.com/arrow2nd/twnyan/twitter"
	"gopkg.in/abiosoft/ishell.v2"
)

func init() {
	favCmd := &ishell.Cmd{
		Name:    "follow",
		Aliases: []string{"fw"},
		Help:    "follow user",
		LongHelp: createLongHelp(
			"Follow user.\nIf you specify a tweet number, the person posting the tweet will be selected.",
			"fw",
			"follow [<userID / tweet number>]",
			"follow arrow_2nd",
		),
		Func: func(c *ishell.Context) {
			reactToUser(c.Args, c.Cmd.Name, twitter.Follow)
		},
	}

	favCmd.AddCmd(&ishell.Cmd{
		Name:    "remove",
		Aliases: []string{"rm"},
		Help:    "unfollow user",
		LongHelp: createLongHelp(
			"Unfollow user.\nIf you specify a tweet number, the person posting the tweet will be selected.",
			"rm",
			"follow remove [<userID / tweet number>]",
			"follow remove arrow_2nd",
		),
		Func: func(c *ishell.Context) {
			reactToUser(c.Args, c.Cmd.Name, twitter.Unfollow)
		},
	})

	shell.AddCmd(favCmd)
}