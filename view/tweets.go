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
	tagStr := ""

	for i := len(tweets) - 1; i >= 0; i-- {
		if shouldShowTweetNum {
			tagStr = fmt.Sprint(i)
		}
		v.showTweet(&tweets[i], tagStr, false)
	}
}

// showTweet ツイートを表示
func (v *View) showTweet(tweets *anaconda.Tweet, tagStr string, isQuote bool) {
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
	tagStr = color.HEXStyle(v.cfg.Color.BoxForground, v.cfg.Color.Accent1).Sprintf(" %s ", tagStr)
	userInfoStr := v.createUserInfoString(&tweets.User)
	postTime, _ := tweets.CreatedAtTime()
	postTimeStr := v.createPostTimeString(postTime)
	favCountStr := v.createCountString(tweets.FavoriteCount, tweets.Favorited, "Fav")
	rtCountStr := v.createCountString(tweets.RetweetCount, tweets.Retweeted, "RT")
	header += fmt.Sprintf("%s %s %s%s%s", tagStr, userInfoStr, postTimeStr, favCountStr, rtCountStr)

	fmt.Printf("%s\n%s", header, v.editTweetText(tweets))

	// QTなら引用元ツイートを表示
	if tweets.QuotedStatus != nil {
		v.showTweet(tweets.QuotedStatus, "↪", true)
		return
	}

	fmt.Print("\n")
}

// editTweetText ツイート文を編集
func (v *View) editTweetText(tweet *anaconda.Tweet) string {
	// 文字をアンエスケープ
	text := html.UnescapeString(tweet.FullText)
	text += "\n"

	// ハッシュタグをハイライト
	if len(tweet.Entities.Hashtags) != 0 {
		for _, h := range tweet.Entities.Hashtags {
			rep := regexp.MustCompile(fmt.Sprintf(`[#＃](%s)([\s　])`, h.Text))
			text = rep.ReplaceAllString(text, color.HEX(v.cfg.Color.Hashtag).Sprintf("#$1$2"))
		}
	}

	// メンションをハイライト
	if len(tweet.Entities.User_mentions) != 0 {
		rep := regexp.MustCompile(`(^|[^\w@#$%&])[@＠](\w+)`)
		text = rep.ReplaceAllString(text, "$1"+color.HEX(v.cfg.Color.Reply).Sprintf("@$2"))
	}

	return text
}

// createCountString いいね・RT数の文字列を作成
func (v *View) createCountString(countNum int, reverseFlg bool, unitStr string) string {
	// 表示色
	colorCode := v.cfg.Color.Favorite
	if unitStr == "RT" {
		colorCode = v.cfg.Color.Retweet
	}

	// カウントが0ならreturn
	if countNum <= 0 {
		return ""
	} else if countNum > 1 {
		unitStr += "s"
	}

	// 文字列作成
	text := " "
	if reverseFlg {
		text += color.HEXStyle(v.cfg.Color.BoxForground, colorCode).Sprintf(" %d%s ", countNum, unitStr)
	} else {
		text += color.HEX(colorCode).Sprintf("%d%s", countNum, unitStr)
	}

	return text
}

// RegisterTweets ツイートを登録
func (v *View) RegisterTweets(t *[]anaconda.Tweet) {
	v.mu.Lock()
	defer v.mu.Unlock()

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
func (v *View) GetDataFromTweetNum(tweetNumStr, dataType string) (string, error) {
	// 数値ではないならエラー
	if !util.IsNumber(tweetNumStr) {
		return "", fmt.Errorf("tweetnumber is invalid")
	}

	tweetNum, _ := strconv.Atoi(tweetNumStr)

	// ツイート番号が範囲外ならエラー
	if tweetNum < 0 || tweetNum > len(v.tweets)-1 {
		return "", errors.New("tweetnumber is out of range")
	}

	// ツイート取得
	tweet := v.tweets[tweetNum]
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
		return "", errors.New("wrong datatype")
	}
}
