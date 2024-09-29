package response

import "time"

type UnInstallInfo struct {
}

type InstallInfo struct {
	TempPath     string `json:"temp_path"`     //软件安装时候临时目录，下载到该目录，然后copy到所属目录
	RootPath     string `json:"root_path"`     //存储根路径
	StoreRoot    string `json:"store_root"`    //云存储根路径
	Pc           string `json:"pc"`            //软件平台
	Name         string `json:"name"`          //软件名称
	FullName     string `json:"full_name"`     //软件名称,带后缀
	User         string `json:"user"`          //所属用户
	DownloadPath string `json:"download_path"` //软件的云端地址
	//InstallPath  string //安装后的所属目录
	Version string `json:"version"` //安装的软件版本

	Other map[string]interface{} `json:"other"`
}

type UpdateVersion struct {
}

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
