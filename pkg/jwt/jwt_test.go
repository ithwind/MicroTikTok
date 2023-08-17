package jwt

import (
	"MicroTikTok/feed/model"
	"fmt"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	u := model.User{
		ID:              1,
		UserName:        "ithwind",
		Avatar:          "www.baidu.com",
		BackgroundImage: "www.baidu.com",
		Signature:       "wwww",
	}

	token := GenerateToken(1, u)
	fmt.Println(token)
}

func TestParseToken(t *testing.T) {
	claims, _ := ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsIklkIjoxLCJOYW1lIjoiaXRod2luZCIsIkF2YXRhciI6Ind3dy5iYWlkdS5jb20iLCJCYWNrZ3JvdW5kSW1hZ2UiOiJ3d3cuYmFpZHUuY29tIiwiU2lnbmF0dXJlIjoid3d3dyIsInN1YiI6Ikl0aFdpbmQiLCJleHAiOjE2OTE3Mzk2ODB9.QtN0eX8I_mmyG5_E5lrzDcDZNcj7ixiAVHV2QkSNmNY")

	fmt.Println(claims.User)
}
