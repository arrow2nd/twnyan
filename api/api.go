package api

import (
	"fmt"

	"github.com/ChimeraCoder/anaconda"
	"github.com/gookit/color"
)

const (
	consumerKey    = "qIrO5Nme1D9UVDV5QCBVyjfZK"
	consumerSecret = "umr6nOFzV3W0AfdQoWPxKSh2ZMEeRgHFih5xQDTlBRO3DoEq8z"
)

// TwitterAPI API構造体
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

// New API構造体を作成
func New() *TwitterAPI {
	tw := &TwitterAPI{
		API:       nil,
		OwnUser:   &anaconda.User{},
		ListNames: []string{},
		ListIDs:   []int64{},
	}
	return tw
}

// Init 初期化
func (ta *TwitterAPI) Init(token, secret string) error {
	var err error
	ta.API = anaconda.NewTwitterApi(token, secret)

	// ユーザー情報を取得
	ta.OwnUser, err = ta.getSelf()
	if err != nil {
		return err
	}
	// リスト情報を取得
	ta.ListNames, ta.ListIDs, err = ta.getLists()
	if err != nil {
		return err
	}

	return nil
}

// Auth 認証
func (ta *TwitterAPI) Auth() (string, string) {
	authAPI := anaconda.NewTwitterApi("", "")

	// 認証URL取得
	uri, cred, err := authAPI.AuthorizationURL("oob")
	if err != nil {
		fmt.Println("Error: Failed to issue the authentication URL")
		panic(err)
	}

	// ロゴ
	color.Red.Println(" __                                     ")
	color.Yellow.Println("|  |_.--.--.--.-----.--.--.---.-.-----.")
	color.Green.Println("|   _|  |  |  |     |  |  |  _  |     |")
	color.Cyan.Println("|____|________|__|__|___  |___._|__|__|")
	color.Blue.Println("                    |_____|            ")
	// URL
	fmt.Println("\n🐈 Go to the following URL to authenticate the application and enter the PIN that is displayed")
	fmt.Printf("[ %s ]\n\n", uri)

	// PIN入力
	pin := ""
	fmt.Print("PIN : ")
	fmt.Scanf("%s", &pin)

	// トークン発行
	cred, _, err = authAPI.GetCredentials(cred, pin)
	if err != nil {
		fmt.Println("Error: Access token could not be obtained")
		panic(err)
	}

	ta.Init(cred.Token, cred.Secret)

	return cred.Token, cred.Secret
}
