package cmd

import (
	"github.com/arrow2nd/ishell"
)

func (c *Cmd) newVersionCmd() *ishell.Cmd {
	return &ishell.Cmd{
		Name:    "version",
		Aliases: []string{"ver"},
		Func: func(c *ishell.Context) {
			c.Printf("twnyan 🐾 ver.%s\n", versionStr)
		},
		Help: "display version",
	}
}
