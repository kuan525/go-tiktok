package v1

import (
	"common/middleware"
	"github.com/kataras/iris/v12/core/router"
	"web_message/internal/controller"
)

// RegisterAuthenticationRouter 需要鉴权
func RegisterAuthenticationRouter(party *router.Party) {
	// 聊天记录 当前登录用户和其他指定用户的聊天消息记录
	(*party).Handle("GET", "/message/chat/ ", middleware.Handler(controller.ConnectController{}.DouyinMessageChat))
	// 消息操作 登录用户对消息的相关操作，目前只支持消息发送
	(*party).Handle("POST", "/message/action/", middleware.Handler(controller.ConnectController{}.DouyinMessageAction))
}

// RegisterConfigRouter 不需要鉴权
func RegisterConfigRouter(party *router.Party) {
}
