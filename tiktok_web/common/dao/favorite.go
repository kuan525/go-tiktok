package dao

import (
	"common/conf"
	"common/models"
)

func (a *Dao) FavoriteByUserIdAndVideoId(userId, videoId int64) error {
	exist, err := conf.Mqcli.Table(FavoriteTable.TableName()).Where("user_id = ? AND video_id = ?", userId, videoId).Exist()
	if err != nil {
		return err
	}
	if exist {
		return nil
	}
	_, err = conf.Mqcli.Table(FavoriteTable.TableName()).Insert(&models.Favorite{
		UserId:  userId,
		VideoId: videoId,
	})
	return err
}

func (a *Dao) CancelFavoriteByUserIdAndVideoId(userId, videoId int64) error {
	exist, err := conf.Mqcli.Table(FavoriteTable.TableName()).Where("user_id = ? AND video_id = ?", userId, videoId).Exist()
	if err != nil {
		return err
	}
	if !exist {
		return nil
	}
	_, err = conf.Mqcli.Table(FavoriteTable.TableName()).Where("user_id = ? AND video_id = ?", userId, videoId).Delete()
	return err
}

func (a *Dao) GetAllFavoriteVideoByUserId(userId int64) ([]int64, error) {

}
