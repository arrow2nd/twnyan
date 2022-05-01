package main

import (
	"github.com/arrow2nd/twnyan/api"
	"github.com/arrow2nd/twnyan/cmd"
	"github.com/arrow2nd/twnyan/config"
)

func main() {
	api := api.New()
	cfg := config.New()

	// 設定ファイルが無いなら認証
	if !cfg.Load() {
		cfg.Cred.Main = api.Auth()
		cfg.Save()
	}

	// TODO: オプションによって認証情報を変えることで、マルチアカウント対応させる
	api.Init(cfg.Cred.Main)

	cmd := cmd.New(cfg, api)
	cmd.Run()
}
