package runner

import (
	"fmt"
	"github.com/liu-cn/runbox/pkg/jsonx"
	"github.com/liu-cn/runbox/sdk/runner/content_type"
	"github.com/liu-cn/runbox/sdk/runner/response"
	"mime"
	"path/filepath"
	"strings"
)

func (c *Context) ResponseFailJSONWithCode(code int, jsonEl interface{}) {
	rsp := &response.CallResponse{
		StatusCode: code,
		Header: map[string]string{
			"Content-Type": content_type.ApplicationJsonCharsetUtf8,
		},
		//ContentType: content_type.ApplicationJsonCharsetUtf8,
		Body: jsonEl}
	c.response(jsonx.String(rsp))
}

func (c *Context) ResponseFailDefaultJSONWithMsg(errMsg string) {
	rsp := &response.CallResponse{
		StatusCode: 200,
		Header: map[string]string{
			"Content-Type": content_type.ApplicationJsonCharsetUtf8,
		},
		//ContentType: content_type.ApplicationJsonCharsetUtf8,
		Body: map[string]interface{}{"msg": errMsg}}
	c.response(jsonx.String(rsp))
}

func (c *Context) ResponseFailTextWithCode(code int, text string) {
	rsp := &response.CallResponse{
		StatusCode: code,
		Header: map[string]string{
			"Content-Type": content_type.TextPlainCharsetUtf8,
		},
		//ContentType: content_type.TextPlainCharsetUtf8,
		Body: text}
	c.response(jsonx.String(rsp))
}
func (c *Context) ResponseOkWithJSON(jsonEl interface{}) {
	rsp := &response.CallResponse{
		StatusCode: 200,
		Header: map[string]string{
			"Content-Type": content_type.ApplicationJsonCharsetUtf8,
		},
		//ContentType: content_type.ApplicationJsonCharsetUtf8,
		Body: jsonEl}
	c.response(jsonx.String(rsp))
}
func (c *Context) ResponseFailParameter() {
	rsp := &response.CallResponse{
		StatusCode: 200,
		Header: map[string]string{
			"Content-Type": content_type.ApplicationJsonCharsetUtf8,
		},
		//ContentType: content_type.ApplicationJsonCharsetUtf8,
		Body: map[string]interface{}{"msg": "参数错误", "code": 10001}}
	c.response(jsonx.String(rsp))
}

type Response struct {
	HttpStatusCode int    // http 响应码
	FilePath       string //如果响应的是是个文件，请给出文件地址
	DeleteFileTime int    //-1 不删除文件，0响应成功后立刻删除文件，>0是时间戳给出具体时间戳，达到该时间戳时刻系统会自动清理该文件

	Header map[string]string
	Body   interface{} //http 响应body
}

// 相对路径转换成绝对路径
func (c *Context) getAbsPath(path string) string {
	fmt.Println("c.WorkPath", c.WorkPath)
	abs := strings.ReplaceAll(path, "./", c.WorkPath+"/")
	return strings.ReplaceAll(abs, "\\", "/")
}

func (c *Context) Response(res Response) error {
	if res.Header == nil {
		res.Header = make(map[string]string) //默认响应json格式
		if res.FilePath != "" {              //如果存在文件返回二进制类型
			path := c.getAbsPath(res.FilePath)
			fmt.Println("abs: ", path)

			res.FilePath = path
			ext := filepath.Ext(res.FilePath)
			mimeType := mime.TypeByExtension(ext)
			fmt.Printf("mimeType: %s\n", mimeType)
			//file, err := os.Open(res.FilePath)
			//if err != nil {
			//	return err
			//}
			//defer file.Close()
			//fileInfo, _ := file.Stat()
			//fileType := http.DetectContentType([]byte(mimeType))
			fmt.Printf("file name :%v fileType:%v \n", res.FilePath, mimeType)
			//默认返回json格式
			res.Header["Content-Type"] = mimeType

		} else {
			res.Header["Content-Type"] = content_type.ApplicationJsonCharsetUtf8
		}
	}
	rsp := &response.CallResponse{
		StatusCode:     res.HttpStatusCode,
		Header:         res.Header,
		Body:           res.Body,
		FilePath:       res.FilePath,
		DeleteFileTime: res.DeleteFileTime,
	}
	c.response(jsonx.String(rsp))
	return nil
}

// ResponseOkWithFile 返回文件
func (c *Context) ResponseOkWithFile(filePath string, deleteFileTime int) error {
	abs, err := filepath.Abs(filePath)
	if err != nil {
		return err
	}

	rsp := &response.CallResponse{
		StatusCode: 200,
		Header: map[string]string{
			"Content-Type": content_type.ApplicationOctetStream,
		},
		DeleteFileTime: deleteFileTime,
		//ContentType: content_type.ApplicationOctetStream,
		FilePath: abs}
	c.response(jsonx.String(rsp))
	return nil
}

func (c *Context) ResponseOkWithText(text string) error {
	rsp := &response.CallResponse{
		StatusCode: 200,

		Header: map[string]string{
			"Content-Type": content_type.TextPlainCharsetUtf8,
		},
		//ContentType: content_type.TextPlainCharsetUtf8,
		Body: text}
	c.response(jsonx.String(rsp))
	return nil
}

func (c *Context) response(text string) {
	fmt.Println("<Response>" + text + "</Response>")
}
