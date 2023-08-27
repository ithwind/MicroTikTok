package logic

import (
	"MicroTikTok/AcessData/mysql"
	"MicroTikTok/feed/rpc/internal/svc"
	"MicroTikTok/feed/rpc/pb/video"
	"context"
	"fmt"
	"testing"
)

func TestFeedLogic_Feed(t *testing.T) {
	mysql.Init()
	var req = video.FeedRequest{
		LatestTime: nil,
		Token:      nil,
	}
	feedResponse, err := NewFeedLogic(context.Background(), &svc.ServiceContext{}).Feed(&req)
	if err != nil {
		return
	}
	fmt.Printf("%v", feedResponse)
	fmt.Println(feedResponse)
}
