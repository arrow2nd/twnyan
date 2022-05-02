package api

import (
	"errors"

	"github.com/ChimeraCoder/anaconda"
	"github.com/garyburd/go-oauth/oauth"
)

const (
	consumerKey    = "qIrO5Nme1D9UVDV5QCBVyjfZK"
	consumerSecret = "umr6nOFzV3W0AfdQoWPxKSh2ZMEeRgHFih5xQDTlBRO3DoEq8z"
)

// Twitter API
type Twitter struct {
	API     *anaconda.TwitterApi
	OwnUser *anaconda.User
	List    map[string]int64
}

func init() {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
}

// New 生成
func New() *Twitter {
	return &Twitter{
		API:     nil,
		OwnUser: nil,
		List:    nil,
	}
}

// Init 初期化
func (tw *Twitter) Init(cred *oauth.Credentials) {
	var err error

	tw.API = anaconda.NewTwitterApi(cred.Token, cred.Secret)
	tw.API.ReturnRateLimitError(true)

	// ユーザー情報を取得
	tw.OwnUser, err = tw.fetchSelfInfo()
	if err != nil {
		panic(tw.createAPIErrorText("", err))
	}

	// リスト情報を取得
	if err := tw.cacheListInfo(); err != nil {
		panic(tw.createAPIErrorText("", err))
	}
}

// Auth アプリケーション認証
func (tw *Twitter) Auth() (*oauth.Credentials, string, error) {
	authAPI := anaconda.NewTwitterApi("", "")

	// 認証URL取得
	uri, cred, err := authAPI.AuthorizationURL("oob")
	if err != nil {
		return nil, "", errors.New("Failed to issue the authentication URL")
	}

	showLogo()

	// URLを表示してPINコードの入力を待つ
	showAuthUrl(uri)
	pin := inputPinCode()

	// トークン発行
	cred, query, err := authAPI.GetCredentials(cred, pin)
	if err != nil {
		return nil, "", errors.New("Access token could not be obtained")
	}

	return cred, query.Get("screen_name"), nil
}
