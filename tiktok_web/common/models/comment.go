package models

type Comment struct {
	BaseModel  `xorm:"extends"` // 当前评论的id
	VideoId    int64            `xorm:"not null comment('视频id') index 'video_id'"`
	UserId     int64            `xorm:"not null comment('评论用户id') index 'user_id'"`
	Content    string           `xorm:"default '' comment('评论内容') index 'content'"`
	CreateDate string           `xorm:"default '' comment('评论发布日期') index 'create_date'"` // 评论发布日期，格式 mm-dd
}

func (p *Comment) TableName() string {
	return "comment"
}
