// Code generated by Kitex v0.4.4. DO NOT EDIT.

package feedservice

import (
	"context"
	douyin_feed "github.com/1037group/dousheng/kitex_gen/douyin_feed"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return feedServiceServiceInfo
}

var feedServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "FeedService"
	handlerType := (*douyin_feed.FeedService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Feed": kitex.NewMethodInfo(feedHandler, newFeedServiceFeedArgs, newFeedServiceFeedResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "douyin_feed",
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

func feedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*douyin_feed.FeedServiceFeedArgs)
	realResult := result.(*douyin_feed.FeedServiceFeedResult)
	success, err := handler.(douyin_feed.FeedService).Feed(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFeedServiceFeedArgs() interface{} {
	return douyin_feed.NewFeedServiceFeedArgs()
}

func newFeedServiceFeedResult() interface{} {
	return douyin_feed.NewFeedServiceFeedResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Feed(ctx context.Context, req *douyin_feed.FeedRequest) (r *douyin_feed.FeedResponse, err error) {
	var _args douyin_feed.FeedServiceFeedArgs
	_args.Req = req
	var _result douyin_feed.FeedServiceFeedResult
	if err = p.c.Call(ctx, "Feed", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
