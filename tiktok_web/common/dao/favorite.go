package dao

import (
	"common/conf"
	"common/models"
)

// FavoriteByUserIdAndVideoId 点赞
func (a *Dao) FavoriteByUserIdAndVideoId(userId, videoId int64) error {
	exist, err := conf.Mqcli.Table(favoriteTable.TableName()).Where("user_id = ? AND video_id = ?", userId, videoId).Exist()
	if err != nil {
		return err
	}
	if exist {
		return nil
	}
	_, err = conf.Mqcli.Table(favoriteTable.TableName()).Insert(&models.Favorite{
		UserId:  userId,
		VideoId: videoId,
	})
	if err != nil {
		return err
	}

	_, err = conf.Mqcli.Table(videoTable.TableName()).Where("id = ?", videoId).Incr("comment_count", 1).Update(&videoTable)
	return err
}

// CancelFavoriteByUserIdAndVideoId 取消点赞
func (a *Dao) CancelFavoriteByUserIdAndVideoId(userId, videoId int64) error {
	exist, err := conf.Mqcli.Table(favoriteTable.TableName()).Where("user_id = ? AND video_id = ?", userId, videoId).Exist()
	if err != nil {
		return err
	}
	if !exist {
		return nil
	}
	_, err = conf.Mqcli.Table(favoriteTable.TableName()).Where("user_id = ? AND video_id = ?", userId, videoId).Delete()
	if err != nil {
		return err
	}

	_, err = conf.Mqcli.Table(videoTable.TableName()).Where("id = ?", videoId).Decr("comment_count", 1).Update(&videoTable)
	return err
}

// GetAllFavoriteVideoByUserId 根据userId获取喜欢的视频列表
func (a *Dao) GetAllFavoriteVideoByUserId(userId int64) ([]int64, error) {
	var FavoriteList []models.Favorite
	_, err := conf.Mqcli.Table(favoriteTable.TableName()).Where("user_id = ?", userId).Get(&FavoriteList)
	var resp []int64
	for _, favorite := range FavoriteList {
		resp = append(resp, favorite.VideoId)
	}

	return resp, err
}

// GetNumFavorite 获取获赞数量
func (a *Dao) GetNumFavorite(videoId int64) int64 {
	var resp models.Video
	_, err := conf.Mqcli.Table(videoTable.TableName()).Where("id = ?", videoTable).Get(&resp)
	if err != nil {
		return int64(-1)
	}
	return resp.FavoriteCount
}

// GetNumUserAllToFavorite 获取当前用户的所有点赞数量
func (a *Dao) GetNumUserAllToFavorite(userId int64) int64 {
	resp, err := conf.Mqcli.Table(favoriteTable.TableName()).Where("user_id = ?", userId).Count(&favoriteTable)
	if err != nil {
		return int64(-1)
	}
	return resp
}

// GetNumUserAllGetFavorite 获取当前用户的所有获赞数量
func (a *Dao) GetNumUserAllGetFavorite(userId int64) int64 {
	var VideoList []models.Video
	_, err := conf.Mqcli.Table(videoTable.TableName()).Where("user_id = ?", userId).Get(&VideoList)
	if err != nil {
		return int64(-1)
	}
	var resp int64
	for _, video := range VideoList {
		resp += video.FavoriteCount
	}
	return resp
}
