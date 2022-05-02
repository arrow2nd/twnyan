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
		Func:    cmd.execAccountListCmd,
		Help:    "List sub-accounts",
		LongHelp: createLongHelp(
			`List sub-accounts added to twnyan.`,
			"ls",
			"account list",
			"",
		),
	})

	accountCmd.AddCmd(&ishell.Cmd{
		Name:      "switch",
		Aliases:   []string{"sw"},
		Completer: cmd.switchAccountNameCompleter,
		Func:      cmd.execAccountSwitchCmd,
		Help:      "switch the account to use",
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
	if len(cmd.cfg.Cred.Sub) == 0 {
		return nil
	}

	items := []string{}

	for name := range cmd.cfg.Cred.Sub {
		items = append(items, name)
	}

	return items
}

func (cmd *Cmd) switchAccountNameCompleter([]string) []string {
	if items := cmd.accountNameCompleter(nil); items != nil {
		return append(items, "main")
	}

	return nil
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
	screenName, err := cmd.parseAccountCmdArgs(c.Args, false)
	if err != nil {
		cmd.showErrorMessage(err.Error())
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

func (cmd *Cmd) execAccountListCmd(c *ishell.Context) {
	if len(cmd.cfg.Cred.Sub) == 0 {
		cmd.showErrorMessage("No sub-accounts")
		return
	}

	for _, name := range cmd.accountNameCompleter(nil) {
		fmt.Printf("- %s\n", name)
	}
}

func (cmd *Cmd) execAccountSwitchCmd(c *ishell.Context) {
	screenName, err := cmd.parseAccountCmdArgs(c.Args, true)
	if err != nil {
		cmd.showErrorMessage(err.Error())
		return
	}

	prevScreenName := cmd.api.OwnUser.ScreenName

	// アカウントを切り替え
	switch screenName {
	case "main":
		cmd.api.Init(cmd.cfg.Cred.Main)
	default:
		cmd.api.Init(cmd.cfg.Cred.Sub[screenName])
	}

	// デフォルトプロンプトを更新
	cmd.setDefaultPrompt()

	msg := fmt.Sprintf("%s -> %s\n", prevScreenName, cmd.api.OwnUser.ScreenName)
	cmd.showMessage("CHANGED", msg, cmd.cfg.Color.Accent3)
}
