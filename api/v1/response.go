package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/runbox/model/response"
	"github.com/liu-cn/runbox/pkg/jsonx"
	"net/http"
	"os"
	"path/filepath"
)

func ResponseJSON(c *gin.Context, call *response.Run) {
	c.Data(200, "application/json; charset=utf-8", []byte(jsonx.JSONString(call.Body)))
	return
}

func ResponseText(c *gin.Context, call *response.Run) {
	c.Data(200, "text/plain; charset=utf-8", []byte(fmt.Sprintf("%v", call.Body)))
	return
}

func ResponseImage(c *gin.Context, call *response.Run) {
	imagePath := call.FilePath

	// 打开图片文件
	file, err := os.Open(imagePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法打开图片文件 err:" + err.Error()})
		return
	}
	defer file.Close()

	// 获取图片文件的类型
	fileInfo, _ := file.Stat()
	//fileType := http.DetectContentType([]byte(fileInfo.Name()))

	// 设置响应头
	c.Header("Content-Type", call.GetContentType())
	c.Header("Content-Disposition", "inline; filename="+fileInfo.Name())

	// 将图片文件作为响应发送给客户端
	http.ServeContent(c.Writer, c.Request, fileInfo.Name(), fileInfo.ModTime(), file)
}

func ResponseOctetStream(c *gin.Context, call *response.Run) {
	if call.HasFile {
		//if call.HasFile
		fileName := filepath.Base(call.FilePath) // 获取文件名
		//todo DeleteFileTime 这里删除文件的逻辑有3种，-1：不删除，0：响应后立即删除，>0的时间戳：引擎会定时去删除这个文件
		if call.FilePath != "" && call.DeleteFileTime == 0 { //说明响应后需要立刻删除文件
			defer os.Remove(call.FilePath)
		}
		if call.FilePath != "" && call.DeleteFileTime > 0 { //说明响应后需要立刻删除文件
			//todo 这里需要把文件记录起来，定时扫描删除
		}
		// 如果请求中有自定义文件名，则使用自定义文件名
		if customFileName := c.Query("filename"); customFileName != "" {
			fileName = customFileName
		}
		c.Writer.Header().Add("Content-Disposition", "attachment; filename="+fileName)
		c.File(call.FilePath)
		return
	}
}

func ResponseVideo(c *gin.Context, call *response.Run) {
	videoPath := call.FilePath

	// 打开图片文件
	file, err := os.Open(videoPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法打开视频文件"})
		return
	}
	defer file.Close()

	// 获取图片文件的类型
	fileInfo, _ := file.Stat()
	//fileType := http.DetectContentType([]byte(fileInfo.Name()))

	// 设置响应头
	c.Header("Content-Type", call.GetContentType())
	c.Header("Content-Disposition", "inline; filename="+fileInfo.Name())

	// 将图片文件作为响应发送给客户端
	http.ServeContent(c.Writer, c.Request, fileInfo.Name(), fileInfo.ModTime(), file)
}
