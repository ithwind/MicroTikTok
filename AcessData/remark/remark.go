package remark

import (
	. "MicroTikTok/AcessData/mysql"
	"time"
)

type Remark struct {
	ID          int64     `gorm:"column:id"`           // Comment ID
	CommentText string    `gorm:"column:comment_text"` // Content
	UserID      int64     `gorm:"column:user_id"`      // User ID
	VideoID     int64     `gorm:"column:video_id"`     // Video ID
	CreateAt    time.Time `gorm:"column:create_at"`    // Create Time
	DelFlag     bool      `gorm:"column:del_flag"`     // 删除标识 1-未删除 0-删除
}

func GetRemarkListByVideoIdAndUserId(videoId int64, userId int64) []Remark {
	var remarks = make([]Remark, 0)
	err := DB.Where("video_id = ? && user_id = ? && del_flag != 0", videoId, userId).Find(&remarks).Error
	if err != nil {
		return nil
	}
	return remarks
}

func CreateRemark(userId, videoId int64, content string) (error, Remark) {
	var remark Remark
	remark.CreateAt = time.Now()
	remark.UserID = userId
	remark.VideoID = videoId
	remark.CommentText = content
	remark.DelFlag = true
	err := DB.Create(&remark).Error

	return err, remark
}

// DeleteRemark 软删除remark即将del_flag置为0
func DeleteRemark(commentId int64) error {
	err := DB.Table("remark").Where("id = ?", commentId).Update("del_flag", 0).Error

	return err
}
