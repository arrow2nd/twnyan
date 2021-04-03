package cmd

import (
	"github.com/arrow2nd/ishell"
)

const ver = "1.2.6"

func (c *Cmd) newVersionCmd() {
	c.shell.AddCmd(&ishell.Cmd{
		Name:    "version",
		Aliases: []string{"ver"},
		Func: func(c *ishell.Context) {
			c.Printf("twnyan🐾 ver.%s\n", ver)
		},
		Help:     "Displays the version",
		LongHelp: "",
	})
}
