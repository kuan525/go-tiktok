package dao

import "common/models"

// InsertComment 返回插入的数据的id（评论id）
func (a *Dao) InsertComment(comment *models.Comment) (int64, error) {

}

func (a *Dao) DeleteCommentByCommentId(commentId int64) error {

}

// GetAllCommentByVideoId 按发布时间倒序。
func (a *Dao) GetAllCommentByVideoId(videoId int64) ([]models.Comment, error) {

}
