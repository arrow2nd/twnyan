package cmd

import (
	"os"
	"os/signal"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/arrow2nd/ishell/v2"
	"github.com/arrow2nd/twnyan/twitter"
)

// AccumulateTweets 蓄積されたツイート
type AccumulateTweets map[int64]anaconda.Tweet

func (cmd *Cmd) newStreamCmd() *ishell.Cmd {
	return &ishell.Cmd{
		Name:    "stream",
		Aliases: []string{"st"},
		Func:    cmd.execStreamCmd,
		Help:    "start a pseudo userstream mode",
		LongHelp: createLongHelp(
			`Provide a one-minute buffer time for displaying the home timeline to simulate real-time updates like the UserStream API.
Ctrl+C to exit.`,
			"st",
			"stream",
			"",
		),
	}
}

func (cmd *Cmd) execStreamCmd(c *ishell.Context) {
	tweets, sinceId, err := cmd.fetchHomeTimelineTweets("25", "")
	if err != nil {
		cmd.showErrorMessage(err.Error())
		return
	}

	// タイムラインを表示
	cmd.view.ShowTweets(*tweets, false)

	// Ctrl+Cでの割り込みを通知
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go cmd.startUserStream(sinceId, quit)

	// 割り込みがあったら終了
	<-quit
	signal.Stop(quit)
	close(quit)
}

// startUserStream  疑似UserStreamを開始
func (cmd *Cmd) startUserStream(startSinceId string, quit chan os.Signal) {
	var (
		accumulateTweets AccumulateTweets
		err              error
		sinceId          = startSinceId
	)

	for {
		for i := 0; i < 60; i++ {
			select {
			case <-quit:
				// 中断
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
		return nil, sinceId, err
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
		count = cmd.config.Option.Counts
	}

	query := twitter.CreateQuery(count)
	if sinceId != "" {
		query.Add("since_id", sinceId)
	}

	tweets, err := cmd.twitter.FetchTimelineTweets(twitter.Home, query)
	if err != nil {
		return nil, sinceId, err
	}

	newSinceId := sinceId
	if len(*tweets) > 0 {
		newSinceId = (*tweets)[0].IdStr
	}

	return tweets, newSinceId, nil
}
