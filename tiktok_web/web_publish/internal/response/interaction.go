package response

type DouyinFavoriteActionResp struct {
	StatusCode int32  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

type DouyinFavoriteListResp struct {
	StatusCode int32  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
	VideoList  Video  `json:"video_list"`  // 用户点赞视频列表
}

type Comment struct {
	Id         int64  `json:"id"`          // 视频评论id
	User       User   `json:"user"`        // 评论用户信息
	Content    string `json:"content"`     // 评论内容
	CreateDate string `json:"create_date"` // 评论发布日期，格式 mm-dd
}

type DouyinCommentActionResp struct {
	StatusCode int32   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string  `json:"status_msg"`  // 返回状态描述
	Comment    Comment `json:"comment"`     // 评论成功返回评论内容，不需要重新拉取整个列表
}

type DouyinCommentListResp struct {
	StatusCode  int32   `json:"status_code"`
	StatusMsg   string  `json:"status_msg"`
	CommentList Comment `json:"comment_list"`
}
