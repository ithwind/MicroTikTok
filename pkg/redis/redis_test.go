package redis

import (
	"MicroTikTok/favoritelist/rpc/pb/favoritelist"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestNewRedisCacheService(t *testing.T) {
	_, err := NewRedisCacheService()
	if err != nil {
		return
	}
}

func TestServiceOfRedis_ExistKey(t *testing.T) {
	service, err := NewRedisCacheService()
	if err != nil {
		fmt.Printf("Error\n")
	}
	flag := service.ExistKey()
	fmt.Println(flag)
}

func TestServiceOfRedis_HashSetRedis(t *testing.T) {
	service, err := NewRedisCacheService()
	if err != nil {
		fmt.Printf("Error\n")
	}
	service.HashSetRedis("1", "48545", "1")
}

func TestServiceOfRedis_HashSetGet(t *testing.T) {
	service, err := NewRedisCacheService()
	if err != nil {
		fmt.Printf("Error\n")
	}
	setGet, err := service.HashSetGet()
	if err != nil {
		return
	}
	for i, v := range setGet {
		parts := strings.Split(i, "-")
		fmt.Println(parts[0], parts[1], v)
	}
}

func TestSetVideoInfo(t *testing.T) {
	service, err := NewRedisCacheService()
	if err != nil {
		fmt.Printf("Error\n")
	}
	// 构造一个示例视频
	video := favoritelist.Video{

		Id: 1001,
		Author: &favoritelist.User{
			Id:              101,
			Name:            "张三",
			FollowCount:     100,
			FollowerCount:   200,
			IsFollow:        false,
			Avatar:          "https://example.com/avatar1.jpg",
			BackgroundImage: "https://example.com/bg1.jpg",
			Signature:       "这个人很懒,没有简介",
			TotalFavorited:  1000,
			WorkCount:       10,
			FavoriteCount:   500,
		},
		PlayUrl:       "https://example.com/video/1001.mp4",
		CoverUrl:      "https://example.com/cover/1001.jpg",
		FavoriteCount: 10,
		CommentCount:  20,
		IsFavorite:    true,
		Title:         "张三去爬山",
	}
	service.SetVideoInfo(fmt.Sprintf(strconv.FormatInt(video.Id, 10)), video)
	fmt.Println(video.Id)
}

func TestGetVideoInfo(t *testing.T) {
	videoid := "1001"
	service, err := NewRedisCacheService()
	if err != nil {
		fmt.Printf("Error\n")
	}
	video, _ := service.GetVideoInfo(videoid)
	fmt.Println(videoid)
	fmt.Println(video)
}
