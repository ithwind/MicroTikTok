package util

func String(str string) *string {
	return &str
}
func GetString(str *string) string {
	return *str
}
