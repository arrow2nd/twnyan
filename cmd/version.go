package cmd

import (
	"github.com/arrow2nd/ishell"
)

const versionStr = "1.4.5"

func (c *Cmd) addVersionCmd() {
	c.shell.AddCmd(&ishell.Cmd{
		Name:    "version",
		Aliases: []string{"ver"},
		Func: func(c *ishell.Context) {
			c.Printf("twnyan🐾 ver.%s\n", versionStr)
		},
		Help: "display version",
	})
}
