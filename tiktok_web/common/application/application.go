package application

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

// PreSetting 注册中间件、定义错误处理
func PreSetting(app *iris.Application) {
	customLogger := logger.New(logger.Config{
		//状态显示状态代码
		Status: true,
		// IP显示请求的远程地址
		IP: true,
		//方法显示http方法
		Method: true,
		// Path显示请求路径
		Path: true,
		// Query将url查询附加到Path。
		Query: true,
	})
	app.Use(
		recover.New(),
		customLogger,
	)

	// ---------------------- 定义错误处理 ------------------------
	app.OnErrorCode(iris.StatusNotFound, customLogger, func(ctx iris.Context) {
		ctx.StatusCode(iris.StatusNotFound)
	})
}
