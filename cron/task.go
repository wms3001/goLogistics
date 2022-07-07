package cron

import (
	"github.com/robfig/cron/v3"
	"logistics/service"
)

func TaskStart() {
	c := cron.New()
	c.AddFunc("*/1 * * * *", service.StatisticsInChannel)
	c.AddFunc("*/1 * * * *", service.StatisticsOutChannel)
	//c.AddFunc("@hourly",      func() { fmt.Println("Every hour") })
	//c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty") })
	c.Start()

}
