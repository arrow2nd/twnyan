package cmd

import (
	"fmt"

	"github.com/arrow2nd/ishell/v2"
	"github.com/garyburd/go-oauth/oauth"
)

func (cmd *Cmd) newAccountCmd() *ishell.Cmd {
	accountCmd := &ishell.Cmd{
		Name:    "account",
		Aliases: []string{"acc"},
		Help:    "manage and switch accounts",
		LongHelp: createLongHelp(
			"Manage and switch accounts used by twnyan.",
			"acc",
			"twnyan account <command>",
			"",
		),
	}

	accountCmd.AddCmd(&ishell.Cmd{
		Name: "add",
		Func: cmd.execAccountAddCmd,
		Help: "add account",
		LongHelp: createLongHelp(
			"Add account to twnyan.",
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
		Help:      "remove account",
		LongHelp: createLongHelp(
			"Remove account from twnyan.",
			"rm",
			"account remove <username>",
			"account remove nekochan",
		),
	})

	accountCmd.AddCmd(&ishell.Cmd{
		Name:    "list",
		Aliases: []string{"ls"},
		Func:    cmd.execAccountListCmd,
		Help:    "List accounts",
		LongHelp: createLongHelp(
			"Lists accounts added to twnyan.",
			"ls",
			"account list",
			"",
		),
	})

	accountCmd.AddCmd(&ishell.Cmd{
		Name:      "switch",
		Aliases:   []string{"sw"},
		Completer: cmd.accountNameCompleter,
		Func:      cmd.execAccountSwitchCmd,
		Help:      "switch the account to use",
		LongHelp: createLongHelp(
			`Switch the account to be used.
(Only available in interactive mode)`,
			"sw",
			"account switch <username>",
			"account switch nekochan",
		),
	})

	return accountCmd
}

func (cmd *Cmd) accountNameCompleter([]string) []string {
	items := []string{"main"}

	for name := range cmd.config.Cred.Sub {
		items = append(items, name)
	}

	return items
}

func (cmd *Cmd) execAccountAddCmd(c *ishell.Context) {
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
	screenName, err := cmd.parseAccountCmdArgs(c.Args)
	if err != nil {
		cmd.showErrorMessage(err.Error())
		return
	}

	// 実行確認
	msg := fmt.Sprintf("Remove account (%s) from twnyan?", screenName)
	if ok := cmd.showExecutionConf(msg); !ok {
		cmd.showMessage("CANCELED", "Interrupted", cmd.config.Color.Accent2)
		return
	}

	// 認証情報を削除
	switch screenName {
	case "main":
		cmd.config.Cred.Main = &oauth.Credentials{}
	default:
		delete(cmd.config.Cred.Sub, screenName)
	}

	cmd.config.Save()

	cmd.showMessage("REMOVED", screenName, cmd.config.Color.Accent3)
}

func (cmd *Cmd) execAccountListCmd(c *ishell.Context) {
	for _, name := range cmd.accountNameCompleter(nil) {
		if name != "main" {
			name = fmt.Sprintf("@%s", name)
		}
		fmt.Printf("- %s\n", name)
	}
}

func (cmd *Cmd) execAccountSwitchCmd(c *ishell.Context) {
	if cmd.checkCommandLineMode() {
		return
	}

	screenName, err := cmd.parseAccountCmdArgs(c.Args)
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

	msg := fmt.Sprintf("%s -> %s", prevScreenName, cmd.twitter.OwnUser.ScreenName)
	cmd.showMessage("SWITCHED", msg, cmd.config.Color.Accent3)
}
