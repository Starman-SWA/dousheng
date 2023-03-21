package redisop

import (
	"dousheng/dal/db"
	"github.com/robfig/cron"
	"time"
)

func WriteRedisToDB() error {
	commentModify := GetAndDeleteAllComments()
	db.MSetComments(commentModify)
	return nil
}

func RunCronJob() {
	c := cron.New()
	err := c.AddFunc("*/10 * * * * ?", func() {
		err := WriteRedisToDB()
		if err != nil {
			panic(err)
		}
	})
	if err != nil {
		panic(err)
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
