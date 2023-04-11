package controller

import (
	"common/middleware"
	"github.com/kataras/iris/v12"
	"web_comment/conf"
	"web_comment/internal/request"
	"web_comment/internal/service"
)

type ConnectController struct {
}

// DouyinCommentAction 评论操作 登录用户对视频进行评论。
func (c ConnectController) DouyinCommentAction(ctx iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinCommentActionHandler, &request.DouyinCommentActionReq{}, conf.Logger)
}

// DouyinCommentList 视频评论列表 查看视频的所有评论，按发布时间倒序。
func (c ConnectController) DouyinCommentList(ctx iris.Context) {
	middleware.AipWrapper(ctx, service.DouyinCommentListHandler, &request.DouyinCommentListReq{}, conf.Logger)
}
