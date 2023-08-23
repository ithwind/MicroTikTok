package service

import (
	"MicroTikTok/AcessData/video"
	cron "MicroTikTok/pkg/corn"
	"MicroTikTok/pkg/redis"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	redisService, err    = redis.NewRedisCacheService()
	cornService          = cron.NewCronService()
	favoriteMap          = make(map[string]bool) // 用户视频点赞记录
	favoriteMapMutex     = sync.Mutex{}          // 用于对 favoriteMap 进行并发安全操作的互斥锁
	pendingFavorites     = make(map[string]bool) // 待存入数据库的点赞记录
	pendingFavoritesLock = sync.Mutex{}          // 用于对 pendingFavorites 进行并发安全操作的互斥锁
)

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
			err := cornService.AddFunc("@every 120s", RedisAndDB)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			cornService.Start()
		}()
	}
	time.Sleep(2 * time.Minute)
}

func RedisAndDB() {
	//取出redis中点赞数据
	redisData, err := redisService.HashSetGet()
	if err != nil {
		fmt.Printf("Error:%v", err)
	}
	fmt.Println("redisData:", len(redisData))
	// 当 Redis 数据为空时，阻塞当前协程等待新的数据
	for i, v := range redisData {
		parts := strings.Split(i, "-")
		userId, _ := strconv.Atoi(parts[0])
		videoId, _ := strconv.Atoi(parts[1])
		key := strconv.Itoa(userId) + "-" + strconv.Itoa(videoId)

		favoriteMapMutex.Lock()
		if favoriteMap[key] {
			fmt.Printf("User %d has already liked video %d\n", userId, videoId)
			favoriteMapMutex.Unlock()
			continue
		}
		favoriteMap[key] = true
		favoriteMapMutex.Unlock()

		fmt.Printf("UserId: %v, VideoId: %v\n", userId, videoId)

		if v == "1" {
			// 点赞操作，将点赞记录存入待存入数据库的点赞记录
			pendingFavoritesLock.Lock()
			pendingFavorites[key] = true
			pendingFavoritesLock.Unlock()
		} else {
			// 取消点赞操作，将取消点赞记录存入待存入数据库的点赞记录
			pendingFavoritesLock.Lock()
			pendingFavorites[key] = false
			pendingFavoritesLock.Unlock()
		}
	}

	// 处理完 Redis 数据后，等待两分钟
	time.Sleep(2 * time.Minute)

	// 将待存入数据库的点赞记录存入数据库
	pendingFavoritesLock.Lock()
	defer pendingFavoritesLock.Unlock()

	for key, value := range pendingFavorites {
		parts := strings.Split(key, "-")
		userId, _ := strconv.Atoi(parts[0])
		videoId, _ := strconv.Atoi(parts[1])

		if value == true {
			// 实际插入数据库的点赞操作
			err := video.NewSetFavorite(int64(userId), int64(videoId))
			if err != nil {
				fmt.Println("InsertError:", err)
			}
		} else if value == false {
			// 实际取消点赞操作，从数据库中删除点赞记录
			err := video.DeleteFavorite(int64(userId), int64(videoId))
			if err != nil {
				fmt.Println("DeleteError:", err)
			}
		}

		// 移除已处理的点赞记录
		delete(pendingFavorites, key)
	}

	cornService.Stop()
}
