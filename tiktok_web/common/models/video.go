package models

type Video struct {
	BaseModel     `xorm:"extends"`
	AuthorId      int64  `xorm:"not null comment('视频作者id') index 'author_id'"`
	PlayUrl       string `xorm:"not null comment('视频url') 'play_url'"`
	CoverUrl      string `xorm:"not null comment('视频封面url') 'cover_url'"`
	FavoriteCount int64  `xorm:"'default 0 comment('总点赞数') 'favorite_count"`
	CommentCount  int64  `xorm:"default 0 comment('总评论数') 'comment_count'"`
	Title         string `xorm:"not null comment('标题') 'title'"`
}

func (p *Video) TableName() string {
	return "video"
}
