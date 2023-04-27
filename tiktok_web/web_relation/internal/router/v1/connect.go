package v1

import (
	"common/middleware"
	"github.com/kataras/iris/v12/core/router"
	"web_relation/internal/controller"
)

// RegisterAuthenticationRouter 需要鉴权
func RegisterAuthenticationRouter(party *router.Party) {
	// 社交接口 实现用户之间的关注关系维护，登录用户能够关注或取关其他用户，同时自己能够看到自己关注过的所有用户列表，以及所有关注自己的用户列表。
	(*party).Handle("POST", "/relation/action/ ", middleware.Handler(controller.ConnectController{}.DouyinRelationAction))
	// 用户关注列表 登录用户关注的所有用户列表。
	(*party).Handle("GET", "/relation/follow/list/", middleware.Handler(controller.ConnectController{}.DouyinRelationFollowList))
	// 用户粉丝列表 所有关注登录用户的粉丝列表。
	(*party).Handle("GET", "/relation/follower/list/", middleware.Handler(controller.ConnectController{}.DouyinRelationFollowerList))
	//// 用户好友列表 所有关注登录用户的粉丝列表。
	//(*party).Handle("GET", "/relation/friend/list/", middleware.Handler(controller.ConnectController{}.DouyinRelationFriendList))
}

// RegisterConfigRouter 不需要鉴权
func RegisterConfigRouter(party *router.Party) {
}
