package service

import (
	"MicroTikTok/AcessData/postgres"
	"testing"
)

func TestFavoriteService(t *testing.T) {
	postgres.Init()
	FavoriteService()
}
