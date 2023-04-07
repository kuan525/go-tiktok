package controller

import (
	"github.com/kataras/iris/v12"
	"go-tiktok/internal/middleware"
	"go-tiktok/internal/service"
)

type ConnectController struct {
}

// DouyinFeed 视频流接口 不限制登录状态，返回按投稿时间倒序的视频列表，视频数由服务端控制，单次最多30个。
func (c ConnectController) DouyinFeed(ctx iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinFeedHandler)
}

// DouyinUserRegister 用户注册接口 新用户注册时提供用户名，密码，昵称即可，用户名需要保证唯一。创建成功后返回用户 id 和权限token.
func (c ConnectController) DouyinUserRegister(ctx iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinUserRegisterHandler)
}

// DouyinUserLogin 用户登录接口 通过用户名和密码进行登录，登录成功后返回用户 id 和权限 token.
func (c ConnectController) DouyinUserLogin(ctx iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinUserLoginHandler)
}

// DouyinUser 用户信息 获取登录用户的 id、昵称，如果实现社交部分的功能，还会返回关注数和粉丝数。
func (c ConnectController) DouyinUser(ctx iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinUserHandler)
}

// DouyinPublishAction 视频投稿 登录用户选择视频上传。
func (c ConnectController) DouyinPublishAction(ctx iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinPublishActionHandler)
}

// DouyinPublishList 发布列表 登录用户的视频发布列表，直接列出用户所有投稿过的视频。
func (c ConnectController) DouyinPublishList(ctx iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinPublishListHandler)
}
