package config

import (
	"fmt"
	"log"

	"github.com/ChimeraCoder/anaconda"
	"github.com/gookit/color"
	"github.com/pkg/browser"
)

// Setup 初期設定
func Setup() {
	// logo
	color.FgRed.Println(" _ ")
	color.FgLightRed.Println("| |___      ___ __  _   _  __ _ _ __")
	color.FgYellow.Println("| __\\ \\ /\\ / / '_ \\| | | |/ _` | '_ \\")
	color.FgGreen.Println("| |_ \\ V  V /| | | | |_| | (_| | | | |")
	color.FgCyan.Println(" \\__| \\_/\\_/ |_| |_|\\__, |\\__,_|_| |_|")
	color.FgLightBlue.Println("                    |___/ \n")

	data := defaultConfig()

	// 認証用URLを生成
	authAPI := anaconda.NewTwitterApi("", "")
	uri, cred, err := authAPI.AuthorizationURL("oob")
	if err != nil {
		log.Fatal(err)
	}

	// 認証ページをブラウザで開く
	color.Notice.Tips("Please enter the PIN code displayed after authenticating the application")
	fmt.Printf("(URL: %s)\n\n", uri)
	browser.OpenURL(uri)

	// PIN入力
	pin := ""
	color.HEX("#FFA6C3").Print("ฅ^•ω•^ฅ : ")
	fmt.Scanf("%s", &pin)

	// トークン発行
	cred, _, err = authAPI.GetCredentials(cred, pin)
	if err != nil {
		log.Fatal(err)
	}
	data.Credentials.Token, data.Credentials.Secret = cred.Token, cred.Secret

	// 保存
	data.Save()
}

func defaultConfig() *Configuration {
	data := Configuration{}

	// 配色
	data.Color.Accent1 = "#EF7701"
	data.Color.Accent2 = "#EF9C02"
	data.Color.Accent3 = "#FCC01D"
	data.Color.BoxFg = "#000000"
	data.Color.UserName = "#FAF8F7"
	data.Color.UserID = "#9C9C9C"
	data.Color.Separator = "#9C9C9C"
	data.Color.Reply = "#4DD0E1"
	data.Color.Hashtag = "#64B5F6"
	data.Color.Fav = "#F06292"
	data.Color.RT = "#4DB6AC"
	data.Color.Verified = "#4E88E5"
	data.Color.Protected = "#787878"
	data.Color.Follow = "#1E88E5"
	data.Color.Block = "#EF5350"
	data.Color.Mute = "#FFF176"

	// デフォルト値
	data.Default.Counts = "25"
	data.Default.Prompt = ": "
	data.Default.DateFormat = "2006/01/02"
	data.Default.TimeFormat = "15:04:05"

	return &data
}
