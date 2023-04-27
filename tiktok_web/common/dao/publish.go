package dao

import (
	"common/conf"
	"common/models"
)

// Insert 插入视频
func (a *Dao) Insert(video *models.Video) error {
	_, err := conf.Mqcli.Table(videoTable.TableName()).Insert(&video)
	if err != nil {
		return err
	}
	_, err = conf.Mqcli.Table(userTable.TableName()).Where("user_id = ?", video.UserId).Incr("work_count", 1).Update(&userTable)
	return err
}

// GetVideoListByUserId 通过userId获取videoId
func (a *Dao) GetVideoListByUserId(userId int64) ([]models.Video, error) {
	var videoList []models.Video
	_, err := conf.Mqcli.Table(videoTable.TableName()).Where("author_id = ?", userId).Get(&videoList)
	return videoList, err
}

// GetVideoByVideoId 通过videoId获取video
func (a *Dao) GetVideoByVideoId(videoId int64) (models.Video, error) {
	var resp models.Video
	_, err := conf.Mqcli.Table(videoTable.TableName()).Where("video_id = ?", videoId).Get(&resp)
	return resp, err
}
