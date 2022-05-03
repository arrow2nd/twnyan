package cmd

import (
	"github.com/arrow2nd/ishell/v2"
)

const version = "1.7.2"

func (c *Cmd) newVersionCmd() *ishell.Cmd {
	return &ishell.Cmd{
		Name:    "version",
		Aliases: []string{"ver"},
		Func: func(c *ishell.Context) {
			c.Printf("🐈 twnyan ver.%s\n", version)
		},
		Help: "show version",
	}
}
