package util

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/mattn/go-runewidth"
	"golang.org/x/term"
)

// GetWindowWidth ウィンドウの幅を取得
func GetWindowWidth() int {
	fd := int(os.Stdout.Fd())

	width, _, err := term.GetSize(fd)
	if err != nil {
		fmt.Println("Error: Could not get window size")
		panic(err)
	}

	return width
}

// TruncateString 文字列を指定した長さに丸める
func TruncateString(str string, width int) string {
	return runewidth.Truncate(str, width, "…")
}

// MatchesRegexp 正規表現が文字列にマッチするか
func MatchesRegexp(reg, str string) bool {
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
func IsSameDate(chkTime time.Time) bool {
	now := time.Now()
	t1 := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	t2 := time.Date(chkTime.Year(), chkTime.Month(), chkTime.Day(), 0, 0, 0, 0, now.Location())

	return t1.Equal(t2)
}

// IsEndLFCode 末尾が改行コードかどうか
func IsEndLFCode(text string) bool {
	return strings.HasSuffix(text, "\n") || strings.HasSuffix(text, "\r")
}
