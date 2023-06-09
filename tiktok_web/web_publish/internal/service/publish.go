package service

import (
	"common/conf"
	"common/dao"
	"common/ip"
	"common/log"
	"common/middleware"
	"common/models"
	"fmt"
	"github.com/kataras/iris/v12"
	"os/exec"
	"strconv"
	"web_publish/internal/request"
	"web_publish/internal/response"
)

var Dao dao.Dao

// DouyinPublishActionHandler 视频投稿 登录用户选择视频上传。
func DouyinPublishActionHandler(ctx iris.Context, reqBody interface{}) {
	req := reqBody.(*request.DouyinPublishActionReq)

	token := req.Token
	title := req.Title
	userId := middleware.GetUserIdByToken(token)

	// 获取文件名，文件类型（文件名中包含了文件格式）
	// file是内容，可以使用流的那一套使用，第二个是元数据包括文件名等等，第三个是错误信息
	file, header, err := ctx.FormFile("data")
	if err != nil {
		log.Logger.Infof(err.Error(), "获取文件失败")
		return
	}
	// 关闭文件
	defer file.Close()
	// 获取文件名
	fileName := header.Filename

	// 生成视频名称 = 配置文件路径 +（雪花id）+ 文件名字（自带后缀）防止出现一样的视频名字
	UserId := middleware.GetSnowflakeId(conf.Cfg.HttpAddr.Port)
	fileName = strconv.Itoa(int(UserId)) + fileName
	// 将新生成的文件名赋值上去，防止出现重复
	header.Filename = fileName

	// 保存文件，第一个是文件元数据，第二个是保存路径
	_, err = ctx.SaveFormFile(header, conf.Cfg.StaticConf.VideoPath)
	if err != nil {
		log.Logger.Infof(err.Error(), "文件存储失败")
		return
	}
	fileUrl := fmt.Sprintf("http://%s:%s/douyin/video/%s", ip.GetIp(conf.Cfg.HttpAddr.NetEnv), conf.Cfg.HttpAddr.Port, fileName)

	// 获取封面名
	coverName := GetAndSaveCover(fileName)
	coverUrl := fmt.Sprintf("http://%s:%s/douyin/cover/%s", ip.GetIp(conf.Cfg.HttpAddr.NetEnv), conf.Cfg.HttpAddr.Port, coverName)

	user, ok := Dao.GetUserByUserId(userId)
	if !ok {
		log.Logger.Infof("访问数据库失败")
		return
	}
	video := response.Video{
		Id: middleware.GetSnowflakeId(conf.Cfg.HttpAddr.Port),
		Author: response.User{
			Id:              user.UserId,
			Name:            user.Name,
			FollowCount:     user.FollowCount,
			FollowerCount:   user.FollowerCount,
			Avatar:          user.Avatar,
			BackgroundImage: user.BackgroundImage,
			Signature:       user.Signature,
			WorkCount:       user.WorkCount,
			IsFollow:        false,
			TotalFavorited:  Dao.GetNumUserAllGetFavorite(user.UserId),
			FavoriteCount:   Dao.GetNumFavorite(user.UserId),
		},
		PlayUrl:       fileUrl,
		CoverUrl:      coverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Title:         title,
	}

	err = Dao.Insert(&models.Video{
		UserId:        video.Author.Id,
		PlayUrl:       video.PlayUrl,
		CoverUrl:      video.CoverUrl,
		FavoriteCount: video.FavoriteCount,
		CommentCount:  video.CommentCount,
		Title:         video.Title,
	})
	if err != nil {
		log.Logger.Infof(err.Error(), "插入数据库失败")
		return
	}

	resp := response.DouyinPublishActionResp{
		StatusCode: 0,
		StatusMsg:  "发布成功",
	}

	err = ctx.JSON(resp)
	if err != nil {
		log.Logger.Infof(err.Error(), "发送失败")
	}
}

// GetAndSaveCover 保存封面并返回封面名称
func GetAndSaveCover(fileName string) string {
	// 截掉视频文件后缀格式
	videoDir := conf.Cfg.StaticConf.VideoPath + fileName
	coverName := fileName[:len(fileName)-6] + ".png"
	coverDir := conf.Cfg.StaticConf.CoverPath + coverName
	cmd := exec.Command(conf.Cfg.StaticConf.Tool+"ffmpeg", "-i", videoDir, "-vframes", "1", coverDir)
	err := cmd.Run()
	if err != nil {
		log.Logger.Infof(err.Error(), "ffmpeg截取封面失败")
	}
	return coverName
}

// DouyinPublishListHandler 发布列表 登录用户的视频发布列表，直接列出用户所有投稿过的视频。
func DouyinPublishListHandler(ctx iris.Context, reqBody interface{}) {
	req := reqBody.(*request.DouyinPublishListReq)
	userId := req.UserId

	// 点赞，评论总数为0
	modelsVideoList, err := Dao.GetVideoListByUserId(userId)
	if err != nil {
		log.Logger.Infof(err.Error(), "访问数据库失败")
		return
	}

	var respVideoList []response.Video
	for _, modelVideo := range modelsVideoList {
		user, ok := Dao.GetUserByUserId(userId)
		if !ok {
			log.Logger.Infof("访问数据库失败")
			return
		}
		respVideo := response.Video{
			Id:            modelVideo.Id,
			PlayUrl:       modelVideo.PlayUrl,
			CoverUrl:      modelVideo.CoverUrl,
			FavoriteCount: modelVideo.FavoriteCount,
			CommentCount:  modelVideo.CommentCount,
			Title:         modelVideo.Title,
			IsFavorite:    Dao.IsFavorite(userId, modelVideo.Id),
			Author: response.User{
				Id:              user.UserId,
				Name:            user.Name,
				FollowCount:     user.FollowCount,
				FollowerCount:   user.FollowerCount,
				Avatar:          user.Avatar,
				BackgroundImage: user.BackgroundImage,
				Signature:       user.Signature,
				WorkCount:       user.WorkCount,
				IsFollow:        false,
				TotalFavorited:  Dao.GetNumUserAllGetFavorite(user.UserId),
				FavoriteCount:   Dao.GetNumFavorite(user.UserId),
			},
		}
		respVideoList = append(respVideoList, respVideo)
	}

	err = ctx.JSON(response.DouyinPublishListResp{
		StatusCode: 0,
		StatusMsg:  "获取成果",
		VideoList:  respVideoList,
	})

	if err != nil {
		log.Logger.Infof(err.Error(), "发送失败")
	}
}
