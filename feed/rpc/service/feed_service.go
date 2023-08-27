package service

import (
	"MicroTikTok/AcessData/modelVo"
	"MicroTikTok/AcessData/user"
	"MicroTikTok/AcessData/video"
	"MicroTikTok/Constant"
	. "MicroTikTok/feed/rpc/pb/video"
	"MicroTikTok/pkg/jwt"
	"MicroTikTok/pkg/util"
	"fmt"
	"sync"
	"time"
)

var currentUser model.UserVo

type FeedService struct {
}

func NewFeedSerVice() *FeedService {
	return &FeedService{}
}

// Feed 获取截止到last_time的不多于30条视频
func (feedService FeedService) Feed(request *FeedRequest) (*FeedResponse, error) {
	/**
	1.若没有last_time则为当前时间
	*/
	resp := &FeedResponse{}
	var lastTime time.Time
	if request.GetLatestTime() == Constant.DefaultLatestTime {
		lastTime = time.Now()
	} else {
		lastTime = time.Unix(*request.LatestTime, 0)
	}
	fmt.Printf("New LastTime: %v\n", lastTime)
	//获取token

	if request.GetToken() != "Default" {
		claim, err := jwt.ParseToken(request.GetToken())
		if err != nil {
			return &FeedResponse{
				StatusCode: Constant.StatusHttpFail,
				StatusMsg:  util.String("token非法"),
				VideoList:  nil,
				NextTime:   nil,
			}, nil
		}
		currentUser = claim.UserVo
	}

	dbVideos, err := video.GetVideosBeforeLastTime(lastTime)
	if err != nil {
		return resp, err
	}
	videos := make([]*Video, 0, Constant.VideoFeedCount)
	feedService.CopyVideos(&dbVideos, &videos)
	var publishTime = time.Now().Unix()
	resp.StatusCode = Constant.StatusHttpOk
	resp.VideoList = videos
	resp.StatusMsg = util.String("获取成功")
	resp.NextTime = &publishTime
	return resp, nil
}

func (feedService FeedService) CopyVideos(rawVideos *[]*video.Video, returnVideos *[]*Video) {
	for _, item := range *rawVideos {
		generateVideo := feedService.GenerateVideo(item)
		*returnVideos = append(*returnVideos, generateVideo)
	}
}

// GenerateVideo 通过数据库中的video生成返回的视频形式
func (feedService FeedService) GenerateVideo(data *video.Video) *Video {
	v := &Video{
		Id:       data.ID,
		PlayUrl:  data.PlayURL,
		CoverUrl: data.CoverURL,
		Title:    data.Title,
	}

	var wg sync.WaitGroup //并发获取参数
	wg.Add(4)
	//通过视频Id获取作者的参数 获取粉丝总数，关注数，评论，喜欢数，作品数，点赞总数
	var userId = user.GetUserIdByVideoId(v.GetId())

	var followCount = user.GetFollowCountByUserId(userId)
	var FollowerCount = user.GetFollowerCountByUserId(userId)
	var TotalFavorited = user.GetTotalFavoriteCount(userId)
	var WorkCount = user.GetWorkCountByUserId(userId)
	var FavoriteCount = user.GetFavoriteCount(userId)

	go func() {
		userInfo := user.GetUserById(userId)
		v.Author = &User{
			Id:              userInfo.ID,
			Name:            userInfo.UserName,
			FollowCount:     &followCount,
			FollowerCount:   &FollowerCount,
			IsFollow:        user.GetIsFavoriteByUserId(userId, v.GetId()),
			Avatar:          &userInfo.Avatar,
			BackgroundImage: &userInfo.BackgroundImage,
			Signature:       &userInfo.Signature,
			TotalFavorited:  &TotalFavorited,
			WorkCount:       &WorkCount,
			FavoriteCount:   &FavoriteCount,
		}
		wg.Done()
	}()

	//获取点赞数
	go func() {
		v.FavoriteCount = video.GetFavoriteCountByVideoId(v.GetId())
		wg.Done()
	}()
	go func() {
		v.CommentCount, _ = video.GetCommentCount(v.GetId())
		wg.Done()
	}()

	//判断视频是否点赞
	go func() {
		v.IsFavorite = user.GetIsFavoriteByUserId(currentUser.ID, v.Id)
		wg.Done()
	}()

	wg.Wait()

	return v
}
