package router

import (
	"github.com/kataras/iris/v12"
	v1 "go-tiktok/internal/router/v1"
)

// InitRouters 初始化路由
func InitRouters(app *iris.Application) {
	appRouter := app.Party("/douyin")
	v1.RegisterConfigRouter(&appRouter)
}
