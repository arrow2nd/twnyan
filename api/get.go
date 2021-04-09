package api

import (
	"errors"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
)

// FetchRelationships ユーザーとの関係性を取得
func (ta *TwitterAPI) FetchRelationships(userID string) ([]string, error) {
	v := url.Values{"user_id": {userID}}

	relationships, err := ta.API.GetFriendshipsLookup(v)
	if err != nil {
		return nil, errors.New(parseAPIErrorMsg(err))
	}

	return relationships[0].Connections, nil
}

// FetchTimelineTweets タイムラインのツイートを取得
func (ta *TwitterAPI) FetchTimelineTweets(category string, query url.Values) (*[]anaconda.Tweet, error) {
	var err error
	timeline := []anaconda.Tweet{}

	switch category {
	case "home":
		timeline, err = ta.API.GetHomeTimeline(query)
	case "mention":
		timeline, err = ta.API.GetMentionsTimeline(query)
	case "user":
		timeline, err = ta.API.GetUserTimeline(query)
	}

	if err != nil {
		return nil, errors.New(parseAPIErrorMsg(err))
	}

	return &timeline, nil
}

// FetchListTweets リストのツイートを取得
func (ta *TwitterAPI) FetchListTweets(listID int64, count string) (*[]anaconda.Tweet, error) {
	query := CreateQuery(count)

	timeline, err := ta.API.GetListTweets(listID, true, query)
	if err != nil {
		return nil, errors.New(parseAPIErrorMsg(err))
	}

	return &timeline, nil
}

// FetchSearchResult 検索結果を取得
func (ta *TwitterAPI) FetchSearchResult(queryStr, count string) (*[]anaconda.Tweet, error) {
	query := CreateQuery(count)
	queryStr += " -filter:retweets"

	// 検索結果を取得
	result, err := ta.API.GetSearch(queryStr, query)
	if err != nil {
		return nil, errors.New(parseAPIErrorMsg(err))
	}

	// 検索結果が0件ならエラーを返す
	if len(result.Statuses) == 0 {
		return nil, errors.New("no tweets found")
	}

	return &result.Statuses, nil
}

// fetchSelfInfo 自分のユーザー情報を取得
func (ta *TwitterAPI) fetchSelfInfo() (*anaconda.User, error) {
	user, err := ta.API.GetSelf(nil)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// createListInfoSlice リスト名とリストIDのスライスを作成
func (ta *TwitterAPI) createListInfoSlice() ([]string, []int64, error) {
	// リストの情報を取得
	lists, err := ta.API.GetListsOwnedBy(ta.OwnUser.Id, nil)
	if err != nil {
		return nil, nil, err
	}

	// リスト名とIDのスライスを作成
	id := make([]int64, len(lists))
	name := make([]string, len(lists))
	for i, l := range lists {
		name[i] = l.Name
		id[i] = l.Id
	}

	return name, id, nil
}
