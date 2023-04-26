package models

type Favorite struct {
	BaseModel `xorm:"extends"`
	UserId    int64 `xorm:"not null comment('执行点赞的user_id') index 'user_id'"`
	VideoId   int64 `xorm:"not null comment('被点赞的video_id') index 'video_id'"`
}

func (p *Favorite) TableName() string {
	return "favorite"
}
