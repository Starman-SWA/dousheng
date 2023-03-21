package db

import (
	"dousheng/cmd/api/biz/model/douyin_api"
	"dousheng/kitex_gen/douyin_feed"
	"strconv"
)

// User
func FeedResponseRpc2Api(m *douyin_feed.FeedResponse) *douyin_api.FeedResponse {
	if m == nil {
		return nil
	}

	var videoList []*douyin_api.Video

	for _, video := range m.VideoList {
		var user douyin_api.User
		if video.Author != nil {
			// favorited count
			favoritedCount := GetFavoritedCountByUser(video.Author.Id)
			// work count
			workCount := len(GetPublishList(video.Author.Id))
			// favorite count
			favoriteCount := GetFavoriteCountByUser(video.Author.Id)

			user = douyin_api.User{
				ID:              video.Author.Id,
				Name:            video.Author.Name,
				FollowCount:     0,
				FollowerCount:   0,
				IsFollow:        false,
				Avatar:          "https://img.zmtc.com/2022/0926/20220926081721426.jpg",
				BackgroundImage: "https://photo.16pic.com/00/54/69/16pic_5469853_b.jpg",
				Signature:       "但做好事，不问前程",
				TotalFavorited:  strconv.FormatInt(favoritedCount, 10),
				WorkCount:       int64(workCount),
				FavoriteCount:   favoriteCount,
			}
		}

		one := douyin_api.Video{
			ID:            video.Id,
			Author:        &user,
			PlayURL:       video.PlayUrl,
			CoverURL:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    video.IsFavorite,
			Title:         video.Title,
		}
		videoList = append(videoList, &one)
	}

	return &douyin_api.FeedResponse{
		StatusCode: m.StatusCode,
		StatusMsg:  m.StatusMsg,
		VideoList:  videoList,
		NextTime:   m.NextTime,
	}
}
