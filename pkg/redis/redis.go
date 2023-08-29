package redis

import (
	"MicroTikTok/favoritelist/rpc/pb/favoritelist"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var tableKey = "video-favorite"

// 定义视频信息表名
var videoTable = "video_info"

type ServiceOfRedis struct {
	client *redis.Client
}

func NewRedisCacheService() (*ServiceOfRedis, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
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

func (r *ServiceOfRedis) VideoKey() bool {
	flag, err := r.client.Exists(context.Background(), videoTable).Result()
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

// SetVideoInfo 存入视频信息
func (r *ServiceOfRedis) SetVideoInfo(videoId string, videoInfo favoritelist.Video) error {

	// 构造哈希表的key
	key := "video:" + videoId

	// 设置视频信息
	ctx := context.Background()
	result, err := r.client.HSet(ctx, videoTable, key, videoInfo).Result()

	if err != nil {
		// 返回错误
		return err
	}

	fmt.Println("Set video success:", result)

	return nil
}

// GetVideoInfo 从redis取出视频信息
func (r *ServiceOfRedis) GetVideoInfo(videoId string) (*favoritelist.Video, error) {

	// 构造redis key
	key := "video:" + videoId

	// 从redis获取视频信息
	ctx := context.Background()
	videoInfo, err := r.client.HGet(ctx, videoTable, key).Result()
	if err != nil {
		return nil, err
	}
	fmt.Printf("redis raw result: %v\n", videoInfo)
	// 反序列化videoInfo到Video对象
	var video favoritelist.Video
	if err = json.Unmarshal([]byte(videoInfo), &video); err != nil {
		return nil, err
	}

	return &video, nil
}

// ClearUserVideoFavorite 清空当前表中数据
func (r *ServiceOfRedis) ClearUserVideoFavorite() {
	err := r.client.Del(context.Background(), tableKey).Err()
	if err != nil {
		panic("删除失败")
	}
}
