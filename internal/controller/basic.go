package controller

import (
	"github.com/kataras/iris/v12"
	"go-tiktok/internal/middleware"
	"go-tiktok/internal/service"
)

type ConnectController struct {
}

func (c ConnectController) DouyinFeed(ctx *iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinFeedHandler)
}

func (c ConnectController) DouyinUserRegister(ctx *iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinUserRegisterHandler)
}

func (c ConnectController) DouyinUserLogin(ctx *iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinUserLoginHandler)
}

func (c ConnectController) DouyinUser(ctx *iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinUserHandler)
}

func (c ConnectController) DouyinPublishAction(ctx *iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinPublishActionHandler)
}

func (c ConnectController) DouyinPublishList(ctx *iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinPublishListHandler)
}
