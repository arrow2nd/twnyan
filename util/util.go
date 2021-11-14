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
		fmt.Fprintln(os.Stderr, "Error: Could not get window size")
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

// AllReplace 該当文字列を全て置換
func AllReplace(str, reg, rep string) string {
	replace := regexp.MustCompile(reg)
	return replace.ReplaceAllString(str, rep)
}

// IndexOf 引数と同じ内容を持つ最初の配列要素の添字を返す
func IndexOf(array []string, str string) (int, bool) {
	for i, value := range array {
		if value == str {
			return i, true
		}
	}

	return 0, false
}

// IsThreeDigitsNumber ３桁までの数値かどうか
func IsThreeDigitsNumber(str string) bool {
	return regexp.MustCompile("^[0-9]{1,3}$").Match([]byte(str))
}

// IsSameDate 今日の日付かどうか
func IsSameDate(chkTime time.Time) bool {
	now := time.Now()
	location := now.Location()
	fixedChkTime := chkTime.In(location)

	t1 := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, location)
	t2 := time.Date(fixedChkTime.Year(), fixedChkTime.Month(), fixedChkTime.Day(), 0, 0, 0, 0, location)

	return t1.Equal(t2)
}

// IsEndLFCode 末尾が改行コードかどうか
func IsEndLFCode(text string) bool {
	return strings.HasSuffix(text, "\n") || strings.HasSuffix(text, "\r")
}
