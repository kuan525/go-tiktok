package controller

import (
	"github.com/kataras/iris/v12"
	"go-tiktok/internal/middleware"
	"go-tiktok/internal/service"
)

func (c ConnectController) DouyinFavoriteAction(ctx *iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinFavoriteActionHandler)
}

func (c ConnectController) DouyinFavoriteList(ctx *iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinFavoriteListHandler)
}

func (c ConnectController) DouyinCommentAction(ctx *iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinCommentActionHandler)
}

func (c ConnectController) DouyinCommentList(ctx *iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinCommentListHandler)
}
