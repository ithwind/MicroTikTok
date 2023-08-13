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

		// 处理上传的文件 将上传文件先保留至本地
		savePath := "feed/uploads/video/" + handler.Filename
		outFile, err := os.Create(savePath)
		defer func(outFile *os.File) {
			err := outFile.Close()
			if err != nil {

			}
		}(outFile)
		_, err = io.Copy(outFile, file)

		if err != nil {
			fmt.Println("Error saving the file:", err)
			return
		}
		fmt.Println("File uploaded and saved to:", savePath)

		// 将上传的文件内容直接写入req.Data字段
		fileData, err := os.ReadFile(savePath)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		req.Data = fileData

		// 手动解析其他表单参数
		req.Title = r.FormValue("title")
		req.Token = r.FormValue("token")
		/*if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			fmt.Printf("Error: %v   Data: %T \n", err, req.Data)
			return
		}*/

		l := logic.NewUploadLogic(r.Context(), svcCtx)
		resp, err := l.Upload(&req, savePath)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
