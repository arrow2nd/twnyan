package twitter

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/ChimeraCoder/anaconda"
	"github.com/arrow2nd/twnyan/util"
)

// DataType 取得するデータの種類
type DataType uint8

const (
	// ScreenName スクリーンネーム
	ScreenName DataType = iota
	// TweetId ツイートID
	TweetId
)

// RegisterTweets ツイートを登録
func (t *Twitter) RegisterTweets(tweets *[]anaconda.Tweet) {
	t.mu.Lock()
	defer t.mu.Unlock()

	// appendで値をコピー
	t.Tweets = append([]anaconda.Tweet{}, *tweets...)
}

// GetDataFromTweetNum ツイート番号から情報を取得
func (t *Twitter) GetDataFromTweetNum(tweetNumStr string, dataType DataType) (string, error) {
	if !util.IsThreeDigitsNumber(tweetNumStr) {
		return "", errors.New("tweet-number is invalid")
	}

	tweetNum, _ := strconv.Atoi(tweetNumStr)

	// ツイート番号が範囲外ならエラー
	if tweetNum < 0 || tweetNum > len(t.Tweets)-1 {
		return "", errors.New("tweet-number is out of range")
	}

	tweet := t.Tweets[tweetNum]

	// リツイートなら元ツイートに置き換え
	if tweet.RetweetedStatus != nil {
		tweet = *tweet.RetweetedStatus
	}

	// 指定されたデータを返す
	switch dataType {
	case ScreenName:
		return tweet.User.ScreenName, nil
	case TweetId:
		return tweet.IdStr, nil
	}

	return "", errors.New("wrong datatype")
}

// GetTweetURL ツイートのURLを取得
func (t *Twitter) GetTweetURL(tweetNumStr string) (string, error) {
	screenName, err := t.GetDataFromTweetNum(tweetNumStr, ScreenName)
	if err != nil {
		return "", err
	}

	tweetId, err := t.GetDataFromTweetNum(tweetNumStr, TweetId)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("https://twitter.com/%s/status/%s", screenName, tweetId), nil
}
