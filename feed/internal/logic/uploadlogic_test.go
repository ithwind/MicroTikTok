package logic

import (
	"github.com/qiniu/go-sdk/v7/sms/bytes"
	"os"
	"testing"
)

func TestNewUploadLogic(t *testing.T) {
	file, _ := os.Open("D:\\goWorkspace\\MicroTikTok\\1.mp4")

	bytes.NewReader(file)
}
