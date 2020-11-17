package cmd

import (
	"fmt"

	"gopkg.in/abiosoft/ishell.v2"
)

func init() {
	shell.AddCmd(&ishell.Cmd{
		Name:    "version",
		Aliases: []string{"ver"},
		Help:    "Displays the version",
		Func: func(c *ishell.Context) {
			fmt.Printf("twnyan ver.%s\n", version)
		},
	})
}
