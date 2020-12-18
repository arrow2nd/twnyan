package twitter

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"path"
	"strconv"
	"strings"

	"github.com/arrow2nd/twnyan/util"
	"github.com/gookit/color"
	"golang.org/x/sync/errgroup"
)

// PostTweet ツイートを投稿する
func PostTweet(status, replyToID string, files []string) {
	// 画像をアップロード
	val, err := upload(files)
	if err != nil {
		color.Error.Prompt("%s", err)
		return
	}

	// リプライ先設定
	if replyToID != "" {
		val.Add("in_reply_to_status_id", replyToID)
		val.Add("auto_populate_reply_metadata", "true")
	}

	// ツイート
	tweet, err := api.PostTweet(status, val)
	if err != nil {
		showAPIErrorString(err)
		return
	}

	util.ShowSuccessMsg("Tweeted", tweet.FullText, cfg.Color.BoxFg, cfg.Color.Accent3)
}

// upload 画像をアップロードする
func upload(files []string) (url.Values, error) {
	var (
		val      = url.Values{}
		fileNum  = len(files)
		mediaIds = make([]string, fileNum)
	)

	// ファイル数チェック
	if fileNum == 0 {
		return val, nil
	} else if fileNum > 4 {
		return nil, errors.New("Up to 4 media can be attached to a tweet")
	}

	eg, ctx := errgroup.WithContext(context.Background())

	for idx, fn := range files {
		// 未対応の拡張子ならスキップ
		if ext := strings.ToLower(path.Ext(fn)); ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".gif" {
			return nil, errors.New("Unsupported extensions (" + ext + ")")
		}

		idx, fn := idx, fn

		eg.Go(func() error {
			select {
			case <-ctx.Done():
				return nil
			default:
				// ファイルを読み込む
				data, err := ioutil.ReadFile(fn)
				if err != nil {
					return errors.New("Failed to load (" + fn + ")")
				}

				// base64にエンコードしてアップロード
				enc := base64.StdEncoding.EncodeToString(data)
				media, err := api.UploadMedia(enc)
				if err != nil {
					return errors.New("Upload failed (" + fn + ")")
				}

				mediaIds[idx] = media.MediaIDString
				return nil
			}
		})
	}

	fmt.Println("Uploading...🐾")

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	val.Add("media_ids", strings.Join(mediaIds, ","))
	return val, nil
}

// DeleteTweet ツイートを削除する
func DeleteTweet(tweetID string) {
	id, _ := strconv.ParseInt(tweetID, 10, 64)
	tweet, err := api.DeleteTweet(id, true)
	if err != nil {
		showAPIErrorString(err)
		return
	}

	util.ShowSuccessMsg("Deleted", tweet.FullText, cfg.Color.BoxFg, cfg.Color.Accent3)
}
