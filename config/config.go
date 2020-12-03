package config

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path"

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
	configPath := getConfigFilePath()
	err = ioutil.WriteFile(configPath, buf, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

// Load ファイル読み込み
func (cfg *Configuration) Load() error {
	configPath := getConfigFilePath()

	// 設定ファイルの存在チェック
	if _, err := os.Stat(configPath); err != nil {
		return errors.New("Not found")
	}

	// 読込
	buf, err := ioutil.ReadFile(configPath)
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
	configPath := getConfigFilePath()

	err := os.Remove(configPath)
	if err != nil {
		color.Error.Tips("Failed to delete the file")
		return
	}

	color.Success.Tips("Configuration files have been deleted")
}

func getConfigFilePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return path.Join(home, ".twnyan.yaml")
}
