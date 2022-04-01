package rss

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Language    string `xml:"language,omitempty"`
	// Copyright           string      `xml:"copyright,omitempty"`
	// ManagingEditor      string      `xml:"managingEditor,omitempty"`
	// WebMaster           string      `xml:"webMaster,omitempty"`
	PubDate       string `xml:"pubDate,omitempty"`
	LastBuildDate string `xml:"lastBuildDate,omitempty"`
	// Categories          []*Category `xml:"categories,omitempty"`
	// Generator           string      `xml:"generator,omitempty"`
	// Docs                string      `xml:"docs,omitempty"`
	// Cloud               *Cloud      `xml:"cloud,omitempty"`
	// TTL                 string      `xml:"ttl,omitempty"`
	// Image               *Image      `xml:"image,omitempty"`
	// Rating              string      `xml:"rating,omitempty"`
	// SkipHours           []string    `xml:"skipHours,omitempty"`
	// SkipDays            []string    `xml:"skipDays,omitempty"`

	// TextInput *TextInput `xml:"textInput,omitempty"`
	Items   []*Item `xml:"items"`
	Version string  `xml:"version"`
}

type Item struct {
	Title       string `xml:"title,omitempty"`
	Link        string `xml:"link,omitempty"`
	Description string `xml:"description,omitempty"`
	// Content     string `xml:"content,omitempty"`
	// Author      string `xml:"author,omitempty"`
	// Categories    []*Category       `xml:"categories,omitempty"`
	// Comments      string            `xml:"comments,omitempty"`
	// Enclosure     *Enclosure        `xml:"enclosure,omitempty"`
	GUID    string `xml:"guid,omitempty"`
	PubDate string `xml:"pubDate,omitempty"`
	// Source        *Source           `xml:"source,omitempty"`
	// Custom        map[string]string `xml:"custom,omitempty"`
}

// type Image struct {
// 	URL         string `xml:"url,omitempty"`
// 	Link        string `xml:"link,omitempty"`
// 	Title       string `xml:"title,omitempty"`
// 	Width       string `xml:"width,omitempty"`
// 	Height      string `xml:"height,omitempty"`
// 	Description string `xml:"description,omitempty"`
// }

// type Enclosure struct {
// 	URL    string `xml:"url,omitempty"`
// 	Length string `xml:"length,omitempty"`
// 	Type   string `xml:"type,omitempty"`
// }

// type Source struct {
// 	Title string `xml:"title,omitempty"`
// 	URL   string `xml:"url,omitempty"`
// }

// type Category struct {
// 	Domain string `xml:"domain,omitempty"`
// 	Value  string `xml:"value,omitempty"`
// }

// type TextInput struct {
// 	Title       string `xml:"title,omitempty"`
// 	Description string `xml:"description,omitempty"`
// 	Name        string `xml:"name,omitempty"`
// 	Link        string `xml:"link,omitempty"`
// }

// type Cloud struct {
// 	Domain            string `xml:"domain,omitempty"`
// 	Port              string `xml:"port,omitempty"`
// 	Path              string `xml:"path,omitempty"`
// 	RegisterProcedure string `xml:"registerProcedure,omitempty"`
// 	Protocol          string `xml:"protocol,omitempty"`
// }
