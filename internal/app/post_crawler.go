package app

import (
	"context"

	"github.com/go-openapi/runtime/middleware"
	"simpleCrawler/internal/generated/models"

	apiPosturLs "simpleCrawler/internal/generated/restapi/operations/post_u_r_ls"
	"simpleCrawler/internal/handlers/crawler"
)

func (srv *Service) PostURLsHandler(params apiPosturLs.PostURLsParams) middleware.Responder {
	var (
		ctx = context.Background()
	)

	return apiPosturLs.NewPostURLsOK().WithPayload(mapPostCrawlerResponse(ctx, params))
}

func mapPostCrawlerResponse(ctx context.Context, params apiPosturLs.PostURLsParams) *models.CrawlerResponse {
	return crawler.GetURLTitles(params.CrawlerRequest.Urls)
}
