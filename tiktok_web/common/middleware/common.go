package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
	"io"
)

type ApiHandler func(ctx iris.Context, req interface{}) error

func AipWrapper(ctx iris.Context, handler ApiHandler, logger *logrus.Logger) {
	// 可以将请求内容处理出来
	var reqBody interface{}

	// io.Copy替代ioutil.ReadAll的原因是后者会一次性把缓冲区占满（溢出），前者是存一点，转一点到内存中
	buf := new(bytes.Buffer)
	_, err := io.Copy(buf, ctx.Request().Body)
	if err != nil {
		logger.Infof("body数据获取失败")
	}
	defer ctx.Request().Body.Close()
	body := buf.Bytes()

	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		logger.Infof("body数据解析失败")
	}

	err = handler(ctx, reqBody)
	if err != nil {
		logger.Error(err, "handler处理错误")
	}
}

func Handler(f func(ctx iris.Context)) iris.Handler {
	return func(original iris.Context) {
		// 可以提前把body单独处理出来
		f(original)
	}
}
