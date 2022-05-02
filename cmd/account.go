package cmd

import (
	"fmt"

	"github.com/arrow2nd/ishell"
)

func (cmd *Cmd) newAccountCmd() *ishell.Cmd {
	accountCmd := &ishell.Cmd{
		Name:    "account",
		Aliases: []string{"acc"},
		Help:    "manage and switch accounts",
		LongHelp: createLongHelp(
			`Manage and switch accounts.`,
			"acc",
			"",
			"",
		),
	}

	accountCmd.AddCmd(&ishell.Cmd{
		Name: "add",
		Func: cmd.execAccountAddCmd,
		Help: "add a sub-account",
		LongHelp: createLongHelp(
			`Add a sub-account to twnyan.`,
			"",
			"account add",
			"",
		),
	})

	accountCmd.AddCmd(&ishell.Cmd{
		Name:      "remove",
		Aliases:   []string{"rm"},
		Completer: cmd.accountNameCompleter,
		Func:      cmd.execAccountRemoveCmd,
		Help:      "remove sub-accounts",
		LongHelp: createLongHelp(
			`Remove sub-accounts from twnyan.`,
			"rm",
			"account remove",
			"",
		),
	})

	accountCmd.AddCmd(&ishell.Cmd{
		Name:    "list",
		Aliases: []string{"ls"},
		Func: func(c *ishell.Context) {
			fmt.Println("run: list")
		},
		Help: "List sub-accounts",
		LongHelp: createLongHelp(
			`List sub-accounts added to twnyan.`,
			"ls",
			"account list",
			"",
		),
	})

	accountCmd.AddCmd(&ishell.Cmd{
		Name:    "switch",
		Aliases: []string{"sw"},
		Func: func(c *ishell.Context) {
			fmt.Println("run: switch")
		},
		Help: "switch the account to use",
		LongHelp: createLongHelp(
			`Switch the account to use.`,
			"sw",
			"account switch",
			"",
		),
	})

	return accountCmd
}

func (cmd *Cmd) accountNameCompleter([]string) []string {
	items := []string{}

	for name := range cmd.cfg.Cred.Sub {
		items = append(items, name)
	}

	return items
}

func (cmd *Cmd) execAccountAddCmd(c *ishell.Context) {
	c.ClearScreen()

	// Auth認証
	newCred, screenName, err := cmd.api.Auth()
	if err != nil {
		cmd.showErrorMessage(err.Error())
		return
	}

	// 追加して保存
	cmd.cfg.Cred.Sub[screenName] = newCred
	cmd.cfg.Save()

	cmd.showMessage("ADDED", screenName, cmd.cfg.Color.Accent3)
}

func (cmd *Cmd) execAccountRemoveCmd(c *ishell.Context) {
	if len(c.Args) == 0 {
		cmd.showWrongArgMessage("account " + c.Cmd.Name)
		return
	}

	// アカウントの存在を確認
	screenName := c.Args[0]
	if _, ok := cmd.cfg.Cred.Sub[screenName]; !ok {
		cmd.showErrorMessage("No account found")
		return
	}

	// 実行確認
	msg := fmt.Sprintf("Delete account (%s) from twnyan?", screenName)
	if ok := cmd.showExecutionConf(msg); !ok {
		cmd.showMessage("CANCELED", "Interrupted", cmd.cfg.Color.Accent2)
		return
	}

	delete(cmd.cfg.Cred.Sub, screenName)
	cmd.cfg.Save()

	cmd.showMessage("REMOVED", screenName, cmd.cfg.Color.Accent3)
}
