package service

import (
	"MicroTikTok/dal/postgres"
	"MicroTikTok/pb/video"
	"MicroTikTok/pkg/util"
	"fmt"
	"testing"
	"time"
)

func TestFeedService_Feed(t *testing.T) {
	postgres.Init()
	var latest = time.Now().Unix()
	request := video.FeedRequest{
		LatestTime: &latest,
		Token:      util.String("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsIklkIjoxLCJOYW1lIjoiaXRod2luZCIsIkF2YXRhciI6Ind3dy5iYWlkdS5jb20iLCJCYWNrZ3JvdW5kSW1hZ2UiOiJ3d3cuYmFpZHUuY29tIiwiU2lnbmF0dXJlIjoid3d3dyIsInN1YiI6Ikl0aFdpbmQiLCJleHAiOjE2OTE3Mzk2ODB9.QtN0eX8I_mmyG5_E5lrzDcDZNcj7ixiAVHV2QkSNmNY"),
	}
	response, err := NewFeedSerVice().Feed(&request)
	if err != nil {
		return
	}
	fmt.Printf("%v", response)
}
