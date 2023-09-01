package ffmpeg

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"log"
	"os"
	"strings"
)

// GenerateCover  生成封面 返回生成封面的地址并且上传封面至服务器
func GenerateCover(videoPath, snapshotPath string, frameNum int) (coverName string, err error) {
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{
			"vframes":     1,
			"format":      "image2",
			"vcodec":      "mjpeg",
			"pix_fmt":     "yuvj420p", // 设置正确的像素格式
			"color_range": "pc",       // 设置正确的像素范围
		}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}

	err = imaging.Save(img, snapshotPath+".png")
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}
	names := strings.Split(snapshotPath, "\\")
	coverName = names[len(names)-1] + ".png"
	return coverName, nil
}
