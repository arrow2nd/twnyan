package view

import (
	"fmt"
	"html"

	"github.com/ChimeraCoder/anaconda"
	"github.com/arrow2nd/twnyan/util"
	"github.com/gookit/color"
	"github.com/mattn/go-runewidth"
)

// ShowUserInfo ユーザー情報を表示
func (v *View) ShowUserInfo(user *anaconda.User, relationships []string) {
	width := util.GetWindowWidth()

	// 基本情報
	userInfo := v.createUserInfoString(user)
	relationshipInfo := v.createRelationshipInfoString(relationships)

	fmt.Printf("%s %s\n", userInfo, relationshipInfo)
	fmt.Print(v.createSeparatorString(false))

	// 各種カウント
	tweetsCount := color.HEX(v.config.Color.Accent1).Sprintf("%d Tweets", user.StatusesCount)
	followingCount := color.HEX(v.config.Color.Accent2).Sprintf("%d Following", user.FriendsCount)
	followersCount := color.HEX(v.config.Color.Accent3).Sprintf("%d Followers", user.FollowersCount)

	fmt.Printf("%s %s %s\n", tweetsCount, followingCount, followersCount)

	// BIO
	bio := runewidth.Wrap(html.UnescapeString(user.Description), width-5)
	if bio != "" {
		fmt.Printf("📄 : %s\n", util.AllReplace(bio, "\n", "\n     "))
	}

	// 場所
	locate := html.UnescapeString(user.Location)
	if locate != "" {
		fmt.Printf("📍 : %s\n", locate)
	}

	// Webサイト
	url := user.URL
	if url != "" {
		fmt.Printf("🔗 : %s\n", url)
	}

	fmt.Print("\n")
}

// createUserInfoString ユーザ基本情報の文字列を作成
func (v *View) createUserInfoString(u *anaconda.User) string {
	screenName := color.HEX(v.config.Color.ScreenName).Sprintf("@%s", u.ScreenName)

	// NOTE: ユーザ名が長すぎると右側のスペースがなくなるので 1/3 に制限
	width := util.GetWindowWidth() / 3
	userName := color.HEX(v.config.Color.UserName).Sprint(
		util.TruncateString(u.Name, width),
	)

	// アカウントの種類
	accountType := ""
	if u.Verified {
		accountType += " ✅"
	}
	if u.Protected {
		accountType += " 🔒"
	}

	return fmt.Sprintf("%s %s%s", userName, screenName, accountType)
}

// createRelationshipInfoString ユーザーとの関係性を表す文字列を作成
func (v *View) createRelationshipInfoString(relationships []string) string {
	relationshipInfo := ""

	for _, str := range relationships {
		switch str {
		case "followed_by":
			relationshipInfo += color.HEX(v.config.Color.FollowedBy).Sprint("Followed by ")
			continue
		case "following":
			relationshipInfo += color.HEX(v.config.Color.Following).Sprint("Following ")
			continue
		case "blocking":
			relationshipInfo += color.HEX(v.config.Color.Block).Sprint("Blocking ")
			continue
		case "muting":
			relationshipInfo += color.HEX(v.config.Color.Mute).Sprint("Muting ")
		}
	}

	return relationshipInfo
}
