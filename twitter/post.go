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

	"golang.org/x/sync/errgroup"
)

// PostTweet ツイートを投稿
func (tw *Twitter) PostTweet(query url.Values, text string) (string, error) {
	tweet, err := tw.API.PostTweet(text, query)
	if err != nil {
		return "", tw.createAPIError("", err)
	}

	return tweet.FullText, nil
}

// DeleteTweet ツイートを削除
func (tw *Twitter) DeleteTweet(tweetId string) (string, error) {
	id, _ := strconv.ParseInt(tweetId, 10, 64)

	tweet, err := tw.API.DeleteTweet(id, true)
	if err != nil {
		return "", tw.createAPIError("", err)
	}

	return tweet.FullText, nil
}

// UploadImage 画像をアップロード
func (tw *Twitter) UploadImage(images []string) (string, error) {
	fileNum := len(images)

	// 画像数が4枚を超えるならエラー
	if fileNum > 4 {
		return "", errors.New("you can attach up to 4 images")
	}

	eg, ctx := errgroup.WithContext(context.Background())
	ch := make(chan string, fileNum)

	for _, filename := range images {
		// サポート外の拡張子ならエラー
		ext := strings.ToLower(path.Ext(filename))
		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".gif" {
			return "", fmt.Errorf("unsupported extensions (%s)", ext)
		}

		// アップロード処理
		filename := filename
		eg.Go(func() error {
			select {
			case <-ctx.Done():
				return nil
			default:
				data, err := ioutil.ReadFile(filename)
				if err != nil {
					return fmt.Errorf("failed to load file (%s)", filename)
				}

				// base64にエンコードしてアップロード
				enc := base64.StdEncoding.EncodeToString(data)
				media, err := tw.API.UploadMedia(enc)
				if err != nil {
					return fmt.Errorf("upload failed (%s)", filename)
				}

				ch <- media.MediaIDString
				return nil
			}
		})
	}

	if err := eg.Wait(); err != nil {
		return "", err
	}

	close(ch)

	// mediaIDをカンマで連結
	mediaIds := ""
	for id := range ch {
		mediaIds += id + ","
	}

	return mediaIds, nil
}
