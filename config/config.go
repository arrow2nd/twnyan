package config

import "github.com/garyburd/go-oauth/oauth"

const (
	crdFile = ".cred.yaml"
	optFile = "option.yaml"
	colFile = "color.yaml"
)

type cred struct {
	Main *oauth.Credentials `yaml:"Main,inline"`
	Sub  map[string]*oauth.Credentials
}

type option struct {
	ConfigDir  string
	Counts     string
	DateFormat string
	TimeFormat string
}

type color struct {
	Accent1      string
	Accent2      string
	Accent3      string
	Error        string
	BoxForground string
	Separator    string
	UserName     string
	ScreenName   string
	Reply        string
	Hashtag      string
	Favorite     string
	Retweet      string
	Verified     string
	Protected    string
	Following    string
	FollowedBy   string
	Block        string
	Mute         string
}

// Config 設定構造体
type Config struct {
	Cred   *cred
	Option *option
	Color  *color
}

// New 構造体を初期化
func New() *Config {
	return &Config{
		Cred: &cred{
			Main: &oauth.Credentials{},
			Sub:  map[string]*oauth.Credentials{},
		},
		Option: &option{
			ConfigDir:  getConfigDir(),
			Counts:     "25",
			DateFormat: "2006/01/02",
			TimeFormat: "15:04:05",
		},
		Color: &color{
			Accent1:      "#e06c75",
			Accent2:      "#c678dd",
			Accent3:      "#56b6c2",
			Error:        "#e06c75",
			BoxForground: "#000000",
			Separator:    "#707070",
			UserName:     "#faf8f7",
			ScreenName:   "#9c9c9c",
			Reply:        "#56b6c2",
			Hashtag:      "#61afef",
			Favorite:     "#e887b9",
			Retweet:      "#98c379",
			Verified:     "#5685d1",
			Protected:    "#787878",
			Following:    "#1877c9",
			FollowedBy:   "#18a0c9",
			Block:        "#e06c75",
			Mute:         "#e5c07b",
		},
	}
}

// Save 保存
func (cfg *Config) Save() {
	cfg.saveYaml(crdFile, cfg.Cred)
	cfg.saveYaml(optFile, cfg.Option)
	cfg.saveYaml(colFile, cfg.Color)
}

// Load 読込
func (cfg *Config) Load() bool {
	if !cfg.configFileExists() {
		return false
	}

	cfg.loadYaml(crdFile, cfg.Cred)
	cfg.loadYaml(optFile, cfg.Option)
	cfg.loadYaml(colFile, cfg.Color)

	return true
}
