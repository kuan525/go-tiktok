package controller

import (
	"github.com/kataras/iris/v12"
	"go-tiktok/internal/middleware"
	"go-tiktok/internal/service"
)

// DouyinRelationAction 社交接口 实现用户之间的关注关系维护，登录用户能够关注或取关其他用户，同时自己能够看到自己关注过的所有用户列表，以及所有关注自己的用户列表。
func (c ConnectController) DouyinRelationAction(ctx iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinRelationActionHandler)
}

// DouyinRelationFollowList 用户关注列表 登录用户关注的所有用户列表。
func (c ConnectController) DouyinRelationFollowList(ctx iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinRelationFollowListHandler)
}

// DouyinRelationFollowerList 用户粉丝列表 所有关注登录用户的粉丝列表。
func (c ConnectController) DouyinRelationFollowerList(ctx iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinRelationFollowerListHandler)
}

// DouyinRelationFriendList 用户好友列表 所有关注登录用户的粉丝列表。
func (c ConnectController) DouyinRelationFriendList(ctx iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinRelationFriendListHandler)
}

// DouyinMessageChat 聊天记录 当前登录用户和其他指定用户的聊天消息记录
func (c ConnectController) DouyinMessageChat(ctx iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinMessageChatHandler)
}

// DouyinMessageAction 消息操作 登录用户对消息的相关操作，目前只支持消息发送
func (c ConnectController) DouyinMessageAction(ctx iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinMessageActionHandler)
}
