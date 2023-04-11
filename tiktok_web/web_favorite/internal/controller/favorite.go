package controller

import (
	"common/middleware"
	"github.com/kataras/iris/v12"
	"web_favorite/conf"
	"web_favorite/internal/request"
	"web_favorite/internal/service"
)

type ConnectController struct {
}

// DouyinFavoriteAction 赞操作 登录用户对视频的点赞和取消点赞操作。
func (c ConnectController) DouyinFavoriteAction(ctx iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinFavoriteActionHandler, &request.DouyinFavoriteActionReq{}, conf.Logger)
}

// DouyinFavoriteList 喜欢列表 登录用户的所有点赞视频。
func (c ConnectController) DouyinFavoriteList(ctx iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinFavoriteListHandler, &request.DouyinFavoriteListReq{}, conf.Logger)
}
