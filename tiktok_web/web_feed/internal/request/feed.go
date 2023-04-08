package request

type DouyinFeedReq struct {
	LastestTime int64  `json:"latest_time"` // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	Token       string `json:"token"`       // 可选参数，登录用户设置
}
