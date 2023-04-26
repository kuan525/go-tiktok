package dao

import (
	"common/conf"
	"common/models"
)

func (a *Dao) Insert(video *models.Video) error {
	_, err := conf.Mqcli.Table(videoTable.TableName()).Insert(&video)
	return err
}

func (a *Dao) GetVideoListByUserId(userId int64) ([]models.Video, error) {
	var videoList []models.Video
	_, err := conf.Mqcli.Table(videoTable.TableName()).Where("author_id = ?", userId).Get(&videoList)
	return videoList, err
}

func (a *Dao) GetVideoByVideoId(videoId int64) (models.Video, error) {

}
