package config

const (
	crdFile = ".cred.yaml"
	optFile = "option.yaml"
	colFile = "color.yaml"
)

// Config 設定構造体
type Config struct {
	Cred   *cred
	Option *option
	Color  *color
}

type cred struct {
	Token  string
	Secret string
}

type option struct {
	ConfigDir  string
	Prompt     string
	Counts     string
	DateFormat string
	TimeFormat string
}

type color struct {
	Accent1      string
	Accent2      string
	Accent3      string
	Dim          string
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
	Follow       string
	Block        string
	Mute         string
}

// New 設定構造体作成
func New() *Config {
	cfg := &Config{
		Cred: &cred{
			Token:  "",
			Secret: "",
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
			Dim:          "#343a44",
			BoxForground: "#000000",
			Separator:    "#9c9c9c",
			UserName:     "#faf8f7",
			ScreenName:   "#9c9c9c",
			Reply:        "#56b6c2",
			Hashtag:      "#61afef",
			Favorite:     "#e887b9",
			Retweet:      "#98c379",
			Verified:     "#5685d1",
			Protected:    "#787878",
			Follow:       "#1877c9",
			Block:        "#e06c75",
			Mute:         "#e5c07b",
		},
	}
	return cfg
}

// Save 保存
func (cfg *Config) Save() {
	saveYaml(cfg.Option.ConfigDir, crdFile, cfg.Cred)
	saveYaml(cfg.Option.ConfigDir, optFile, cfg.Option)
	saveYaml(cfg.Option.ConfigDir, colFile, cfg.Color)
}

// Load 読込
func (cfg *Config) Load() bool {
	if !configFileExists(cfg.Option.ConfigDir) {
		return false
	}
	loadYaml(cfg.Option.ConfigDir, crdFile, cfg.Cred)
	loadYaml(cfg.Option.ConfigDir, optFile, cfg.Option)
	loadYaml(cfg.Option.ConfigDir, colFile, cfg.Color)
	return true
}
