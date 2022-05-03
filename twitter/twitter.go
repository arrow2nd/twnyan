package twitter

import (
	"errors"
	"sync"

	"github.com/ChimeraCoder/anaconda"
	"github.com/garyburd/go-oauth/oauth"
)

const (
	consumerKey    = "qIrO5Nme1D9UVDV5QCBVyjfZK"
	consumerSecret = "umr6nOFzV3W0AfdQoWPxKSh2ZMEeRgHFih5xQDTlBRO3DoEq8z"
)

// Twitter APIと付随するデータの管理
type Twitter struct {
	// API TwitterAPI
	API *anaconda.TwitterApi
	// OwnUser ユーザの情報
	OwnUser *anaconda.User
	// Tweets 操作対象のツイート
	Tweets []anaconda.Tweet
	// List リスト情報
	List map[string]int64
	mu   sync.Mutex
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
		Tweets:  []anaconda.Tweet{},
		List:    nil,
		mu:      sync.Mutex{},
	}
}

// Init 初期化
func (tw *Twitter) Init(cred *oauth.Credentials) {
	var err error

	tw.API = anaconda.NewTwitterApi(cred.Token, cred.Secret)
	tw.API.ReturnRateLimitError(true)

	// ユーザ情報を取得
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
