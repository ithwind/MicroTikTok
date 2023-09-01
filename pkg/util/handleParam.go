package util

func String(str string) *string {
	return &str
}
func GetString(str *string) string {
	return *str
}

func Int(num int64) *int64 {
	return &num
}

func GetInt(num *int64) int64 {
	return *num
}
