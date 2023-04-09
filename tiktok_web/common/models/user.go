package models

type User struct {
	BaseModel       `xorm:"extends"`
	UserId          int64  `xorm:"not null comment('用户id') index unique 'user_id'"`
	UserName        string `xorm:"not null comment('用户名') index 'username'"`
	Password        string `xorm:"not null comment('用户密码') 'password'"`
	Name            string `xorm:"default '' comment('用户名称') index 'name'"`
	FollowCount     int64  `xorm:"default 0 comment('关注总数') 'follow_count'"`
	FollowerCount   int64  `xorm:"default 0 comment('粉丝总数') 'follower_count'"`
	Avatar          string `xorm:"default '' comment('用户头像') 'avatar'"`
	BackgroundImage string `xorm:"default '' comment('用户个人页顶部大图') 'background_image'"`
	Signature       string `xorm:"default '' comment('个人简介') 'signature'"`
	WorkCount       int64  `xorm:"default 0 comment('作品数量') 'work_count'"`
}

func (p *User) TableName() string {
	return "user"
}
