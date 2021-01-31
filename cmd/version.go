package cmd

import (
	"fmt"

	"gopkg.in/abiosoft/ishell.v2"
)

const ver = "2.0.0"

func (c *Cmd) newVersionCmd() {
	c.shell.AddCmd(&ishell.Cmd{
		Name:    "version",
		Aliases: []string{"ver"},
		Func: func(c *ishell.Context) {
			fmt.Printf("twnyan🐾 ver.%s\n", ver)
		},
		Help:     "Displays the version",
		LongHelp: "",
	})
}
