package response

type Message struct {
	Id         int64  `json:"id"`           // 消息id
	ToUserId   int64  `json:"to_user_id"`   // 该消息接收者的id
	FromUserId int64  `json:"from_user_id"` // 该消息发送者的id
	Content    string `json:"content"`      // 消息内容
	CreateTime string `json:"create_time"`  // 消息创建时间
}

type DouyinMessageChatResp struct {
	StatusCode  int32   `json:"status_code"`  // 状态码，0-成功，其他值-失败
	StatusMsg   string  `json:"status_msg"`   // 返回状态描述
	MessageList Message `json:"message_list"` // 消息列表
}

type DouyinMessageActionResp struct {
	StatusCode int32  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}
