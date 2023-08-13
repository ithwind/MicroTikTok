package logic

import (
	"MicroTikTok/feed/rpc/internal/svc"
	"MicroTikTok/feed/rpc/pb/video"
	"context"
	"testing"
)

func TestNewUploadLogic(t *testing.T) {

	var req video.PublishActionRequest
	req.Title = "asfd"
	_, err := NewUploadLogic(context.Background(), &svc.ServiceContext{}).Upload(&req)
	if err != nil {
		return
	}

}
