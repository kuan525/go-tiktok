package v1

import (
	"common/middleware"
	"github.com/kataras/iris/v12/core/router"
	"web_publish/internal/controller"
)

// RegisterAuthenticationRouter 需要鉴权
func RegisterAuthenticationRouter(party *router.Party) {
	// 视频投稿 登录用户选择视频上传。
	(*party).Handle("POST", "/publish/action/", middleware.Handler(controller.ConnectController{}.DouyinPublishAction))
	// 发布列表 登录用户的视频发布列表，直接列出用户所有投稿过的视频。
	(*party).Handle("GET", "/publish/list/ ", middleware.Handler(controller.ConnectController{}.DouyinPublishList))
}

// RegisterConfigRouter 不需要鉴权
func RegisterConfigRouter(party *router.Party) {
}
