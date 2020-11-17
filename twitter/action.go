package twitter

import (
	"fmt"
	"strconv"
)

// Favorite いいねする
func Favorite(tweetID string) {
	id, _ := strconv.ParseInt(tweetID, 10, 64)
	tweet, err := api.Favorite(id)
	if err != nil {
		showAPIErrorString(err)
		return
	}
	showSuccessMsg(tweet.FullText, "Favorited!:", cfg.Color.Fav)
}

// Unfavorite いいねを取り消す
func Unfavorite(tweetID string) {
	id, _ := strconv.ParseInt(tweetID, 10, 64)
	tweet, err := api.Unfavorite(id)
	if err != nil {
		showAPIErrorString(err)
		return
	}
	showSuccessMsg(tweet.FullText, "Unfavorited:", cfg.Color.Fav)
}

// Retweet リツイートする
func Retweet(tweetID string) {
	id, _ := strconv.ParseInt(tweetID, 10, 64)
	tweet, err := api.Retweet(id, true)
	if err != nil {
		showAPIErrorString(err)
		return
	}
	showSuccessMsg(tweet.FullText, "Retweeted!:", cfg.Color.RT)
}

// UnRetweet リツイートを取り消す
func UnRetweet(tweetID string) {
	id, _ := strconv.ParseInt(tweetID, 10, 64)
	tweet, err := api.UnRetweet(id, true)
	if err != nil {
		showAPIErrorString(err)
		return
	}
	showSuccessMsg(tweet.FullText, "UnRetweeted:", cfg.Color.RT)
}

// Follow フォローする
func Follow(name string) {
	user, err := api.FollowUser(name)
	if err != nil {
		showAPIErrorString(err)
		return
	}
	text := fmt.Sprintf("%s (@%s)", user.Name, user.ScreenName)
	showSuccessMsg(text, "Followed!:", cfg.Color.Follow)
}

// Unfollow フォローを解除する
func Unfollow(name string) {
	user, err := api.UnfollowUser(name)
	if err != nil {
		showAPIErrorString(err)
		return
	}
	text := fmt.Sprintf("%s (@%s)", user.Name, user.ScreenName)
	showSuccessMsg(text, "Unfollowed:", cfg.Color.Follow)
}

// Block ブロックする
func Block(name string) {
	user, err := api.BlockUser(name, nil)
	if err != nil {
		showAPIErrorString(err)
		return
	}
	text := fmt.Sprintf("%s (@%s)", user.Name, user.ScreenName)
	showSuccessMsg(text, "Blocked:", cfg.Color.Block)
}

// Unblock ブロックを解除する
func Unblock(name string) {
	user, err := api.UnblockUser(name, nil)
	if err != nil {
		showAPIErrorString(err)
		return
	}
	text := fmt.Sprintf("%s (@%s)", user.Name, user.ScreenName)
	showSuccessMsg(text, "Unblocked:", cfg.Color.Block)
}

// Mute ミュートする
func Mute(name string) {
	user, err := api.MuteUser(name, nil)
	if err != nil {
		showAPIErrorString(err)
		return
	}
	text := fmt.Sprintf("%s (@%s)", user.Name, user.ScreenName)
	showSuccessMsg(text, "Muted:", cfg.Color.Mute)
}

// Unmute ミュートを解除する
func Unmute(name string) {
	user, err := api.UnmuteUser(name, nil)
	if err != nil {
		showAPIErrorString(err)
		return
	}
	text := fmt.Sprintf("%s (@%s)", user.Name, user.ScreenName)
	showSuccessMsg(text, "Unmuted:", cfg.Color.Mute)
}
