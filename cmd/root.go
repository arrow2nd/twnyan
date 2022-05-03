package cmd

import (
	"os"

	"github.com/arrow2nd/ishell"
	"github.com/arrow2nd/twnyan/config"
	"github.com/arrow2nd/twnyan/twitter"
	"github.com/arrow2nd/twnyan/view"
)

// Cmd コマンド管理
type Cmd struct {
	shell   *ishell.Shell
	config  *config.Config
	twitter *twitter.Twitter
	view    *view.View
}

// New 作成
func New(c *config.Config, t *twitter.Twitter) *Cmd {
	nc := &Cmd{
		shell:   ishell.New(),
		config:  c,
		twitter: t,
		view:    view.New(c),
	}

	// TODO: 後でAuth処理と統合
	nc.init()

	return nc
}

// init 初期化
func (cmd *Cmd) init() {
	cmd.setDefaultPrompt()

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

	// コマンドエラー時の表示を設定
	cmd.shell.NotFound(func(c *ishell.Context) {
		cmd.showErrorMessage("command not found: " + c.Args[0])
	})
}

// Run 実行
func (cmd *Cmd) Run() {
	// 対話モードで実行
	if len(os.Args) <= 1 {
		cmd.shell.Process("timeline")
		cmd.shell.Run()
		return
	}

	// 直接実行
	if err := cmd.shell.Process(os.Args[1:]...); err != nil {
		os.Exit(1)
	}
}
