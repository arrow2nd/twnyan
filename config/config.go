package config

const (
	crdFile = ".cred.yaml"
	optFile = "option.yaml"
	colFile = "color.yaml"
)

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

// New 構造体を初期化
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

	return cfg
}

// Save 保存
func (cfg *Config) Save() {
	saveYAML(cfg.Option.ConfigDir, crdFile, cfg.Cred)
	saveYAML(cfg.Option.ConfigDir, optFile, cfg.Option)
	saveYAML(cfg.Option.ConfigDir, colFile, cfg.Color)
}

// Load 読込
func (cfg *Config) Load() bool {
	if !configFileExists(cfg.Option.ConfigDir) {
		return false
	}

	loadYAML(cfg.Option.ConfigDir, crdFile, cfg.Cred)
	loadYAML(cfg.Option.ConfigDir, optFile, cfg.Option)
	loadYAML(cfg.Option.ConfigDir, colFile, cfg.Color)

	return true
}
