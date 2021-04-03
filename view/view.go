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

// View 表示
type View struct {
	tweets []anaconda.Tweet
	cfg    *config.Config
	mu     sync.Mutex
}

// New 表示
func New(c *config.Config) *View {
	v := &View{
		tweets: []anaconda.Tweet{},
		cfg:    c,
	}
	return v
}

// createPostTimeStr 投稿時刻
func (v *View) createPostTimeStr(t time.Time) string {
	postTime := ""

	// 今日の時刻だった場合、日付を省く
	if util.IsSameDate(t) {
		postTime = t.Local().Format(v.cfg.Option.TimeFormat)
	} else {
		format := fmt.Sprintf("%s %s", v.cfg.Option.DateFormat, v.cfg.Option.TimeFormat)
		postTime = t.Local().Format(format)
	}

	return color.HEX(v.cfg.Color.Accent2).Sprint(postTime)
}

// createSeparatorStr セパレータ
func (v *View) createSeparatorStr(space bool) string {
	width := util.GetWindowWidth()
	if space {
		width -= 2
	}

	sep := color.HEX(v.cfg.Color.Separator).Sprintf("%s\n", strings.Repeat("-", width))
	if space {
		return " " + sep
	}

	return sep
}
