package video

import (
	"MicroTikTok/AcessData/mysql"
	"fmt"
	"testing"
	"time"
)

func TestGetVideosBeforeLastTime(t *testing.T) {
	mysql.Init()
	videos, err := GetVideosBeforeLastTime(time.Now())

	if err != nil {
		return
	}
	for _, video := range videos {
		fmt.Println(video)
	}
}

func TestGetFavoriteCountByVideoId(t *testing.T) {
	mysql.Init()
	println(GetFavoriteCountByVideoId(2))
}

func TestUpdateVideo(t *testing.T) {
	mysql.Init()
	newVideo := &Video{
		PlayURL:     "fsfsdf",
		CoverURL:    "fsdsfsf",
		PublishTime: time.Now(),
		Title:       "sdwres",
	}
	UpdateVideo(newVideo)
}

func TestQueryVideosByUserId(t *testing.T) {
	mysql.Init()
	fmt.Println(QueryVideoIdsByUserId(1))
}

func TestGetVideoById(t *testing.T) {
	mysql.Init()
	fmt.Println(GetVideoById(1))
}

func TestGetPublishTest(t *testing.T) {
	mysql.Init()
	fmt.Println(GetPublishList(1))
}

func TestNewSetFavorite(t *testing.T) {
	mysql.Init()
	err := NewSetFavorite(565, 5656)
	if err != nil {
		return
	}
}

func TestDeleteFavorite(t *testing.T) {
	mysql.Init()
	err := DeleteFavorite(1, 57)
	if err != nil {
		return
	}
}

func TestGetCommentCount(t *testing.T) {
	mysql.Init()
	commentCount, err := GetCommentCount(1)
	if err != nil {
		return
	}

	fmt.Println(commentCount)
}
