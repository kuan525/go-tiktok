package models

type Relation struct {
	BaseModel  `xorm:"extends"`
	UserIdTo   int64 `xorm:"not null comment('关注者id') index 'user_id_to'"`
	BeFollowed int64 `xorm:"not null comment('被关注者id') index 'be_followed'"`
}

func (p *Relation) TableName() string {
	return "relation"
}
