package main

import (
	"github.com/arrow2nd/twnyan/api"
	"github.com/arrow2nd/twnyan/cmd"
	"github.com/arrow2nd/twnyan/config"
)

func main() {
	api := api.New()
	cfg := config.New()

	if !cfg.Load() {
		// 設定ファイルが無いなら認証する
		cfg.Cred.Token, cfg.Cred.Secret = api.Auth()
		cfg.Save()
	} else {
		api.Init(cfg.Cred.Token, cfg.Cred.Secret)
	}

	cmd := cmd.New(cfg, api)
	cmd.Run()
}
