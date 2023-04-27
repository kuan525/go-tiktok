package dao

import "common/models"

// GetMessageByFromUserIdAndPreMsgTime 根据发送userId和上次新消息时间获取所有message
func (a *Dao) GetMessageByFromUserIdAndPreMsgTime(userId, PreMsgTime int64) ([]models.Message, error) {

}

func (a *Dao) CreateMessage(message *models.Message) error {

}
