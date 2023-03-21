package db

import (
	"dousheng/kitex_gen/douyin_feed"
	"dousheng/kitex_gen/douyin_user"
	"dousheng/pkg/configs/sqlmodel"
)

// Video pack video info
func Video(m *sqlmodel.Video, n *sqlmodel.User, isLike bool) *douyin_feed.Video {
	if m == nil {
		return nil
	}

	return &douyin_feed.Video{
		Id:            m.VideoId,
		Author:        User(n),
		PlayUrl:       m.VideoPlayUrl,
		CoverUrl:      m.VideoCoverUrl,
		FavoriteCount: m.VideoFavoriteCount,
		CommentCount:  m.VideoCommentCount,
		IsFavorite:    isLike,
		Title:         m.VideoTitle,
	}
}

// User pack user info
func User(m *sqlmodel.User) *douyin_user.User {
	if m == nil {
		return nil
	}

	return &douyin_user.User{
		Id:            m.UserId,
		Name:          m.UserName,
		FollowCount:   &m.UserFollowCount,
		FollowerCount: &m.UserFollowerCount,
		IsFollow:      false,
	}
}

// Videos pack list of videos info
func Videos(ms []*sqlmodel.Video, ns map[int64]sqlmodel.User, userId *int64) []*douyin_feed.Video {
	videos := make([]*douyin_feed.Video, 0)

	if userId == nil {
		for _, m := range ms {
			user := ns[m.UserId]
			if video := Video(m, &user, false); video != nil {
				videos = append(videos, video)
			}
		}
	} else {
		likeStates := MGetLikeStates(ms, userId)

		for i, m := range ms {
			user := ns[m.UserId]
			if video := Video(m, &user, likeStates[i]); video != nil {
				videos = append(videos, video)
			}
		}
	}

	return videos
}
