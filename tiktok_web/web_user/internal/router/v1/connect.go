package v1

import (
	"common/middleware"
	"github.com/kataras/iris/v12/core/router"
	"web_user/internal/controller"
)

// RegisterAuthenticationRouter 需要鉴权
func RegisterAuthenticationRouter(party *router.Party) {
	// 用户登录接口 通过用户名和密码进行登录，登录成功后返回用户 id 和权限 token.
	(*party).Handle("POST", "/user/login/", middleware.Handler(controller.ConnectController{}.DouyinUserLogin))
	// 用户信息 获取登录用户的 id、昵称，如果实现社交部分的功能，还会返回关注数和粉丝数。
	(*party).Handle("GET", "/user/", middleware.Handler(controller.ConnectController{}.DouyinUser))
}

// RegisterConfigRouter 不需要鉴权
func RegisterConfigRouter(party *router.Party) {
	// 用户注册接口 新用户注册时提供用户名，密码，昵称即可，用户名需要保证唯一。创建成功后返回用户 id 和权限token.
	(*party).Handle("POST", "/user/register/", middleware.Handler(controller.ConnectController{}.DouyinUserRegister))
}
