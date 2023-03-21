package redisop

import (
	"dousheng/dal/db"
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"strings"
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

func Like(userId int64, videoId int64) error {
	pipe := conn.Db.TxPipeline()
	cmd := pipe.HGet("dousheng_video_user_like", getKeyFromUserAndVideo(userId, videoId))

	if cmd.Val() == "1" {
		return nil
	}
	pipe.HSet("dousheng_video_user_like", getKeyFromUserAndVideo(userId, videoId), "1")
	pipe.HIncrBy("dousheng_video_like_count", strconv.FormatInt(videoId, 10), 1)
	pipe.Exec()

	return nil
}

func Unlike(userId int64, videoId int64) error {
	pipe := conn.Db.TxPipeline()
	cmd := pipe.HGet("dousheng_video_user_like", getKeyFromUserAndVideo(userId, videoId))
	if cmd.Val() == "2" {
		return nil
	}
	pipe.HSet("dousheng_video_user_like", getKeyFromUserAndVideo(userId, videoId), "2")
	pipe.HIncrBy("dousheng_video_like_count", strconv.FormatInt(videoId, 10), -1)
	pipe.Exec()

	return nil
}

func GetAndDeleteAllLikes() (err error, videoLikeCounts []db.VideoLikeCount, userVideoLikes []db.UserVideoLike) {
	pipe := conn.Db.TxPipeline()
	cmdVideoLikeCount := pipe.HGetAll("dousheng_video_like_count")
	cmdVideoUserLike := pipe.HGetAll("dousheng_video_user_like")
	pipe.Exec()

	pipe = conn.Db.TxPipeline()
	videoLikeCountMap := cmdVideoLikeCount.Val()
	for videoIdStr, countStr := range videoLikeCountMap {
		videoId, _ := strconv.Atoi(videoIdStr)
		count, _ := strconv.Atoi(countStr)
		videoLikeCounts = append(videoLikeCounts, db.VideoLikeCount{VideoId: int64(videoId), LikeCount: int64(count)})
		pipe.HDel("dousheng_video_like_count", videoIdStr)
	}

	videoUserLikeMap := cmdVideoUserLike.Val()
	for videoUserStr, likeStr := range videoUserLikeMap {
		videoUserSplits := strings.Split(videoUserStr, "::")
		userId, _ := strconv.Atoi(videoUserSplits[0])
		videoId, _ := strconv.Atoi(videoUserSplits[1])
		isLike, _ := strconv.Atoi(likeStr)

		userVideoLikes = append(userVideoLikes, db.UserVideoLike{UserId: int64(userId), VideoId: int64(videoId), IsLike: isLike})

		pipe.HDel("dousheng_video_user_like", videoUserStr)
	}

	pipe.Exec()
	return
}
