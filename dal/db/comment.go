package db

import (
	"dousheng/kitex_gen/douyin_comment"
	"dousheng/kitex_gen/douyin_user"
	"dousheng/pkg/configs/sqlmodel"
	"fmt"
	"strconv"
)

type CommentModify struct {
	Comment sqlmodel.Comment
	Delete  bool
}

func MSetComments(commentModify []CommentModify) {
	for _, cm := range commentModify {
		if cm.Delete {
			DB.Delete(&sqlmodel.Comment{}, cm.Comment.CommentId)
		} else {
			DB.Create(&sqlmodel.Comment{
				VideoId:        cm.Comment.VideoId,
				UserId:         cm.Comment.UserId,
				CommentContent: cm.Comment.CommentContent,
				Ctime:          cm.Comment.Ctime,
				Utime:          cm.Comment.Utime,
			})
		}
	}
}

func GetComments(videoId int64) (resp []*douyin_comment.Comment) {
	var comments []sqlmodel.Comment
	DB.Order("ctime desc").Where("video_id = ?", videoId).Find(&comments)

	var userIds []int64
	var users []*sqlmodel.User
	for _, c := range comments {
		userIds = append(userIds, c.UserId)
	}
	users, _ = MGetUserByIDAllowRepetition(nil, userIds)
	fmt.Println(len(users))
	fmt.Println(len(comments))

	for i, c := range comments {
		// favorited count
		favoritedCount := GetFavoritedCountByUser(users[i].UserId)
		// work count
		workCount := len(GetPublishList(users[i].UserId))
		// favorite count
		favoriteCount := GetFavoriteCountByUser(users[i].UserId)

		resp = append(resp, &douyin_comment.Comment{
			Id: c.CommentId,
			User: &douyin_user.User{
				Id:             users[i].UserId,
				Name:           users[i].UserName,
				FollowCount:    &users[i].UserFollowCount,
				FollowerCount:  &users[i].UserFollowerCount,
				IsFollow:       false,
				TotalFavorited: favoritedCount,
				WorkCount:      int64(workCount),
				FavoriteCount:  favoriteCount,
			},
			Content:    c.CommentContent,
			CreateDate: strconv.Itoa(int(c.Ctime.Month())) + "-" + strconv.Itoa(c.Ctime.Day()),
		})
	}

	return
}
