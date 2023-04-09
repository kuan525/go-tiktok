package dao

import (
	"common/middleware"
	"common/models"
	"web_user/conf"
	"web_user/internal/response"
)

type UserDao struct {
}

var table models.User

// Register 注册用户
func (a *UserDao) Register(username, password string, userId int64) bool {
	_, err := conf.Mqcli.Table(table.TableName()).Insert(&models.User{
		UserId:   userId,
		UserName: username,
		Password: password,
	})
	if err == nil {
		return true
	} else {
		return false
	}
}

// UserIsExistByUsername 查看当前用户是否存在
func (a *UserDao) UserIsExistByUsername(userName string) bool {
	ok, err := conf.Mqcli.Table(table.TableName()).Where("username = ?", userName).Exist()
	if err != nil {
		conf.Logger.Infof(err.Error(), "dao:查询数据库错误")
	}
	return ok
}

// GetTokenAndUserIdByUsernameAndPassword 根据用户名和密码得到token和userId
func (a *UserDao) GetTokenAndUserIdByUsernameAndPassword(username, password string) (string, int64, bool) {
	var user models.User
	ok, err := conf.Mqcli.Table(table.TableName()).Where("username = ? AND password = ?", username, password).Get(&user)
	if err != nil {
		conf.Logger.Infof(err.Error(), "dao:查询数据库错误\"")
		return "", 0, false
	}
	if !ok {
		conf.Logger.Infof("dao:查询数据库错误\"")
		return "", 0, false
	}

	token, err := middleware.GenerateToken(user.UserId)
	if err != nil {
		conf.Logger.Infof(err.Error(), "dao:token申请失败")
		return "", 0, false
	}

	return token, user.UserId, true
}

// GetUserByUserId 根据userId得到用户信息
func (a *UserDao) GetUserByUserId(userId int64) (*response.User, bool) {
	var resp response.User
	var user models.User
	ok, err := conf.Mqcli.Table(table.TableName()).Where("user_id = ?", userId).Get(&user)
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
