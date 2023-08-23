package logic

import (
	"MicroTikTok/chat/api/internal/svc"
	"MicroTikTok/chat/api/internal/types"
	"MicroTikTok/chat/rpc/pb/chat"
	"MicroTikTok/pkg/jwt"
	"context"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatActionLogic {
	return &ChatActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatActionLogic) ChatAction(req *types.ChatActionRequest) (*types.ChatActionResponse, error) {
	//解析当前用户的id
	claims, err := jwt.ParseToken(req.Token)
	currentUserId := claims.UserId
	request := chat.ChatActionRequest{
		UserId:   strconv.Itoa(currentUserId),
		ToUserId: strconv.FormatInt(req.ToUserId, 10),
		Content:  req.Content,
	}
	resp, err := l.svcCtx.ChatRpc.ChatAction(l.ctx, &request)
	if err != nil && resp.Status == false {
		return nil, err
	}
	return &types.ChatActionResponse{
		StatusCode: 0,
		StatusMsg:  "发送消息成功",
	}, nil
}
