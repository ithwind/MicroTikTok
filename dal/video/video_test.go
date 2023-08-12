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
