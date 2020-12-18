package twitter

import (
	"fmt"
	"strconv"

	"github.com/arrow2nd/twnyan/util"
)

// Favorite いいねする
func Favorite(tweetID string) {
	id, _ := strconv.ParseInt(tweetID, 10, 64)
	tweet, err := api.Favorite(id)
	if err != nil {
		showAPIErrorString(err)
		return
	}
	util.ShowSuccessMsg("Favorited", tweet.FullText, cfg.Color.BoxFg, cfg.Color.Fav)
}

// Unfavorite いいねを取り消す
func Unfavorite(tweetID string) {
	id, _ := strconv.ParseInt(tweetID, 10, 64)
	tweet, err := api.Unfavorite(id)
	if err != nil {
		showAPIErrorString(err)
		return
	}
	util.ShowSuccessMsg("Unfavorited", tweet.FullText, cfg.Color.BoxFg, cfg.Color.Fav)
}

// Retweet リツイートする
func Retweet(tweetID string) {
	id, _ := strconv.ParseInt(tweetID, 10, 64)
	tweet, err := api.Retweet(id, true)
	if err != nil {
		showAPIErrorString(err)
		return
	}
	util.ShowSuccessMsg("Retweeted", tweet.FullText, cfg.Color.BoxFg, cfg.Color.RT)
}

// UnRetweet リツイートを取り消す
func UnRetweet(tweetID string) {
	id, _ := strconv.ParseInt(tweetID, 10, 64)
	tweet, err := api.UnRetweet(id, true)
	if err != nil {
		showAPIErrorString(err)
		return
	}
	util.ShowSuccessMsg("UnRetweeted", tweet.FullText, cfg.Color.BoxFg, cfg.Color.RT)
}

// Follow フォローする
func Follow(name string) {
	user, err := api.FollowUser(name)
	if err != nil {
		showAPIErrorString(err)
		return
	}
	text := fmt.Sprintf("%s (@%s)", user.Name, user.ScreenName)
	util.ShowSuccessMsg("Followed", text, cfg.Color.BoxFg, cfg.Color.Follow)
}

// Unfollow フォローを解除する
func Unfollow(name string) {
	user, err := api.UnfollowUser(name)
	if err != nil {
		showAPIErrorString(err)
		return
	}
	text := fmt.Sprintf("%s (@%s)", user.Name, user.ScreenName)
	util.ShowSuccessMsg("Unfollowed", text, cfg.Color.BoxFg, cfg.Color.Follow)
}

// Block ブロックする
func Block(name string) {
	user, err := api.BlockUser(name, nil)
	if err != nil {
		showAPIErrorString(err)
		return
	}
	text := fmt.Sprintf("%s (@%s)", user.Name, user.ScreenName)
	util.ShowSuccessMsg("Blocked", text, cfg.Color.BoxFg, cfg.Color.Block)
}

// Unblock ブロックを解除する
func Unblock(name string) {
	user, err := api.UnblockUser(name, nil)
	if err != nil {
		showAPIErrorString(err)
		return
	}
	text := fmt.Sprintf("%s (@%s)", user.Name, user.ScreenName)
	util.ShowSuccessMsg("Unblocked", text, cfg.Color.BoxFg, cfg.Color.Block)
}

// Mute ミュートする
func Mute(name string) {
	user, err := api.MuteUser(name, nil)
	if err != nil {
		showAPIErrorString(err)
		return
	}
	text := fmt.Sprintf("%s (@%s)", user.Name, user.ScreenName)
	util.ShowSuccessMsg("Muted", text, cfg.Color.BoxFg, cfg.Color.Mute)
}

// Unmute ミュートを解除する
func Unmute(name string) {
	user, err := api.UnmuteUser(name, nil)
	if err != nil {
		showAPIErrorString(err)
		return
	}
	text := fmt.Sprintf("%s (@%s)", user.Name, user.ScreenName)
	util.ShowSuccessMsg("Unmuted", text, cfg.Color.BoxFg, cfg.Color.Mute)
}
