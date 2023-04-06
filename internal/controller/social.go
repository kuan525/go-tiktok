package controller

import (
	"github.com/kataras/iris/v12"
	"go-tiktok/internal/middleware"
	"go-tiktok/internal/service"
)

func (c ConnectController) DouyinRelationAction(ctx iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinRelationActionHandler)
}

func (c ConnectController) DouyinRelationFollowList(ctx iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinRelationFollowListHandler)
}

func (c ConnectController) DouyinRelationFollowerList(ctx iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinRelationFollowerListHandler)
}

func (c ConnectController) DouyinRelationFriendList(ctx iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinRelationFriendListHandler)
}

func (c ConnectController) DouyinMessageChat(ctx iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinMessageChatHandler)
}

func (c ConnectController) DouyinMessageAction(ctx iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinMessageActionHandler)
}
