package user

import (
	"MicroTikTok/dal/postgres"
	"fmt"
	"testing"
)

func TestGetIsFavoriteByUserId(t *testing.T) {
	postgres.Init()
	a := GetIsFavoriteByUserId(1, 1)

	fmt.Println(a)
}

func TestGetWorkCountByUserId(t *testing.T) {
	postgres.Init()
	fmt.Println(GetWorkCountByUserId(1))
}

func TestGetUserIdByVideoId(t *testing.T) {
	postgres.Init()
	fmt.Println(GetUserIdByVideoId(1))
}

func TestGetFollowCountByUserId(t *testing.T) {
	postgres.Init()
	fmt.Println(GetFollowCountByUserId(1))
}

func TestGetFollowerCountByUserId(t *testing.T) {
	postgres.Init()
	fmt.Println(GetFollowerCountByUserId(1))
}

func TestGetFavoriteCount(t *testing.T) {
	postgres.Init()

	fmt.Println(GetFavoriteCount(1))
}

func TestGetTotalFavoriteCount(t *testing.T) {
	postgres.Init()
	fmt.Println(GetTotalFavoriteCount(1))
}

func TestAddUserVideoTable(t *testing.T) {
	postgres.Init()
	err := AddUserVideoTable(99, 50)
	if err != nil {
		return
	}
}

func TestGetUserById(t *testing.T) {
	postgres.Init()
	fmt.Println(GetUserById(1))
}
