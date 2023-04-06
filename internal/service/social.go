package service

import "github.com/kataras/iris/v12"

// DouyinRelationActionHandler 社交接口 实现用户之间的关注关系维护，登录用户能够关注或取关其他用户，同时自己能够看到自己关注过的所有用户列表，以及所有关注自己的用户列表。
func DouyinRelationActionHandler(ctx iris.Context) error {
	var err error
	return err
}

// DouyinRelationFollowListHandler 用户关注列表 登录用户关注的所有用户列表。
func DouyinRelationFollowListHandler(ctx iris.Context) error {
	var err error
	return err
}

// DouyinRelationFollowerListHandler 用户粉丝列表 所有关注登录用户的粉丝列表。
func DouyinRelationFollowerListHandler(ctx iris.Context) error {
	var err error
	return err
}

// DouyinRelationFriendListHandler 用户好友列表 所有关注登录用户的粉丝列表。
func DouyinRelationFriendListHandler(ctx iris.Context) error {
	var err error
	return err
}

// DouyinMessageChatHandler 聊天记录 当前登录用户和其他指定用户的聊天消息记录
func DouyinMessageChatHandler(ctx iris.Context) error {
	var err error
	return err
}

// DouyinMessageActionHandler 消息操作 登录用户对消息的相关操作，目前只支持消息发送
func DouyinMessageActionHandler(ctx iris.Context) error {
	var err error
	return err
}
