package cmd

import (
	"fmt"
	"os"

	"github.com/arrow2nd/twnyan/api"
	"github.com/arrow2nd/twnyan/config"
	"github.com/arrow2nd/twnyan/view"
	"gopkg.in/abiosoft/ishell.v2"
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
		view:  view.New(c),
	}
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
	cmd.newFollowCmd()
	cmd.newBlockCmd()
	cmd.newMuteCmd()
	cmd.newOpenCmd()
	cmd.newVersionCmd()
	// プロンプト設定
	prompt := fmt.Sprintf("@%s: ", cmd.api.OwnUser.ScreenName)
	cmd.shell.SetPrompt(prompt)
}

// Run 実行
func (cmd *Cmd) Run() {
	if len(os.Args) > 1 {
		cmd.shell.Process(os.Args[1:]...)
	} else {
		cmd.shell.Process("timeline")
		cmd.shell.Run()
	}
}
