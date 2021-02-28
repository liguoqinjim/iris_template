package datasource

import "github.com/robfig/cron/v3"

func initCron() {
	Cron = cron.New(cron.WithSeconds())

	Cron.Start()
}
