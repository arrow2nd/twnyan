package main

import (
	"github.com/arrow2nd/twnyan/api"
	"github.com/arrow2nd/twnyan/cmd"
	"github.com/arrow2nd/twnyan/config"
)

func main() {
	var err error

	api := api.New()
	cfg := config.New()

	// TODO: ここらへんをまとめてcmd.Init()内に移動したい

	// 設定ファイルが無いなら認証
	if !cfg.Load() {
		cfg.Cred.Main, _, err = api.Auth()
		if err != nil {
			panic(err)
		}
		cfg.Save()
	}

	// TODO: オプションによって認証情報を変えることで、マルチアカウント対応させる
	api.Init(cfg.Cred.Main)

	cmd := cmd.New(cfg, api)
	cmd.Run()
}
