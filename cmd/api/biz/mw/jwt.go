package mw

import (
	"dousheng/cmd/api/biz/model/douyin_api"
	"dousheng/pkg/consts"
	"dousheng/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/kitex/pkg/klog"
	"net/http"
	"time"

	"context"
	"dousheng/cmd/api/biz/rpc"
	"dousheng/kitex_gen/douyin_user"
	"github.com/hertz-contrib/jwt"
)

// TODO code review

var JwtMiddleware *jwt.HertzJWTMiddleware
var err error

func InitJWT() {
	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Key:           []byte(consts.SecretKey),
		TokenLookup:   "form: token, query: token",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		IdentityKey:   consts.IdentityKey,
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return &douyin_user.User{
				Id: int64(claims[consts.IdentityKey].(float64)),
			}
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					consts.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var err error
			var req douyin_api.UserLoginRequest
			if err = c.BindAndValidate(&req); err != nil {
				klog.CtxErrorf(ctx, err.Error())
				return "", jwt.ErrMissingLoginValues
			}
			if len(req.Username) == 0 || len(req.Password) == 0 {
				klog.CtxErrorf(ctx, err.Error())
				return "", jwt.ErrMissingLoginValues
			}
			userId, err := rpc.CheckUser(ctx, &douyin_user.UserLoginRequest{
				Username: req.Username,
				Password: req.Password,
			})
			if err == nil {
				c.Set("user_id", userId)
			}
			return userId, err
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, utils.H{
				"status_code": errno.Success.ErrCode,
				"token":       token,
				"user_id":     c.Keys["user_id"],
				"status_msg":  "success",
			})
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusOK, utils.H{
				//"code":    errno.AuthorizationFailedErr.ErrCode,
				"code":    code,
				"message": message,
			})
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			switch t := e.(type) {
			case errno.ErrNo:
				return t.ErrMsg
			default:
				return t.Error()
			}
		},
	})
	if err != nil {
		panic(err)
	}
}
