package jwt

import (
	"MicroTikTok/AcessData/modelVo"
	"fmt"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	u := model.UserVo{
		ID:              2,
		UserName:        "ithwind",
		Avatar:          "www.baidu.com",
		BackgroundImage: "www.baidu.com",
		Signature:       "wwww",
	}

	token := GenerateToken(1, u)
	fmt.Println(token)
}

func TestParseToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImlkIjoxLCJuYW1lIjoiaXRod2luZCIsImF2YXRhciI6Ind3dy5iYWlkdS5jb20iLCJiYWNrZ3JvdW5kX2ltYWdlIjoid3d3LmJhaWR1LmNvbSIsInNpZ25hdHVyZSI6Ind3d3ciLCJzdWIiOiJJdGhXaW5kIiwiZXhwIjoxNjkyNzc3MDY4fQ.ygpX_3JVnouPwR29TzctNeNgmZLJ1r-I0-BDZPoCDLw"
	claims, _ := ParseToken(token)

	fmt.Println(claims.UserVo)
}
