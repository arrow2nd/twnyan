package api

import (
	"errors"
	"strconv"
)

// Favorite いいねする
func (tw *TwitterAPI) Favorite(tweetID string) error {
	id, _ := strconv.ParseInt(tweetID, 10, 64)
	_, err := tw.API.Favorite(id)
	if err != nil {
		return errors.New(parseAPIError(err))
	}
	return nil
}

// Unfavorite いいねを取り消す
func (tw *TwitterAPI) Unfavorite(tweetID string) error {
	id, _ := strconv.ParseInt(tweetID, 10, 64)
	_, err := tw.API.Unfavorite(id)
	if err != nil {
		return errors.New(parseAPIError(err))
	}
	return nil
}

// Retweet リツイートする
func (tw *TwitterAPI) Retweet(tweetID string) error {
	id, _ := strconv.ParseInt(tweetID, 10, 64)
	_, err := tw.API.Retweet(id, true)
	if err != nil {
		return errors.New(parseAPIError(err))
	}
	return nil
}

// UnRetweet リツイートを取り消す
func (tw *TwitterAPI) UnRetweet(tweetID string) error {
	id, _ := strconv.ParseInt(tweetID, 10, 64)
	_, err := tw.API.UnRetweet(id, true)
	if err != nil {
		return errors.New(parseAPIError(err))
	}
	return nil
}

// Follow フォローする
func (tw *TwitterAPI) Follow(screenname string) error {
	_, err := tw.API.FollowUser(screenname)
	if err != nil {
		return errors.New(parseAPIError(err))
	}
	return nil
}

// Unfollow フォローを解除する
func (tw *TwitterAPI) Unfollow(screenname string) error {
	_, err := tw.API.UnfollowUser(screenname)
	if err != nil {
		return errors.New(parseAPIError(err))
	}
	return nil
}

// Block ブロックする
func (tw *TwitterAPI) Block(screenname string) error {
	_, err := tw.API.BlockUser(screenname, nil)
	if err != nil {
		return errors.New(parseAPIError(err))
	}
	return nil
}

// Unblock ブロックを解除する
func (tw *TwitterAPI) Unblock(screenname string) error {
	_, err := tw.API.UnblockUser(screenname, nil)
	if err != nil {
		return errors.New(parseAPIError(err))
	}
	return nil
}

// Mute ミュートする
func (tw *TwitterAPI) Mute(screenname string) error {
	_, err := tw.API.MuteUser(screenname, nil)
	if err != nil {
		return errors.New(parseAPIError(err))
	}
	return nil
}

// Unmute ミュートを解除する
func (tw *TwitterAPI) Unmute(screenname string) error {
	_, err := tw.API.UnmuteUser(screenname, nil)
	if err != nil {
		return errors.New(parseAPIError(err))
	}
	return nil
}
