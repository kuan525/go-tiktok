package service

import "github.com/kataras/iris/v12"

// DouyinMessageChatHandler 聊天记录 当前登录用户和其他指定用户的聊天消息记录
func DouyinMessageChatHandler(ctx iris.Context, reqBody interface{}) {
}

// DouyinMessageActionHandler 消息操作 登录用户对消息的相关操作，目前只支持消息发送
func DouyinMessageActionHandler(ctx iris.Context, reqBody interface{}) {
}
