package app

import (
	"github.com/robfig/cron/v3"
)

func InitCronJobs() {
	c := cron.New(cron.WithSeconds()) // 支持到秒
	// 每 30 秒调用一次外部 API
	c.AddFunc("*/30 * * * * *", CallExternalAPI)
	c.Start()
}
