package controller

import (
	"common/middleware"
	"github.com/kataras/iris/v12"
	"web_relation/internal/request"
	"web_relation/internal/service"
)

type ConnectController struct {
}

// DouyinRelationAction 社交接口 实现用户之间的关注关系维护，登录用户能够关注或取关其他用户，同时自己能够看到自己关注过的所有用户列表，以及所有关注自己的用户列表。
func (c ConnectController) DouyinRelationAction(ctx iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinRelationActionHandler, &request.DouyinRelationActionReq{})
}

// DouyinRelationFollowList 用户关注列表 登录用户关注的所有用户列表。
func (c ConnectController) DouyinRelationFollowList(ctx iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinRelationFollowListHandler, &request.DouyinRelationFollowListReq{})
}

// DouyinRelationFollowerList 用户粉丝列表 所有关注登录用户的粉丝列表。
func (c ConnectController) DouyinRelationFollowerList(ctx iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinRelationFollowerListHandler, &request.DouyinRelationFollowerListReq{})
}

// DouyinRelationFriendList 用户好友列表 所有关注登录用户的粉丝列表。
func (c ConnectController) DouyinRelationFriendList(ctx iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinRelationFriendListHandler, &request.DouyinRelationFriendListReq{})
}
