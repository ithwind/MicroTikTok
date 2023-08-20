package logic

import (
	_ "MicroTikTok/user/model/dao/model"
	"MicroTikTok/user/rpc/internal/svc"
	"MicroTikTok/user/rpc/rpc/user"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	table := l.svcCtx.BkModel.User
	res, err := table.WithContext(l.ctx).Where(table.ID.Eq(int32(in.UserId))).Debug().First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Error(100, "用户不存在")
		}
		return nil, status.Error(500, err.Error())
	}
	//作品
	UserVideotable := l.svcCtx.BkModel.UserVideo
	WorkCount, err := UserVideotable.WithContext(l.ctx).Count()
	if WorkCount != 0 {
		video, _ := UserVideotable.WithContext(l.ctx).Where(UserVideotable.UserID.Eq(int32(in.UserId))).Count()
		WorkCount = WorkCount + video
	}
	//点赞
	favorite := l.svcCtx.BkModel.UserVideoFavorite
	Favorite, err := favorite.WithContext(l.ctx).Count()
	if Favorite != 0 {
		FavoriteCount, _ := favorite.WithContext(l.ctx).Where(favorite.LikedID.Eq(int32(in.UserId))).Count()
		Favorite += FavoriteCount
	}
	Favoriter, err := favorite.WithContext(l.ctx).Count()
	if Favoriter != 0 {
		FavoriterCount, _ := favorite.WithContext(l.ctx).Where(favorite.LikerID.Eq(int32(in.UserId))).Count()
		Favoriter += FavoriterCount
	}
	//关注
	followtable := l.svcCtx.BkModel.UserFollow
	follow, err := followtable.WithContext(l.ctx).Count()
	if follow != 0 {
		followCount, _ := followtable.WithContext(l.ctx).Where(followtable.FollowedID.Eq(int32(in.UserId))).Count() //粉丝总数
		follow += followCount
	}
	followr, err := followtable.WithContext(l.ctx).Count()
	if followr != 0 {
		followerCount, _ := followtable.WithContext(l.ctx).Where(followtable.FollowerID.Eq(int32(in.UserId))).Count() //关注总数
		followr += followerCount
	}

	u := user.User{
		Id:              int64(res.ID),
		Name:            res.Name,
		Avatar:          &res.Avatar,
		WorkCount:       &WorkCount,
		FavoriteCount:   &Favorite,
		FollowCount:     &followr,
		FollowerCount:   &follow,
		BackgroundImage: &res.BackgroundImage,
		Signature:       &res.Signature,
		TotalFavorited:  &Favoriter,
	}

	statusMsg := "返回成功"
	return &user.UserInfoResponse{
		StatusCode: 200,
		StatusMsg:  &statusMsg,
		UserInfo:   &u,
	}, nil
}
