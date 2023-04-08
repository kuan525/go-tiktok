package response

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

type DouyinUserRegisterResp struct {
	StatusCode int32  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
	UserId     int64  `json:"user_id"`     // 用户id
	Token      string `json:"token"`       // 用户鉴权token
}

type DouyinUserLoginResp struct {
	StatusCode int32  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
	UserId     string `json:"user_id"`     // 用户id
	Token      string `json:"token"`       // 用户鉴权token
}

type DouyinUserResp struct {
	StatusCode int32  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
	User       User   `json:"user"`        // 用户信息
}
