package model

type User struct {
	ID              int64  `json:"id"`
	UserName        string `json:"name"`
	Password        string `json:"password"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
}
