package kafkaop

import (
	"dousheng/cmd/comment/redisop"
	"fmt"
	"github.com/Shopify/sarama"
	"strconv"
	"strings"
	"sync"
)

var consumer sarama.Consumer
var wg sync.WaitGroup

func init() {
	var err error

	consumer, err = sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		fmt.Println("consumer connect err:", err)
		return
	}
}

func ConsumeComment() {
	//获取 kafka 主题
	partitions, err := consumer.Partitions("dousheng_comment")
	if err != nil {
		fmt.Println("get partitions failed, err:", err)
		return
	}

	for _, p := range partitions {
		//sarama.OffsetNewest：从当前的偏移量开始消费，sarama.OffsetOldest：从最老的偏移量开始消费
		partitionConsumer, err := consumer.ConsumePartition("dousheng_comment", p, sarama.OffsetNewest)
		if err != nil {
			fmt.Println("partitionConsumer err:", err)
			continue
		}
		wg.Add(1)
		go func() {
			for m := range partitionConsumer.Messages() {
				fmt.Printf("key: %s, text: %s, offset: %d\n", string(m.Key), string(m.Value), m.Offset)
				tokens := strings.Split(string(m.Value), "::")
				userId, _ := strconv.ParseInt(tokens[0], 10, 64)
				videoId, _ := strconv.ParseInt(tokens[1], 10, 64)
				if tokens[2] == "1" {
					// 发布
					err := redisop.Comment(userId, videoId, tokens[3], tokens[5])
					if err != nil {
						panic(err)
					}
				} else if tokens[2] == "2" {
					// 删除
					err := redisop.UnComment(userId, videoId, tokens[4], tokens[5])
					if err != nil {
						panic(err)
					}
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
