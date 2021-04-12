package api

import (
	"fmt"

	"github.com/ChimeraCoder/anaconda"
)

const (
	consumerKey    = "qIrO5Nme1D9UVDV5QCBVyjfZK"
	consumerSecret = "umr6nOFzV3W0AfdQoWPxKSh2ZMEeRgHFih5xQDTlBRO3DoEq8z"
)

type TwitterAPI struct {
	API       *anaconda.TwitterApi
	OwnUser   *anaconda.User
	ListNames []string
	ListIDs   []int64
}

func init() {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
}

// New 構造体を初期化
func New() *TwitterAPI {
	api := &TwitterAPI{
		API:       nil,
		OwnUser:   &anaconda.User{},
		ListNames: []string{},
		ListIDs:   []int64{},
	}

	return api
}

// Init 初期化処理
func (tw *TwitterAPI) Init(token, secret string) error {
	var err error

	// TwitterApi構造体を作成
	tw.API = anaconda.NewTwitterApi(token, secret)
	tw.API.ReturnRateLimitError(true)

	// ユーザー情報を取得
	tw.OwnUser, err = tw.fetchSelfInfo()
	if err != nil {
		return err
	}

	// リスト情報を取得
	tw.ListNames, tw.ListIDs, err = tw.createListInfoSlice()
	if err != nil {
		return err
	}

	return nil
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

	// URLを表示してPINコードを入力
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
