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
