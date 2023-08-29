package service

import (
	cron "MicroTikTok/pkg/corn"
	"MicroTikTok/pkg/redis"
	"fmt"
	"sync"
)

var (
	redisService, err       = redis.NewRedisCacheService()
	cornService             = cron.NewCronService()
	pendingFavoriteListLock = sync.Mutex{} // 用于对 pendingFavoriteList 进行并发安全操作的互斥锁
)

func FavoriteListService() {
	if err != nil {
		fmt.Printf("连接redis失败\n")
	}

}
