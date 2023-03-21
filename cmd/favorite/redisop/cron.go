package redisop

import (
	"dousheng/dal/db"
	"github.com/robfig/cron"
	"time"
)

func WriteRedisToDB() error {
	err, videoLikeCounts, userVideoLikes := GetAndDeleteAllLikes()
	if err != nil {
		return err
	}
	err = db.MSetLikes(videoLikeCounts, userVideoLikes)
	if err != nil {
		return err
	}
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
