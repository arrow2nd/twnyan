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

// createAPIErrorMsg エラーメッセージを作成
func (tw *TwitterAPI) createAPIErrorMsg(resourceName string, err error) string {
	// エラー文字列からメッセージを抽出
	bytes := []byte(err.Error())
	result := regexp.MustCompile(`"(message|error)":"([^"]+)"`).FindSubmatch(bytes)
	if len(result) <= 0 {
		return ""
	}

	errMsg := fmt.Sprintf("%s", result[2])

	// レート制限なら解除時刻を追加
	if errMsg == "Rate limit exceeded" && resourceName != "" {
		resetTimeStr := tw.fetchRateLimitResetTime(resourceName)
		errMsg += fmt.Sprintf(" (Reset Time : %s)", resetTimeStr)
	}

	return errMsg
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
