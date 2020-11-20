package twitter

import (
	"fmt"
	"html"
	"regexp"
	"strings"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/arrow2nd/twnyan/util"
	"github.com/gookit/color"
	"github.com/mattn/go-runewidth"
)

// showTweet ツイートを表示
func showTweet(i int, tw *anaconda.Tweet) {
	width := util.GetWindowWidth()
	status := ""
	space := ""
	isQuote := (i == -1)

	// 引用リツイートの場合セパレータを挿入
	if isQuote {
		status += createSeparator(true)
		space = " "
	}
	// リツイート元のツイートに置換
	if tw.RetweetedStatus != nil {
		status += color.HEX(cfg.Color.RT).Sprintf("%sRT by %s (@%s)\n", space, tw.User.Name, tw.User.ScreenName)
		tw = tw.RetweetedStatus
	}
	// リプライ先表示
	if tw.InReplyToScreenName != "" {
		status += color.HEX(cfg.Color.Reply).Sprintf("%sReply to @%s\n", space, tw.InReplyToScreenName)
	}

	// ツイート番号
	index := ""
	if !isQuote {
		index = color.HEXStyle(cfg.Color.BoxFg, cfg.Color.Accent1).Sprintf(" %d ", i)
	}

	// ユーザー情報
	user := createUserString(&tw.User)

	// ツイートテキスト
	text := runewidth.Wrap(html.UnescapeString(tw.FullText), width-2)
	util.Replace(&text, "\n", "\n ")

	// ハッシュタグハイライト
	rep := regexp.MustCompile("[#＃]([\\w\u05be\u05f3\u05f4]*[\\p{L}_]+[\\w\u05be\u05f3\u05f4]*)")
	text = rep.ReplaceAllString(text, color.HEX(cfg.Color.Hashtag).Sprintf("#$1"))
	// メンションハイライト
	rep = regexp.MustCompile("[@＠]([\\w]+)")
	text = rep.ReplaceAllString(text, color.HEX(cfg.Color.Reply).Sprintf("@$1"))

	// いいね・RT
	reaction := ""
	rCnt := []int{tw.FavoriteCount, tw.RetweetCount}
	rFlag := []bool{tw.Favorited, tw.Retweeted}
	rUnit := []string{"fav", "RT"}
	rCol := []string{cfg.Color.Fav, cfg.Color.RT}
	for i := 0; i < 2; i++ {
		if rCnt[i] <= 0 {
			continue
		} else if rCnt[i] > 1 {
			rUnit[i] += "s"
		}
		reaction += " "
		if rFlag[i] {
			reaction += color.HEXStyle(cfg.Color.BoxFg, rCol[i]).Sprintf(" %d%s ", rCnt[i], rUnit[i])
		} else {
			reaction += color.HEX(rCol[i]).Sprintf("%d%s", rCnt[i], rUnit[i])
		}
	}

	// 投稿時間
	postTime := ""
	createdAtTime, _ := tw.CreatedAtTime()
	if util.IsSameDate(createdAtTime) {
		postTime = createdAtTime.Local().Format(cfg.Default.TimeFormat)
	} else {
		format := fmt.Sprintf("%s %s", cfg.Default.DateFormat, cfg.Default.TimeFormat)
		postTime = createdAtTime.Local().Format(format)
	}
	postTime = color.HEX(cfg.Color.Accent2).Sprint(postTime)

	// 表示
	if !isQuote {
		fmt.Print("\n")
	}
	fmt.Printf("%s%s %s %s%s\n", status, index, user, postTime, reaction)
	color.HEX(cfg.Color.Text).Printf(" %s\n", text)

	// 引用リツイート
	if tw.QuotedStatus != nil {
		showTweet(-1, tw.QuotedStatus)
	}
}

// showUserInfo ユーザー情報を表示
func showUserInfo(user *anaconda.User) {
	width := util.GetWindowWidth()

	// ユーザー情報
	userStr := createUserString(user)
	// 関係性
	status := getFriendships(user)

	// 各種数値データ
	twCnt := color.HEX(cfg.Color.Accent1).Sprintf("%d tweets", user.StatusesCount)
	followCnt := color.HEX(cfg.Color.Accent2).Sprintf("%d following", user.FriendsCount)
	followerCnt := color.HEX(cfg.Color.Accent3).Sprintf("%d followers", user.FollowersCount)

	// 自己紹介
	desc := runewidth.Wrap(html.UnescapeString(user.Description), width-12)
	if desc == "" {
		desc = "none"
	} else {
		util.Replace(&desc, "\n", "\n           ")
	}
	// 場所
	location := user.Location
	if location == "" {
		location = "none"
	}
	// Webサイト
	website := user.URL
	if website == "" {
		website = "none"
	}
	// アカウント作成日
	tm, _ := time.Parse(time.RubyDate, user.CreatedAt)
	format := fmt.Sprintf("%s %s", cfg.Default.DateFormat, cfg.Default.TimeFormat)
	createdAt := tm.Local().Format(format)

	// 表示
	fmt.Printf("\n %s %s\n", userStr, status)
	fmt.Print(createSeparator(false))
	fmt.Printf("     data: %s  %s  %s\n", twCnt, followCnt, followerCnt)
	fmt.Printf("     desc: %s\n", desc)
	fmt.Printf(" location: %s\n", location)
	fmt.Printf("  website: %s\n", website)
	fmt.Printf("   joined: %s\n", createdAt)
}

// createUserString ユーザー情報の文字列を作成
func createUserString(user *anaconda.User) string {
	// ユーザー名
	name := color.HEX(cfg.Color.UserName).Sprint(user.Name)
	screenName := color.HEX(cfg.Color.UserID).Sprintf("(@%s)", user.ScreenName)

	// ユーザータイプ
	badge := ""
	if user.Verified {
		badge += color.HEX(cfg.Color.Verified).Sprint(" *verified*")
	}
	if user.Protected {
		badge += color.HEX(cfg.Color.Protected).Sprint(" *protected*")
	}

	return fmt.Sprintf("%s %s%s", name, screenName, badge)
}

// createSeparator セパレータを作成
func createSeparator(hasPutSpace bool) string {
	width := util.GetWindowWidth() - 2
	sep := "-"
	if hasPutSpace {
		width /= 2
		sep += " "
	}
	return color.HEX(cfg.Color.Separator).Sprintf(" %s\n", strings.Repeat(sep, width))
}

// showSuccessMsg 処理完了メッセージを表示
func showSuccessMsg(text, tips, hex string) {
	util.Replace(&text, "[\n\r]", "")
	text = html.UnescapeString(text)
	cutText := util.CutString(text, util.GetWindowWidth()-len(tips)-1)
	tips = color.HEX(hex).Sprint(tips)
	fmt.Printf("%s %s\n", tips, cutText)
}

// showAPIErrorString APIのエラーを表示
func showAPIErrorString(err error) {
	bytes := []byte(err.Error())
	errMsg := regexp.MustCompile("\"(message|error)\":\\s*\"(.+)\"").FindSubmatch(bytes)
	if len(errMsg) == 0 {
		color.Error.Tips(err.Error())
		return
	}
	color.Error.Tips("%s", errMsg[2])
}
