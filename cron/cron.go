package main

import (
	"gin-Blog/service"
	"log"
	"time"

	"github.com/robfig/cron"
)

func main() {
	log.Println("Starting...")

	c := cron.New()
	err := c.AddFunc("* * * * * *", func() {
		log.Println("Run service.CleanAllTag...")
		service.CronTag()
	})
	if err != nil {
		log.Fatalln(err)
	}
	err = c.AddFunc("* * * * * *", func() {
		log.Println("Run service.CleanAllArticle...")
		service.CronArticle()
	})
	if err != nil {
		log.Fatalln(err)
	}
	c.Start()
	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
