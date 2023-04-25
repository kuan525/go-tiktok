package dao

import (
	"common/conf"
	"common/log"
	"common/models"
	"web_publish/internal/response"
)

type PublishDao struct {
}

var (
	userTable  models.User
	videoTable models.Video
)

// GetUserByUserId 根据userId得到用户信息
func (a *PublishDao) GetUserByUserId(userId int64) (*response.User, bool) {
	var resp response.User
	var user models.User
	ok, err := conf.Mqcli.Table(userTable.TableName()).Where("user_id = ?", userId).Get(&user)
	if err != nil {
		log.Logger.Infof(err.Error(), "dao:查询数据库失败")
		return &response.User{}, false
	}
	if !ok {
		log.Logger.Infof("dao:查询数据库失败")
		return &response.User{}, false
	}
	resp.Id = userId
	resp.Name = user.Name
	resp.FollowCount = user.FollowCount
	resp.FollowerCount = user.FollowerCount
	resp.IsFollow = false // 还不知道是干什么的
	resp.Avatar = user.Avatar
	resp.BackgroundImage = user.BackgroundImage
	resp.Signature = user.Signature
	resp.TotalFavorited = 0 // 未知
	resp.WorkCount = user.WorkCount
	resp.FollowCount = 0 // 未知

	return &resp, true
}

func (a *PublishDao) Insert(video *response.Video) error {
	_, err := conf.Mqcli.Table(videoTable.TableName()).Insert(&models.Video{
		AuthorId:      video.Id,
		PlayUrl:       video.PlayUrl,
		CoverUrl:      video.CoverUrl,
		FavoriteCount: video.FavoriteCount,
		CommentCount:  video.CommentCount,
		Title:         video.Title,
	})
	return err
}

func (a *PublishDao) GetVideoListByUserId(userId int64) ([]response.Video, error) {
	var videoList []response.Video
	_, err := conf.Mqcli.Table(videoTable.TableName()).Where("author_id = ?", userId).Get(&videoList)
	return videoList, err
}
