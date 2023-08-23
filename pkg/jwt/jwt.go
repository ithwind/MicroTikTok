package jwt

import (
	. "MicroTikTok/AcessData/modelVo"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type MyClaims struct {
	UserId int
	UserVo
	jwt.RegisteredClaims
}

// 签名密钥
const key = "ithWind"

// GenerateToken /*生成token*/
func GenerateToken(id int, user UserVo) string {
	claims := MyClaims{
		UserId: id,
		UserVo: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)), //设置超时时间
			Subject:   "IthWind",
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(key))

	fmt.Println("GenerateError:", err)
	return token
}

// ParseToken /*解析token*/
func ParseToken(tokenStr string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	fmt.Println("ParseError:", err)
	if claims, ok := token.Claims.(*MyClaims); ok {
		return claims, nil
	} else {
		return nil, err
	}

}
