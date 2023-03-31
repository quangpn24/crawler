package main

import (
	"context"
	"crawler/conf"
	"crawler/pkg/crawler"
	"crawler/pkg/handler"
	"crawler/pkg/repo"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

func main() {
	conf.SetEnv()
	//connect with PostgresSQL
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s connect_timeout=5 sslmode=disable",
		conf.GetConfig().DBHost,
		conf.GetConfig().DBPort,
		conf.GetConfig().DBUser,
		conf.GetConfig().DBName,
		conf.GetConfig().DBPass,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		panic(err)
	}
	ctx := context.Background()

	r := repo.NewPGRepo(db)
	crawlHandler := crawler.NewCrawlHandler(r)
	cronjob := handler.NewCronJobHandlers(ctx, *crawlHandler)
	cronjob.StartCron()

	migration := handler.NewMigrationHandler(db)

	route := gin.Default()
	route.POST("/internal/migrate", migration.Migrate)
	route.Run(":8088")

	os.Clearenv()
}
