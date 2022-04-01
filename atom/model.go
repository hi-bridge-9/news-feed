package atom

type Feed struct {
	Title   string `xml:"title,omitempty"`
	ID      string `xml:"id,omitempty"`
	Updated string `xml:"updated,omitempty"`
	// Subtitle      string      `xml:"subtitle,omitempty"`
	// Links         []*Link     `xml:"links,omitempty"`
	Language string `xml:"language,omitempty"`
	// Generator     *Generator  `xml:"generator,omitempty"`
	// Icon          string      `xml:"icon,omitempty"`
	// Logo          string      `xml:"logo,omitempty"`
	// Rights        string      `xml:"rights,omitempty"`
	// Contributors  []*Person   `xml:"contributors,omitempty"`
	// Authors       []*Person   `xml:"authors,omitempty"`
	// Categories []*Category `xml:"categories,omitempty"`
	Entries    []*Entry    `xml:"entries"`
	Version    string      `xml:"version"`
}

// Entry is an Atom Entry
type Entry struct {
	Title   string `xml:"title,omitempty"`
	ID      string `xml:"id,omitempty"`
	Updated string `xml:"updated,omitempty"`
	// Summary string `xml:"summary,omitempty"`
	// Authors         []*Person   `xml:"authors,omitempty"`
	// Contributors    []*Person   `xml:"contributors,omitempty"`
	// Categories []*Category `xml:"categories,omitempty"`
	Links      []*Link     `xml:"links,omitempty"`
	// Rights          string      `xml:"rights,omitempty"`
	Published string `xml:"published,omitempty"`
	// Source          *Source     `xml:"source,omitempty"`
	// Content         *Content    `xml:"content,omitempty"`
}

// type Category struct {
// 	Term   string `xml:"term,omitempty"`
// 	Scheme string `xml:"scheme,omitempty"`
// 	Label  string `xml:"label,omitempty"`
// }

// type Person struct {
// 	Name  string `xml:"name,omitempty"`
// 	Email string `xml:"email,omitempty"`
// 	URI   string `xml:"uri,omitempty"`
// }

type Link struct {
	Href     string `xml:"href,omitempty"`
	// Hreflang string `xml:"hreflang,omitempty"`
	// Rel      string `xml:"rel,omitempty"`
	// Type     string `xml:"type,omitempty"`
	// Title    string `xml:"title,omitempty"`
	// Length   string `xml:"length,omitempty"`
}

// type Content struct {
// 	Src   string `xml:"src,omitempty"`
// 	Type  string `xml:"type,omitempty"`
// 	Value string `xml:"value,omitempty"`
// }

// type Generator struct {
// 	Value   string `xml:"value,omitempty"`
// 	URI     string `xml:"uri,omitempty"`
// 	Version string `xml:"version,omitempty"`
// }

// type Source struct {
// 	Title         string      `xml:"title,omitempty"`
// 	ID            string      `xml:"id,omitempty"`
// 	Updated       string      `xml:"updated,omitempty"`
// 	UpdatedParsed *time.Time  `xml:"updatedParsed,omitempty"`
// 	Subtitle      string      `xml:"subtitle,omitempty"`
// 	Links         []*Link     `xml:"links,omitempty"`
// 	Generator     *Generator  `xml:"generator,omitempty"`
// 	Icon          string      `xml:"icon,omitempty"`
// 	Logo          string      `xml:"logo,omitempty"`
// 	Rights        string      `xml:"rights,omitempty"`
// 	Contributors  []*Person   `xml:"contributors,omitempty"`
// 	Authors       []*Person   `xml:"authors,omitempty"`
// 	Categories    []*Category `xml:"categories,omitempty"`
// }
