package main

import (
	"log"
	"simpleCrawler/internal/generated/restapi"
	"simpleCrawler/internal/generated/restapi/operations"

	"simpleCrawler/internal/app"
	"simpleCrawler/internal/config"
	apiPosturLs "simpleCrawler/internal/generated/restapi/operations/post_crawlers"

	"github.com/go-openapi/loads"
)

func main() {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	srv := app.New()
	api := operations.NewSimpleCrawlerAPI(swaggerSpec)

	api.PostCrawlersPostCrawlersHandler = apiPosturLs.PostCrawlersHandlerFunc(srv.PostURLsHandler)
	api.ServerShutdown = srv.OnShutdown
	server := restapi.NewServer(api)
	defer server.Shutdown()

	cfg, err := config.InitConfig("simple_crawler")
	if err != nil {
		log.Fatalln(err)
	}

	server.ConfigureAPI()

	server.Port = cfg.HTTPBindPort
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
