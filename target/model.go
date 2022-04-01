package target

type Info struct {
	Category string `json:"category"`
	Sites    []Site `json:"sites"`
}

type Site struct {
	Name    string `json:"name"`
	TopURL  string `json:"top_url"`
	FeedURL string `json:"feed_url"`
}
