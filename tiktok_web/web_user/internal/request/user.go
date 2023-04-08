package request

type DouyinUserRegisterReq struct {
	UserName string `json:"username"` // 注册用户名，最长32个字符
	Password string `json:"password"` // 密码，最长32个字符
}

type DouyinUserLoginReq struct {
	UserName string `json:"username"` // 登录用户名
	Password string `json:"password"` // 登录密码
}

type DouyinUserReq struct {
	UserId int64  `json:"user_id"` // 用户id
	Token  string `json:"token"`   // 用户鉴权token
}
