package view

import (
	"fmt"
	"html"

	"github.com/ChimeraCoder/anaconda"
	"github.com/arrow2nd/twnyan/util"
	"github.com/gookit/color"
	"github.com/mattn/go-runewidth"
)

// DrawUser ユーザー情報描画
func (v *View) DrawUser(u *anaconda.User, c []string) {
	width := util.GetWindowWidth()

	// ユーザー情報
	userInfo := v.createUserStr(u)
	// 関係情報
	connection := v.createConnectionStr(c)
	// BIO
	bio := runewidth.Wrap(html.UnescapeString(u.Description), width-5)
	// 場所
	locate := html.UnescapeString(u.Location)
	// URL
	url := u.URL
	// count
	tweetsCount := color.HEX(v.cfg.Color.Accent1).Sprintf("%d Tweets", u.StatusesCount)
	followingCount := color.HEX(v.cfg.Color.Accent2).Sprintf("%d Following", u.FriendsCount)
	followersCount := color.HEX(v.cfg.Color.Accent3).Sprintf("%d Followers", u.FollowersCount)

	// 描画
	fmt.Printf("%s %s\n", userInfo, connection)
	fmt.Print(v.createSeparator())
	fmt.Printf("%s %s %s\n", tweetsCount, followingCount, followersCount)
	if bio != "" {
		util.AllReplace(&bio, "\n", "\n     ")
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

func (v *View) createUserStr(u *anaconda.User) string {
	// ユーザー名、スクリーンネーム
	name := v.truncateUserName(u.Name)
	name = color.HEX(v.cfg.Color.UserName).Sprint(name)
	screenName := color.HEX(v.cfg.Color.ScreenName).Sprintf("@%s", u.ScreenName)

	// バッジ
	badge := ""
	if u.Verified {
		badge += color.HEX(v.cfg.Color.Verified).Sprint(" verified")
	}
	if u.Protected {
		badge += color.HEX(v.cfg.Color.Protected).Sprint(" protected")
	}

	// 結合
	text := fmt.Sprintf("%s %s%s", name, screenName, badge)
	return text
}

func (v *View) truncateUserName(un string) string {
	width := util.GetWindowWidth()
	return runewidth.Truncate(un, width/2, "…")
}

func (v *View) createConnectionStr(c []string) string {
	connection := ""
	for _, str := range c {
		switch str {
		case "followed_by":
			connection += color.HEX(v.cfg.Color.FollowedBy).Sprint("Followed by ")
		case "following":
			connection += color.HEX(v.cfg.Color.Following).Sprint("Following ")
		case "blocking":
			connection += color.HEX(v.cfg.Color.Block).Sprint("Blocking ")
		case "muting":
			connection += color.HEX(v.cfg.Color.Mute).Sprint("Muting ")
		}
	}
	return connection
}
