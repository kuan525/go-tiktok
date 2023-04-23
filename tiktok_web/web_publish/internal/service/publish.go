package service

import (
	"common/middleware"
	"github.com/kataras/iris/v12"
	"os/exec"
	"strconv"
	"web_publish/conf"
	"web_publish/internal/dao"
	"web_publish/internal/request"
	"web_publish/internal/response"
)

var publishDao dao.PublishDao

// DouyinPublishActionHandler 视频投稿 登录用户选择视频上传。
func DouyinPublishActionHandler(ctx iris.Context, reqBody interface{}) {
	req := reqBody.(*request.DouyinPublishActionReq)

	token := req.Token
	title := req.Title
	userId := middleware.GetUserIdByToken(token)

	// 获取文件名，文件类型（文件名中包含了文件格式）
	file, header, err := ctx.FormFile("data")
	if err != nil {
		conf.Logger.Infof(err.Error(), "获取文件失败")
		return
	}
	defer file.Close()
	fileName := header.Filename

	// 生成视频名称 = 配置文件路径 +（雪花id）+ 文件名字（自带后缀）
	UserId := middleware.GetSnowflakeId(conf.Cfg.HttpAddr.Port)
	fileName = strconv.Itoa(int(UserId)) + fileName
	filepath := conf.Cfg.StaticConf.VoidePath + fileName

	// 保存文件
	_, err = ctx.SaveFormFile(header, filepath)
	if err != nil {
		conf.Logger.Infof(err.Error(), "文件存储失败")
		return
	}

	// 获取封面名， 并保存，传如视频路径
	coverpath := GetAndSaveCover(fileName)

	// 将视频信息（url）存入数据库
	user, ok := publishDao.GetUserByUserId(userId)
	if !ok {
		conf.Logger.Infof("访问数据库失败")
	}
	video := response.Video{
		Id:            middleware.GetSnowflakeId(conf.Cfg.HttpAddr.Port),
		Author:        *user,
		PlayUrl:       filepath,
		CoverUrl:      coverpath,
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Title:         title,
	}

	publishDao.Insert(&video)

	err = ctx.JSON(video)
	if err != nil {
		conf.Logger.Infof(err.Error(), "发送失败")
	}
}

func GetAndSaveCover(fileName string) string {
	// 截掉视频文件后缀格式
	videoDir := conf.Cfg.StaticConf.VoidePath + fileName
	coverDir := conf.Cfg.StaticConf.CoverPath + fileName[:len(fileName)-8] + ".png"
	cmd := exec.Command(conf.Cfg.StaticConf.Tool+"ffmpeg", "-i", videoDir, "-vframes", "1", coverDir)
	err := cmd.Run()
	if err != nil {
		conf.Logger.Infof(err.Error(), "ffmpeg截取封面失败")
	}
	return coverDir
}

// DouyinPublishListHandler 发布列表 登录用户的视频发布列表，直接列出用户所有投稿过的视频。
func DouyinPublishListHandler(ctx iris.Context, reqBody interface{}) {
}
