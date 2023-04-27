package dao

import (
	"common/conf"
	"common/models"
	"sort"
)

// GetNumComment 获取评论数量
func (a *Dao) GetNumComment(videoId int64) int64 {
	var resp models.Video
	ok, err := conf.Mqcli.Table(videoTable.TableName()).Where("id = ?", videoId).Get(&resp)
	if !ok || err != nil {
		return int64(-1)
	}
	return resp.CommentCount
}

// InsertComment 返回插入的数据的id（评论id）
func (a *Dao) InsertComment(comment *models.Comment) (int64, error) {
	id, err := conf.Mqcli.Table(commonTable.TableName()).Insert(comment)
	if err != nil {
		return id, err
	}
	_, err = conf.Mqcli.Table(videoTable.TableName()).Where("id = ?", comment.VideoId).Incr("comment_count", 1).Update(&videoTable)
	return id, err
}

// DeleteCommentByCommentId 根据CommentId删除
func (a *Dao) DeleteCommentByCommentId(commentId int64) error {
	var comment models.Comment
	_, err := conf.Mqcli.Table(commonTable.TableName()).Where("comment_id = ?", commentId).Get(&comment)
	if err != nil {
		return err
	}
	_, err = conf.Mqcli.Table(commonTable.TableName()).Where("comment_id = ?", commentId).Delete()
	if err != nil {
		return err
	}
	_, err = conf.Mqcli.Table(videoTable.TableName()).Where("id = ?", comment.VideoId).Decr("comment_count", 1).Update(&videoTable)
	return err
}

// GetAllCommentByVideoId 按发布时间倒序。
func (a *Dao) GetAllCommentByVideoId(videoId int64) ([]models.Comment, error) {
	var resp []models.Comment
	_, err := conf.Mqcli.Table(commonTable.TableName()).Where("video_id = ?", videoId).Get(&resp)
	sort.Slice(resp, func(i, j int) bool {
		return resp[i].CreateDate > resp[j].CreateDate
	})
	return resp, err
}
