package controller

import (
	"github.com/kataras/iris/v12"
	"go-tiktok/internal/middleware"
	"go-tiktok/internal/service"
)

// DouyinFavoriteAction 赞操作 登录用户对视频的点赞和取消点赞操作。
func (c ConnectController) DouyinFavoriteAction(ctx iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinFavoriteActionHandler)
}

// DouyinFavoriteList 喜欢列表 登录用户的所有点赞视频。
func (c ConnectController) DouyinFavoriteList(ctx iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinFavoriteListHandler)
}

// DouyinCommentAction 评论操作 登录用户对视频进行评论。
func (c ConnectController) DouyinCommentAction(ctx iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinCommentActionHandler)
}

// DouyinCommentList 视频评论列表 查看视频的所有评论，按发布时间倒序。
func (c ConnectController) DouyinCommentList(ctx iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinCommentListHandler)
}
