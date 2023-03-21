package main

import (
	"context"
	"dousheng/cmd/favorite/kafkaop"
	"dousheng/dal/db"
	douyin_favorite "dousheng/kitex_gen/douyin_favorite"
	"dousheng/kitex_gen/douyin_user"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// FavoriteAction implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteAction(ctx context.Context, req *douyin_favorite.FavoriteActionRequest) (resp *douyin_favorite.FavoriteActionResponse, err error) {
	// TODO: Your code here...
	err = kafkaop.ProduceLike(ctx, req.UserId, req.VideoId, req.ActionType)
	if err != nil {
		return nil, err
	}

	resp = douyin_favorite.NewFavoriteActionResponse()
	return resp, nil
}

// FavoriteList implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteList(ctx context.Context, req *douyin_favorite.FavoriteListRequest) (resp *douyin_favorite.FavoriteListResponse, err error) {
	// TODO: Your code here...
	videos := db.GetFavoriteList(req.UserId)
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

	resp = douyin_favorite.NewFavoriteListResponse()
	resp.VideoList = videos
	return resp, nil
}
