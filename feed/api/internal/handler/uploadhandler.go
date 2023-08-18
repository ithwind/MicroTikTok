package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"MicroTikTok/feed/api/internal/logic"
	"MicroTikTok/feed/api/internal/svc"
	"MicroTikTok/feed/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func uploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PublishActionRequest
		file, handler, err := r.FormFile("data")
		fmt.Println("Error1:", err)
		// 处理上传的文件 将上传文件先保留至本地
		savePath := "uploads/video/" + handler.Filename
		outFile, err := os.Create(savePath)
		fmt.Println("Error2:", err)
		defer func(outFile *os.File) {
			err := outFile.Close()
			if err != nil {
				return
			}
		}(outFile)
		_, err = io.Copy(outFile, file)
		fmt.Println("Error3:", err)
		if err != nil {
			fmt.Println("Error saving the file:", err)
			return
		}
		fmt.Println("File uploaded and saved to:", savePath)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		fileData, err := os.ReadFile(savePath)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		req.Data = fileData

		// 手动解析其他表单参数
		req.Title = r.FormValue("title")
		req.Token = r.FormValue("token")

		l := logic.NewUploadLogic(r.Context(), svcCtx)
		fmt.Printf("Data: %v, Title: %v, Token: %v, Path:%v\n", len(req.Data), req.Title, req.Token, savePath)
		resp, err := l.Upload(&req, savePath)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
