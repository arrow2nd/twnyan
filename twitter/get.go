package twitter

import (
	"net/url"

	"github.com/ChimeraCoder/anaconda"
	"github.com/gookit/color"
)

// GetLists リストの一覧を取得
func GetLists() ([]string, []int64) {
	// 自分のユーザーIDを取得
	me, err := api.GetSelf(nil)
	if err != nil {
		showAPIErrorString(err)
		return nil, nil
	}

	// リストを取得
	list, err := api.GetListsOwnedBy(me.Id, nil)
	if err != nil {
		showAPIErrorString(err)
		return nil, nil
	}

	// リスト名とIDのスライスを作成
	listID := make([]int64, len(list))
	listName := make([]string, len(list))
	for i, l := range list {
		listName[i] = l.Name
		listID[i] = l.Id
	}

	return listName, listID
}

// getFriendships ユーザーとの関係を取得
func getFriendships(u *anaconda.User) string {
	v := url.Values{"user_id": {u.IdStr}}
	friendships, err := api.GetFriendshipsLookup(v)
	if err != nil {
		showAPIErrorString(err)
		return ""
	}

	status := ""
	for _, v := range friendships[0].Connections {
		switch v {
		case "followed_by":
			status += color.HEXStyle(cfg.Color.BoxFg, cfg.Color.Follow).Sprint(" followed-by ") + " "
		case "following":
			status += color.HEXStyle(cfg.Color.BoxFg, cfg.Color.Follow).Sprintf(" %s ", v) + " "
		case "blocking":
			status += color.HEXStyle(cfg.Color.BoxFg, cfg.Color.Block).Sprintf(" %s ", v) + " "
		case "muting":
			status += color.HEXStyle(cfg.Color.BoxFg, cfg.Color.Mute).Sprintf(" %s ", v) + " "
		}
	}

	return status
}

// getTimeline タイムラインを取得
func getTimeline(mode string, v url.Values) []anaconda.Tweet {
	var (
		tl  []anaconda.Tweet
		err error
	)

	switch mode {
	case "home":
		tl, err = api.GetHomeTimeline(v)
	case "mention":
		tl, err = api.GetMentionsTimeline(v)
	case "user":
		tl, err = api.GetUserTimeline(v)
	}

	if err != nil {
		showAPIErrorString(err)
		return nil
	}

	return tl
}
