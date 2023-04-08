package router

import (
	"common/middleware"
	"github.com/kataras/iris/v12"
	v1 "web_user/internal/router/v1"
)

// InitRouters 初始化路由
func InitRouters(app *iris.Application) {
	// 需要鉴权
	appAuthenticationRouter := app.Party("/douyin", middleware.Auth)
	v1.RegisterAuthenticationRouter(&appAuthenticationRouter)

	// 不需要鉴权
	appConfRouter := app.Party("/douyin")
	v1.RegisterConfigRouter(&appConfRouter)
}
