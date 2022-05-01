package api

import (
	"fmt"

	"github.com/ChimeraCoder/anaconda"
)

const (
	consumerKey    = "qIrO5Nme1D9UVDV5QCBVyjfZK"
	consumerSecret = "umr6nOFzV3W0AfdQoWPxKSh2ZMEeRgHFih5xQDTlBRO3DoEq8z"
)

type list struct {
	Names []string
	IDs   []int64
}

// TwitterAPI API構造体
type TwitterAPI struct {
	API     *anaconda.TwitterApi
	OwnUser *anaconda.User
	List    list
}

func init() {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
}

// New 生成
func New() *TwitterAPI {
	return &TwitterAPI{
		API:     nil,
		OwnUser: &anaconda.User{},
		List:    list{},
	}
}

// Init 初期化
func (tw *TwitterAPI) Init(token, secret string) {
	var err error

	tw.API = anaconda.NewTwitterApi(token, secret)
	tw.API.ReturnRateLimitError(true)

	// ユーザー情報を取得
	tw.OwnUser, err = tw.fetchSelfInfo()
	if err != nil {
		panic(tw.createAPIErrorMsg("", err))
	}

	// リスト情報を取得
	if err := tw.createListInfoSlice(); err != nil {
		panic(tw.createAPIErrorMsg("", err))
	}
}

// Auth アプリケーション認証
func (tw *TwitterAPI) Auth() (string, string) {
	authAPI := anaconda.NewTwitterApi("", "")

	// 認証URL取得
	uri, cred, err := authAPI.AuthorizationURL("oob")
	if err != nil {
		fmt.Println("Error: Failed to issue the authentication URL")
		panic(err)
	}

	// URLを表示してPINコードの入力を待つ
	showLogo()
	showAuthUrl(uri)
	pin := inputPinCode()

	// トークン発行
	cred, _, err = authAPI.GetCredentials(cred, pin)
	if err != nil {
		fmt.Println("Error: Access token could not be obtained")
		panic(err)
	}

	tw.Init(cred.Token, cred.Secret)
	return cred.Token, cred.Secret
}
