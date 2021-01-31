package api

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
)

// GetFriendships ユーザーとの関係を取得
func (ta *TwitterAPI) GetFriendships(u *anaconda.User) (string, error) {
	v := url.Values{"user_id": {u.IdStr}}
	friendships, err := ta.API.GetFriendshipsLookup(v)
	if err != nil {
		return "", errors.New(parseAPIError(err))
	}

	status := ""
	for _, v := range friendships[0].Connections {
		switch v {
		case "followed_by":
			status += fmt.Sprintf("[blue]Followed by[while] ")
		case "following":
			status += fmt.Sprintf("[blue]Following[while] ")
		case "blocking":
			status += fmt.Sprintf("[red]Blocking[white] ")
		case "muting":
			status += fmt.Sprintf("[yellow]Muting[white] ")
		}
	}

	return status, nil
}

// GetTimeline タイムラインを取得
func (ta *TwitterAPI) GetTimeline(mode string, v url.Values) (*[]anaconda.Tweet, error) {
	var (
		timeline []anaconda.Tweet
		err      error
	)

	switch mode {
	case "home":
		timeline, err = ta.API.GetHomeTimeline(v)
	case "mention":
		timeline, err = ta.API.GetMentionsTimeline(v)
	case "user":
		timeline, err = ta.API.GetUserTimeline(v)
	}

	if err != nil {
		return nil, errors.New(parseAPIError(err))
	}

	return &timeline, nil
}

// GetListTimeline リストタイムラインを取得
func (ta *TwitterAPI) GetListTimeline(listID int64, count string) (*[]anaconda.Tweet, error) {
	v := CreateURLValues(count)
	timeline, err := ta.API.GetListTweets(listID, true, v)
	if err != nil {
		return nil, errors.New(parseAPIError(err))
	}
	return &timeline, nil
}

// GetSearchResult 検索結果を取得
func (ta *TwitterAPI) GetSearchResult(query, count string) (*[]anaconda.Tweet, error) {
	v := CreateURLValues(count)
	query += " -filter:retweets"
	result, err := ta.API.GetSearch(query, v)
	if err != nil {
		return nil, errors.New(parseAPIError(err))
	}
	return &result.Statuses, nil
}

// getSelf 自分のユーザー情報を取得
func (ta *TwitterAPI) getSelf() (*anaconda.User, error) {
	user, err := ta.API.GetSelf(nil)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// getLists リストの一覧を取得
func (ta *TwitterAPI) getLists() ([]string, []int64, error) {
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
