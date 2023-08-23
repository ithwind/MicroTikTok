package model

type UserVo struct {
	ID              int64  `json:"id"`
	UserName        string `json:"name"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
}
