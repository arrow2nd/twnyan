package cmd

import (
	"os"

	"github.com/arrow2nd/ishell"
	"github.com/arrow2nd/twnyan/api"
	"github.com/arrow2nd/twnyan/config"
	"github.com/arrow2nd/twnyan/view"
)

// Cmd コマンド
type Cmd struct {
	shell *ishell.Shell
	cfg   *config.Config
	api   *api.TwitterAPI
	view  *view.View
}

// New コマンド構造体作成
func New(c *config.Config, a *api.TwitterAPI) *Cmd {
	nc := &Cmd{
		shell: ishell.New(),
		cfg:   c,
		api:   a,
		view:  nil,
	}
	nc.view = view.New(nc.cfg)
	return nc
}

// Init 初期化
func (cmd *Cmd) Init() {
	// コマンド登録
	cmd.newTweetCmd()
	cmd.newReplyCmd()
	cmd.newTimelineCmd()
	cmd.newMentionCmd()
	cmd.newListCmd()
	cmd.newSearchCmd()
	cmd.newUserCmd()
	cmd.newFaoriteCmd()
	cmd.newRetweetCmd()
	cmd.newQuoteCmd()
	cmd.newFollowCmd()
	cmd.newBlockCmd()
	cmd.newMuteCmd()
	cmd.newOpenCmd()
	cmd.newVersionCmd()
	// プロンプト設定
	cmd.setDefaultPrompt()
}

// Run 実行
func (cmd *Cmd) Run() {
	if len(os.Args) > 1 {
		err := cmd.shell.Process(os.Args[1:]...)
		if err != nil {
			cmd.drawErrorMessage(err.Error())
		}
	} else {
		cmd.shell.Process("timeline")
		cmd.shell.Run()
	}
}
