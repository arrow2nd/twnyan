package cmd

import (
	"github.com/arrow2nd/ishell"
)

const ver = "1.3.0"

func (c *Cmd) addVersionCmd() {
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
