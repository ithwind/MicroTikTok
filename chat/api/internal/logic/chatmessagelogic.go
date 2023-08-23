package logic

import (
	"MicroTikTok/chat/rpc/pb/chat"
	"MicroTikTok/pkg/jwt"
	"context"
	"fmt"
	"strconv"

	"MicroTikTok/chat/api/internal/svc"
	"MicroTikTok/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatMessageLogic {
	return &ChatMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatMessageLogic) ChatMessage(req *types.ChatMessageRequest) (*types.ChatMessageResponse, error) {
	// todo: add your logic here and delete this line
	fmt.Println("Token:", req.Token)
	fmt.Println("UserId:", req.ToUserId)
	fmt.Println("TimeStamp:", req.PreMsgTime)
	var resp types.ChatMessageResponse
	//解析token
	claims, err := jwt.ParseToken(req.Token)
	if err != nil {
		resp.StatusCode = "-1"
		resp.StatusMsg = "解析token失败"
		resp.MessageList = nil
		return &resp, err
	}
	fromUserId := claims.UserId
	response, err := l.svcCtx.ChatRpc.ChatMessage(l.ctx, &chat.ChatMessageRequest{
		PreMsgTime: req.PreMsgTime,
		FromUserId: strconv.Itoa(fromUserId),
		ToUserId:   req.ToUserId,
	})
	if err != nil || response.Status == false {
		resp.StatusCode = "-1"
		resp.StatusMsg = "rpc连接失败"
		resp.MessageList = nil
		return &resp, err
	}
	var chatMessages []types.Message

	for _, record := range response.MessageList {
		chatMessage := convertChatRecordToMessage(record)
		chatMessages = append(chatMessages, chatMessage)
	}
	fmt.Println("chatMessages:", chatMessages)
	resp.StatusCode = "0"
	resp.StatusMsg = "获取聊天记录成功"
	resp.MessageList = chatMessages
	return &resp, nil
}
func convertChatRecordToMessage(message *chat.Message) types.Message {
	return types.Message{
		Id:         message.Id,
		ToUserId:   message.ToUserId,
		FromUserId: message.FromUserId,
		Content:    message.Content,
		CreateTime: message.CreateTime,
	}
}
