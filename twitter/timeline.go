package twitter

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"

	"github.com/ChimeraCoder/anaconda"
	"github.com/arrow2nd/twnyan/util"
	"github.com/gookit/color"
)

// Tweets ツイート
type Tweets struct {
	tweets []anaconda.Tweet
}

// DrawTweets ツイートを表示
func (t *Tweets) DrawTweets() {
	if len(t.tweets) <= 0 {
		color.Info.Tips("Couldn't find a tweet")
		return
	}
	for i := len(t.tweets) - 1; i >= 0; i-- {
		showTweet(i, &t.tweets[i])
	}
	fmt.Print("\n")
}

// DrawUserInfo ユーザー情報を表示
func (t *Tweets) DrawUserInfo(idx int) {
	if len(t.tweets) <= 0 || len(t.tweets) <= idx {
		return
	}
	showUserInfo(&t.tweets[idx].User)
	fmt.Print("\n")
}

// GetDataFromTweetNum ツイート番号から指定したデータを取得
func (t *Tweets) GetDataFromTweetNum(numStr, dataType string) (string, error) {
	if !util.IsNumber(numStr) {
		return "", fmt.Errorf("'%s' is not a number", numStr)
	}

	// 数値に変換
	num, _ := strconv.Atoi(numStr)
	if num < 0 || num > len(t.tweets)-1 {
		return "", errors.New("Tweet number is out of range")
	}
	tw := t.tweets[num]

	// リツイート元のツイートに置換
	if tw.RetweetedStatus != nil {
		tw = *tw.RetweetedStatus
	}

	switch dataType {
	case "ScreenName":
		return tw.User.ScreenName, nil
	case "TweetID":
		return tw.IdStr, nil
	default:
		return "", errors.New("Wrong dataType")
	}
}

// GetTweetURL ツイートのURLを取得
func (t *Tweets) GetTweetURL(numStr string) (string, error) {
	name, err := t.GetDataFromTweetNum(numStr, "ScreenName")
	if err != nil {
		return "", err
	}
	ID, _ := t.GetDataFromTweetNum(numStr, "TweetID")
	return fmt.Sprintf("https://twitter.com/%s/status/%s", name, ID), nil
}

// LoadHomeTL ホームタイムラインを読み込む
func (t *Tweets) LoadHomeTL(count string) error {
	v := createURLValues(count)
	tl := getTimeline("home", v)
	if tl == nil {
		return errors.New("error")
	}
	t.add(tl)
	return nil
}

// LoadMentionTL 自分宛てのメンションを読み込む
func (t *Tweets) LoadMentionTL(count string) error {
	v := createURLValues(count)
	tl := getTimeline("mention", v)
	if tl == nil {
		return errors.New("error")
	}
	t.add(tl)
	return nil
}

// LoadUserTL ユーザータイムラインを読み込む
func (t *Tweets) LoadUserTL(user, count string) error {
	v := createURLValues(count)
	v.Add("exclude_replies", "true")
	if user != "" {
		v.Add("screen_name", user)
	}

	tl := getTimeline("user", v)
	if tl == nil {
		return errors.New("error")
	}

	t.add(tl)
	return nil
}

// LoadListTL リストのタイムラインを読み込む
func (t *Tweets) LoadListTL(listID int64, count string) error {
	v := createURLValues(count)

	tl, err := api.GetListTweets(listID, true, v)
	if err != nil {
		showAPIErrorString(err)
		return errors.New("error")
	}

	t.add(tl)
	return nil
}

// LoadSearchResult 検索結果を読み込む
func (t *Tweets) LoadSearchResult(query, count string) error {
	v := createURLValues(count)
	query += " -filter:retweets"

	res, err := api.GetSearch(query, v)
	if err != nil {
		showAPIErrorString(err)
		return errors.New("error")
	}

	t.add(res.Statuses)
	return nil
}

func (t *Tweets) add(tl []anaconda.Tweet) {
	t.tweets = make([]anaconda.Tweet, len(tl))
	copy(t.tweets, tl)
}

func createURLValues(cnt string) url.Values {
	v := url.Values{}
	v.Add("tweet_mode", "extended")
	v.Add("count", cnt)
	return v
}
