package video

import (
	"MicroTikTok/constant"
	. "MicroTikTok/dal/postgres"
	. "MicroTikTok/pkg/util"
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
	videos := make([]*Video, constant.VideoFeedCount)
	lastTime = ConvertTimeFormat(lastTime.Format("2006-01-02 15:04:05"))
	err := DB.Where("publish_time <= ?", lastTime).Order("publish_time desc").Limit(constant.VideoFeedCount).Find(&videos).Error
	if err != nil {
		return videos, err
	}
	return videos, nil
}

func GetFavoriteCountByVideoId(videoId int64) int64 {
	var count int64
	DB.Table("user_video_favorite").Where("video_id = ?", videoId).Select("user_id").Count(&count)

	return count
}
