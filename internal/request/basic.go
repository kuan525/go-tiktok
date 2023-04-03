package request

type DouyinFeedReq struct {
	LastestTime int64  `json:"latest_time"` // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	Token       string `json:"token"`       // 可选参数，登录用户设置
}

type DouyinUserRegisterReq struct {
	UserName string `json:"userName"` // 注册用户名，最长32个字符
	Password string `json:"password"` // 密码，最长32个字符
}

type DouyinUserLoginReq struct {
	UserName string `json:"username"` // 登录用户名
	Password string `json:"password"` // 登录密码
}

type DouyinUserReq struct {
	UserId string `json:"user_id"` // 用户id
	Token  string `json:"token"`   // 用户鉴权token
}

type DouyinPublishActionReq struct {
	Token string `json:"token"` // 用户鉴权token
	Data  []byte `json:"data"`  // 视频数据
	Title string `json:"title"` // 视频标题
}

type DouyinPublishListReq struct {
	UserId int64  `json:"userId"` // 用户id
	Token  string `json:"token"`  // 用户鉴权token
}
