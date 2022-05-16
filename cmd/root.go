package cmd

import (
	"fmt"
	"os"

	"github.com/arrow2nd/ishell/v2"
	"github.com/arrow2nd/twnyan/config"
	"github.com/arrow2nd/twnyan/twitter"
	"github.com/arrow2nd/twnyan/view"
	"github.com/spf13/pflag"
)

// Cmd コマンド管理
type Cmd struct {
	shell   *ishell.Shell
	flagSet *pflag.FlagSet
	config  *config.Config
	twitter *twitter.Twitter
	view    *view.View
}

// New 作成
func New() *Cmd {
	config := config.New()

	return &Cmd{
		shell:   ishell.New(),
		flagSet: pflag.NewFlagSet("twnyan", pflag.ContinueOnError),
		config:  config,
		twitter: twitter.New(),
		view:    view.New(config),
	}
}

// Init 初期化
func (cmd *Cmd) Init() {
	var err error

	// コマンド・フラグを登録
	cmd.registerCommands()
	cmd.registerFlags()

	// 設定ファイル読み込み
	ok := cmd.config.Load()

	// メインアカウントがない場合、認証
	if !ok || cmd.config.Cred.Main.Token == "" || cmd.config.Cred.Main.Secret == "" {
		if cmd.config.Cred.Main, _, err = cmd.twitter.Auth(); err != nil {
			cmd.showErrorMessage(err.Error())
			os.Exit(1)
		}
		cmd.config.Save()
	}

	// フラグをパース
	if err := cmd.flagSet.Parse(os.Args[1:]); err != nil {
		cmd.showErrorMessage(err.Error())
		os.Exit(1)
	}

	// ログイン
	if err = cmd.login(); err != nil {
		cmd.showErrorMessage(err.Error())
		os.Exit(1)
	}

	cmd.setDefaultPrompt()
}

// Run 実行
func (cmd *Cmd) Run() {
	// ヘルプの表示
	if ok, _ := cmd.flagSet.GetBool("help"); ok {
		fmt.Print(cmd.shell.HelpText())
		cmd.flagSet.Usage()
		return
	}

	// 対話モードで実行
	if cmd.flagSet.NArg() == 0 {
		cmd.shell.Process("timeline")
		cmd.shell.Run()
		return
	}

	// 直接実行
	if err := cmd.shell.Process(cmd.flagSet.Args()...); err != nil {
		os.Exit(1)
	}
}

func (cmd *Cmd) registerFlags() {
	// フラグを登録
	cmd.flagSet.StringP("account", "a", "", "specify the account to use")
	cmd.flagSet.BoolP("help", "h", false, "display help with options")

	// フラグのヘルプ表示
	cmd.flagSet.Usage = func() {
		fmt.Println("Flags:")
		cmd.flagSet.PrintDefaults()
		fmt.Print("\n")
	}
}

func (cmd *Cmd) registerCommands() {
	// コマンドを登録
	cmd.shell.AddCmd(cmd.newAccountCmd())
	cmd.shell.AddCmd(cmd.newTweetCmd())
	cmd.shell.AddCmd(cmd.newReplyCmd())
	cmd.shell.AddCmd(cmd.newTimelineCmd())
	cmd.shell.AddCmd(cmd.newMentionCmd())
	cmd.shell.AddCmd(cmd.newListCmd())
	cmd.shell.AddCmd(cmd.newSearchCmd())
	cmd.shell.AddCmd(cmd.newUserCmd())
	cmd.shell.AddCmd(cmd.newLikeCmd())
	cmd.shell.AddCmd(cmd.newRetweetCmd())
	cmd.shell.AddCmd(cmd.newLikertCmd())
	cmd.shell.AddCmd(cmd.newQuoteCmd())
	cmd.shell.AddCmd(cmd.newFollowCmd())
	cmd.shell.AddCmd(cmd.newBlockCmd())
	cmd.shell.AddCmd(cmd.newMuteCmd())
	cmd.shell.AddCmd(cmd.newOpenCmd())
	cmd.shell.AddCmd(cmd.newStreamCmd())
	cmd.shell.AddCmd(cmd.newVersionCmd())

	// コマンドエラー時の表示
	cmd.shell.NotFound(func(c *ishell.Context) {
		cmd.showErrorMessage(fmt.Sprintf("command not found: %s", c.Args[0]))
	})
}

func (cmd *Cmd) login() error {
	screenName, _ := cmd.flagSet.GetString("account")

	// メインアカウント
	if screenName == "" {
		cmd.twitter.Init(cmd.config.Cred.Main)
		return nil
	}

	// サブアカウント
	if cred, ok := cmd.config.Cred.Sub[screenName]; ok {
		cmd.twitter.Init(cred)
		return nil
	}

	return fmt.Errorf("account does not exist: %s", screenName)
}
