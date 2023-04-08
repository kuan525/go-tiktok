package v1

import (
	"common/middleware"
	"github.com/kataras/iris/v12/core/router"
	"web_comment/internal/controller"
)

// RegisterAuthenticationRouter 需要鉴权
func RegisterAuthenticationRouter(party *router.Party) {
	// 评论操作 登录用户对视频进行评论。
	(*party).Handle("POST", "/comment/action/", middleware.Handler(controller.ConnectController{}.DouyinCommentAction))
	// 视频评论列表 查看视频的所有评论，按发布时间倒序。
	(*party).Handle("GET", "/comment/list/", middleware.Handler(controller.ConnectController{}.DouyinCommentList))
}

// RegisterConfigRouter 不需要鉴权
func RegisterConfigRouter(party *router.Party) {
}
