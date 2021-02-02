package api

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
func (ta *TwitterAPI) PostTweet(val url.Values, status string) (string, error) {
	tweet, err := ta.API.PostTweet(status, val)
	if err != nil {
		return "", errors.New(parseAPIError(err))
	}
	return tweet.FullText, nil
}

// DeleteTweet ツイートを削除
func (ta *TwitterAPI) DeleteTweet(tweetID string) (string, error) {
	id, _ := strconv.ParseInt(tweetID, 10, 64)
	tweet, err := ta.API.DeleteTweet(id, true)
	if err != nil {
		return "", errors.New(parseAPIError(err))
	}
	return tweet.FullText, nil
}

// UploadImage 画像をアップロード
func (ta *TwitterAPI) UploadImage(files []string) (string, error) {
	fileNum := len(files)
	mediaIds := make([]string, fileNum)

	// ファイル数チェック
	if fileNum > 4 {
		return "", errors.New("You can attach up to 4 images")
	}

	eg, ctx := errgroup.WithContext(context.Background())

	for idx, filename := range files {
		// 拡張子をチェック
		if ext := strings.ToLower(path.Ext(filename)); ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".gif" {
			return "", fmt.Errorf("Unsupported extensions (%s)", ext)
		}

		// アップロード処理
		idx, filename := idx, filename
		eg.Go(func() error {
			select {
			case <-ctx.Done():
				return nil
			default:
				// ファイルを読み込む
				data, err := ioutil.ReadFile(filename)
				if err != nil {
					return fmt.Errorf("Failed to load file (%s)", filename)
				}
				// base64にエンコードしてアップロード
				enc := base64.StdEncoding.EncodeToString(data)
				media, err := ta.API.UploadMedia(enc)
				if err != nil {
					return fmt.Errorf("Upload failed (%s)", filename)
				}
				mediaIds[idx] = media.MediaIDString
				return nil
			}
		})
	}

	// 待機
	if err := eg.Wait(); err != nil {
		return "", err
	}

	// カンマで連結
	return strings.Join(mediaIds, ","), nil
}
