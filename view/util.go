package view

import (
	"fmt"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/arrow2nd/twnyan/util"
	"github.com/gookit/color"
)

// createUserStr ユーザー情報文字列作成
func (v *View) createUserStr(u *anaconda.User) string {
	// ユーザー名、スクリーンネーム
	name := color.HEX(v.cfg.Color.UserName).Sprint(u.Name)
	screenName := color.HEX(v.cfg.Color.ScreenName).Sprintf("@%s", u.ScreenName)
	// バッジ
	badge := ""
	if u.Verified {
		badge += color.HEX(v.cfg.Color.Verified).Sprint(" verified")
	}
	if u.Protected {
		badge += color.HEX(v.cfg.Color.Protected).Sprint(" protected")
	}
	// 結合
	text := fmt.Sprintf("%s %s%s", name, screenName, badge)
	return text
}

// createTimeStr 投稿時刻文字列作成
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

// createReactionCountStr リアクション数文字列作成
func (v *View) createReactionCountStr(count int, flg bool, unit string) string {
	// 表示色
	colorCode := v.cfg.Color.Favorite
	if unit == "RT" {
		colorCode = v.cfg.Color.Retweet
	}
	// カウントが0なら処理を終了
	if count <= 0 {
		return ""
	} else if count > 1 {
		unit += "s"
	}
	// 文字列作成
	text := " "
	if flg {
		text += color.HEXStyle(v.cfg.Color.BoxForground, colorCode).Sprintf(" %d%s ", count, unit)
	} else {
		text += color.HEX(colorCode).Sprintf("%d%s", count, unit)
	}
	return text
}
