package middleware

import (
	"encoding/json"
	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
)

type ApiHandler func(ctx iris.Context, req *[]byte)

func AipWrapper(ctx iris.Context, handler ApiHandler, logger *logrus.Logger) {
	// 可以将请求内容处理出来
	var reqBody interface{}
	err := ctx.ReadJSON(&reqBody)
	if err != nil {
		logger.Infof(err.Error(), "body请求读取错误")
	}

	date, _ := json.Marshal(reqBody)

	handler(ctx, &date)
}

func Handler(f func(ctx iris.Context)) iris.Handler {
	return func(original iris.Context) {
		// 可以提前把body单独处理出来
		f(original)
	}
}
