package feed

import "time"

type Feed struct {
	Title       string
	Description string
	Link        string
	FeedLink    string
	// Links           []string
	Updated         string
	UpdatedParsed   *time.Time
	// Author          *Person
	// Authors         []*Person
	// Language        string
	// Image           *Image
	// Copyright       string
	// Generator       string
	// Categories      []string
	// Custom          map[string]string
	Items   []Item
	Type    string
	Version string
}

type Item struct {
	Title       string
	Description string
	// Content         string
	Link string
	// Links           []string
	// Updated         string
	// UpdatedParsed   *time.Time
	Published       string
	PublishedParsed *time.Time
	// Author          *Person
	// Authors         []*Person
	GUID string
	// Image           *Image
	// Categories      []string
	// Enclosures      []*Enclosure
	// Custom          map[string]string
}

// type Person struct {
// 	Name  string
// 	Email string
// }

// type Image struct {
// 	URL   string
// 	Title string
// }

// type Enclosure struct {
// 	URL    string
// 	Length string
// 	Type   string
// }


// 新しいニュース用

type Output struct {
	SiteTitle string
	SiteURL   string
	Articles  []Item
}

type Article struct {
	Title           string
	URL             string
	UpdatedParsed   *time.Time
	PublishedParsed *time.Time
}


// ソート用構造体

type ByPublishedParsed []Item
