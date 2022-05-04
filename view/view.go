package view

import (
	"fmt"
	"strings"
	"time"

	"github.com/arrow2nd/twnyan/config"
	"github.com/arrow2nd/twnyan/util"
	"github.com/gookit/color"
)

// View 表示処理
type View struct {
	config *config.Config
}

// New 生成
func New(c *config.Config) *View {
	return &View{
		config: c,
	}
}

// createCreatedAtString 投稿時刻の文字列を作成
func (v *View) createCreatedAtString(postTime time.Time) string {
	format := ""

	// 今日の日付なら時刻のみを表示
	if util.IsSameDate(postTime) {
		format = v.config.Option.TimeFormat
	} else {
		format = fmt.Sprintf("%s %s", v.config.Option.DateFormat, v.config.Option.TimeFormat)
	}

	return color.HEX(v.config.Color.Accent2).Sprint(
		postTime.Local().Format(format),
	)
}

// createSeparatorString セパレータ文字列を作成
func (v *View) createSeparatorString(hasInsertSpace bool) string {
	width := util.GetWindowWidth()
	if hasInsertSpace {
		width -= 2
	}

	sep := color.HEX(v.config.Color.Separator).Sprintf("%s\n", strings.Repeat("-", width))
	if hasInsertSpace {
		return " " + sep
	}

	return sep
}
