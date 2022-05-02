package cmd

import (
	"github.com/arrow2nd/ishell"
)

func (cmd *Cmd) newFollowCmd() *ishell.Cmd {
	// follow
	followCmd := &ishell.Cmd{
		Name:    "follow",
		Aliases: []string{"fw"},
		Func: func(c *ishell.Context) {
			cmd.actionOnUser("FOLLOWED", c.Cmd.Name, cmd.cfg.Color.Following, c.Args, cmd.api.Follow)
		},
		Help: "follow a user",
		LongHelp: createLongHelp(
			`Follow a user.
If you specify a tweet-number, the person posting the tweet will be selected.`,
			"fw",
			"follow [<username / tweet-number>]",
			"follow arrow_2nd\n  follow 2",
		),
	}

	// follow remove
	followCmd.AddCmd(&ishell.Cmd{
		Name:    "remove",
		Aliases: []string{"rm"},
		Func: func(c *ishell.Context) {
			cmd.actionOnUser("UNFOLLOWED", "follow "+c.Cmd.Name, cmd.cfg.Color.Following, c.Args, cmd.api.Unfollow)
		},
		Help: "unfollow a user",
		LongHelp: createLongHelp(
			`Unfollow a user.
If you specify a tweet-number, the person posting the tweet will be selected.`,
			"rm",
			"follow remove [<username / tweet-number>]",
			"follow remove arrow_2nd\n  follow rm 2",
		),
	})

	return followCmd
}
