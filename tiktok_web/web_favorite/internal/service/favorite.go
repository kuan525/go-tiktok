package service

import (
	"common/dao"
	"common/log"
	"common/middleware"
	"github.com/kataras/iris/v12"
	"web_favorite/internal/request"
	"web_favorite/internal/response"
)

var Dao dao.Dao

// DouyinFavoriteActionHandler 赞操作 登录用户对视频的点赞和取消点赞操作。
func DouyinFavoriteActionHandler(ctx iris.Context, reqBody interface{}) {
	req := reqBody.(*request.DouyinFavoriteActionReq)
	userId := middleware.GetUserIdByToken(req.Token)

	if req.ActionType == 1 {
		err := Dao.FavoriteByUserIdAndVideoId(userId, req.VideoId)
		if err != nil {
			log.Logger.Infof(err.Error(), "查询数据库失败")
			return
		}
		err = ctx.JSON(response.DouyinFavoriteActionResp{
			StatusCode: 0,
			StatusMsg:  "点赞成功",
		})
		if err != nil {
			log.Logger.Infof(err.Error(), "发送失败")
		}
	} else {
		err := Dao.CancelFavoriteByUserIdAndVideoId(userId, req.VideoId)
		if err != nil {
			log.Logger.Infof(err.Error(), "查询数据库失败")
			return
		}
		err = ctx.JSON(response.DouyinFavoriteActionResp{
			StatusCode: 0,
			StatusMsg:  "取消点赞成功",
		})
		if err != nil {
			log.Logger.Infof(err.Error(), "发送失败")
		}
	}
}

// DouyinFavoriteListHandler 喜欢列表 登录用户的所有点赞视频。
func DouyinFavoriteListHandler(ctx iris.Context, reqBody interface{}) {
	req := reqBody.(*request.DouyinFavoriteListReq)
	userId := req.UserId

	favoriteVideo, err := Dao.GetAllFavoriteVideoByUserId(userId)
	if err != nil {
		log.Logger.Infof(err.Error(), "访问数据库失败")
		return
	}

	var respVideoList []response.Video
	for _, videoId := range favoriteVideo {
		modelVideo, err := Dao.GetVideoByVideoId(videoId)
		if err != nil {
			log.Logger.Infof(err.Error(), "访问数据库失败")
			return
		}
		authorId := modelVideo.UserId
		authorModelUser, ok := Dao.GetUserByUserId(authorId)
		if !ok {
			log.Logger.Infof("访问数据库失败")
			return
		}
		authorResqUser := response.User{
			Name:            authorModelUser.Name,
			Id:              authorModelUser.UserId,
			FollowerCount:   authorModelUser.FollowerCount,
			FollowCount:     authorModelUser.FollowCount,
			IsFollow:        Dao.IsFollow(userId, authorId),
			Avatar:          authorModelUser.Avatar,
			BackgroundImage: authorModelUser.BackgroundImage,
			Signature:       authorModelUser.Signature,
			TotalFavorited:  Dao.GetNumUserAllGetFavorite(authorId),
			WorkCount:       authorModelUser.WorkCount,
			FavoriteCount:   Dao.GetNumUserAllToFavorite(authorId),
		}
		respVideo := response.Video{
			Id:            modelVideo.Id,
			Author:        authorResqUser,
			PlayUrl:       modelVideo.PlayUrl,
			CoverUrl:      modelVideo.CoverUrl,
			FavoriteCount: modelVideo.FavoriteCount,
			CommentCount:  modelVideo.CommentCount,
			IsFavorite:    Dao.IsFavorite(userId, modelVideo.Id),
			Title:         modelVideo.Title,
		}
		respVideoList = append(respVideoList, respVideo)
	}

	err = ctx.JSON(response.DouyinFavoriteListResp{
		StatusCode: 0,
		StatusMsg:  "获取成功",
		VideoList:  respVideoList,
	})
	if err != nil {
		log.Logger.Infof(err.Error(), "发送失败")
	}
}
