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

// ShowRegisteredTweets 登録済みのツイートを表示
func (v *View) ShowRegisteredTweets() {
	v.ShowTweetsFromArray(v.tweets, true)
}

// ShowTweetsFromArray 配列からツイートを表示
func (v *View) ShowTweetsFromArray(tweets []anaconda.Tweet, shouldShowTweetNum bool) {
	tagText := " "

	for i := len(tweets) - 1; i >= 0; i-- {
		if shouldShowTweetNum {
			tagText = fmt.Sprintf(" %d ", i)
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
		header += color.HEX(v.cfg.Color.Retweet).Sprintf("RT by %s @%s\n", util.TruncateString(tweets.User.Name, halfWidth), tweets.User.ScreenName)
		tweets = tweets.RetweetedStatus
	}
	// リプライなら宛先を追加
	if tweets.InReplyToScreenName != "" {
		header += color.HEX(v.cfg.Color.Reply).Sprintf("Reply to @%s\n", tweets.InReplyToScreenName)
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
	tweetText := html.UnescapeString(tweet.FullText)
	tweetText += "\n"

	// ハッシュタグをハイライト
	if len(tweet.Entities.Hashtags) != 0 {
		// 半角・全角英数字、漢字、平仮名、全角・半角片仮名、伸ばし棒、アンダースコア
		rep := regexp.MustCompile(`[#＃]([A-Za-z0-9Ａ-Ｚａ-ｚ０-９\x{4E00}-\x{9FFF}\x{3005}-\x{3007}ぁ-ヶｦ-ﾟー～_]+)`)
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
func (v *View) createCountString(countNum int, reverseFlg bool, unitStr string) string {
	colorCode := v.cfg.Color.Favorite
	if unitStr == "RT" {
		colorCode = v.cfg.Color.Retweet
	}

	if countNum <= 0 {
		return ""
	} else if countNum > 1 {
		unitStr += "s"
	}

	countStr := " "

	// フラグが立っていれば文字を反転する
	if reverseFlg {
		countStr += color.HEXStyle(v.cfg.Color.BoxForground, colorCode).Sprintf(" %d%s ", countNum, unitStr)
	} else {
		countStr += color.HEX(colorCode).Sprintf("%d%s", countNum, unitStr)
	}

	return countStr
}

// RegisterTweets ツイートを登録
func (v *View) RegisterTweets(tweets *[]anaconda.Tweet) {
	v.mu.Lock()
	defer v.mu.Unlock()

	tmp := []anaconda.Tweet{}
	v.tweets = append(tmp, *tweets...)
}

// GetTweetURL ツイートのURLを取得
func (v *View) GetTweetURL(tweetNumStr string) (string, error) {
	screenName, err := v.GetDataFromTweetNum(tweetNumStr, "screenName")
	if err != nil {
		return "", err
	}

	tweetID, err := v.GetDataFromTweetNum(tweetNumStr, "tweetID")
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("https://twitter.com/%s/status/%s", screenName, tweetID), nil
}

// GetDataFromTweetNum ツイート番号から情報を取得
func (v *View) GetDataFromTweetNum(tweetNumStr, dataType string) (string, error) {
	if !util.IsThreeDigitsNumber(tweetNumStr) {
		return "", errors.New("tweetnumber is invalid")
	}

	tweetNum, _ := strconv.Atoi(tweetNumStr)

	// ツイート番号が範囲外ならエラー
	if tweetNum < 0 || tweetNum > len(v.tweets)-1 {
		return "", errors.New("tweetnumber is out of range")
	}

	// ツイートを取得
	tweet := v.tweets[tweetNum]
	if tweet.RetweetedStatus != nil {
		tweet = *tweet.RetweetedStatus
	}

	// 指定されたデータを返す
	switch dataType {
	case "screenName":
		return tweet.User.ScreenName, nil
	case "tweetID":
		return tweet.IdStr, nil
	}

	return "", errors.New("wrong datatype")
}
