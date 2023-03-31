package handler

import (
	"context"
	"crawler/pkg/crawler"
	"crawler/pkg/utils"
	"github.com/carlescere/scheduler"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type CronJobHandler struct {
	CronJobInfo      map[string]map[string]interface{}
	CronJob          map[string]scheduler.Job
	CronScheduleInfo map[string]map[string]interface{}
	CronSchedule     map[string]*time.Timer
	CrawlHandler     crawler.CrawlHandler
	Ctx              context.Context
}

func NewCronJobHandlers(ctx context.Context, handler crawler.CrawlHandler) *CronJobHandler {
	return &CronJobHandler{Ctx: ctx, CrawlHandler: handler}
}

func (h *CronJobHandler) StartCron() {
	Job := make(map[string]scheduler.Job)
	JobDetail := make(map[string]map[string]interface{})

	JobTask := func() {
		//var err error
		log.Println("================= Start the job ===================")
		h.CrawlHandler.Start(h.Ctx, utils.SITES)
		log.Println("================= End the job ===================")
	}

	job, err := scheduler.Every(24).Hours().Run(JobTask)
	if err != nil {
		log.Println("While starting cron got error : " + err.Error())
	}
	Job["ADMIN"] = *job
	CronJobDetail := make(map[string]interface{})
	CronJobDetail["time_start"] = time.Now().Format("2006-01-02 15:04:05")
	CronJobDetail["frequency"] = "Every 24 hours"
	CronJobDetail["description"] = "Crawl data from blog"

	JobDetail["ADMIN"] = CronJobDetail
	h.CronJob = Job
	h.CronJobInfo = JobDetail

}
func (h *CronJobHandler) ListCronJobInfo(c *gin.Context) {

}

func (h *CronJobHandler) ListCronScheduleInfo(r *gin.Context) {

}

func (h *CronJobHandler) ClearAllCronFunction() {
	for key, element := range h.CronJob {
		if key != "ADMIN" {
			element.Quit <- true
			delete(h.CronJob, key)
			delete(h.CronJobInfo, key)
		}
	}
}

func (h *CronJobHandler) ClearAllScheduleFunction() {
	for key, element := range h.CronSchedule {
		if key != "ADMIN" {
			element.Stop()
			delete(h.CronSchedule, key)
			delete(h.CronScheduleInfo, key)
		}
	}
}
