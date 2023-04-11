package controller

import (
	"common/middleware"
	"github.com/kataras/iris/v12"
	"web_feed/conf"
	"web_feed/internal/request"
	"web_feed/internal/service"
)

type ConnectController struct {
}

// DouyinFeed 视频流接口 不限制登录状态，返回按投稿时间倒序的视频列表，视频数由服务端控制，单次最多30个。
func (c ConnectController) DouyinFeed(ctx iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinFeedHandler, &request.DouyinFeedReq{}, conf.Logger)
}
