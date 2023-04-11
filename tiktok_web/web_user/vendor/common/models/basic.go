package models

import "time"

type BaseModel struct {
	Id        int64     `xorm:"pk autoincr not null comment('自增主键') int(11) 'id'"`
	CreatedAt time.Time `xorm:"created  comment('创建时间') datetime 'created_at'"`
	UpdatedAt time.Time `xorm:"updated  comment('更新时间') datetime 'updated_at'"`
	DeletedAt time.Time `xorm:"deleted  comment('删除时间') datetime 'deleted_at'"`
}
