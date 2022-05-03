package twitter

import (
	"errors"
	"net/url"
	"time"

	"github.com/ChimeraCoder/anaconda"
)

type timelineCategory uint8

const (
	Home timelineCategory = iota
	Mention
	User
)

// FetchUserInfo ユーザ情報を取得
func (tw *Twitter) FetchUserInfo(screenName string) (*anaconda.User, error) {
	user, err := tw.API.GetUsersShow(screenName, nil)
	if err != nil {
		return nil, tw.createAPIError("/users/show", err)
	}

	return &user, nil
}

// FetchRelationships ユーザとの関係性を取得
func (tw *Twitter) FetchRelationships(userId string) ([]string, error) {
	query := url.Values{"user_id": {userId}}

	relationships, err := tw.API.GetFriendshipsLookup(query)
	if err != nil {
		return nil, tw.createAPIError("/statuses/lookup", err)
	}

	return relationships[0].Connections, nil
}

// FetchTimelineTweets タイムラインのツイートを取得
func (tw *Twitter) FetchTimelineTweets(category timelineCategory, query url.Values) (*[]anaconda.Tweet, error) {
	var (
		timeline     []anaconda.Tweet
		resourceName string
		err          error
	)

	switch category {
	case Home:
		resourceName = "/statuses/home_timeline"
		timeline, err = tw.API.GetHomeTimeline(query)
	case Mention:
		resourceName = "/statuses/mentions_timeline"
		timeline, err = tw.API.GetMentionsTimeline(query)
	case User:
		resourceName = "/statuses/user_timeline"
		timeline, err = tw.API.GetUserTimeline(query)
	default:
		return nil, errors.New("category is wrong")
	}

	if err != nil {
		return nil, tw.createAPIError(resourceName, err)
	}

	return &timeline, nil
}

// FetchListTweets リストのツイートを取得
func (tw *Twitter) FetchListTweets(listId int64, count string) (*[]anaconda.Tweet, error) {
	query := CreateQuery(count)

	timeline, err := tw.API.GetListTweets(listId, true, query)
	if err != nil {
		return nil, tw.createAPIError("/lists/statuses", err)
	}

	return &timeline, nil
}

// FetchSearchResult 検索結果を取得
func (tw *Twitter) FetchSearchResult(queryStr, count string) (*[]anaconda.Tweet, error) {
	query := CreateQuery(count)

	// 検索結果を取得
	result, err := tw.API.GetSearch(queryStr+" -filter:retweets", query)
	if err != nil {
		return nil, tw.createAPIError("/search/tweets", err)
	}

	// 検索結果が0件ならエラー
	if len(result.Statuses) == 0 {
		return nil, errors.New("no tweets found")
	}

	return &result.Statuses, nil
}

// fetchRateLimitResetTime レート制限の解除時刻を取得
func (tw *Twitter) fetchRateLimitResetTime(resouceName string) string {
	rateLimit, err := tw.API.GetRateLimits([]string{"statuses", "lists", "search", "users"})
	if err != nil {
		return ""
	}

	for _, resources := range rateLimit.Resources {
		for name, baseResource := range resources {
			if name == resouceName {
				return time.Unix(int64(baseResource.Reset), 0).Local().Format("15:04:05")
			}
		}
	}

	return ""
}

// fetchSelfInfo 自分のユーザ情報を取得
func (tw *Twitter) fetchSelfInfo() (*anaconda.User, error) {
	user, err := tw.API.GetSelf(nil)
	if err != nil {
		return nil, tw.createAPIError("", err)
	}

	return &user, nil
}

// cacheListInfo リスト情報を取得してキャッシュ
func (tw *Twitter) cacheListInfo() error {
	// リスト情報を取得
	lists, err := tw.API.GetListsOwnedBy(tw.OwnUser.Id, nil)
	if err != nil {
		return err
	}

	tw.List = map[string]int64{}

	// リスト情報を追加
	for _, ls := range lists {
		tw.List[ls.Name] = ls.Id
	}

	return nil
}
