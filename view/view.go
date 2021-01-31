package view

import (
	"github.com/ChimeraCoder/anaconda"
	"github.com/arrow2nd/twnyan/config"
)

// View 表示
type View struct {
	tweets []anaconda.Tweet
	cfg    *config.Config
}

// New 表示
func New(c *config.Config) *View {
	v := &View{
		tweets: []anaconda.Tweet{},
		cfg:    c,
	}
	return v
}
