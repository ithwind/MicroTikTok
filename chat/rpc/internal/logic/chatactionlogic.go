package logic

import (
	chat2 "MicroTikTok/AcessData/chat"
	"MicroTikTok/chat/rpc/internal/svc"
	"MicroTikTok/chat/rpc/pb/chat"
	"MicroTikTok/pkg/AIChat"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
)

type ChatActionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChatActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatActionLogic {
	return &ChatActionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChatActionLogic) ChatAction(in *chat.ChatActionRequest) (*chat.ChatActionResponse, error) {
	// todo: add your logic here and delete this line
	var status = true
	fmt.Println("UserId:", in.UserId, "ToUserId:", in.ToUserId, "Content:", in.Content)
	fromUserId, _ := strconv.Atoi(in.UserId)
	toUserId, _ := strconv.Atoi(in.ToUserId)
	content := in.Content
	//判断聊天用户
	//AI对象
	if in.ToUserId == "0" {
		//获取AI的返回内容
		responseContent := AIChat.Chat(in.Content)
		/**
		1.先将用户的聊天信息的存入数据数据库
		2.将AI响应的信息存入数据库 发送者和收发者进行对调
		*/
		// 将聊天记录缓存到内存中
		chat2.CacheChatRecord(int64(fromUserId), int64(toUserId), content)
		chat2.CacheChatRecord(int64(toUserId), int64(fromUserId), responseContent)

		// 在协程中批量插入数据库
		go func() {
			err := chat2.InsertCachedRecords()
			if err != nil {
				status = false
				fmt.Println("Error", err)
			}
		}()
	} else {
		//将用户的发送聊天信息存入数据库
		err := chat2.CreateMessage(int64(fromUserId), int64(toUserId), content)
		if err != nil {
			status = false
			fmt.Println("Error:", err)
		}
	}
	return &chat.ChatActionResponse{Status: status}, nil
}
