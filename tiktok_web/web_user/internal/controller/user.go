package controller

import (
	"common/middleware"
	"github.com/kataras/iris/v12"
	"web_user/conf"
	"web_user/internal/service"
)

type ConnectController struct {
}

// DouyinUserRegister 用户注册接口 新用户注册时提供用户名，密码，昵称即可，用户名需要保证唯一。创建成功后返回用户 id 和权限token.
func (c ConnectController) DouyinUserRegister(ctx iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinUserRegisterHandler, conf.Logger)
}

// DouyinUserLogin 用户登录接口 通过用户名和密码进行登录，登录成功后返回用户 id 和权限 token.
func (c ConnectController) DouyinUserLogin(ctx iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinUserLoginHandler, conf.Logger)
}

// DouyinUser 用户信息 获取登录用户的 id、昵称，如果实现社交部分的功能，还会返回关注数和粉丝数。
func (c ConnectController) DouyinUser(ctx iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinUserHandler, conf.Logger)
}
