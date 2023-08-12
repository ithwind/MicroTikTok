package util

import (
	"strconv"
	"time"
)

// ConvertTimeFormat ConvertString 将字符串转换为时间
// ConvertTimeFormat 时间格式转换
func ConvertTimeFormat(timeString string) time.Time {
	// 找到第一个 "-" 字符的位置
	separatorIndex := 0
	for i, ch := range timeString {
		if ch == '-' {
			separatorIndex = i
			break
		}
	}
	year, _ := strconv.Atoi(timeString[:separatorIndex])
	month, _ := strconv.Atoi(timeString[separatorIndex+1 : 7])
	day, _ := strconv.Atoi(timeString[8:10])
	hour, _ := strconv.Atoi(timeString[11:13])
	minute, _ := strconv.Atoi(timeString[14:16])
	second, _ := strconv.Atoi(timeString[17:19])
	latestTime := time.Date(year, time.Month(month), day, hour, minute, second, 0, time.Local)
	return latestTime
}
