package crawler

import (
	"golang.org/x/net/html"
	"net/http"
	"simpleCrawler/internal/generated/models"
	"sync"
)

func GetURLTitles(urls []string) *models.CrawlerResponse {
	wg := &sync.WaitGroup{}
	var result models.CrawlerResponse
	for _, url := range urls {
		wg.Add(1)
		go getTitle(url, wg, &result)
	}
	wg.Wait()
	return &result
}

func getTitle(url string, wg *sync.WaitGroup, result *models.CrawlerResponse) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	tkn := html.NewTokenizer(resp.Body)
	var isTitle bool
	for {
		tt := tkn.Next()
		switch {
		case tt == html.ErrorToken:
			return
		case tt == html.StartTagToken:
			t := tkn.Token()
			isTitle = t.Data == "title"
		case tt == html.TextToken:
			t := tkn.Token()
			if isTitle {
				result.Items = append(result.Items, t.Data)
				return
			}
		}
	}
}
