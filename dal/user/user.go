package user

import (
	"MicroTikTok/constant"
	. "MicroTikTok/dal/postgres"
	"MicroTikTok/dal/video"
)

type User struct {
	ID              int64  `json:"id"`
	UserName        string `json:"user_name"`
	Password        string `json:"password"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
}

func (User) TableName() string {
	return constant.UserTableName
}

func GetUserIdByVideoId(videoId int64) int64 {
	var userId int64
	DB.Table("user_video").Where("video_id = ?", videoId).Select("user_id").Find(&userId)

	return userId
}

func GetUserById(userId int64) (*User, error) {
	var user User
	if err := DB.Where("id = ?", userId).Find(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func GetIsFavoriteByUserId(userId int64, videoId int64) bool {
	var videoIds []int64
	DB.Table("user_video_favorite").Where("user_id = ?", userId).Find(&videoIds)

	for _, id := range videoIds {
		if id == videoId {
			return true
		}
	}
	return false
}

func GetWorkCountByUserId(userId int64) int64 {
	var cnt int64
	DB.Table("user_video").Where("user_id = ?", userId).Select("video_id").Count(&cnt)

	return cnt
}

// GetFollowCountByUserId 获取当前用户的关注数
func GetFollowCountByUserId(userId int64) int64 {
	var count int64
	DB.Table("user_follow").Where("follower_id = ?", userId).Select("user_id").Count(&count)

	return count
}

// GetFollowerCountByUserId 获取当前用户的粉丝数
func GetFollowerCountByUserId(userId int64) int64 {
	var count int64
	DB.Table("user_follow").Where("user_id = ?", userId).Select("follower_id").Count(&count)

	return count
}

func GetFavoriteCount(userId int64) int64 {
	var count int64
	DB.Table("user_video_favorite").Where("user_id = ?", userId).Select("video_id").Count(&count)

	return count
}

// GetTotalFavoriteCount 获赞数量
func GetTotalFavoriteCount(userId int64) int64 {
	/**
	1.获取所有当前用户的视频
	2.统计所有视频的点赞数
	*/
	var videoIds []int64
	var totalCount int64 = 0
	DB.Table("user_video").Where("user_id = ?", userId).Select("video_id").Find(&videoIds)

	for _, videoId := range videoIds {
		totalCount += video.GetFavoriteCountByVideoId(videoId)
	}
	return totalCount
}

// AddUserVideoTable 增加user_video
func AddUserVideoTable(userId int64, videoId int64) error {

	result := DB.Table("user_video").Create(map[string]interface{}{
		"user_id":  userId,
		"video_id": videoId,
	})
	return result.Error
}
