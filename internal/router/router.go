package router

import (
	"github.com/kataras/iris/v12"
	"go-tiktok/internal/middleware"
	v1 "go-tiktok/internal/router/v1"
)

// InitRouters 初始化路由
func InitRouters(app *iris.Application) {
	appAuthenticationRouter := app.Party("/douyin", middleware.Auth)
	v1.RegisterAuthenticationRouter(appAuthenticationRouter)

	// 可以路由组前缀一样，但是要是后面也一样的话，后面的会覆盖前面的
	appConfRouter := app.Party("/douyin")
	v1.RegisterConfigRouter(appConfRouter)
}
