package logic

import (
	chat2 "MicroTikTok/AcessData/chat"
	"context"
	"fmt"
	"strconv"
	"time"

	"MicroTikTok/chat/rpc/internal/svc"
	"MicroTikTok/chat/rpc/pb/chat"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChatMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatMessageLogic {
	return &ChatMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChatMessageLogic) ChatMessage(in *chat.ChatMessageRequest) (*chat.ChatMessageResponse, error) {
	fmt.Println("RpcParam:", in.FromUserId, "RpcParam:", in.ToUserId, "RpcParam:", in.PreMsgTime)

	/**
	1.获取聊天记录
	2.形成返回List
	*/
	fromUserId, _ := strconv.Atoi(in.FromUserId)
	toUserId, _ := strconv.Atoi(in.ToUserId)
	//查询大于当前时间的聊天记录
	messages, err := chat2.QueryMessagesByFromUserIdAndToUserId(int64(fromUserId), int64(toUserId), time.Unix(in.PreMsgTime, 0))
	var chatMessageList []*chat.Message
	for _, message := range *messages {
		msg := convertChatRecordToMessage(&message)
		chatMessageList = append(chatMessageList, msg)
	}
	if err != nil {
		return nil, err
	}

	return &chat.ChatMessageResponse{
		Status:      true,
		MessageList: chatMessageList,
	}, nil
}

func convertChatRecordToMessage(record *chat2.RecordChat) *chat.Message {
	return &chat.Message{
		Id:         record.Id,
		FromUserId: record.FromUserId,
		ToUserId:   record.ToUserId,
		Content:    record.Content,
		CreateTime: record.CreatedAt.Unix(),
	}
}
