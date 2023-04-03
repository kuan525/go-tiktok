package v1

import "github.com/kataras/iris/v12/core/router"

func RegisterConfigRouter(party *router.Party) {
	// 视频流接口 不限制登录状态，返回按投稿时间倒序的视频列表，视频数由服务端控制，单次最多30个。
	(*party).Handle("GET", "/feed/", nil)
	// 用户注册接口 新用户注册时提供用户名，密码，昵称即可，用户名需要保证唯一。创建成功后返回用户 id 和权限token.
	(*party).Handle("POST", "/user/register/", nil)
	// 用户登录接口 通过用户名和密码进行登录，登录成功后返回用户 id 和权限 token.
	(*party).Handle("POST", "/user/login/", nil)
	// 用户信息 获取登录用户的 id、昵称，如果实现社交部分的功能，还会返回关注数和粉丝数。
	(*party).Handle("GET", "/user/", nil)
	// 视频投稿 登录用户选择视频上传。
	(*party).Handle("POST", "/publish/action/", nil)
	// 发布列表 登录用户的视频发布列表，直接列出用户所有投稿过的视频。
	(*party).Handle("GET", "/publish/list/ ", nil)

	// 赞操作 登录用户对视频的点赞和取消点赞操作。
	(*party).Handle("POST", "/favorite/action/", nil)
	// 喜欢列表 登录用户的所有点赞视频。
	(*party).Handle("GET", "/favorite/list/", nil)
	// 评论操作 登录用户对视频进行评论。
	(*party).Handle("POST", "/comment/action/", nil)
	// 视频评论列表 查看视频的所有评论，按发布时间倒序。
	(*party).Handle("GET", "/comment/list/", nil)

	// 社交接口 实现用户之间的关注关系维护，登录用户能够关注或取关其他用户，同时自己能够看到自己关注过的所有用户列表，以及所有关注自己的用户列表。
	(*party).Handle("POST", "/relation/action/ ", nil)
	// 用户关注列表 登录用户关注的所有用户列表。
	(*party).Handle("GET", "/relation/follow/list/", nil)
	// 用户粉丝列表 所有关注登录用户的粉丝列表。
	(*party).Handle("GET", "/relation/follower/list/", nil)
	// 用户好友列表 所有关注登录用户的粉丝列表。
	(*party).Handle("GET", "/relation/friend/list/", nil)

	// 聊天记录 当前登录用户和其他指定用户的聊天消息记录
	(*party).Handle("GET", "/message/chat/ ", nil)
	// 消息操作 登录用户对消息的相关操作，目前只支持消息发送
	(*party).Handle("POST", "/relation/friend/list/", nil)
}
