package controller

import (
	"github.com/kataras/iris/v12"
	"go-tiktok/internal/middle"
	"go-tiktok/internal/service"
)

func (c ConnectController) DouyinRelationAction(ctx *iris.Context) {
	middle.AipWrapper(ctx, service.DouyinRelationActionHandler)
}

func (c ConnectController) DouyinRelationFollowList(ctx *iris.Context) {
	middle.AipWrapper(ctx, service.DouyinRelationFollowListHandler)
}

func (c ConnectController) DouyinRelationFollowerList(ctx *iris.Context) {
	middle.AipWrapper(ctx, service.DouyinRelationFollowerListHandler)
}

func (c ConnectController) DouyinRelationFriendList(ctx *iris.Context) {
	middle.AipWrapper(ctx, service.DouyinRelationFriendListHandler)
}

func (c ConnectController) DouyinMessageChat(ctx *iris.Context) {
	middle.AipWrapper(ctx, service.DouyinMessageChatHandler)
}

func (c ConnectController) DouyinMessageAction(ctx *iris.Context) {
	middle.AipWrapper(ctx, service.DouyinMessageActionHandler)
}
