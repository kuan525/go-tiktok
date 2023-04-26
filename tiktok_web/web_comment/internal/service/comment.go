package service

import (
	"common/dao"
	"common/log"
	"common/middleware"
	"common/models"
	"github.com/kataras/iris/v12"
	"time"
	"web_comment/internal/request"
	"web_comment/internal/response"
)

var Dao dao.Dao

// DouyinCommentActionHandler 评论操作 登录用户对视频进行评论。
func DouyinCommentActionHandler(ctx iris.Context, reqBody interface{}) {
	req := reqBody.(*request.DouyinCommentActionReq)
	userId := middleware.GetUserIdByToken(req.Token)
	if req.ActionType == 1 { // 发布
		modelComment := models.Comment{
			VideoId:    req.VideoId,
			UserId:     userId,
			Content:    req.CommentText,
			CreateDate: time.Now().Format("01-02"),
		}
		commonId, err := Dao.InsertComment(&modelComment)
		if err != nil {
			log.Logger.Infof(err.Error(), "写入评论失败")
			return
		}
		modelUser, ok := Dao.GetUserByUserId(userId)
		if !ok {
			log.Logger.Infof("通过userId获取user失败")
			return
		}
		AuthorId, err := Dao.GetUserIdByVideoId(req.VideoId)
		if err != nil {
			log.Logger.Infof(err.Error(), "通过VideoId获取UserId失败")
			return
		}
		respUser := response.User{
			Id:              modelUser.UserId,
			Name:            modelUser.Name,
			FollowCount:     modelUser.FollowCount,
			FollowerCount:   modelUser.FollowerCount,
			Avatar:          modelUser.Avatar,
			BackgroundImage: modelUser.BackgroundImage,
			Signature:       modelUser.Signature,
			WorkCount:       modelUser.WorkCount,
			IsFollow:        Dao.IsFollow(AuthorId, userId),
			TotalFavorited:  Dao.GetNumUserAllGetFavorite(userId),
			FavoriteCount:   Dao.GetNumFavorite(userId),
		}
		err = ctx.JSON(response.Comment{
			Id:         commonId,
			User:       respUser,
			Content:    req.CommentText,
			CreateDate: modelComment.CreateDate,
		})
		if err != nil {
			log.Logger.Infof(err.Error(), "发送失败")
			return
		}
	} else { // 删除
		err := Dao.DeleteCommentByCommentId(req.CommentId)
		if err != nil {
			log.Logger.Infof(err.Error(), "删除失败")
			return
		}
		err = ctx.JSON(response.DouyinCommentActionResp{
			StatusCode: 1,
			StatusMsg:  "删除成功",
		})
		if err != nil {
			log.Logger.Infof(err.Error(), "删除失败")
			return
		}
	}
}

// DouyinCommentListHandler 视频评论列表 查看视频的所有评论，按发布时间倒序。
func DouyinCommentListHandler(ctx iris.Context, reqBody interface{}) {
	req := reqBody.(*request.DouyinCommentListReq)
	userId := middleware.GetUserIdByToken(req.Token)

	// 按发布时间排序
	modelComments, err := Dao.GetAllCommentByVideoId(req.VideoId)
	if err != nil {
		log.Logger.Infof(err.Error(), "访问数据库获取所有评论失败")
		return
	}

	var respCommentList []response.Comment
	for i := 0; i < len(modelComments); i++ {
		modelUser, ok := Dao.GetUserByUserId(modelComments[i].UserId)
		if !ok {
			log.Logger.Infof("通过userId获取user失败")
			return
		}
		respUser := response.User{
			Id:              modelUser.UserId,
			Name:            modelUser.Name,
			FollowCount:     modelUser.FollowCount,
			FollowerCount:   modelUser.FollowerCount,
			Avatar:          modelUser.Avatar,
			BackgroundImage: modelUser.BackgroundImage,
			Signature:       modelUser.Signature,
			WorkCount:       modelUser.WorkCount,
			// 这里是当前访问的用户是否对改评论的用户关注，不是对视频作者
			IsFollow:       Dao.IsFollow(modelComments[i].UserId, userId),
			TotalFavorited: Dao.GetNumUserAllGetFavorite(userId),
			FavoriteCount:  Dao.GetNumFavorite(userId),
		}
		respComment := response.Comment{
			Id:         modelComments[i].Id,
			User:       respUser,
			Content:    modelComments[i].Content,
			CreateDate: modelComments[i].CreateDate,
		}
		respCommentList = append(respCommentList, respComment)
	}
	err = ctx.JSON(&response.DouyinCommentListResp{
		StatusCode:  1,
		StatusMsg:   "获取成功",
		CommentList: respCommentList,
	})
	if err != nil {
		log.Logger.Infof(err.Error(), "发送失败")
		return
	}
}
