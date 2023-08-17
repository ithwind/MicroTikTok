package video

import (
	"MicroTikTok/dal/postgres"
	"fmt"
	"testing"
	"time"
)

func TestGetVideosBeforeLastTime(t *testing.T) {
	postgres.Init()
	videos, err := GetVideosBeforeLastTime(time.Now())
	if err != nil {
		return
	}
	fmt.Printf("%v", videos)
}

func TestGetFavoriteCountByVideoId(t *testing.T) {
	postgres.Init()
	println(GetFavoriteCountByVideoId(2))
}

func TestUpdateVideo(t *testing.T) {
	postgres.Init()
	newVideo := &Video{
		PlayURL:     "fsfsdf",
		CoverURL:    "fsdsfsf",
		PublishTime: time.Now(),
		Title:       "sdwres",
	}
	UpdateVideo(newVideo)
}

func TestQueryVideosByUserId(t *testing.T) {
	postgres.Init()
	fmt.Println(QueryVideoIdsByUserId(1))
}

func TestGetVideoById(t *testing.T) {
	postgres.Init()
	fmt.Println(GetVideoById(1))
}

func TestGetPublishTest(t *testing.T) {
	postgres.Init()
	fmt.Println(GetPublishList(1))
}

func TestNewSetFavorite(t *testing.T) {
	postgres.Init()
	err := NewSetFavorite(565, 5656)
	if err != nil {
		return
	}
}
