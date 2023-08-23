package chat

import (
	. "MicroTikTok/AcessData/postgres"
	"MicroTikTok/Constant"
	"fmt"
	"sync"
	"time"
)

type ChatRecord struct {
	Id         int64     `gorm:"column:id"`
	FromUserId int64     `gorm:"column:from_user_id"`
	ToUserId   int64     `gorm:"column:to_user_id"`
	Content    string    `gorm:"column:content"`
	CreateTime time.Time `gorm:"column:created_at"`
}

var chatRecordCache = make(map[int64][]ChatRecord)
var chatRecordCacheMutex sync.Mutex

func (ChatRecord) TableName() string {
	return Constant.ChatTableName
}

// CacheChatRecord 将聊天记录缓存到内存中
func CacheChatRecord(fromUserId, toUserId int64, content string) {
	chatRecordCacheMutex.Lock()
	defer chatRecordCacheMutex.Unlock()

	chatRecordCache[fromUserId] = append(chatRecordCache[fromUserId], ChatRecord{
		FromUserId: fromUserId,
		ToUserId:   toUserId,
		Content:    content,
	})
}

// InsertCachedRecords 批量插入缓存的聊天记录到数据库
func InsertCachedRecords() error {
	chatRecordCacheMutex.Lock()
	defer chatRecordCacheMutex.Unlock()

	for userId, records := range chatRecordCache {
		for _, record := range records {
			err := CreateMessage(record.FromUserId, record.ToUserId, record.Content)
			if err != nil {
				return err
			}
		}
		// 清空已插入数据库的记录
		chatRecordCache[userId] = nil
	}

	return nil
}

func CreateMessage(fromUserId, toUserId int64, content string) error {
	fmt.Println("From:", fromUserId, "To:", toUserId, "Content:", content)
	chatMessage := ChatRecord{
		FromUserId: fromUserId,
		ToUserId:   toUserId,
		Content:    content,
		CreateTime: time.Now(),
	}
	err := DB.Select("from_user_id", "to_user_id", "content", "created_at").Create(&chatMessage).Error
	fmt.Println("CreateError:", err)
	return err
}

func QueryMessagesByFromUserIdAndToUserId(fromUserId int64, toUserId int64, preMsgTime time.Time) (*[]ChatRecord, error) {
	//var count int64
	var chatMessages *[]ChatRecord
	err := DB.Table(Constant.ChatTableName).Where("from_user_id = ? And to_user_id = ? And created_at > ?", fromUserId, toUserId, preMsgTime).Find(&chatMessages).Error

	return chatMessages, err
}
