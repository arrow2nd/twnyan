package util

import (
	"errors"
	"log"
	"os"
	"regexp"
	"time"

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

// CutString 文字列を指定した長さに丸める
func CutString(str string, width int) string {
	return runewidth.Truncate(str, width, "...")
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

// FetchStringSpecifiedType 引数から指定した形の文字列を取り出す
func FetchStringSpecifiedType(args []string, aType ...string) ([]string, error) {
	if l := len(args); l <= 0 || l > len(aType) {
		return nil, errors.New("No arguments or too many")
	}

	results := make([]string, len(aType))
	for _, v := range args {
		for i, t := range aType {
			if results[i] != "" {
				continue
			}
			isNum := IsNumber(v)
			if t == "num" && isNum || t == "str" && !isNum {
				results[i] = v
				break
			}
		}
	}
	return results, nil
}

// IsSameDate 今日の日付かどうか
func IsSameDate(a time.Time) bool {
	t := time.Now()
	t1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	t2 := time.Date(a.Year(), a.Month(), a.Day(), 0, 0, 0, 0, t.Location())
	return t1.Equal(t2)
}
