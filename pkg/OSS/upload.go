package OSS

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"os"
	"time"
)

// Upload /**
func Upload(fileType string, filePath string) (bool, error, string) {
	// 密钥存储
	accessKey := "0DrrfXE_AykDNMb6mPZWB2VA1n5S7Ik5u4krtdLJ"
	secretKey := "E8vKp6fpPgrYx7ub7F2dIQeGje-9ylt-Y-YBNAWt"

	//获取当前文件路径
	localFile := filePath
	bucket := "test-tik-tok"

	//获取当前时间
	NowTime := time.Now().Format("200601021504")
	key := "test-TikTok/" + fileType + "/" + NowTime

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	resumeUploader := storage.NewResumeUploaderV2(&cfg)
	ret := storage.PutRet{}
	recorder, err := storage.NewFileRecorder(os.TempDir())
	// 可选配置
	putExtra := storage.RputV2Extra{
		Recorder: recorder,
	}
	//上传传来的文件

	err = resumeUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		return false, nil, NowTime
	}
	return true, nil, NowTime
}
