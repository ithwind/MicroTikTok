package service

import (
	"MicroTikTok/dal/video"
	cron "MicroTikTok/pkg/corn"
	"MicroTikTok/pkg/redis"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var redisService, err = redis.NewRedisCacheService()
var cornService = cron.NewCronService()

func FavoriteService() {
	/**
	1.获取redis数据
	*/
	fmt.Println("+++++++++++++++++++++++++++")

	if err != nil {
		fmt.Printf("连接redis失败\n")
	}
	//判断redis是否存在点赞数据
	flag := redisService.ExistKey()
	fmt.Println("Flag:", flag)
	//存在
	if flag == true {
		// 在协程中执行异步操作
		go func() {
			err := cornService.AddFunc("@every 5s", RedisAndDB)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			cornService.Start()
		}()
	}
	time.Sleep(5 * time.Second)
}

func RedisAndDB() {
	//取出redis中点赞数据
	redisData, err := redisService.HashSetGet()
	fmt.Println("redisData:", len(redisData))
	// 当 Redis 数据为空时，阻塞当前协程等待新的数据
	for i, v := range redisData {
		fmt.Println(i, v)
		parts := strings.Split(i, "-")
		userId, _ := strconv.Atoi(parts[0])
		videoId, _ := strconv.Atoi(parts[1])
		fmt.Printf("UserId: %v, VideoId: %v\n", userId, videoId)
		//点赞操作 插入user_video_favorite表
		if v == "1" {
			err := video.NewSetFavorite(int64(userId), int64(videoId))
			if err != nil {
				fmt.Println("InsertError:", err)
				return
			}
		}
	}
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	//处理完redis数据数据清除
	redisService.ClearUserVideoFavorite()

	cornService.Stop()
}
