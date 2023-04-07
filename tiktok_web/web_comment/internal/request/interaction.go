package request

type DouyinFavoriteActionReq struct {
	Token      string `json:"token"`       // 用户鉴权token
	VideoId    int64  `json:"video_id"`    // 视频id
	ActionType int32  `json:"action_type"` // 1-点赞，2-取消点赞
}

type DouyinFavoriteListReq struct {
	UserId int64  `json:"user_id"` // 用户id
	Token  string `json:"token"`   // 用户鉴权token
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
