package chat

import (
	"MicroTikTok/AcessData/mysql"
	"fmt"
	"testing"
	"time"
)

func TestCreateMessage(t *testing.T) {
	mysql.Init()
	err := CreateMessage(1, 0, "bb")
	if err != nil {
		return
	}
	fmt.Println(err)
}

func TestQueryMessagesByFromUserIdAndToUserId(t *testing.T) {
	mysql.Init()
	count, err := QueryMessagesByFromUserIdAndToUserId(1, 0, time.Unix(1692688320, 0))
	if err != nil {
		return
	}
	fmt.Println(count)
}
