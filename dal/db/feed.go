package db

import (
	"context"
	"dousheng/pkg/configs/sqlmodel"
	"github.com/cloudwego/kitex/pkg/klog"
	"time"
)

// MGetVideos multiple get list of videos
func MGetVideos(ctx context.Context, last_time *int64) ([]*sqlmodel.Video, error) {
	klog.CtxInfof(ctx, "[MGetVideos] last_time: %+v\n", last_time)

	var res []*sqlmodel.Video

	if last_time == nil || *last_time == 0 {
		cur_time := time.Now().UnixMilli()
		klog.CtxInfof(ctx, "cur_time %+v", cur_time)
		klog.CtxInfof(ctx, "cur_time %+v", cur_time)

		last_time = &cur_time
	}

	query := sqlmodel.SQL_VIDEO_UTIME + " <= ?"
	if err := DB.WithContext(ctx).Order(sqlmodel.SQL_VIDEO_UTIME+" desc").Find(&res, query, time.UnixMilli(*last_time)).Error; err != nil {
		klog.CtxInfof(ctx, "[MGetVideos] res: %+v\n", res)
		return res, err
	}
	return res, nil
}
