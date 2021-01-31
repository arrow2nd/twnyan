package view

import (
	"errors"
	"fmt"
	"html"
	"strconv"

	"github.com/ChimeraCoder/anaconda"
	"github.com/arrow2nd/twnyan/util"
	"github.com/gookit/color"
	"github.com/mattn/go-runewidth"
)

// DrawTweets ツイートを描画
func (v *View) DrawTweets() {
	for i := len(v.tweets) - 1; i >= 0; i-- {
		v.drawTweet(i, &v.tweets[i])
	}
}

func (v *View) drawTweet(i int, t *anaconda.Tweet) {
	header := ""
	width := util.GetWindowWidth()
	// RTなら元のツイートに置換
	if t.RetweetedStatus != nil {
		header += color.HEX(v.cfg.Color.Retweet).Sprintf("RT by %s @%s\n", t.User.Name, t.User.ScreenName)
		t = t.RetweetedStatus
	}
	// リプライなら宛先を追加
	if t.InReplyToScreenName != "" {
		header += color.HEX(v.cfg.Color.Reply).Sprintf("Reply to @%s\n", t.InReplyToScreenName)
	}
	// ヘッダー作成
	userInfo := v.createUserStr(&t.User)
	pt, _ := t.CreatedAtTime()
	postTime := v.createTimeStr(pt)
	fav := v.createReactionCountStr(t.FavoriteCount, t.Favorited, "Fav")
	rt := v.createReactionCountStr(t.RetweetCount, t.Retweeted, "RT")
	header += fmt.Sprintf("%s %s %s%s", userInfo, postTime, fav, rt)
	// テキスト作成
	text := runewidth.Wrap(html.UnescapeString(t.FullText), width)
	// 表示
	fmt.Printf("%s\n%s\n\n", header, text)
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
	// ツイートを取得
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
