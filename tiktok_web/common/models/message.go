package models

type Message struct {
	BaseModel  `xorm:"extends"` // 当前评论的id
	FromUserId int64            `xorm:"not null comment('该消息发送者的id') index 'from_user_id'"`
	ToUserId   int64            `xorm:"not null comment('该消息接收者的id') index 'to_user_id'"`
	Content    string           `xorm:"default '' comment('消息内容') index 'content'"`
	CreateTime int64            `xorm:"not null comment('消息创建时间') index 'create_time'"`
}

func (p *Message) TableName() string {
	return "message"
}
