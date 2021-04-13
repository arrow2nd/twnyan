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

	// ユーザー情報
	userInfo := v.createUserInfoString(user)
	relationshipInfo := v.createRelationshipInfoString(relationships)
	bio := runewidth.Wrap(html.UnescapeString(user.Description), width-5)
	locate := html.UnescapeString(user.Location)
	url := user.URL

	// 各種カウント
	tweetsCount := color.HEX(v.cfg.Color.Accent1).Sprintf("%d Tweets", user.StatusesCount)
	followingCount := color.HEX(v.cfg.Color.Accent2).Sprintf("%d Following", user.FriendsCount)
	followersCount := color.HEX(v.cfg.Color.Accent3).Sprintf("%d Followers", user.FollowersCount)

	fmt.Printf("%s %s\n", userInfo, relationshipInfo)
	fmt.Print(v.createSeparatorString(false))
	fmt.Printf("%s %s %s\n", tweetsCount, followingCount, followersCount)

	if bio != "" {
		bio = util.AllReplace(bio, "\n", "\n     ")
		fmt.Printf("📄 : %s\n", bio)
	}
	if locate != "" {
		fmt.Printf("📍 : %s\n", locate)
	}
	if url != "" {
		fmt.Printf("🔗 : %s\n", url)
	}

	fmt.Print("\n")
}

// createUserInfoString ユーザー情報の文字列を作成
func (v *View) createUserInfoString(u *anaconda.User) string {
	halfWidth := util.GetWindowWidth() / 3

	// ユーザー名、スクリーンネーム
	userName := util.TruncateString(u.Name, halfWidth)
	userName = color.HEX(v.cfg.Color.UserName).Sprint(userName)
	screenName := color.HEX(v.cfg.Color.ScreenName).Sprintf("@%s", u.ScreenName)

	// アカウントタイプ
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
			relationshipInfo += color.HEX(v.cfg.Color.FollowedBy).Sprint("Followed by ")
		case "following":
			relationshipInfo += color.HEX(v.cfg.Color.Following).Sprint("Following ")
		case "blocking":
			relationshipInfo += color.HEX(v.cfg.Color.Block).Sprint("Blocking ")
		case "muting":
			relationshipInfo += color.HEX(v.cfg.Color.Mute).Sprint("Muting ")
		}
	}

	return relationshipInfo
}
