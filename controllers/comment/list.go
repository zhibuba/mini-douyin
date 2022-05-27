package comment

type CommentListResponse struct {
	CommentList []Comment `json:"comment_list"`// 评论列表
	StatusCode  int64     `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg   *string   `json:"status_msg"`  // 返回状态描述
}