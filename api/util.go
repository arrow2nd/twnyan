package api

import (
	"fmt"
	"net/url"
	"regexp"

	"github.com/gookit/color"
)

// CreateQuery クエリを作成
func CreateQuery(count string) url.Values {
	q := url.Values{}
	q.Add("tweet_mode", "extended")
	q.Add("count", count)

	return q
}

// parseAPIErrorMsg エラーメッセージをパース
func parseAPIErrorMsg(err error) string {
	bytes := []byte(err.Error())
	errMsg := regexp.MustCompile(`"(message|error)":"([^"]+)"`).FindSubmatch(bytes)

	return fmt.Sprintf("%s", errMsg[2])
}

// showLogo ロゴを表示
func showLogo() {
	color.Red.Println(" __                                     ")
	color.Yellow.Println("|  |_.--.--.--.-----.--.--.---.-.-----.")
	color.Green.Println("|   _|  |  |  |     |  |  |  _  |     |")
	color.Cyan.Println("|____|________|__|__|___  |___._|__|__|")
	color.Blue.Println("                    |_____|            ")
}

// showAuthUrl 認証URLを表示
func showAuthUrl(uri string) {
	fmt.Println("\n🐈 Go to the following URL to authenticate the application and enter the PIN that is displayed")
	fmt.Printf("[ %s ]\n\n", uri)
}

// inputPinCode PINコードを入力
func inputPinCode() string {
	pin := ""

	fmt.Print("PIN : ")
	fmt.Scanf("%s", &pin)

	return pin
}
