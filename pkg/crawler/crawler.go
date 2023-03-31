package crawler

import (
	"context"
	"crawler/pkg/repo"
	"crawler/pkg/utils"
	"fmt"
)

type CrawlHandler struct {
	repo repo.PGInterface
}

func NewCrawlHandler(repo repo.PGInterface) *CrawlHandler {
	return &CrawlHandler{repo: repo}
}

func (c *CrawlHandler) Start(ctx context.Context, infos []utils.InfoSite) {
	for _, value := range infos {
		switch value.SiteName {
		case utils.SITE_NAME_200LAB:
			c.Crawl200labData(value.URL)
		case utils.SITE_NAME_VIBLO:
			c.CrawlVibloData(ctx, value.URL)
		default:
			fmt.Println("Currently we do not support this site yet!")
		}
	}
}
