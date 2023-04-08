package request

type User struct {
	Name            string `json:"name"`             // 用户id
	Id              int64  `json:"id"`               // 用户名称
	FollowCount     int64  `json:"follow_count"`     // 关注总数
	FollowerCount   int64  `json:"follower_count"`   // 粉丝总数
	IsFollow        bool   `json:"is_follow"`        // true-已关注，false-未关注
	Avatar          string `json:"avatar"`           //用户头像
	BackgroundImage string `json:"background_image"` //用户个人页顶部大图
	Signature       string `json:"signature"`        //个人简介
	TotalFavorited  int64  `json:"total_favorited"`  //获赞数量
	WorkCount       int64  `json:"work_count"`       //作品数量
	FavoriteCount   int64  `json:"favorite_count"`   //点赞数量
}

type DouyinCommentActionReq struct {
	Token       string `json:"token"`        // 用户鉴权token
	VideoId     int64  `json:"video_id"`     // 视频id
	ActionType  int32  `json:"action_type"`  // 1-发布评论，2-删除评论
	CommentText string `json:"comment_text"` // 用户填写的评论内容，在action_type=1的时候使用
	CommentId   int64  `json:"comment_id"`   // 要删除的评论id，在action_type=2的时候使用
}

type DouyinCommentListReq struct {
	Token   string `json:"token"`    // 用户鉴权token
	VideoId int64  `json:"video_id"` // 视频id
}
