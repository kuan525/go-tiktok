package middleware

import (
	"github.com/kataras/iris/v12"
	"go-tiktok/conf"
)

type ApiHandler func(ctx *iris.Context) error

func AipWrapper(ctx *iris.Context, handler ApiHandler) {
	// 待添加，参数校验等
	err := handler(ctx)
	if err != nil {
		conf.HandleLogsErr(err, "handler处理错误")
	}
}

func Handler(f func(ctx *iris.Context)) iris.Handler {
	return func(original iris.Context) {
		// 可以提前把body单独处理出来
		f(&original)
	}
}
