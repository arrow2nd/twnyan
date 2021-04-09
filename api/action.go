package api

import (
	"errors"
	"fmt"
	"strconv"
)

// Favorite いいねする
func (tw *TwitterAPI) Favorite(tweetID string) (string, error) {
	id, _ := strconv.ParseInt(tweetID, 10, 64)

	tweet, err := tw.API.Favorite(id)
	if err != nil {
		return "", errors.New(parseAPIErrorMsg(err))
	}

	return tweet.FullText, nil
}

// Unfavorite いいねを取り消す
func (tw *TwitterAPI) Unfavorite(tweetID string) (string, error) {
	id, _ := strconv.ParseInt(tweetID, 10, 64)

	tweet, err := tw.API.Unfavorite(id)
	if err != nil {
		return "", errors.New(parseAPIErrorMsg(err))
	}

	return tweet.FullText, nil
}

// Retweet リツイートする
func (tw *TwitterAPI) Retweet(tweetID string) (string, error) {
	id, _ := strconv.ParseInt(tweetID, 10, 64)

	tweet, err := tw.API.Retweet(id, true)
	if err != nil {
		return "", errors.New(parseAPIErrorMsg(err))
	}

	return tweet.RetweetedStatus.FullText, nil
}

// UnRetweet リツイートを取り消す
func (tw *TwitterAPI) UnRetweet(tweetID string) (string, error) {
	id, _ := strconv.ParseInt(tweetID, 10, 64)

	tweet, err := tw.API.UnRetweet(id, true)
	if err != nil {
		return "", errors.New(parseAPIErrorMsg(err))
	}

	return tweet.FullText, nil
}

// Follow フォローする
func (tw *TwitterAPI) Follow(screenname string) (string, error) {
	user, err := tw.API.FollowUser(screenname)
	if err != nil {
		return "", errors.New(parseAPIErrorMsg(err))
	}

	str := fmt.Sprintf("%s @%s", user.Name, user.ScreenName)

	return str, nil
}

// Unfollow フォローを解除する
func (tw *TwitterAPI) Unfollow(screenname string) (string, error) {
	user, err := tw.API.UnfollowUser(screenname)
	if err != nil {
		return "", errors.New(parseAPIErrorMsg(err))
	}

	str := fmt.Sprintf("%s @%s", user.Name, user.ScreenName)

	return str, nil
}

// Block ブロックする
func (tw *TwitterAPI) Block(screenname string) (string, error) {
	user, err := tw.API.BlockUser(screenname, nil)
	if err != nil {
		return "", errors.New(parseAPIErrorMsg(err))
	}

	str := fmt.Sprintf("%s @%s", user.Name, user.ScreenName)

	return str, nil
}

// Unblock ブロックを解除する
func (tw *TwitterAPI) Unblock(screenname string) (string, error) {
	user, err := tw.API.UnblockUser(screenname, nil)
	if err != nil {
		return "", errors.New(parseAPIErrorMsg(err))
	}

	str := fmt.Sprintf("%s @%s", user.Name, user.ScreenName)

	return str, nil
}

// Mute ミュートする
func (tw *TwitterAPI) Mute(screenname string) (string, error) {
	user, err := tw.API.MuteUser(screenname, nil)
	if err != nil {
		return "", errors.New(parseAPIErrorMsg(err))
	}

	str := fmt.Sprintf("%s @%s", user.Name, user.ScreenName)

	return str, nil
}

// Unmute ミュートを解除する
func (tw *TwitterAPI) Unmute(screenname string) (string, error) {
	user, err := tw.API.UnmuteUser(screenname, nil)
	if err != nil {
		return "", errors.New(parseAPIErrorMsg(err))
	}

	str := fmt.Sprintf("%s @%s", user.Name, user.ScreenName)

	return str, nil
}
