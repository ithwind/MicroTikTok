package video

import (
	"MicroTikTok/constant"
	. "MicroTikTok/dal/postgres"
	"MicroTikTok/favorite/model"
	. "MicroTikTok/pkg/util"
	"fmt"
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
	return constant.VideoTableName
}

func GetVideosBeforeLastTime(lastTime time.Time) ([]*Video, error) {
	videos := make([]*Video, 0, constant.VideoFeedCount)
	lastTime = ConvertTimeFormat(lastTime.Format("2006-01-02 15:04:05"))
	err := DB.Where("publish_time <= ?", lastTime).Order("publish_time desc").Limit(constant.VideoFeedCount).Find(&videos).Error
	if err != nil {
		fmt.Printf("Error: %v", err)
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
		fmt.Printf("Error: %v", err)
	}
}

func QueryVideoIdsByUserId(userId int64) []int64 {
	videoIds := make([]int64, 0, 30)
	DB.Table("user_video_favorite").Where("user_id = ?", userId).Select("video_id").Find(&videoIds)

	return videoIds
}

func GetVideoById(videoId int64) *Video {
	var video *Video
	DB.Table("video").Where("id = ?", videoId).Find(&video)

	return video
}

func GetPublishList(userId int64) ([]*Video, error) {
	videoIds := QueryVideoIdsByUserId(userId)
	fmt.Printf("UserIds : %v\n", videoIds)
	videos := make([]*Video, 0, constant.VideoFeedCount)
	for _, videoId := range videoIds {
		video := GetVideoById(videoId)
		fmt.Printf("videoId : %v  video: %v \n", videoId, video)
		videos = append(videos, video)
	}

	return videos, nil
}

func NewSetFavorite(userId int64, videoId int64) error {
	var userVideoFavorite model.UserVideoFavorite
	userVideoFavorite.VideoId = videoId
	userVideoFavorite.UserId = userId
	fmt.Printf("InsertParam: %v %v\n", userId, videoId)
	err := DB.Table("user_video_favorite").Create(&userVideoFavorite).Error

	if err != nil {
		fmt.Println("AddError:", err)
		return err
	}
	return nil
}

func DeleteFavorite(userId int64, videoId int64) error {
	err := DB.Table("user_video_favorite").Where("user_id = ? And video_id = ?", userId, videoId).Delete(&model.UserVideoFavorite{}).Error
	fmt.Printf("DeleteParam: %v, %v", userId, videoId)
	return err
}
