package sqlmodel

type UserVideoLike struct {
	UserId  int64 `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`
	VideoId int64 `gorm:"column:video_id" db:"video_id" json:"video_id" form:"video_id"`
	IsLike  int   `gorm:"column:is_like" db:"is_like" json:"is_like" form:"is_like"`
}

func (UserVideoLike) TableName() string {
	return "user_video_like"
}

const SQL_USER_VIDEO_LIKE_USER_ID = "user_id"
const SQL_USER_VIDEO_LIKE_VIDEO_ID = "video_id"
const SQL_USER_VIDEO_LIKE_IS_LIKE = "is_like"
