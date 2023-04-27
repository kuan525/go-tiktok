package request

type DouyinMessageChatReq struct {
	Token  string `json:"token"`   // 用户鉴权token
	ToUser int64  `json:"to_user"` // 对方用户id
	// 返回报文是string format格式的时间，默认当作是前端自己处理的（文档定死）
	PreMsgTime int64 `json:"pre_msg_time"` //上次最新消息的时间（新增字段-apk更新中）(默认时间戳吧)
}

type DouyinMessageActionReq struct {
	Token      string `json:"token"`       // 用户鉴权token
	ToUserId   int64  `json:"to_user_id"`  // 对方用户id
	ActionType int32  `json:"action_type"` // 1-发送消息
	Content    string `json:"content"`     // 消息内容
}
