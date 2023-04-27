package dao

import (
	"common/conf"
	"common/log"
	"common/middleware"
	"common/models"
)

type Dao struct {
}

var (
	userTable     models.User
	videoTable    models.Video
	favoriteTable models.Favorite
	relationTable models.Relation
	commonTable   models.Comment
	messageTable  models.Message
)

// Register 注册用户
func (a *Dao) Register(username, password string, userId int64) bool {
	_, err := conf.Mqcli.Table(userTable.TableName()).Insert(&models.User{
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
func (a *Dao) UserIsExistByUsername(userName string) bool {
	ok, err := conf.Mqcli.Table(userTable.TableName()).Where("username = ?", userName).Exist()
	if err != nil {
		log.Logger.Infof(err.Error(), "dao:查询数据库错误")
	}
	return ok
}

// GetTokenAndUserIdByUsernameAndPassword 根据用户名和密码得到token和userId
func (a *Dao) GetTokenAndUserIdByUsernameAndPassword(username, password string) (string, int64, bool) {
	var user models.User
	ok, err := conf.Mqcli.Table(userTable.TableName()).Where("username = ? AND password = ?", username, password).Get(&user)
	if err != nil {
		log.Logger.Infof(err.Error(), "dao:查询数据库错误\"")
		return "", 0, false
	}
	if !ok {
		log.Logger.Infof("dao:查询数据库错误\"")
		return "", 0, false
	}

	token, err := middleware.GenerateToken(user.UserId)
	if err != nil {
		log.Logger.Infof(err.Error(), "dao:token申请失败")
		return "", 0, false
	}

	return token, user.UserId, true
}

// GetUserByUserId 根据userId得到用户信息
func (a *Dao) GetUserByUserId(userId int64) (*models.User, bool) {
	var user models.User
	ok, err := conf.Mqcli.Table(userTable.TableName()).Where("user_id = ?", userId).Get(&user)
	if err != nil {
		log.Logger.Infof(err.Error(), "dao:查询数据库失败")
		return &models.User{}, false
	}
	if !ok {
		log.Logger.Infof("dao:查询数据库失败")
		return &models.User{}, false
	}
	return &user, true
}
