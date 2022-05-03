package cmd

import (
	"fmt"

	"github.com/arrow2nd/ishell/v2"
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
	if len(cmd.config.Cred.Sub) == 0 {
		return nil
	}

	items := []string{}

	for name := range cmd.config.Cred.Sub {
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
	newCred, screenName, err := cmd.twitter.Auth()
	if err != nil {
		cmd.showErrorMessage(err.Error())
		return
	}

	// 追加して保存
	cmd.config.Cred.Sub[screenName] = newCred
	cmd.config.Save()

	cmd.showMessage("ADDED", screenName, cmd.config.Color.Accent3)
}

func (cmd *Cmd) execAccountRemoveCmd(c *ishell.Context) {
	screenName, err := cmd.parseAccountCmdArgs(c.Args, false)
	if err != nil {
		cmd.showErrorMessage(err.Error())
	}

	// 実行確認
	msg := fmt.Sprintf("Delete account (%s) from twnyan?", screenName)
	if ok := cmd.showExecutionConf(msg); !ok {
		cmd.showMessage("CANCELED", "Interrupted", cmd.config.Color.Accent2)
		return
	}

	delete(cmd.config.Cred.Sub, screenName)
	cmd.config.Save()

	cmd.showMessage("REMOVED", screenName, cmd.config.Color.Accent3)
}

func (cmd *Cmd) execAccountListCmd(c *ishell.Context) {
	if len(cmd.config.Cred.Sub) == 0 {
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

	prevScreenName := cmd.twitter.OwnUser.ScreenName

	// アカウントを切り替え
	switch screenName {
	case "main":
		cmd.twitter.Init(cmd.config.Cred.Main)
	default:
		cmd.twitter.Init(cmd.config.Cred.Sub[screenName])
	}

	// デフォルトプロンプトを更新
	cmd.setDefaultPrompt()

	msg := fmt.Sprintf("%s -> %s\n", prevScreenName, cmd.twitter.OwnUser.ScreenName)
	cmd.showMessage("SWITCHED", msg, cmd.config.Color.Accent3)
}
