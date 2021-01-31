package view

import (
	"fmt"
	"strings"
	"time"

	"github.com/arrow2nd/twnyan/util"
	"github.com/gookit/color"
)

func (v *View) createTimeStr(t time.Time) string {
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

func (v *View) createSeparator() string {
	width := util.GetWindowWidth()
	return color.HEX(v.cfg.Color.Separator).Sprintf("%s\n", strings.Repeat("-", width))
}
