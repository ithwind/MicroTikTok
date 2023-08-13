package service

import (
	"MicroTikTok/constant"
	"MicroTikTok/dal/user"
	"MicroTikTok/dal/video"
	. "MicroTikTok/feed/api/internal/types"
	"MicroTikTok/model"
	"MicroTikTok/pkg/jwt"
	"fmt"
	"sync"
	"time"
)

var currentUser model.User

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
	if *request.LastTime == 0 {
		lastTime = time.Now()
	} else {
		lastTime = time.Unix(*request.LastTime, 0)
	}
	if *request.Token != "Default" {
		claim, err := jwt.ParseToken(*request.Token)
		if err != nil {
			return &FeedResponse{
				StatusCode: 400,
				StatusMsg:  "token非法",
				VideoList:  nil,
				NextTime:   0,
			}, nil
		}
		currentUser = claim.User
	}

	dbVideos, err := video.GetVideosBeforeLastTime(lastTime)
	if err != nil || len(dbVideos) == 0 {
		fmt.Println(err)
		return resp, err
	}

	videos := make([]Video, 0, constant.VideoFeedCount)
	var publishTime = dbVideos[len(dbVideos)-1].PublishTime
	feedService.CopyVideos(&dbVideos, &videos)
	resp.StatusCode = 200
	resp.VideoList = videos
	resp.StatusMsg = "获取成功"
	resp.NextTime = publishTime.Unix()
	return resp, nil
}

func (feedService FeedService) GetPublishList(userId int64) []Video {
	videos := make([]Video, 0, constant.VideoFeedCount)
	dbVideos, _ := video.GetPublishList(userId)
	feedService.CopyVideos(&dbVideos, &videos)

	return videos
}

func (feedService FeedService) CopyVideos(rawVideos *[]*video.Video, returnVideos *[]Video) {
	for _, item := range *rawVideos {
		generateVideo := feedService.GenerateVideo(item)
		*returnVideos = append(*returnVideos, generateVideo)
	}
}

// GenerateVideo 通过数据库中的video生成返回的视频形式
func (feedService FeedService) GenerateVideo(data *video.Video) Video {
	v := &Video{
		Id:       data.ID,
		PlayUrl:  data.PlayURL,
		CoverUrl: data.CoverURL,
		Title:    data.Title,
	}

	var wg sync.WaitGroup //并发获取参数
	wg.Add(4)
	//通过视频Id获取作者的参数 获取粉丝总数，关注数，评论，喜欢数，作品数，点赞总数
	var userId = user.GetUserIdByVideoId(v.Id)

	var followCount = user.GetFollowCountByUserId(userId)
	var FollowerCount = user.GetFollowerCountByUserId(userId)
	var TotalFavorited = user.GetTotalFavoriteCount(userId)
	var WorkCount = user.GetWorkCountByUserId(userId)
	var FavoriteCount = user.GetFavoriteCount(userId)

	go func() {
		userInfo, _ := user.GetUserById(userId)
		v.Author = Author{
			Id:              userInfo.ID,
			Name:            userInfo.UserName,
			FollowCount:     followCount,
			FollowerCount:   FollowerCount,
			IsFollow:        user.GetIsFavoriteByUserId(userId, v.Id),
			Avatar:          userInfo.Avatar,
			BackgroundImage: userInfo.BackgroundImage,
			Signature:       userInfo.Signature,
			TotalFavorited:  TotalFavorited,
			WorkCount:       WorkCount,
			FavoriteCount:   FavoriteCount,
		}
		wg.Done()
	}()

	//获取点赞数
	go func() {
		v.FavoriteCount = video.GetFavoriteCountByVideoId(v.Id)
		wg.Done()
	}()

	//todo 获取评论数
	go func() {
		v.CommentCount = 0
		wg.Done()
	}()

	//判断视频是否点赞
	go func() {
		v.IsFavorite = user.GetIsFavoriteByUserId(currentUser.ID, v.Id)
		wg.Done()
	}()

	wg.Wait()

	return *v
}
