package db

import (
	"dousheng/kitex_gen/douyin_feed"
	"dousheng/pkg/configs/sqlmodel"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func MSetLikes(videoLikeCounts []VideoLikeCount, userVideoLikes []UserVideoLike) error {
	// modify videoCount to video
	for _, vc := range videoLikeCounts {
		//Db.Model(xy).Where("id = ? ", id).Update("sign_up_num", gorm.Expr("sign_up_num+ ?", 1))
		DB.Model(&sqlmodel.Video{}).
			Where("video_id = ?", vc.VideoId).
			Update(sqlmodel.SQL_VIDEO_VIDEO_FAVORITE_COUNT,
				gorm.Expr(sqlmodel.SQL_VIDEO_VIDEO_FAVORITE_COUNT+"+ ?", vc.LikeCount))
	}

	// set userVideoLike
	for _, uvl := range userVideoLikes {
		userVideoLike := sqlmodel.UserVideoLike{UserId: uvl.UserId, VideoId: uvl.VideoId, IsLike: uvl.IsLike}

		//db.Clauses(clause.OnConflict{
		//	Columns:   []clause.Column{{Name: "id"}},
		//	DoUpdates: clause.AssignmentColumns([]string{"name", "age"}),
		//}).Create(&users)
		DB.Clauses(clause.OnConflict{
			DoUpdates: clause.AssignmentColumns([]string{sqlmodel.SQL_USER_VIDEO_LIKE_IS_LIKE}),
		}).Create(&userVideoLike)
	}

	return nil
}

func MGetLikeStates(videos []*sqlmodel.Video, userId *int64) []bool {
	var uvl []*sqlmodel.UserVideoLike
	var likeStates []bool

	for _, video := range videos {
		DB.Where("user_id = ? AND video_id = ?", *userId, video.VideoId).Find(&uvl)
		if len(uvl) == 0 {
			likeStates = append(likeStates, false)
		} else if uvl[0].IsLike == 1 {
			likeStates = append(likeStates, true)
		} else {
			likeStates = append(likeStates, false)
		}
	}

	return likeStates
}

func GetFavoriteCountByUser(userId int64) int64 {
	var uvl []*sqlmodel.UserVideoLike
	DB.Where("user_id = ?", userId).Find(&uvl)

	var count int64
	for _, uv := range uvl {
		if uv.IsLike == 1 {
			count += 1
		}
	}

	return count
}

func GetFavoritedCountByUser(userId int64) int64 {
	var videos []*sqlmodel.Video
	DB.Where("user_id = ?", userId).Find(&videos)

	var count int64
	for _, video := range videos {
		count += video.VideoFavoriteCount
	}
	return count
}

func GetFavoriteList(userId int64) (videoList []*douyin_feed.Video) {
	var uvl []*sqlmodel.UserVideoLike
	DB.Where("user_id = ?", userId).Find(&uvl)

	var video_ids []int64
	for _, uv := range uvl {
		if uv.IsLike == 1 {
			video_ids = append(video_ids, uv.VideoId)
		}
	}

	var videos []*sqlmodel.Video
	DB.Where("video_id IN ?", video_ids).Find(&videos)

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
