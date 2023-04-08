package request

type DouyinPublishActionReq struct {
	Token string `json:"token"` // 用户鉴权token
	Data  []byte `json:"data"`  // 视频数据
	Title string `json:"title"` // 视频标题
}

type DouyinPublishListReq struct {
	UserId int64  `json:"user_id"` // 用户id
	Token  string `json:"token"`   // 用户鉴权token
}
