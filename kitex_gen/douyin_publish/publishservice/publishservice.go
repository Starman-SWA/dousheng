// Code generated by Kitex v0.4.4. DO NOT EDIT.

package publishservice

import (
	"context"
	douyin_publish "dousheng/kitex_gen/douyin_publish"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return publishServiceServiceInfo
}

var publishServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "PublishService"
	handlerType := (*douyin_publish.PublishService)(nil)
	methods := map[string]kitex.MethodInfo{
		"PublishAction": kitex.NewMethodInfo(publishActionHandler, newPublishServicePublishActionArgs, newPublishServicePublishActionResult, false),
		"PublishList":   kitex.NewMethodInfo(publishListHandler, newPublishServicePublishListArgs, newPublishServicePublishListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "douyin_publish",
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

func publishActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*douyin_publish.PublishServicePublishActionArgs)
	realResult := result.(*douyin_publish.PublishServicePublishActionResult)
	success, err := handler.(douyin_publish.PublishService).PublishAction(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPublishServicePublishActionArgs() interface{} {
	return douyin_publish.NewPublishServicePublishActionArgs()
}

func newPublishServicePublishActionResult() interface{} {
	return douyin_publish.NewPublishServicePublishActionResult()
}

func publishListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*douyin_publish.PublishServicePublishListArgs)
	realResult := result.(*douyin_publish.PublishServicePublishListResult)
	success, err := handler.(douyin_publish.PublishService).PublishList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPublishServicePublishListArgs() interface{} {
	return douyin_publish.NewPublishServicePublishListArgs()
}

func newPublishServicePublishListResult() interface{} {
	return douyin_publish.NewPublishServicePublishListResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) PublishAction(ctx context.Context, req *douyin_publish.PublishActionRequest) (r *douyin_publish.PublishActionResponse, err error) {
	var _args douyin_publish.PublishServicePublishActionArgs
	_args.Req = req
	var _result douyin_publish.PublishServicePublishActionResult
	if err = p.c.Call(ctx, "PublishAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PublishList(ctx context.Context, req *douyin_publish.PublishListRequest) (r *douyin_publish.PublishListResponse, err error) {
	var _args douyin_publish.PublishServicePublishListArgs
	_args.Req = req
	var _result douyin_publish.PublishServicePublishListResult
	if err = p.c.Call(ctx, "PublishList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
