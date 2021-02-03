package cmd

import (
	"gopkg.in/abiosoft/ishell.v2"
)

const ver = "2.0.1"

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
