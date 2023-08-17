package redis

import (
	"fmt"
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
