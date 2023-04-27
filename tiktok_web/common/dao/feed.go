package dao

import (
	"common/conf"
	"common/models"
	"time"
)

// GetVideos 返回的视频，限制最新投稿时间，要返回这个时间之前的
func (a *Dao) GetVideos(LastestTime int64) ([]models.Video, int64, error) {
	// 单次最多30个
	var feedTime time.Time
	if LastestTime != 0 {
		feedTime = time.Unix(LastestTime, 0)
	} else {
		feedTime = time.Now()
	}

	var modelsVideo []models.Video
	err := conf.Mqcli.Table(videoTable.TableName()).Desc("created_at").Where("created_at < ?", feedTime).Limit(30).Find(&modelsVideo)

	var timeLastest time.Time

	for _, video := range modelsVideo {
		if video.CreatedAt.Before(timeLastest) {
			timeLastest = video.CreatedAt
		}
	}

	return modelsVideo, timeLastest.Unix(), err
}

// IsFavorite true-已点赞，false-未点赞
func (a *Dao) IsFavorite(userId int64, videoId int64) bool {
	ok, err := conf.Mqcli.Table(favoriteTable.TableName()).Where("user_id = ? AND video_id = ?", userId, videoId).Exist()
	if err != nil {
		return false
	}
	return ok
}

// IsFollow true-已关注，false-未关注
func (a *Dao) IsFollow(Author, Readers int64) bool {
	ok, err := conf.Mqcli.Table(relationTable.TableName()).Where("user_id = ? AND be_followed = ?", Readers, Author).Exist()
	if err != nil {
		return false
	}
	return ok
}

// GetUserIdByVideoId 通过VideoId获取UserId
func (a *Dao) GetUserIdByVideoId(videoId int64) (int64, error) {
	var resp models.Video
	_, err := conf.Mqcli.Table(videoTable.TableName()).Where("video_id = ?", videoId).Get(&resp)
	return resp.UserId, err
}
