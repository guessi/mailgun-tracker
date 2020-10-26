package main

import (
	"github.com/gin-gonic/gin"
	"github.com/guessi/mailgun-tracker/config"
	"github.com/guessi/mailgun-tracker/pkg/mailgun"
	"github.com/robfig/cron/v3"
)

var cfg config.Config

func main() {
	cfg = config.LoadConfig()

	// cron
	cron := cron.New()
	_, err := cron.AddFunc(cfg.CronConfig.CheckPeriod, func() {
		mailgun.CheckBounce(cfg.MailgunConfig, cfg.SlackConfig)
	})
	if err == nil {
		// FIXME: should handle error here
		cron.Start()
	}

	// http server setup
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.POST("/health", mailgun.HealthHandler)
		v1.POST("/mailgun/permanent-failure", func(ctx *gin.Context) {
			mailgun.PermanentFailureHandler(ctx, cfg.MailgunConfig, cfg.SlackConfig)
		})
	}

	//nolint
	r.Run(":8080")
}
