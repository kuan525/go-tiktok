package dao

import (
	"common/models"
	"web_publish/conf"
	"web_publish/internal/response"
)

type PublishDao struct {
}

var userTable models.User

// GetUserByUserId 根据userId得到用户信息
func (a *PublishDao) GetUserByUserId(userId int64) (*response.User, bool) {
	var resp response.User
	var user models.User
	ok, err := conf.Mqcli.Table(userTable.TableName()).Where("user_id = ?", userId).Get(&user)
	if err != nil {
		conf.Logger.Infof(err.Error(), "dao:查询数据库失败")
		return &response.User{}, false
	}
	if !ok {
		conf.Logger.Infof("dao:查询数据库失败")
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

func (a *PublishDao) Insert(video *response.Video) {

}
