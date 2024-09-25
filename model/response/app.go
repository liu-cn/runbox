package response

import "time"

//type Response struct {
//	Code    int         `json:"code"`
//	Message string      `json:"message"`
//	Data    interface{} `json:"data"`
//}

type Run struct {
	//StatusCode  int         `json:"status_code"`
	//Msg         string      `json:"msg"`
	//ContentType string      `json:"content_type"`
	//HasFile     bool        `json:"has_file"`
	//FilePath    string      `json:"path"`
	//DeleteFile  bool        `json:"delete_file"`
	//Data        interface{} `json:"data"`
	//
	//CallCostTime     time.Duration `json:"-"`
	//ResponseMetaData string        `json:"-"`

	StatusCode int    `json:"status_code"`
	Msg        string `json:"msg"`
	//ContentType    string      `json:"content_type"`
	HasFile        bool        `json:"has_file"`
	FilePath       string      `json:"path"`
	DeleteFile     bool        `json:"delete_file"`
	DeleteFileTime int         `json:"delete_file_time"` //-1 不删除文件，0响应成功后立刻删除文件，>0是时间戳给出具体时间戳，达到该时间戳时刻系统会自动清理该文件
	Body           interface{} `json:"data"`

	Header map[string]string `json:"header"` // response header

	//meta data
	CallCostTime     time.Duration `json:"-"` //
	ResponseMetaData string        `json:"-"`
}

func (r *Run) GetContentType() string {
	if r.Header != nil {
		return r.Header["Content-Type"]
	}
	return ""
}
