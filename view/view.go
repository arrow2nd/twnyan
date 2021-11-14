package view

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/arrow2nd/twnyan/config"
	"github.com/arrow2nd/twnyan/util"
	"github.com/gookit/color"
)

type View struct {
	tweets []anaconda.Tweet
	cfg    *config.Config
	mu     sync.Mutex
}

// New 構造体を初期化
func New(c *config.Config) *View {
	return &View{
		tweets: []anaconda.Tweet{},
		cfg:    c,
	}
}

// createCreatedAtString 投稿時刻の文字列を作成
func (v *View) createCreatedAtString(postTime time.Time) string {
	postTimeStr := ""

	// 今日の時刻なら、日付を省略する
	if util.IsSameDate(postTime) {
		postTimeStr = postTime.Local().Format(v.cfg.Option.TimeFormat)
	} else {
		format := fmt.Sprintf("%s %s", v.cfg.Option.DateFormat, v.cfg.Option.TimeFormat)
		postTimeStr = postTime.Local().Format(format)
	}

	return color.HEX(v.cfg.Color.Accent2).Sprint(postTimeStr)
}

// createSeparatorString セパレータ文字列を作成
func (v *View) createSeparatorString(hasInsertSpace bool) string {
	width := util.GetWindowWidth()
	if hasInsertSpace {
		width -= 2
	}

	sep := color.HEX(v.cfg.Color.Separator).Sprintf("%s\n", strings.Repeat("-", width))
	if hasInsertSpace {
		return " " + sep
	}

	return sep
}
