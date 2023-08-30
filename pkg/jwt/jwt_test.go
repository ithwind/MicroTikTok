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
	claims, _ := ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjIsImlkIjoyLCJuYW1lIjoiMjkxMzEzNTY3MkBxcS5jb20iLCJwYXNzd29yZCI6IiIsImF2YXRhciI6ImF2YXRhciIsImJhY2tncm91bmRfaW1hZ2UiOiJiYWNrZ3JvdW5kaW1hZ2UiLCJzaWduYXR1cmUiOiLor6XnlKjmiLfku4DkuYjpg73msqHnlZnkuIsiLCJzdWIiOiJJdGhXaW5kIiwiZXhwIjoxNjkyNDcyMzY5fQ.o4C4g_uNKzvDMHvZqb2gLow4YhZCAvjTtSsuO3fJ7Lo")

	fmt.Println(claims.User)
}
