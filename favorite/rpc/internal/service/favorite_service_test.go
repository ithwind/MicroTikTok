package service

import (
	"MicroTikTok/AcessData/mysql"
	"testing"
)

func TestFavoriteService(t *testing.T) {
	mysql.Init()
	FavoriteService()
}
