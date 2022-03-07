package app

import (
	"context"
	"golang.org/x/net/html"
	"net/http"
	"simpleCrawler/internal/generated/models"
	apiPosturLs "simpleCrawler/internal/generated/restapi/operations/post_u_r_ls"
	"sync"

	"github.com/go-openapi/runtime/middleware"
)

func (srv *Service) PostURLsHandler(params apiPosturLs.PostURLsParams) middleware.Responder {
	var (
		ctx = context.Background()
	)

	return apiPosturLs.NewPostURLsOK().WithPayload(mapPostCrawlerResponse(ctx, params))
}

func mapPostCrawlerResponse(ctx context.Context, params apiPosturLs.PostURLsParams) *models.CrawlerResponse {
	return GetURLTitles(params.CrawlerRequest.Urls)
}

func GetURLTitles(urls []string) *models.CrawlerResponse {
	var result models.CrawlerResponse
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go getTitle(url, &wg, &result)
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
