package db

import (
	"context"
	"dousheng/pkg/configs/sqlmodel"
	"github.com/cloudwego/kitex/pkg/klog"
)

// MGetUserById multiple get list of user info
func MGetUserByID(ctx context.Context, userIDs []int64) ([]*sqlmodel.User, error) {
	if ctx != nil {
		klog.CtxInfof(ctx, "[MGetUserByID] userIDs: %+v\n", userIDs)
	}

	res := make([]*sqlmodel.User, 0)
	if len(userIDs) == 0 {
		return res, nil
	}

	query := sqlmodel.SQL_USER_USER_ID + " in ?"
	if err := DB.WithContext(ctx).Where(query, userIDs).Find(&res).Error; err != nil {
		klog.CtxInfof(ctx, "[MGetUserByID] res: %+v\n", res)
		return res, err
	}
	return res, nil
}

func MGetUserByIDAllowRepetition(ctx context.Context, userIDs []int64) ([]*sqlmodel.User, error) {
	if ctx != nil {
		klog.CtxInfof(ctx, "[MGetUserByID] userIDs: %+v\n", userIDs)
	}

	res := make([]*sqlmodel.User, 0)
	var user sqlmodel.User
	if len(userIDs) == 0 {
		return res, nil
	}
	for _, id := range userIDs {
		DB.Where("user_id = ?", id).Find(&user)
		var u sqlmodel.User
		u = user
		res = append(res, &u)
	}

	return res, nil
}

// QueryUser get user by userName
func GetUserByUserName(ctx context.Context, userName string) ([]*sqlmodel.User, error) {
	klog.CtxInfof(ctx, "[GetUserByUserName] userName: %+v\n", userName)

	res := make([]*sqlmodel.User, 0)

	query := sqlmodel.SQL_USER_USER_NAME + " = ?"
	if err := DB.WithContext(ctx).Where(query, userName).Find(&res).Error; err != nil {
		klog.CtxErrorf(ctx, err.Error())
		return nil, err
	}
	return res, nil
}

// CreateUser create user info
func CreateUser(ctx context.Context, user *sqlmodel.User) error {
	klog.CtxInfof(ctx, "[CreateUser] userName: %+v\n", user)
	return DB.WithContext(ctx).Create(user).Error
}
