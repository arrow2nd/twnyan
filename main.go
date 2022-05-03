package main

import (
	"github.com/arrow2nd/twnyan/cmd"
	"github.com/arrow2nd/twnyan/config"
	"github.com/arrow2nd/twnyan/twitter"
)

func main() {
	var err error

	twitter := twitter.New()
	config := config.New()

	// TODO: ここらへんをまとめてcmd.Init()内に移動したい

	// 設定ファイルが無いなら認証
	if !config.Load() {
		config.Cred.Main, _, err = twitter.Auth()
		if err != nil {
			panic(err)
		}
		config.Save()
	}

	// TODO: オプションによって認証情報を変えることで、マルチアカウント対応させる
	twitter.Init(config.Cred.Main)

	cmd := cmd.New(config, twitter)
	cmd.Run()
}
