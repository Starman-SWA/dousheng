// Code generated by Kitex v0.4.4. DO NOT EDIT.
package userservice

import (
	douyin_user "github.com/1037group/dousheng/kitex_gen/douyin_user"
	server "github.com/cloudwego/kitex/server"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler douyin_user.UserService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
