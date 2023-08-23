// Code generated by goctl. DO NOT EDIT.
package types

type Message struct {
	Id         int64  `json:"id"`
	ToUserId   int64  `json:"to_user_id"`
	FromUserId int64  `json:"from_user_id"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}

type ChatActionRequest struct {
	Token      string `form:"token"`
	ToUserId   int64  `form:"to_user_id"`
	Content    string `form:"content"`
	ActionType string `form:"action_type"`
}

type ChatActionResponse struct {
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
}

type ChatMessageRequest struct {
	Token      string `form:"token"`
	ToUserId   string `form:"to_user_id"`
	PreMsgTime int64  `form:"pre_msg_time"`
}

type ChatMessageResponse struct {
	StatusCode  string    `json:"status_code"`  // 状态码，0-成功，其他值-失败
	StatusMsg   string    `json:"status_msg"`   // 返回状态描述
	MessageList []Message `json:"message_list"` // 用户消息列表
}
