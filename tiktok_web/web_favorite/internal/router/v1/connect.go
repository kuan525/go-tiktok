package v1

import (
	"common/middleware"
	"github.com/kataras/iris/v12/core/router"
	"web_favorite/internal/controller"
)

// RegisterAuthenticationRouter 需要鉴权
func RegisterAuthenticationRouter(party *router.Party) {
	// 赞操作 登录用户对视频的点赞和取消点赞操作。
	(*party).Handle("POST", "/favorite/action/", middleware.Handler(controller.ConnectController{}.DouyinFavoriteAction))
	// 喜欢列表 登录用户的所有点赞视频。
	(*party).Handle("GET", "/favorite/list/", middleware.Handler(controller.ConnectController{}.DouyinFavoriteList))
}

// RegisterConfigRouter 不需要鉴权
func RegisterConfigRouter(party *router.Party) {
}
