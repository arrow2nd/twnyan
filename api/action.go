package api

import (
	"strconv"
)

// Favorite いいねする
func (tw *Twitter) Favorite(tweetId string) (string, error) {
	id, _ := strconv.ParseInt(tweetId, 10, 64)

	tweet, err := tw.API.Favorite(id)
	if err != nil {
		return "", tw.createAPIError("", err)
	}

	return tweet.FullText, nil
}

// Unfavorite いいねを取り消す
func (tw *Twitter) Unfavorite(tweetId string) (string, error) {
	id, _ := strconv.ParseInt(tweetId, 10, 64)

	tweet, err := tw.API.Unfavorite(id)
	if err != nil {
		return "", tw.createAPIError("", err)
	}

	return tweet.FullText, nil
}

// Retweet リツイートする
func (tw *Twitter) Retweet(tweetId string) (string, error) {
	id, _ := strconv.ParseInt(tweetId, 10, 64)

	tweet, err := tw.API.Retweet(id, true)
	if err != nil {
		return "", tw.createAPIError("", err)
	}

	return tweet.RetweetedStatus.FullText, nil
}

// UnRetweet リツイートを取り消す
func (tw *Twitter) UnRetweet(tweetId string) (string, error) {
	id, _ := strconv.ParseInt(tweetId, 10, 64)

	tweet, err := tw.API.UnRetweet(id, true)
	if err != nil {
		return "", tw.createAPIError("", err)
	}

	return tweet.FullText, nil
}

// Follow フォローする
func (tw *Twitter) Follow(screenName string) (string, error) {
	user, err := tw.API.FollowUser(screenName)
	if err != nil {
		return "", tw.createAPIError("", err)
	}

	return tw.createUserInfoText(user.Name, user.ScreenName), nil
}

// Unfollow フォローを解除する
func (tw *Twitter) Unfollow(screenName string) (string, error) {
	user, err := tw.API.UnfollowUser(screenName)
	if err != nil {
		return "", tw.createAPIError("", err)
	}

	return tw.createUserInfoText(user.Name, user.ScreenName), nil
}

// Block ブロックする
func (tw *Twitter) Block(screenName string) (string, error) {
	user, err := tw.API.BlockUser(screenName, nil)
	if err != nil {
		return "", tw.createAPIError("", err)
	}

	return tw.createUserInfoText(user.Name, user.ScreenName), nil
}

// Unblock ブロックを解除する
func (tw *Twitter) Unblock(screenName string) (string, error) {
	user, err := tw.API.UnblockUser(screenName, nil)
	if err != nil {
		return "", tw.createAPIError("", err)
	}

	return tw.createUserInfoText(user.Name, user.ScreenName), nil
}

// Mute ミュートする
func (tw *Twitter) Mute(screenName string) (string, error) {
	user, err := tw.API.MuteUser(screenName, nil)
	if err != nil {
		return "", tw.createAPIError("", err)
	}

	return tw.createUserInfoText(user.Name, user.ScreenName), nil
}

// Unmute ミュートを解除する
func (tw *Twitter) Unmute(screenName string) (string, error) {
	user, err := tw.API.UnmuteUser(screenName, nil)
	if err != nil {
		return "", tw.createAPIError("", err)
	}

	return tw.createUserInfoText(user.Name, user.ScreenName), nil
}
