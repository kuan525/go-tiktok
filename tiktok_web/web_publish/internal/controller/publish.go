package controller

import (
	"common/middleware"
	"github.com/kataras/iris/v12"
	"web_publish/internal/request"
	"web_publish/internal/service"
)

type ConnectController struct {
}

// DouyinPublishAction 视频投稿 登录用户选择视频上传。
func (c ConnectController) DouyinPublishAction(ctx iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinPublishActionHandler, &request.DouyinPublishActionReq{})
}

// DouyinPublishList 发布列表 登录用户的视频发布列表，直接列出用户所有投稿过的视频。
func (c ConnectController) DouyinPublishList(ctx iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinPublishListHandler, &request.DouyinPublishListReq{})
}
