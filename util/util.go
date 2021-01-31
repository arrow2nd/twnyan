package util

import (
	"fmt"
	"html"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/gookit/color"
	"github.com/mattn/go-runewidth"
	"golang.org/x/crypto/ssh/terminal"
)

// GetWindowWidth ウィンドウの幅を取得
func GetWindowWidth() int {
	fd := int(os.Stdout.Fd())
	w, _, err := terminal.GetSize(fd)
	if err != nil {
		log.Fatal(err)
	}
	return w
}

// ShowSuccessMsg 処理完了メッセージを表示
func ShowSuccessMsg(tips, text, fg, bg string) {
	AllReplace(&text, "[\n\r]", "")
	text = html.UnescapeString(text)
	cutText := CutString(text, GetWindowWidth()-len(tips)-2)
	tips = color.HEXStyle(fg, bg).Sprintf(" %s ", tips)
	fmt.Printf("%s %s\n", tips, cutText)
}

// CutString 文字列を指定した長さに丸める
func CutString(str string, width int) string {
	return runewidth.Truncate(str, width, "…")
}

// ChkRegexp 文字列が含まれるかどうか
func ChkRegexp(reg, str string) bool {
	return regexp.MustCompile(reg).Match([]byte(str))
}

// IsNumber ３桁までの数値かどうか
func IsNumber(str string) bool {
	return regexp.MustCompile("^[0-9]{1,3}$").Match([]byte(str))
}

// AllReplace 該当文字列を全て置換
func AllReplace(str *string, reg, rep string) {
	replace := regexp.MustCompile(reg)
	*str = replace.ReplaceAllString(*str, rep)
}

// IndexOf 文字列配列内を検索
func IndexOf(array []string, str string) int {
	for i, v := range array {
		if v == str {
			return i
		}
	}
	return -1
}

// IsSameDate 今日の日付かどうか
func IsSameDate(a time.Time) bool {
	t := time.Now()
	t1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	t2 := time.Date(a.Year(), a.Month(), a.Day(), 0, 0, 0, 0, t.Location())
	return t1.Equal(t2)
}
