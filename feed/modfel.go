package feed

import (
	"time"

	"github.com/mmcdole/gofeed"
)

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
