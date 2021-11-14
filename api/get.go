package api

import (
	"errors"
	"net/url"
	"time"

	"github.com/ChimeraCoder/anaconda"
)

// FetchRelationships ユーザーとの関係性を取得
func (tw *TwitterAPI) FetchRelationships(userID string) ([]string, error) {
	values := url.Values{"user_id": {userID}}

	relationships, err := tw.API.GetFriendshipsLookup(values)
	if err != nil {
		return nil, errors.New(tw.createAPIErrorMsg("/statuses/lookup", err))
	}

	return relationships[0].Connections, nil
}

// FetchTimelineTweets タイムラインのツイートを取得
func (tw *TwitterAPI) FetchTimelineTweets(category string, query url.Values) (*[]anaconda.Tweet, error) {
	var (
		timeline     []anaconda.Tweet
		resourceName string
		err          error
	)

	switch category {
	case "home":
		resourceName = "/statuses/home_timeline"
		timeline, err = tw.API.GetHomeTimeline(query)
	case "mention":
		resourceName = "/statuses/mentions_timeline"
		timeline, err = tw.API.GetMentionsTimeline(query)
	case "user":
		resourceName = "/statuses/user_timeline"
		timeline, err = tw.API.GetUserTimeline(query)
	default:
		return nil, errors.New("category is wrong")
	}

	if err != nil {
		return nil, errors.New(tw.createAPIErrorMsg(resourceName, err))
	}

	return &timeline, nil
}

// FetchListTweets リストのツイートを取得
func (tw *TwitterAPI) FetchListTweets(listID int64, count string) (*[]anaconda.Tweet, error) {
	query := CreateQuery(count)

	timeline, err := tw.API.GetListTweets(listID, true, query)
	if err != nil {
		return nil, errors.New(tw.createAPIErrorMsg("/lists/statuses", err))
	}

	return &timeline, nil
}

// FetchSearchResult 検索結果を取得
func (tw *TwitterAPI) FetchSearchResult(queryStr, count string) (*[]anaconda.Tweet, error) {
	query := CreateQuery(count)
	queryStr += " -filter:retweets"

	// 検索結果を取得
	result, err := tw.API.GetSearch(queryStr, query)
	if err != nil {
		return nil, errors.New(tw.createAPIErrorMsg("/search/tweets", err))
	}

	// 検索結果が0件ならエラーを返す
	if len(result.Statuses) == 0 {
		return nil, errors.New("no tweets found")
	}

	return &result.Statuses, nil
}

// fetchRateLimitResetTime レート制限の解除時刻を取得
func (tw *TwitterAPI) fetchRateLimitResetTime(resouceName string) string {
	rateLimitRes, err := tw.API.GetRateLimits([]string{"statuses", "lists", "search"})
	if err != nil {
		return ""
	}

	for _, resources := range rateLimitRes.Resources {
		for name, baseResource := range resources {
			if name == resouceName {
				return time.Unix(int64(baseResource.Reset), 0).Local().Format("15:04:05")
			}
		}
	}

	return ""
}

// fetchSelfInfo 自分のユーザー情報を取得
func (tw *TwitterAPI) fetchSelfInfo() (*anaconda.User, error) {
	user, err := tw.API.GetSelf(nil)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// createListInfoSlice リスト名とリストIDのスライスを作成
func (tw *TwitterAPI) createListInfoSlice() ([]string, []int64, error) {
	// リストの情報を取得
	lists, err := tw.API.GetListsOwnedBy(tw.OwnUser.Id, nil)
	if err != nil {
		return nil, nil, err
	}

	// リスト名とIDのスライスを作成
	listIDs := make([]int64, len(lists))
	listNames := make([]string, len(lists))

	for i, ls := range lists {
		listNames[i] = ls.Name
		listIDs[i] = ls.Id
	}

	return listNames, listIDs, nil
}
