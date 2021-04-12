package cmd

import (
	"os"
	"os/signal"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/arrow2nd/ishell"
	"github.com/arrow2nd/twnyan/api"
)

func (cmd *Cmd) addStreamCmd() {
	cmd.shell.AddCmd(&ishell.Cmd{
		Name:    "stream",
		Aliases: []string{"st"},
		Func:    cmd.streamCmd,
		Help:    "start a pseudo-UserStream",
		LongHelp: createLongHelp(
			"After accumulating up to 250 tweets in the first minute, the tweets will be displayed with a one-minute delay, just like the UserStream API.",
			"st",
			"stream",
			"",
		),
	})
}

func (cmd *Cmd) streamCmd(c *ishell.Context) {
	// タイムラインを表示
	tweets, sinceId, err := cmd.fetchHomeTimelineTweets("25", "")
	if err != nil {
		cmd.showErrorMessage(err.Error())
		return
	}
	cmd.view.ShowTweetsFromArray(*tweets, false)

	// Ctrl+Cが入力されたら通知する
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	// 疑似ストリーム開始
	go cmd.startUserStream(sinceId, quit)

	// 通知が来たら終了
	<-quit
	signal.Stop(quit)
	close(quit)
}

// startUserStream  疑似userStreamを開始
func (cmd *Cmd) startUserStream(startSinceId string, quit chan os.Signal) {
	var (
		accumulateTweets AccumulateTweets
		err              error
	)
	sinceId := startSinceId

	for {
		for i := 0; i < 60; i++ {
			select {
			case <-quit:
				// 中断された
				return
			default:
				// 蓄積したツイートを表示
				accumulateTweets = cmd.showAccumulateTweet(accumulateTweets)
			}
			time.Sleep(1 * time.Second)
		}

		// 次の60秒間に表示するツイートを取得
		accumulateTweets, sinceId, err = cmd.createNewAccumulateTweets(sinceId)
		if err != nil {
			cmd.showErrorMessage(err.Error())
		}
	}
}

// showAccumulateTweet 蓄積されたツイートを表示
func (cmd *Cmd) showAccumulateTweet(accumulateTweets AccumulateTweets) AccumulateTweets {
	nowTime := time.Now().Local().Unix()

	for displayTiming, tweet := range accumulateTweets {
		if displayTiming <= nowTime {
			cmd.view.ShowTweet(&tweet, " ", false)
			delete(accumulateTweets, displayTiming)
		}
	}

	return accumulateTweets
}

// createNewAccumulateTweets 新しい蓄積データを作成する
func (cmd *Cmd) createNewAccumulateTweets(sinceId string) (AccumulateTweets, string, error) {
	tweets, newSinceId, err := cmd.fetchHomeTimelineTweets("200", sinceId)
	if err != nil {
		return nil, "", err
	}

	// 画面に表示するタイミング（時間）をキーとした連想配列
	accumulateTweets := make(AccumulateTweets, len(*tweets))

	// 投稿時間の1分後をキーにして連想配列を作成
	for _, tweet := range *tweets {
		createdAtTime, _ := tweet.CreatedAtTime()
		displayTiming := createdAtTime.Add(1 * time.Minute).Local().Unix()
		accumulateTweets[displayTiming] = tweet
	}

	return accumulateTweets, newSinceId, nil
}

// fetchHomeTimelineTweets ホームタイムラインのツイートを取得
func (cmd *Cmd) fetchHomeTimelineTweets(count, sinceId string) (*[]anaconda.Tweet, string, error) {
	// 取得件数が指定されていないならデフォルト値を代入
	if count == "" {
		count = cmd.cfg.Option.Counts
	}

	query := api.CreateQuery(count)
	if sinceId != "" {
		query.Add("since_id", sinceId)
	}

	tweets, err := cmd.api.FetchTimelineTweets("home", query)
	if err != nil {
		return nil, "", err
	}

	newSinceId := sinceId
	if len(*tweets) > 0 {
		newSinceId = (*tweets)[0].IdStr
	}

	return tweets, newSinceId, nil
}
