package crawler

import (
	"context"
	"crawler/pkg/model"
	"crawler/pkg/repo"
	"crawler/pkg/utils"
	"fmt"
	"github.com/gocolly/colly"
	"strconv"
	"strings"
)

func (c *CrawlHandler) CrawlVibloData(ctx context.Context, url string) {
	collector := colly.NewCollector()
	var numOfPage int
	collector.OnHTML(".pagination", func(element *colly.HTMLElement) {
		e := element.DOM.Find("li:not(.page-item)")
		numOfPage, _ = strconv.Atoi(strings.TrimSpace(e.Last().Text()))
	})
	collector.Visit(url)

	jobsChan := make(chan string, numOfPage*utils.PAGE_SIZE_VIBLO)
	quit := make(chan bool)

	// assign job for worker
	for i := 1; i <= utils.NUM_OF_WORKER; i++ {
		go func(i int, jobs <-chan string, quit <-chan bool) {
			for {
				select {
				case url, _ := <-jobs:
					CrawlBlog(ctx, url, c.repo)
				case <-quit:
					return
				}
			}
		}(i, jobsChan, quit)
	}

	//add url into pool
	clt := colly.NewCollector()
	clt.OnRequest(func(request *colly.Request) {
		fmt.Println("Visting:", request.URL)
	})
	//get post's url
	clt.OnHTML(".post-feed > .post-feed-item .link", func(element *colly.HTMLElement) {
		jobsChan <- utils.URL_VIBLO + element.Attr("href")
	})
	
	for i := 1; i <= numOfPage; i++ {
		newUrl := fmt.Sprintf("%s?page=%d", url, i)
		clt.Visit(newUrl)
	}

	quit <- true
	close(jobsChan)
}
func CrawlBlog(ctx context.Context, url string, repo repo.PGInterface) {
	clt := colly.NewCollector()
	clt.OnRequest(func(request *colly.Request) {
		fmt.Println("Read blog: " + url)
	})
	clt.OnHTML("article.post-content", func(element *colly.HTMLElement) {
		post := model.Post{
			URLOrigin: url,
		}
		post.HTMLData, _ = element.DOM.Find(".md-contents").Html()
		post.Title = element.ChildText(".article-content__title")
		post.Source = utils.SITE_NAME_VIBLO
		post.User = element.ChildText(".post-author__info a")
		//inert into DB
		repo.InsertPost(ctx, post, nil)
	})
	clt.Visit(url)
}
