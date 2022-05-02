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
)

// DataType 取得するデータの種類
type DataType uint8

const (
	ScreenName DataType = iota
	TweetId
)

// Unit 単位
type Unit string

const (
	Fav Unit = "Fav"
	RT  Unit = "RT"
)

// ShowRegisteredTweets 登録済みのツイートを表示
func (v *View) ShowRegisteredTweets() {
	v.ShowTweetsFromArray(v.tweets, true)
}

// ShowTweetsFromArray 配列からツイートを表示
func (v *View) ShowTweetsFromArray(tweets []anaconda.Tweet, shouldShowTweetNum bool) {
	// ツイートがない
	if len(tweets) == 0 {
		fmt.Print("\n-- No tweets --\n\n")
		return
	}

	tagText := ""

	// 逆順 (上から下に向けて最新) で出力
	for i := len(tweets) - 1; i >= 0; i-- {
		if shouldShowTweetNum {
			tagText = fmt.Sprintf(" %d ", i)
		} else {
			tagText = " "
		}

		v.ShowTweet(&tweets[i], tagText, false)
	}
}

// ShowTweet ツイートを表示
func (v *View) ShowTweet(tweets *anaconda.Tweet, tagText string, isQuote bool) {
	header := ""
	halfWidth := util.GetWindowWidth() / 2

	// QTならセパレータを挿入
	if isQuote {
		header += v.createSeparatorString(true)
	}

	// RTなら元のツイートに置換
	if tweets.RetweetedStatus != nil {
		header += color.HEX(v.cfg.Color.Retweet).Sprintf(
			"RT by %s @%s\n",
			util.TruncateString(tweets.User.Name, halfWidth),
			tweets.User.ScreenName,
		)

		tweets = tweets.RetweetedStatus
	}

	// リプライなら宛先を追加
	if tweets.InReplyToScreenName != "" {
		header += color.HEX(v.cfg.Color.Reply).Sprintf(
			"Reply to @%s\n",
			tweets.InReplyToScreenName,
		)
	}

	// ヘッダー文字列を作成
	tagStr := color.HEXStyle(v.cfg.Color.BoxForground, v.cfg.Color.Accent1).Sprint(tagText)
	userInfoStr := v.createUserInfoString(&tweets.User)
	createdAt, _ := tweets.CreatedAtTime()
	createdAtStr := v.createCreatedAtString(createdAt)
	favCountStr := v.createCountString(tweets.FavoriteCount, tweets.Favorited, "Fav")
	rtCountStr := v.createCountString(tweets.RetweetCount, tweets.Retweeted, "RT")

	header += fmt.Sprintf("%s %s %s%s%s", tagStr, userInfoStr, createdAtStr, favCountStr, rtCountStr)

	fmt.Printf("%s\n%s", header, v.editTweetText(tweets))

	// QTなら引用元ツイートを表示
	if tweets.QuotedStatus != nil {
		v.ShowTweet(tweets.QuotedStatus, " ", true)
		return
	}

	fmt.Print("\n")
}

// editTweetText ツイート文を編集
func (v *View) editTweetText(tweet *anaconda.Tweet) string {
	// 文字をアンエスケープ
	tweetText := html.UnescapeString(tweet.FullText) + "\n"

	// ハッシュタグをハイライト
	if len(tweet.Entities.Hashtags) != 0 {
		rep := regexp.MustCompile(`[#＃](\S+)`)
		tweetText = rep.ReplaceAllString(tweetText, color.HEX(v.cfg.Color.Hashtag).Sprintf("#$1"))
	}

	// メンションをハイライト
	if len(tweet.Entities.User_mentions) != 0 {
		rep := regexp.MustCompile(`(^|[^\w@#$%&])[@＠](\w+)`)
		tweetText = rep.ReplaceAllString(tweetText, "$1"+color.HEX(v.cfg.Color.Reply).Sprintf("@$2"))
	}

	return tweetText
}

// createCountString いいね・RT数の文字列を作成
func (v *View) createCountString(countNum int, isReacted bool, unit Unit) string {
	// 表示色
	colorCode := v.cfg.Color.Favorite
	if unit == RT {
		colorCode = v.cfg.Color.Retweet
	}

	if countNum <= 0 {
		return ""
	} else if countNum > 1 {
		// カウントが1以上なら複数形にする
		unit += "s"
	}

	countStr := " "

	// リアクション済みなら文字色と背景色を反転させる
	if isReacted {
		countStr += color.HEXStyle(v.cfg.Color.BoxForground, colorCode).Sprintf(" %d%s ", countNum, unit)
	} else {
		countStr += color.HEX(colorCode).Sprintf("%d%s", countNum, unit)
	}

	return countStr
}

// RegisterTweets ツイートを登録
func (v *View) RegisterTweets(tweets *[]anaconda.Tweet) {
	v.mu.Lock()
	defer v.mu.Unlock()

	v.tweets = append([]anaconda.Tweet{}, *tweets...)
}

// GetTweetURL ツイートのURLを取得
func (v *View) GetTweetURL(tweetNumStr string) (string, error) {
	screenName, err := v.GetDataFromTweetNum(tweetNumStr, ScreenName)
	if err != nil {
		return "", err
	}

	tweetId, err := v.GetDataFromTweetNum(tweetNumStr, TweetId)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("https://twitter.com/%s/status/%s", screenName, tweetId), nil
}

// GetDataFromTweetNum ツイート番号から情報を取得
func (v *View) GetDataFromTweetNum(tweetNumStr string, dataType DataType) (string, error) {
	if !util.IsThreeDigitsNumber(tweetNumStr) {
		return "", errors.New("tweet-number is invalid")
	}

	tweetNum, _ := strconv.Atoi(tweetNumStr)

	// ツイート番号が範囲外ならエラー
	if tweetNum < 0 || tweetNum > len(v.tweets)-1 {
		return "", errors.New("tweet-number is out of range")
	}

	// ツイートを取得
	tweet := v.tweets[tweetNum]
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
