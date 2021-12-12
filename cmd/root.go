package cmd

import (
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/arrow2nd/ishell"
	"github.com/arrow2nd/twnyan/api"
	"github.com/arrow2nd/twnyan/config"
	"github.com/arrow2nd/twnyan/view"
)

const versionStr = "1.7.0"

type Cmd struct {
	shell *ishell.Shell
	cfg   *config.Config
	api   *api.TwitterAPI
	view  *view.View
}

type AccumulateTweets map[int64]anaconda.Tweet

// New 構造体を初期化
func New(c *config.Config, a *api.TwitterAPI) *Cmd {
	nc := &Cmd{
		shell: ishell.New(),
		cfg:   c,
		api:   a,
		view:  view.New(c),
	}

	nc.init()
	return nc
}

// init 初期化
func (cmd *Cmd) init() {
	cmd.setDefaultPrompt()

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
	// 引数があるなら直接実行
	if len(os.Args) > 1 {
		if err := cmd.shell.Process(os.Args[1:]...); err != nil {
			cmd.showErrorMessage(err.Error())
		}
		os.Exit(0)
	}

	// 対話モードで実行
	cmd.shell.Process("timeline")
	cmd.shell.Run()
}
