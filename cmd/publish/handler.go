package main

import (
	"context"
	"dousheng/cmd/publish/util"
	"dousheng/dal/db"
	douyin_publish "dousheng/kitex_gen/douyin_publish"
	"dousheng/kitex_gen/douyin_user"
	"dousheng/obss"
	"dousheng/pkg/consts"
	"github.com/google/uuid"
	"os"
)

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

// PublishAction implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishAction(ctx context.Context, req *douyin_publish.PublishActionRequest) (resp *douyin_publish.PublishActionResponse, err error) {
	// TODO: Your code here...
	// 生成UUID
	u := uuid.New()
	us := u.String()
	videoPath := consts.VideoDir + us

	// 将视频保存到本地
	err = os.WriteFile(videoPath, req.Data, 0444)
	if err != nil {
		return nil, err
	}

	// 生成缩略图
	var snapshotPath string
	snapshotPath, err = util.GetSnapshot(videoPath, us, 0)
	if err != nil {
		return nil, err
	}

	// 将视频和缩略图上传到云端
	videoKey := us + "_video"
	snapshotKey := us + "_snapshot"
	err = obss.PutFile(videoKey, videoPath)
	if err != nil {
		return nil, err
	}
	err = obss.PutFile(snapshotKey, snapshotPath)
	if err != nil {
		return nil, err
	}

	// 添加数据库
	err = db.PublishActionToDB(ctx, req, obss.GenGetURL(videoKey), obss.GenGetURL(snapshotKey))
	if err != nil {
		return nil, err
	}

	// 构造返回值
	resp = douyin_publish.NewPublishActionResponse()
	return resp, nil
}

// PublishList implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishList(ctx context.Context, req *douyin_publish.PublishListRequest) (resp *douyin_publish.PublishListResponse, err error) {
	// TODO: Your code here...
	videos := db.GetPublishList(req.UserId)
	users, _ := db.MGetUserByID(ctx, []int64{req.UserId})

	// favorited count
	favoritedCount := db.GetFavoritedCountByUser(req.UserId)
	// work count
	workCount := len(db.GetPublishList(req.UserId))
	// favorite count
	favoriteCount := db.GetFavoriteCountByUser(req.UserId)

	for i, _ := range videos {
		videos[i].Author = &douyin_user.User{
			Id:             users[0].UserId,
			Name:           users[0].UserName,
			FollowCount:    &users[0].UserFollowCount,
			FollowerCount:  &users[0].UserFollowerCount,
			IsFollow:       false,
			TotalFavorited: favoritedCount,
			WorkCount:      int64(workCount),
			FavoriteCount:  favoriteCount,
		}
	}

	resp = douyin_publish.NewPublishListResponse()
	resp.VideoList = videos
	return resp, nil
}
