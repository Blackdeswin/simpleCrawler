package app

import (
	"github.com/go-openapi/runtime/middleware"
	"simpleCrawler/internal/generated/models"

	apiPostCrawlers "simpleCrawler/internal/generated/restapi/operations/post_crawlers"
	"simpleCrawler/internal/handlers/crawler"
)

func (srv *Service) PostURLsHandler(params apiPostCrawlers.PostCrawlersParams) middleware.Responder {
	return apiPostCrawlers.NewPostCrawlersOK().WithPayload(mapPostCrawlerResponse(params))
}

func mapPostCrawlerResponse(params apiPostCrawlers.PostCrawlersParams) *models.CrawlerResponse {
	return crawler.GetURLTitles(params.CrawlerRequest.Urls)
}
