package v1

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/runbox/model/request"
	"github.com/liu-cn/runbox/model/response"
	"github.com/liu-cn/runbox/pkg/jsonx"
	"github.com/liu-cn/runbox/pkg/stringsx"
	"github.com/liu-cn/runbox/runner"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

// GetSoftLogs softRunTime 程序执行耗时，callCostTime：调用+执行的耗时
func GetSoftLogs(callResponse *response.Call) {
	softRunTime := ""
	softRunTimeList := stringsx.ParserHtmlTagContent(callResponse.ResponseMetaData, "UserCost")
	if len(softRunTimeList) > 0 {
		softRunTime = softRunTimeList[0]
	}
	logs := stringsx.ParserHtmlTagContent(callResponse.ResponseMetaData, "Logger")
	for _, v := range logs {
		mp := make(map[string]interface{})
		err := json.Unmarshal([]byte(v), &mp)
		if err != nil {
			continue
		}
		msg := mp["msg"]
		delete(mp, "msg")
		mp["soft_run_time"] = softRunTime
		mp["call_cost_time"] = callResponse.CallCostTime.String()
		logrus.WithFields(mp).Info(msg)
		//	这里可以对接消息队列存储日志数据
	}
}

func GetCall(c *gin.Context) {
	var (
		req request.Call
		err error
	)
	req.Soft = c.Param("soft")
	req.Command = c.Param("command")
	req.Method = "GET"
	req.User = c.Param("user")

	if req.Soft == "" {
		response.FailWithMessage("soft不能为空", c)
		return
	}
	if req.Command == "" {
		response.FailWithMessage("Command不能为空", c)
		return
	}

	if req.User == "" {
		response.FailWithMessage("请登录后访问", c)

	}
	if req.Data == nil {
		req.Data = make(map[string]interface{})
	}
	queryMap := c.Request.URL.Query()
	for k, values := range queryMap {
		if len(values) > 1 {
			req.Data[k] = values
		} else {
			req.Data[k] = values[0]
		}
	}
	//todo 验证用户是否有调用该应用的权限
	j, err := req.RequestJSON()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	req.ReqBody = j
	//todo
	run := runner.NewRunner(nil)
	req.RequestJsonPath = req.GetRequestFilePath(run.GetInstallPath())

	err = jsonx.SaveFile(req.RequestJsonPath, req) //
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//todo 请求参数需要删除
	//defer os.Remove(req.RequestJsonPath)

	call, err := run.Call(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	GetSoftLogs(call)
	if call.StatusCode != 200 {
		c.JSON(call.StatusCode, gin.H{
			"msg": call.Data,
		})
		return
	}

	if call.ContentType == "file" {
		if call.HasFile {
			//if call.HasFile
			fileName := filepath.Base(call.FilePath) // 获取文件名
			if call.DeleteFile {
				defer os.Remove(call.FilePath)
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

	if call.ContentType == "json" {
		c.Data(200, "application/json; charset=utf-8", []byte(jsonx.JSONString(call.Data)))
	}
	if call.ContentType == "text" {
		c.Data(200, "text/plain; charset=utf-8", []byte(fmt.Sprintf("%v", call.Data)))
	}

}
