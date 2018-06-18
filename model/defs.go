package model


// SiteMap Struct
type SitemapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

// Links inside the sitemap Struct
type News struct {
	Titles     []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	Locations []string `xml:"url>loc"`
}

// Map structure of the Sitemap
type NewsMap struct {
	Keyword string
	Location string
}

// Struct for Mapping Strign to NewsMap objects
type NewsAggPage struct {
	Title string
	News map[string]NewsMap
}