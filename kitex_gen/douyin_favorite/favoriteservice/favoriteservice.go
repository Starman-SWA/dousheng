// Code generated by Kitex v0.4.4. DO NOT EDIT.

package favoriteservice

import (
	"context"
	douyin_favorite "dousheng/kitex_gen/douyin_favorite"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return favoriteServiceServiceInfo
}

var favoriteServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "FavoriteService"
	handlerType := (*douyin_favorite.FavoriteService)(nil)
	methods := map[string]kitex.MethodInfo{
		"FavoriteAction": kitex.NewMethodInfo(favoriteActionHandler, newFavoriteServiceFavoriteActionArgs, newFavoriteServiceFavoriteActionResult, false),
		"FavoriteList":   kitex.NewMethodInfo(favoriteListHandler, newFavoriteServiceFavoriteListArgs, newFavoriteServiceFavoriteListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "douyin_favorite",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func favoriteActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*douyin_favorite.FavoriteServiceFavoriteActionArgs)
	realResult := result.(*douyin_favorite.FavoriteServiceFavoriteActionResult)
	success, err := handler.(douyin_favorite.FavoriteService).FavoriteAction(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFavoriteServiceFavoriteActionArgs() interface{} {
	return douyin_favorite.NewFavoriteServiceFavoriteActionArgs()
}

func newFavoriteServiceFavoriteActionResult() interface{} {
	return douyin_favorite.NewFavoriteServiceFavoriteActionResult()
}

func favoriteListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*douyin_favorite.FavoriteServiceFavoriteListArgs)
	realResult := result.(*douyin_favorite.FavoriteServiceFavoriteListResult)
	success, err := handler.(douyin_favorite.FavoriteService).FavoriteList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFavoriteServiceFavoriteListArgs() interface{} {
	return douyin_favorite.NewFavoriteServiceFavoriteListArgs()
}

func newFavoriteServiceFavoriteListResult() interface{} {
	return douyin_favorite.NewFavoriteServiceFavoriteListResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) FavoriteAction(ctx context.Context, req *douyin_favorite.FavoriteActionRequest) (r *douyin_favorite.FavoriteActionResponse, err error) {
	var _args douyin_favorite.FavoriteServiceFavoriteActionArgs
	_args.Req = req
	var _result douyin_favorite.FavoriteServiceFavoriteActionResult
	if err = p.c.Call(ctx, "FavoriteAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) FavoriteList(ctx context.Context, req *douyin_favorite.FavoriteListRequest) (r *douyin_favorite.FavoriteListResponse, err error) {
	var _args douyin_favorite.FavoriteServiceFavoriteListArgs
	_args.Req = req
	var _result douyin_favorite.FavoriteServiceFavoriteListResult
	if err = p.c.Call(ctx, "FavoriteList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
