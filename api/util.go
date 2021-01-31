package api

import (
	"fmt"
	"net/url"
	"regexp"
)

// CreateURLValues クエリを作成
func CreateURLValues(cnt string) url.Values {
	v := url.Values{}
	v.Add("tweet_mode", "extended")
	v.Add("count", cnt)
	return v
}

func parseAPIError(err error) string {
	bytes := []byte(err.Error())
	errMsg := regexp.MustCompile("\"(message|error)\":\\s*\"(.+)\"").FindSubmatch(bytes)
	return fmt.Sprint(errMsg[2])
}
