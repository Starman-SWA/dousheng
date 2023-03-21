package redisop

import (
	"dousheng/dal/db"
	"dousheng/pkg/configs/sqlmodel"
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"strings"
	"time"
)

// RedisSingleObj 定义一个RedisSingleObj结构体
type RedisSingleObj struct {
	RedisHost string
	RedisPort uint16
	RedisAuth string
	Database  int
	Db        *redis.Client
}

var conn *RedisSingleObj

func init() {
	// 实例化RedisSingleObj结构体
	conn = &RedisSingleObj{
		RedisHost: "localhost",
		RedisPort: 6379,
		RedisAuth: "yourpass",
	}

	// 初始化连接 Single Redis 服务端
	err := conn.InitSingleRedis()
	if err != nil {
		panic(err)
	}

	go RunCronJob()
}

// InitSingleRedis 结构体InitSingleRedis方法: 用于初始化redis数据库
func (r *RedisSingleObj) InitSingleRedis() (err error) {
	// Redis连接格式拼接
	redisAddr := fmt.Sprintf("%s:%d", r.RedisHost, r.RedisPort)
	// Redis 连接对象: NewClient将客户端返回到由选项指定的Redis服务器。
	r.Db = redis.NewClient(&redis.Options{
		Addr:        redisAddr,   // redis服务ip:port
		Password:    r.RedisAuth, // redis的认证密码
		DB:          r.Database,  // 连接的database库
		IdleTimeout: 300,         // 默认Idle超时时间
		PoolSize:    100,         // 连接池
	})
	fmt.Printf("Connecting Redis : %v\n", redisAddr)

	// 验证是否连接到redis服务端
	res, err := r.Db.Ping().Result()
	if err != nil {
		fmt.Printf("Connect Failed! Err: %v\n", err)
		return err
	} else {
		fmt.Printf("Connect Successful! Ping => %v\n", res)
		return nil
	}
}

func getKeyFromUserAndVideo(userId int64, videoId int64) string {
	return strconv.FormatInt(userId, 10) + "::" + strconv.FormatInt(videoId, 10)
}

func genCommentMsg(userId int64, videoId int64, comment_text string, timestamp string) string {
	return strconv.FormatInt(userId, 10) + "::" + strconv.FormatInt(videoId, 10) + "::1::" + comment_text + "::" + timestamp
}

func genUnCommentMsg(userId int64, videoId int64, comment_id string, timestamp string) string {
	return strconv.FormatInt(userId, 10) + "::" + strconv.FormatInt(videoId, 10) + "::2::" + comment_id + "::" + timestamp
}

func Comment(userId int64, videoId int64, comment_text string, timestamp string) error {
	conn.Db.RPush("dousheng_comment", genCommentMsg(userId, videoId, comment_text, timestamp))

	return nil
}

func UnComment(userId int64, videoId int64, comment_id string, timestamp string) error {
	conn.Db.RPush("dousheng_comment", genUnCommentMsg(userId, videoId, comment_id, timestamp))

	return nil
}

func GetAndDeleteAllComments() (commentModify []db.CommentModify) {
	for {
		cmd := conn.Db.LPop("dousheng_comment")
		if cmd.Err() != nil {
			break
		}
		s := cmd.Val()
		ss := strings.Split(s, "::")
		if ss[2] == "1" {
			userId, _ := strconv.Atoi(ss[0])
			videoId, _ := strconv.Atoi(ss[1])
			timestamp, _ := strconv.Atoi(ss[4])
			datetime := time.Unix(int64(timestamp), 0)

			commentModify = append(commentModify,
				db.CommentModify{
					Comment: sqlmodel.Comment{
						VideoId:        int64(videoId),
						UserId:         int64(userId),
						CommentContent: ss[3],
						Ctime:          datetime,
						Utime:          datetime,
					},
					Delete: false})
		} else if ss[2] == "2" {
			commentId, _ := strconv.Atoi(ss[3])
			commentModify = append(commentModify,
				db.CommentModify{
					Comment: sqlmodel.Comment{CommentId: int64(commentId)},
					Delete:  true,
				})
		}
	}
	return
}
