package db

type VideoLikeCount struct {
	VideoId   int64
	LikeCount int64
}

type UserVideoLike struct {
	UserId  int64
	VideoId int64
	IsLike  int
}
