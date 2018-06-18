package handlers

import (
	"net/http"
	"io/ioutil"
	"encoding/xml"
	"../model"
	"html/template"
)

func NewsAggHandler(w http.ResponseWriter, r *http.Request) {
	var s model.SitemapIndex
	var n model.News
	newsMap := make(map[string]model.NewsMap)

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	// string_body := string(bytes)
	// fmt.Println(string_body)
	resp.Body.Close()
	xml.Unmarshal(bytes, &s)

	for _, loc := range s.Locations {
		resp, _ := http.Get(loc)
		bytes, _ := ioutil.ReadAll(resp.Body)
		xml.Unmarshal(bytes, &n)

		for i, _ := range n.Titles {
			newsMap[n.Titles[i]] = model.NewsMap{n.Keywords[i], n.Locations[i]}
		}
	}

	p := model.NewsAggPage{Title: "Washington Post News Aggregator", News: newsMap}
	t, _ := template.ParseFiles("templates/newsaggtemplate.html")
	t.Execute(w, p)
}
