// Code generated by goctl. DO NOT EDIT.
package types

type RegisterRequest struct {
	Password string `form:"password"` // 密码，最长32个字符
	Username string `form:"username"` // 注册用户名，最长32个字符
}

type RegisterResponse struct {
	StatusCode int64   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg"`  // 返回状态描述
	Token      string  `json:"token"`       // 用户鉴权token
	UserID     int64   `json:"user_id"`     // 用户id
}

type LoginRequest struct {
	Password string `form:"password"` // 登录密码
	Username string `form:"username"` // 登录用户名
}

type LoginResponse struct {
	StatusCode int64   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg"`  // 返回状态描述
	Token      *string `json:"token"`       // 用户鉴权token
	UserID     *int64  `json:"user_id"`     // 用户id
}

type UserInfoRequest struct {
	Token  string `form:"token"`   // 用户鉴权token
	UserID int64  `form:"user_id"` // 用户id
}

type UserInfoResponse struct {
	StatusCode int64   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg"`  // 返回状态描述
	User       *User   `json:"user"`        // 用户信息
}

type User struct {
	Avatar          string `json:"avatar"`           // 用户头像
	BackgroundImage string `json:"background_image"` // 用户个人页顶部大图
	FavoriteCount   int64  `json:"favorite_count"`   // 喜欢数
	FollowCount     int64  `json:"follow_count"`     // 关注总数
	FollowerCount   int64  `json:"follower_count"`   // 粉丝总数
	ID              int64  `json:"id"`               // 用户id
	Name            string `json:"name"`             // 用户名称
	Signature       string `json:"signature"`        // 个人简介
	TotalFavorited  int64  `json:"total_favorited"`  // 获赞数量
	WorkCount       int64  `json:"work_count"`       // 作品数
}
