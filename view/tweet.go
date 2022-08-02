package view

import (
	"fmt"
	"html"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/ChimeraCoder/anaconda"
	"github.com/arrow2nd/twnyan/util"
	"github.com/gookit/color"
)

// Unit 単位
type Unit string

const (
	// Fav お気に入り
	Fav Unit = "Fav"
	// RT リツイート
	RT Unit = "RT"
)

// ShowTweets ツイートを一覧表示
func (v *View) ShowTweets(tweets []anaconda.Tweet, shouldShowTweetNum bool) {
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

	// QTならセパレータを挿入
	if isQuote {
		header += v.createSeparatorString(true)
	}

	// RTなら元のツイートに置換
	if tweets.RetweetedStatus != nil {
		width := util.GetWindowWidth() / 2

		header += color.HEX(v.config.Color.Retweet).Sprintf(
			"RT by %s @%s\n",
			util.TruncateString(tweets.User.Name, width),
			tweets.User.ScreenName,
		)

		tweets = tweets.RetweetedStatus
	}

	// リプライなら宛先を追加
	if tweets.InReplyToScreenName != "" {
		header += color.HEX(v.config.Color.Reply).Sprintf(
			"Reply to @%s\n",
			tweets.InReplyToScreenName,
		)
	}

	// ヘッダを作成
	tagStr := color.HEXStyle(v.config.Color.BoxForground, v.config.Color.Accent1).Sprint(tagText)
	userInfoStr := v.createUserInfoString(&tweets.User)
	createdAt, _ := tweets.CreatedAtTime()
	createdAtStr := v.createCreatedAtString(createdAt)
	favCountStr := v.createCountString(tweets.FavoriteCount, tweets.Favorited, "Fav")
	rtCountStr := v.createCountString(tweets.RetweetCount, tweets.Retweeted, "RT")

	header += fmt.Sprintf("%s %s %s%s%s", tagStr, userInfoStr, createdAtStr, favCountStr, rtCountStr)

	fmt.Printf("%s\n%s", header, v.processTweetText(tweets))

	// QTなら引用元ツイートを表示
	if tweets.QuotedStatus != nil {
		v.ShowTweet(tweets.QuotedStatus, " ", true)
		return
	}

	fmt.Print("\n")
}

// processTweetText ツイート文を加工
func (v *View) processTweetText(tweet *anaconda.Tweet) string {
	// 文字をアンエスケープ
	text := html.UnescapeString(tweet.FullText) + "\n"

	// 全角記号を置換
	text = strings.ReplaceAll(text, "＃", "#")
	text = strings.ReplaceAll(text, "＠", "@")

	// ハッシュタグをハイライト
	if len(tweet.Entities.Hashtags) != 0 {
		text = v.highlightHashTags(text, &tweet.Entities)
	}

	// メンションをハイライト
	if len(tweet.Entities.User_mentions) != 0 {
		rep := regexp.MustCompile(`(^|[^\w@#$%&])@(\w+)`)
		text = rep.ReplaceAllString(text, "$1"+color.HEX(v.config.Color.Reply).Sprintf("@$2"))
	}

	return text
}

// highlightHashTags ハッシュタグをハイライト
func (v *View) highlightHashTags(text string, entities *anaconda.Entities) string {
	result := ""
	runes := []rune(text)
	end := 0

	for _, hashtag := range entities.Hashtags {
		hashtagText := fmt.Sprintf("#%s", hashtag.Text)

		// NOTE: URLや絵文字を多く含むツイートなどで、ハッシュタグの開始位置が後方にズレていることがあるので
		//       +1 して意図的にズラした後、ハッシュタグ全文が見つかるまで開始位置を前方に移動することで正しい位置を見つける

		begin := hashtag.Indices[0] + 1
		textLength := utf8.RuneCountInString(hashtag.Text) + 1

		for ; begin > end; begin-- {
			e := begin + textLength

			if l := len(runes); e > l {
				continue
			}

			if string(runes[begin:e]) == hashtagText {
				break
			}
		}

		// 前方の文とハイライトされたハッシュタグを結合
		result += string(runes[end:begin]) + color.HEX(v.config.Color.Hashtag).Sprint(hashtagText)

		// ハッシュタグの終了位置
		end = begin + utf8.RuneCountInString(hashtagText)
	}

	// 残りの文を結合
	if len(runes) > end {
		result += string(runes[end:])
	}

	return result
}

// createCountString いいね・RT数の文字列を作成
func (v *View) createCountString(countNum int, isReacted bool, unit Unit) string {
	// 表示色
	colorCode := v.config.Color.Favorite
	if unit == RT {
		colorCode = v.config.Color.Retweet
	}

	if countNum <= 0 {
		return ""
	} else if countNum > 1 {
		// カウントが1以上なら複数形にする
		unit += "s"
	}

	result := " "

	// リアクション済みなら文字色と背景色を反転させる
	if isReacted {
		result += color.HEXStyle(v.config.Color.BoxForground, colorCode).Sprintf(" %d%s ", countNum, unit)
	} else {
		result += color.HEX(colorCode).Sprintf("%d%s", countNum, unit)
	}

	return result
}
