package service

import (
	"MicroTikTok/dal/postgres"
	"testing"
)

func TestFavoriteService(t *testing.T) {
	postgres.Init()
	FavoriteService()
}
