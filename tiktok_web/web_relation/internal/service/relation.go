package service

import "github.com/kataras/iris/v12"

// DouyinRelationActionHandler 社交接口 实现用户之间的关注关系维护，登录用户能够关注或取关其他用户，同时自己能够看到自己关注过的所有用户列表，以及所有关注自己的用户列表。
func DouyinRelationActionHandler(ctx iris.Context, reqBody interface{}) {

}

// DouyinRelationFollowListHandler 用户关注列表 登录用户关注的所有用户列表。
func DouyinRelationFollowListHandler(ctx iris.Context, reqBody interface{}) {
}

// DouyinRelationFollowerListHandler 用户粉丝列表 所有关注登录用户的粉丝列表。
func DouyinRelationFollowerListHandler(ctx iris.Context, reqBody interface{}) {
}

// DouyinRelationFriendListHandler 用户好友列表 所有关注登录用户的粉丝列表。
func DouyinRelationFriendListHandler(ctx iris.Context, reqBody interface{}) {
}
