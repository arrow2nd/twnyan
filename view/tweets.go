package view

import (
	"errors"
	"fmt"
	"html"
	"regexp"
	"strconv"

	"github.com/ChimeraCoder/anaconda"
	"github.com/arrow2nd/twnyan/util"
	"github.com/gookit/color"
	"github.com/mattn/go-runewidth"
)

// DrawTweets ツイート描画
func (v *View) DrawTweets() {
	for i := len(v.tweets) - 1; i >= 0; i-- {
		v.drawTweet(i, false, &v.tweets[i])
	}
}

func (v *View) drawTweet(i int, isQuote bool, t *anaconda.Tweet) {
	header := ""

	// QTならセパレータを挿入
	if isQuote {
		header += v.createSeparatorStr(true)
	}
	// RTなら元のツイートに置換
	if t.RetweetedStatus != nil {
		header += color.HEX(v.cfg.Color.Retweet).Sprintf("RT by %s @%s\n", v.truncateUserName(t.User.Name), t.User.ScreenName)
		t = t.RetweetedStatus
	}
	// リプライなら宛先を追加
	if t.InReplyToScreenName != "" {
		header += color.HEX(v.cfg.Color.Reply).Sprintf("Reply to @%s\n", t.InReplyToScreenName)
	}

	// ツイート番号
	index := "↪"
	if !isQuote {
		index = color.HEXStyle(v.cfg.Color.BoxForground, v.cfg.Color.Accent1).Sprintf(" %d ", i)
	}
	// ユーザー情報
	userInfo := v.createUserStr(&t.User)
	// 投稿時刻
	pt, _ := t.CreatedAtTime()
	postTime := v.createPostTimeStr(pt)
	// いいね、RT数
	fav := v.createReactionCountStr(t.FavoriteCount, t.Favorited, "Fav")
	rt := v.createReactionCountStr(t.RetweetCount, t.Retweeted, "RT")
	// ヘッダー
	header += fmt.Sprintf("%s %s %s%s%s", index, userInfo, postTime, fav, rt)
	// ツイート内容
	text := v.createTweetText(t)

	// 表示
	v.shell.Printf("%s\n%s", header, text)
	// QTなら引用元を表示
	if t.QuotedStatus != nil {
		v.drawTweet(0, true, t.QuotedStatus)
		return
	}
	v.shell.Print("\n")
}

func (v *View) createTweetText(t *anaconda.Tweet) string {
	width := util.GetWindowWidth()

	// 画面幅でワープ
	text := runewidth.Wrap(html.UnescapeString(t.FullText), width)
	text += "\n"

	// ハッシュタグをハイライト
	if len(t.Entities.Hashtags) != 0 {
		for _, h := range t.Entities.Hashtags {
			rep := regexp.MustCompile(fmt.Sprintf("[#＃](%s)([\\s])", h.Text))
			text = rep.ReplaceAllString(text, color.HEX(v.cfg.Color.Hashtag).Sprintf("#$1$2"))
		}
	}
	// メンションをハイライト
	if len(t.Entities.User_mentions) != 0 {
		rep := regexp.MustCompile("(^|[^\\w@#$%&])[@＠](\\w+)")
		text = rep.ReplaceAllString(text, "$1"+color.HEX(v.cfg.Color.Reply).Sprintf("@$2"))
	}

	return text
}

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

// RegisterTweets ツイートを登録
func (v *View) RegisterTweets(t *[]anaconda.Tweet) {
	tmp := []anaconda.Tweet{}
	v.tweets = append(tmp, *t...)
}

// GetTweetURL ツイートのURLを取得
func (v *View) GetTweetURL(numStr string) (string, error) {
	name, err := v.GetDataFromTweetNum(numStr, "screenname")
	if err != nil {
		return "", err
	}
	ID, _ := v.GetDataFromTweetNum(numStr, "tweetID")
	return fmt.Sprintf("https://twitter.com/%s/status/%s", name, ID), nil
}

// GetDataFromTweetNum ツイート番号から情報を取得
func (v *View) GetDataFromTweetNum(numStr, dataType string) (string, error) {
	// ツイート番号かチェック
	if !util.IsNumber(numStr) {
		return "", fmt.Errorf("tweetnumber is invalid")
	}
	// ツイート番号が範囲内かチェック
	num, _ := strconv.Atoi(numStr)
	if num < 0 || num > len(v.tweets)-1 {
		return "", errors.New("tweetnumber is out of range")
	}

	// ツイート取得
	tweet := v.tweets[num]
	if tweet.RetweetedStatus != nil {
		tweet = *tweet.RetweetedStatus
	}

	// 指定されたデータを返す
	switch dataType {
	case "screenname":
		return tweet.User.ScreenName, nil
	case "tweetID":
		return tweet.IdStr, nil
	default:
		return "", errors.New("Wrong dataType")
	}
}
