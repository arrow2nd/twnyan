package config

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/arrow2nd/twnyan/util"
	"github.com/gookit/color"
	"gopkg.in/yaml.v2"
)

// Configuration 設定情報
type Configuration struct {
	Credentials Credentials
	Color       ColorScheme
	Default     DefaultValue
}

// Credentials トークン情報
type Credentials struct {
	Token  string
	Secret string
}

// ColorScheme 配色
type ColorScheme struct {
	Accent1   string
	Accent2   string
	Accent3   string
	BoxFg     string
	UserName  string
	UserID    string
	Separator string
	Reply     string
	Hashtag   string
	Fav       string
	RT        string
	Verified  string
	Protected string
	Follow    string
	Block     string
	Mute      string
}

// DefaultValue デフォルト値
type DefaultValue struct {
	Counts     string
	Prompt     string
	DateFormat string
	TimeFormat string
}

// Save ファイル保存
func (cfg *Configuration) Save() {
	// byteに変換
	buf, err := yaml.Marshal(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// 保存
	filePath := getFilePath()
	err = ioutil.WriteFile(filePath, buf, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

// Load ファイル読み込み
func (cfg *Configuration) Load() error {
	// 設定ファイルの存在チェック
	filePath := getFilePath()
	if _, err := os.Stat(filePath); err != nil {
		return errors.New("Not found")
	}

	// 読込
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// 構造体にマッピング
	err = yaml.Unmarshal(buf, cfg)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// Remove ファイル削除
func (cfg *Configuration) Remove() {
	// 実行確認
	if !util.ExecConfirmation("Delete the configuration file. Are you sure?", "Deletion canceled") {
		return
	}

	// ファイルパスを取得
	filePath := getFilePath()

	// 削除
	err := os.Remove(filePath)
	if err != nil {
		color.Error.Prompt("Failed to delete the file")
		return
	}

	util.ShowSuccessMsg("Success", "Configuration files have been deleted", cfg.Color.BoxFg, cfg.Color.Accent3)
}

// getFilePath 設定ファイルのパスを取得
func getFilePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return path.Join(home, ".twnyan.yaml")
}
