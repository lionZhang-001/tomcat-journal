package main

import (
	"github.com/robfig/cron"
)

func main() {

	//定时任务
	c := cron.New()
	c.AddFunc("0 */5 * * * *", job)

	c.Start()

	select {}

}