package handlers

import (
	"net/http"
	"io/ioutil"
	"encoding/xml"
	"../model"
	"html/template"
	"sync"
)

var wg sync.WaitGroup

func newsFetchFromLinks(c chan model.News, Location string) {
	defer wg.Done()
	var n model.News

	resp, _ := http.Get(Location)
	bytes, _ := ioutil.ReadAll(resp.Body)
	xml.Unmarshal(bytes, &n)
	resp.Body.Close()

	c <- n
}

func NewsAggHandler(w http.ResponseWriter, r *http.Request) {
	var s model.SitemapIndex
	newsMap := make(map[string]model.NewsMap)

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	// string_body := string(bytes)
	// fmt.Println(string_body)
	xml.Unmarshal(bytes, &s)
	resp.Body.Close()
	queue := make(chan model.News, 30)

	for _, Location := range s.Locations {
		wg.Add(1)
		go newsFetchFromLinks(queue, Location)
	}
	wg.Wait()
	close(queue)

	for elem := range queue {
		for i := range elem.Titles {
			newsMap[elem.Titles[i]] = model.NewsMap{Keyword: elem.Keywords[i], Location: elem.Locations[i]}
		}
	}

	p := model.NewsAggPage{Title: "Washington Post News Aggregator", News: newsMap}
	t, _ := template.ParseFiles("templates/newsaggtemplate.html")
	t.Execute(w, p)
}
