package video

import (
	. "MicroTikTok/AcessData/postgres"
	"MicroTikTok/Constant"
	"MicroTikTok/favorite/model"
	. "MicroTikTok/pkg/util"
	"context"
	"github.com/zeromicro/go-zero/core/logc"
	"time"
)

type Video struct {
	ID          int64     `gorm:"column:id"`
	PlayURL     string    `gorm:"column:play_url"`
	CoverURL    string    `gorm:"column:cover_url"`
	PublishTime time.Time `gorm:"column:publish_time"`
	Title       string    `gorm:"column:title"`
}

func (Video) TableName() string {
	return Constant.VideoTableName
}

func GetVideosBeforeLastTime(lastTime time.Time) ([]*Video, error) {
	videos := make([]*Video, 0, Constant.VideoFeedCount)
	lastTime = ConvertTimeFormat(lastTime.Format("2006-01-02 15:04:05"))
	err := DB.Where("publish_time <= ?", lastTime).Order("publish_time desc").Limit(Constant.VideoFeedCount).Find(&videos).Error
	if err != nil {
		logc.Error(context.Background(), err)
		return videos, err
	}
	return videos, nil
}

func GetFavoriteCountByVideoId(videoId int64) int64 {
	var count int64
	DB.Table("user_video_favorite").Where("video_id = ?", videoId).Select("user_id").Count(&count)

	return count
}

func UpdateVideo(addVideo *Video) {
	err := DB.Create(addVideo).Error
	if err != nil {
		logc.Error(context.Background(), err)
	}
}

func QueryVideoIdsByUserId(userId int64) []int64 {
	videoIds := make([]int64, 0, 30)
	DB.Table("user_video").Where("user_id = ?", userId).Select("video_id").Find(&videoIds)

	return videoIds
}

func GetVideoById(videoId int64) *Video {
	var video *Video
	DB.Table("video").Where("id = ?", videoId).Find(&video)

	return video
}

func GetPublishList(userId int64) ([]*Video, error) {
	videoIds := QueryVideoIdsByUserId(userId)
	videos := make([]*Video, 0, Constant.VideoFeedCount)
	for _, videoId := range videoIds {
		video := GetVideoById(videoId)
		videos = append(videos, video)
	}

	return videos, nil
}

func NewSetFavorite(userId int64, videoId int64) error {
	var userVideoFavorite model.UserVideoFavorite
	userVideoFavorite.VideoId = videoId
	userVideoFavorite.UserId = userId
	err := DB.Table("user_video_favorite").Create(&userVideoFavorite).Error

	if err != nil {
		logc.Error(context.Background(), err)
		return err
	}
	return nil
}

func DeleteFavorite(userId int64, videoId int64) error {
	err := DB.Table("user_video_favorite").Where("user_id = ? And video_id = ?", userId, videoId).Delete(&model.UserVideoFavorite{}).Error
	return err
}

func GetCommentCount(videoId int64) (int64, error) {
	var commentCount int64
	err := DB.Table("comment").Where("video_id = ?", videoId).Count(&commentCount).Error

	return commentCount, err
}
