package dao

import (
	"common/conf"
	"common/models"
)

// GetMessageByFromUserIdAndPreMsgTime 根据发送userId和上次新消息时间获取所有message
func (a *Dao) GetMessageByFromUserIdAndPreMsgTime(userId, PreMsgTime int64) ([]models.Message, error) {
	var resp []models.Message
	_, err := conf.Mqcli.Table(messageTable.TableName()).Where("user_id = ? AND create_time >= ?", userId, PreMsgTime).Get(&resp)
	return resp, err
}

// CreateMessage 新增聊天
func (a *Dao) CreateMessage(message *models.Message) error {
	_, err := conf.Mqcli.Table(commonTable.TableName()).Insert(message)
	return err
}
