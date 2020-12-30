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
	data.Color.Accent1 = "#E06C75"
	data.Color.Accent2 = "#C678DD"
	data.Color.Accent3 = "#56B6C2"
	data.Color.BoxFg = "#000000"
	data.Color.UserName = "#FAF8F7"
	data.Color.UserID = "#9C9C9C"
	data.Color.Separator = "#9C9C9C"
	data.Color.Reply = "#56B6C2"
	data.Color.Hashtag = "#61AFEF"
	data.Color.Fav = "#E887B9"
	data.Color.RT = "#98C379"
	data.Color.Verified = "#5685D1"
	data.Color.Protected = "#787878"
	data.Color.Follow = "#1877C9"
	data.Color.Block = "#E06C75"
	data.Color.Mute = "#E5C07B"

	// デフォルト値
	data.Default.Counts = "25"
	data.Default.Prompt = ": "
	data.Default.DateFormat = "2006/01/02"
	data.Default.TimeFormat = "15:04:05"

	return &data
}
