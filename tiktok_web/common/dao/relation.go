package dao

import (
	"common/conf"
	"common/models"
)

// AddRelation userIdTo关注BeFollowed
func (a *Dao) AddRelation(userIdTo, BeFollowed int64) error {
	_, err := conf.Mqcli.Table(relationTable.TableName()).Insert(&models.Relation{
		UserIdTo:   userIdTo,
		BeFollowed: BeFollowed,
	})
	if err != nil {
		return err
	}
	_, err = conf.Mqcli.Table(userTable.TableName()).Where("user_id = ?", userIdTo).Incr("follow_count", 1).Update(&userTable)
	if err != nil {
		return err
	}
	_, err = conf.Mqcli.Table(userTable.TableName()).Where("user_id = ?", BeFollowed).Incr("follower_count", 1).Update(&userTable)
	return err
}

// DeleteRelation 取消关注
func (a *Dao) DeleteRelation(userIdTo, BeFollowed int64) error {
	_, err := conf.Mqcli.Table(relationTable.TableName()).Where("user_id_to = ? AND be_followed = ?", userIdTo, BeFollowed).Delete()
	if err != nil {
		return err
	}
	_, err = conf.Mqcli.Table(userTable.TableName()).Where("user_id = ?", userIdTo).Decr("follow_count", 1).Update(&userTable)
	if err != nil {
		return err
	}
	_, err = conf.Mqcli.Table(userTable.TableName()).Where("user_id = ?", BeFollowed).Decr("follower_count", 1).Update(&userTable)
	return err
}

// GetFollowListByUserId 登录用户关注的所有用户列表
func (a *Dao) GetFollowListByUserId(userId int64) ([]int, error) {
	var resp []int
	_, err := conf.Mqcli.Table(relationTable.TableName()).Where("user_id_to = ?", userId).Get(&resp)
	return resp, err
}

// GetFollowerListerByUserId 通过userId得到FollowList
func (a *Dao) GetFollowerListerByUserId(userId int64) ([]int, error) {
	var resp []int
	_, err := conf.Mqcli.Table(relationTable.TableName()).Where("be_followed = ?", userId).Get(&resp)
	return resp, err
}
