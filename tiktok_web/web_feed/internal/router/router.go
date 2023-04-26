package router

import (
	"common/middleware"
	"github.com/kataras/iris/v12"
	v1 "web_feed/internal/router/v1"
)

// InitRouters 初始化路由
func InitRouters(app *iris.Application) {
	// 不需要鉴权
	appConfRouter := app.Party("/douyin")
	appConfRouter.HandleDir("/video", "../../../static/video")
	appConfRouter.HandleDir("/cover", "../../../static/cover")
	v1.RegisterConfigRouter(&appConfRouter)

	// 需要鉴权
	appAuthenticationRouter := app.Party("/douyin", middleware.Auth)
	v1.RegisterAuthenticationRouter(&appAuthenticationRouter)
}
