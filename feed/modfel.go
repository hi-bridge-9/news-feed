package feed

import (
	"time"

	"github.com/mmcdole/gofeed"
)

// 入力用

type Tartget struct {
	Category string `json:"category"`
	Sites    []Site `json:"sites"`
}

type Site struct {
	Name    string `json:"name"`
	TopURL  string `json:"top_url"`
	FeedURL string `json:"feed_url"`
}

// 出力用

type News struct {
	SiteTitle string
	SiteURL   string
	Articles  []gofeed.Item
}

type Article struct {
	Title           string
	URL             string
	UpdatedParsed   *time.Time
	PublishedParsed *time.Time
}

// ソート用構造体

type ByPublishedParsed []*gofeed.Item

