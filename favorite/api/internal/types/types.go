// Code generated by goctl. DO NOT EDIT.
package types

type FavoriteRequest struct {
	Token      string `form:"token"`       //用户鉴权oken
	VideoId    int64  `form:"video_id"`    //视频id
	ActionType string `form:"action_type"` //1-点赞 2-未点赞
}

type FavoriteResponse struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}
