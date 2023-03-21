package kafkaop

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/cloudwego/kitex/pkg/klog"
	"strconv"
	"time"
)

var Client sarama.SyncProducer

// 消息写入kafka
func init() {
	var err error

	//初始化配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	//生产者
	Client, err = sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		fmt.Println("producer close,err:", err)
		return
	}

}

func ProduceComment(ctx context.Context, userId int64, videoId int64, actionType int32, comment_text string, comment_id string) error {
	klog.CtxInfof(ctx, "[ProduceComment]: userId==%v, videoId==%v, actionType==%v\n", userId, videoId, actionType)
	//创建消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "dousheng_comment"
	msg.Value = sarama.StringEncoder(
		strconv.FormatInt(userId, 10) +
			"::" +
			strconv.FormatInt(videoId, 10) +
			"::" +
			strconv.FormatInt(int64(actionType), 10) +
			"::" +
			comment_text +
			"::" +
			comment_id +
			"::" +
			strconv.FormatInt(time.Now().Unix(), 10),
	)
	//发送消息
	partition, offset, err := Client.SendMessage(msg)
	fmt.Println("producer sent msg, partition:%v, offset:%v", partition, offset)
	return err
}
