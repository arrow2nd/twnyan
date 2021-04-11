package cmd

import (
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/arrow2nd/ishell"
	"github.com/arrow2nd/twnyan/api"
	"github.com/arrow2nd/twnyan/config"
	"github.com/arrow2nd/twnyan/view"
)

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
		view:  nil,
	}
	nc.view = view.New(nc.cfg)

	return nc
}

// Init 初期化
func (cmd *Cmd) Init() {
	// コマンドを登録
	cmd.addTweetCmd()
	cmd.addReplyCmd()
	cmd.addTimelineCmd()
	cmd.addMentionCmd()
	cmd.addListCmd()
	cmd.addSearchCmd()
	cmd.addUserCmd()
	cmd.addLikeCmd()
	cmd.addRetweetCmd()
	cmd.addQuoteCmd()
	cmd.addFollowCmd()
	cmd.addBlockCmd()
	cmd.addMuteCmd()
	cmd.addOpenCmd()
	cmd.addVersionCmd()
	cmd.addStreamCmd()

	// プロンプトを設定
	cmd.setDefaultPrompt()

	// コマンドエラーの表示を設定
	cmd.shell.NotFound(func(c *ishell.Context) {
		cmd.showErrorMessage("command not found: " + c.ReadLine())
	})
}

// Run 実行
func (cmd *Cmd) Run() {
	// コマンドライン引数がある
	if len(os.Args) > 1 {
		err := cmd.shell.Process(os.Args[1:]...)
		if err != nil {
			cmd.showErrorMessage(err.Error())
		}
		os.Exit(0)
	}

	// 対話モードで実行
	cmd.shell.Process("timeline")
	cmd.shell.Run()
}
