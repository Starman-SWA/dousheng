package db

import (
	"context"
	"dousheng/kitex_gen/douyin_feed"
	"dousheng/kitex_gen/douyin_publish"
	"dousheng/pkg/configs/sqlmodel"
	"github.com/cloudwego/kitex/pkg/klog"
	"time"
)

func PublishActionToDB(ctx context.Context, req *douyin_publish.PublishActionRequest, videoURL string, coverURL string) error {
	klog.CtxInfof(ctx, "[PublishActionToDB] req.Title: %+v\n", req.Title)

	video := sqlmodel.Video{
		UserId:             req.UserId,
		VideoPlayUrl:       videoURL,
		VideoCoverUrl:      coverURL,
		VideoFavoriteCount: 0,
		VideoCommentCount:  0,
		VideoTitle:         req.Title,
		Ctime:              time.Now(),
		Utime:              time.Now()}
	result := DB.Select(
		"user_id",
		"video_play_url",
		"video_cover_url",
		"video_favorite_count",
		"video_comment_count",
		"video_title",
		"ctime",
		"utime").Create(&video)
	return result.Error
}

func GetPublishList(userId int64) (videoList []*douyin_feed.Video) {
	var videos []*sqlmodel.Video
	DB.Where("user_id = ?", userId).Find(&videos)

	likeStates := MGetLikeStates(videos, &userId)

	for i, video := range videos {
		videoList = append(videoList, &douyin_feed.Video{
			Id:            video.VideoId,
			Author:        nil,
			PlayUrl:       video.VideoPlayUrl,
			CoverUrl:      video.VideoCoverUrl,
			FavoriteCount: video.VideoFavoriteCount,
			CommentCount:  video.VideoCommentCount,
			IsFavorite:    likeStates[i],
			Title:         video.VideoTitle,
		})
	}

	return
}
