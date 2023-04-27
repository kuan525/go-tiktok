package request

type DouyinRelationActionReq struct {
	Token      string `json:"token"`       // 用户鉴权token
	ToUser     int64  `json:"toUser"`      // 对方用户id
	ActionType int32  `json:"action_type"` // 1-关注，2-取消关注
}

type DouyinRelationFollowListReq struct {
	UserId int64  `json:"user_id"` // 用户id
	Token  string `json:"token"`   // 用户鉴权token
}

type DouyinRelationFollowerListReq struct {
	UserId int64  `json:"user_id"` // 用户id
	Token  string `json:"token"`   // 用户鉴权token
}

//type DouyinRelationFriendListReq struct {
//	UserId int64  `json:"user_id"` // 用户id
//	Token  string `json:"token"`   // 用户鉴权token
//}
