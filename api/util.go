package api

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"

	"github.com/gookit/color"
)

// CreateQuery クエリを作成
func CreateQuery(count string) url.Values {
	return url.Values{
		"tweet_mode": {"extended"},
		"count":      {count},
	}
}

// createUserInfoText ユーザー情報の文字列を作成
func (tw *TwitterAPI) createUserInfoText(name, screenName string) string {
	return fmt.Sprintf("%s @%s", name, screenName)
}

// createAPIErrorText APIのエラーメッセージを作成
func (tw *TwitterAPI) createAPIErrorText(resource string, err error) string {
	bytes := []byte(err.Error())

	// エラー文字列からメッセージを抽出
	result := regexp.MustCompile(`"(message|error)":"([^"]+)"`).FindSubmatch(bytes)
	if len(result) <= 0 {
		return ""
	}

	errMsg := string(result[2])

	// レート制限なら解除時刻を追加
	if errMsg == "Rate limit exceeded" && resource != "" {
		resetTimeStr := tw.fetchRateLimitResetTime(resource)
		errMsg += fmt.Sprintf(" (Reset Time : %s)", resetTimeStr)
	}

	return errMsg
}

// createAPIError APIのエラーを作成
func (tw *TwitterAPI) createAPIError(resource string, err error) error {
	return errors.New(tw.createAPIErrorText(resource, err))
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
	fmt.Println("\n🐈  Go to the following URL to authenticate the application and enter the PIN that is displayed")
	fmt.Printf("[ %s ]\n\n", uri)
}

// inputPinCode PINコードを入力
func inputPinCode() string {
	fmt.Print("PIN : ")

	pin := ""
	fmt.Scanf("%s", &pin)

	return pin
}
