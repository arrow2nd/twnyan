package cmd

import (
	"os"

	"github.com/arrow2nd/twnyan/config"
	"gopkg.in/abiosoft/ishell.v2"
)

func init() {
	configCmd := &ishell.Cmd{
		Name: "config",
		Help: "manipulation of configuration files",
	}

	configCmd.AddCmd(&ishell.Cmd{
		Name:     "remove",
		Help:     "delete the configuration file",
		LongHelp: "Deletes the configuration file.\nQuit the program after deleting the file.",
		Func: func(c *ishell.Context) {
			cfg.Remove()
			os.Exit(0)
		},
	})

	configCmd.AddCmd(&ishell.Cmd{
		Name:     "reset",
		Help:     "regenerate the configuration file",
		LongHelp: "Regenerate the configuration file.\nRedo the application authentication.",
		Func: func(c *ishell.Context) {
			config.Setup()
		},
	})

	shell.AddCmd(configCmd)
}
