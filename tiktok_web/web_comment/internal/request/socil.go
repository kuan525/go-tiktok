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

type DouyinRelationFriendListReq struct {
	UserId int64  `json:"user_id"` // 用户id
	Token  string `json:"token"`   // 用户鉴权token
}

type DouyinMessageChatReq struct {
	Token      string `json:"token"`        // 用户鉴权token
	ToUser     int64  `json:"to_user"`      // 对方用户id
	PreMsgTime int64  `json:"pre_msg_time"` //上次最新消息的时间（新增字段-apk更新中）
}

type DouyinMessageActionReq struct {
	Token      string `json:"token"`       // 用户鉴权token
	ToUserId   int64  `json:"to_user_id"`  // 对方用户id
	ActionType int32  `json:"action_type"` // 1-发送消息
	Content    string `json:"content"`     // 消息内容
}
