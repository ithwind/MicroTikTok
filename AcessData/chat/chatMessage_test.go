package chat

import (
	"MicroTikTok/AcessData/postgres"
	"fmt"
	"testing"
	"time"
)

func TestCreateMessage(t *testing.T) {
	postgres.Init()
	err := CreateMessage(1, 0, "aaa")
	if err != nil {
		return
	}
	fmt.Println(err)
}

func TestQueryMessagesByFromUserIdAndToUserId(t *testing.T) {
	postgres.Init()
	count, err := QueryMessagesByFromUserIdAndToUserId(1, 0, time.Unix(1692688320, 0))
	if err != nil {
		return
	}
	fmt.Println(count)
}
