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
		Help:    "",
		LongHelp: createLongHelp(
			"",
			"st",
			"stream",
			"stream",
		),
	})
}

func (cmd *Cmd) streamCmd(c *ishell.Context) {
	// タイムラインを取得して表示
	tweets, err := cmd.fetchHomeTimelineTweets("25", "")
	if err != nil {
		cmd.showErrorMessage(err.Error())
		return
	}
	cmd.view.ShowTweetsFromArray(*tweets, false)

	// Ctrl+Cが入力されたら通知する
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	// 疑似ストリーム開始
	go cmd.startUserStream(quit)

	// 通知を待つ
	<-quit
	signal.Stop(quit)
	close(quit)
}

// startUserStream  疑似userStream
func (cmd *Cmd) startUserStream(quit chan os.Signal) {
	var (
		accumulateTweets AccumulateTweets
		err              error
	)
	sinceId := ""

	for {
		for i := 0; i < 60; i++ {
			select {
			case <-quit:
				return
			default:
				accumulateTweets = cmd.showAccumulateTweet(accumulateTweets)
			}
			time.Sleep(1 * time.Second)
		}

		// 蓄積する分のツイートを取得
		accumulateTweets, sinceId, err = cmd.fetchAccumulateTweets(sinceId)
		if err != nil {
			cmd.showErrorMessage(err.Error())
		}
	}
}

// showAccumulateTweet 蓄積されたツイートを表示
func (cmd *Cmd) showAccumulateTweet(accumulateTweets AccumulateTweets) AccumulateTweets {
	nowTime := time.Now().Local().Unix()

	for time, tweet := range accumulateTweets {
		// fmt.Printf("%d = %d\n", time, nowTime)
		if time <= nowTime {
			cmd.view.ShowTweet(&tweet, " ", false)
			delete(accumulateTweets, time)
		}
	}

	return accumulateTweets
}

// fetchAccumulateTweets 蓄積されたツイートを取得
func (cmd *Cmd) fetchAccumulateTweets(sinceId string) (AccumulateTweets, string, error) {
	tweets, err := cmd.fetchHomeTimelineTweets("50", sinceId)
	if err != nil {
		return nil, "", err
	}

	// 1分後の投稿時間をキーにした連想配列を作成
	accumulateTweets := make(AccumulateTweets, len(*tweets))
	for _, tweet := range *tweets {
		createdAtTime, _ := tweet.CreatedAtTime()
		createdAtUnixTime := createdAtTime.Add(1 * time.Minute).Local().Unix()
		accumulateTweets[createdAtUnixTime] = tweet
	}

	newSinceId := (*tweets)[0].IdStr
	// fmt.Printf("ok! : %s\n", newSinceId)

	return accumulateTweets, newSinceId, nil
}

// fetchHomeTimelineTweets タイムラインを取得
func (cmd *Cmd) fetchHomeTimelineTweets(count, sinceId string) (*[]anaconda.Tweet, error) {
	if count == "" {
		count = cmd.cfg.Option.Counts
	}

	query := api.CreateQuery(count)
	if sinceId != "" {
		query.Add("since_id", sinceId)
	}

	tweets, err := cmd.api.FetchTimelineTweets("home", query)
	if err != nil {
		return nil, err
	}

	return tweets, nil
}
