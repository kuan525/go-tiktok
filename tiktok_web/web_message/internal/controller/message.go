package controller

import (
	"common/middleware"
	"github.com/kataras/iris/v12"
	"web_message/internal/service"
)

type ConnectController struct {
}

// DouyinMessageChat 聊天记录 当前登录用户和其他指定用户的聊天消息记录
func (c ConnectController) DouyinMessageChat(ctx iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinMessageChatHandler)
}

// DouyinMessageAction 消息操作 登录用户对消息的相关操作，目前只支持消息发送
func (c ConnectController) DouyinMessageAction(ctx iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinMessageActionHandler)
}
