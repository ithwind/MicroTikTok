package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var tableKey = "video-favorite"

type ServiceOfRedis struct {
	client *redis.Client
}

func NewRedisCacheService() (*ServiceOfRedis, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "172.17.24.94:6379",
		Password: "",
		DB:       0,
	})

	// 检查连接是否正常
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return &ServiceOfRedis{
		client: client,
	}, nil
}

func (r *ServiceOfRedis) Set(key, value string, expiration time.Duration) error {
	err := r.client.Set(context.Background(), key, value, expiration).Err()
	return err
}

// SetRedisValue 存入token key值为username value为token
func (r *ServiceOfRedis) SetRedisValue(username string, token string) {
	var ctx = context.Background()

	r.client.Set(ctx, username, token, 30*time.Minute)
}

// HashSetRedis 将点赞操作存入redis 存入形式为userId-videoId favoriteType
func (r *ServiceOfRedis) HashSetRedis(userId string, videoId string, favoriteType string) {
	var ctx = context.Background()
	var key = userId + "-" + videoId
	result, err := r.client.HSet(ctx, tableKey, key, favoriteType).Result()
	fmt.Println("Error:", err)
	if err != nil {
		return
	}
	fmt.Println("InsertRedis:", result)

}

// GetValueByUserName 通过username取token
func (r *ServiceOfRedis) GetValueByUserName(username string) (string, error) {
	var ctx = context.Background()
	token, err := r.client.Get(ctx, username).Result()
	return token, err
}

func (r *ServiceOfRedis) ExistKey() bool {
	flag, err := r.client.Exists(context.Background(), tableKey).Result()
	if err != nil || flag == 0 {
		return false
	}
	return true
}

func (r *ServiceOfRedis) HashSetGet() (map[string]string, error) {
	result, err := r.client.HGetAll(context.Background(), tableKey).Result()
	if err != nil {
		fmt.Println("获取哈希表所有字段和值失败：", err)
		return nil, err
	}
	return result, nil
}

// ClearUserVideoFavorite 清空当前表中数据
func (r *ServiceOfRedis) ClearUserVideoFavorite() {
	err := r.client.Del(context.Background(), tableKey).Err()
	if err != nil {
		panic("删除失败")
	}
}
