package service

import (
	"common/dao"
	"common/log"
	"common/middleware"
	"github.com/kataras/iris/v12"
	"web_feed/internal/request"
	"web_feed/internal/response"
)

var Dao dao.Dao

// DouyinFeedHandler 视频流接口 不限制登录状态，返回按投稿时间倒序的视频列表，视频数由服务端控制，单次最多30个。
func DouyinFeedHandler(ctx iris.Context, reqBody interface{}) {
	req := reqBody.(*request.DouyinFeedReq)

	// 和当前视频是否点赞有关，以及next_time有关
	var userId int64
	if req.Token != "" {
		userId = middleware.GetUserIdByToken(req.Token)
	}

	var nextTime int64
	modelsVideos, nextTime, err := Dao.GetVideos(req.LastestTime)
	if err != nil {
		log.Logger.Infof(err.Error(), "获取视频失败")
		return
	}

	var respVideos []response.Video
	for _, video := range modelsVideos {
		modelsUserId := video.UserId
		// 查user表，这个结构体是model中的user，要变成
		modelsUser, ok := Dao.GetUserByUserId(modelsUserId)
		if !ok {
			log.Logger.Infof("访问数据库失败")
			return
		}

		modelsVideo := response.Video{
			Id: video.Id,
			Author: response.User{
				Name:            modelsUser.Name,
				Id:              modelsUser.UserId,
				FollowCount:     modelsUser.FollowerCount,
				FollowerCount:   modelsUser.FollowerCount,
				Avatar:          modelsUser.Avatar,
				BackgroundImage: modelsUser.BackgroundImage,
				Signature:       modelsUser.Signature,
				WorkCount:       modelsUser.WorkCount,
				IsFollow:        Dao.IsFollow(modelsUser.UserId, userId),
				TotalFavorited:  Dao.GetNumUserAllGetFavorite(modelsUser.UserId),
				FavoriteCount:   Dao.GetNumUserAllToFavorite(modelsUser.UserId),
			},
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: Dao.GetNumFavorite(video.Id),
			CommentCount:  Dao.GetNumComment(video.Id),
			IsFavorite:    Dao.IsFavorite(userId, video.Id),
			Title:         video.Title,
		}
		respVideos = append(respVideos, modelsVideo)
	}

	err = ctx.JSON(response.DouyinFeedResp{
		StatusCode: 0,
		StatusMsg:  "获取成果",
		VideoList:  respVideos,
		NextTime:   nextTime,
	})

	if err != nil {
		log.Logger.Infof(err.Error(), "发送失败")
	}
}
