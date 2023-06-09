package middleware

import (
	"common/log"
	"github.com/kataras/iris/v12"
)

type ApiHandler func(ctx iris.Context, reqBody interface{})

func AipWrapper(ctx iris.Context, handler ApiHandler, reqBody interface{}) {
	// 可以将请求内容处理出来
	err := ctx.ReadJSON(reqBody)
	if err != nil {
		log.Logger.Infof(err.Error(), "请求Body失败")
	}

	handler(ctx, reqBody)
}

func Handler(f func(ctx iris.Context)) iris.Handler {
	return func(original iris.Context) {
		// 可以提前把body单独处理出来
		f(original)
	}
}
