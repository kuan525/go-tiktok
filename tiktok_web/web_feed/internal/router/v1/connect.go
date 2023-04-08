package v1

import (
	"common/middleware"
	"github.com/kataras/iris/v12/core/router"
	"web_feed/internal/controller"
)

// RegisterAuthenticationRouter 需要鉴权
func RegisterAuthenticationRouter(party *router.Party) {
}

// RegisterConfigRouter 不需要鉴权
func RegisterConfigRouter(party *router.Party) {
	// 视频流接口 不限制登录状态，返回按投稿时间倒序的视频列表，视频数由服务端控制，单次最多30个。
	(*party).Handle("GET", "/feed/", middleware.Handler(controller.ConnectController{}.DouyinFeed))
}
