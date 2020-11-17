package twitter

import (
	"github.com/ChimeraCoder/anaconda"
	"github.com/arrow2nd/twnyan/config"
)

const (
	consumerKey    = "qIrO5Nme1D9UVDV5QCBVyjfZK"
	consumerSecret = "umr6nOFzV3W0AfdQoWPxKSh2ZMEeRgHFih5xQDTlBRO3DoEq8z"
)

var (
	api *anaconda.TwitterApi
	cfg *config.ConfigData
)

func init() {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
}

// SetConfig 設定
func SetConfig(config *config.ConfigData) {
	// API
	api = anaconda.NewTwitterApi(config.Credentials.Token, config.Credentials.Secret)

	// 設定情報
	cfg = config
}
