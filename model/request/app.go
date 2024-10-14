package request

import (
	"encoding/json"
	"fmt"
	"github.com/liu-cn/runbox/model"
	"time"
)

type MetaData struct {
	WorkPath   string `json:"work_path"`
	RunnerType string `json:"runner_type"`
	Version    string `json:"version"`
	Command    string `json:"command"` //命令
	User       string `json:"user"`    //软件所属的用户
	Soft       string `json:"soft"`    //软件名
	OssPath    string `json:"oss_path"`
}

type RollbackVersion struct {
	RunnerConf *model.Runner `json:"runner_conf"`
	OldVersion string        `json:"old_version"`
}

type Run struct {
	MetaData MetaData `json:"meta_data"`

	User    string   `json:"user"`    //软件所属的用户
	Soft    string   `json:"soft"`    //软件名
	Type    string   `json:"type"`    //软件类型
	Command string   `json:"command"` //命令
	Method  string   `json:"method"`  //请求方式
	Files   []string `json:"files"`

	Body            map[string]interface{} `json:"body"`              //请求json
	RequestJsonPath string                 `json:"request_json_path"` //请求参数存储路径

	//UpdateVersion   bool                   `json:"update_version"`    //此时正处于版本更新的状态

	//ReqBody         string
}

func (c *Run) IsOpenCommand() bool {
	return c.Command == "_cloud_func" || c.Command == "_docs_info_text"
}

func (c *Run) RequestJSON() (string, error) {
	j, err := json.Marshal(c.Body)
	if err != nil {
		return "", err
	}
	return string(j), nil
}

func (c *Run) GetRequestFilePath(callerPath string) string {
	reqJson := callerPath + fmt.Sprintf("/.request/%v_%v.json",
		c.Soft, time.Now().UnixNano())
	return reqJson
}
