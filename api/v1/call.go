package v1

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/runbox/engine"
	"github.com/liu-cn/runbox/model"
	"github.com/liu-cn/runbox/model/request"
	"github.com/liu-cn/runbox/model/response"
	"github.com/liu-cn/runbox/pkg/jsonx"
	"github.com/liu-cn/runbox/pkg/stringsx"
	"github.com/liu-cn/runbox/runner"
	xerrors "github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

type RunBox struct {
	Engine *engine.RunBox
}

// GetSoftLogs softRunTime 程序执行耗时，callCostTime：调用+执行的耗时
func GetSoftLogs(callResponse *response.Call) {
	//todo 这里应该记录请求用户信息
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
		//	todo 这里可以对接消息队列存储日志数据
	}
}

func (r *RunBox) GetCall(c *gin.Context) {
	var (
		req request.Call
		err error
	)
	req.Method = "GET"               //请求方式
	req.Soft = c.Param("soft")       //软件名称
	req.User = c.Param("user")       //应用所属租户
	req.Command = c.Param("command") //命令

	if req.Soft == "" {
		response.FailWithHttpStatus(c, 400, "soft不能为空")
		return
	}
	if req.Command == "" {
		response.FailWithHttpStatus(c, 400, "Command不能为空")
		return
	}
	if req.User == "" {
		response.FailWithHttpStatus(c, 400, "user不能为空")
		return
	}

	runnerMetaData := c.Request.Header.Get("Runner-Meta-Data")
	if runnerMetaData == "" {
		response.FailWithHttpStatus(c, 403, "请携带元数据信息")
		return
	}

	var runnerMeta model.Runner
	err = json.Unmarshal([]byte(runnerMetaData), &runnerMeta)
	if err != nil {
		response.FailWithHttpStatus(c, 403, "请携带正确的元数据信息")
		return
	}
	err = runnerMeta.Check()
	if err != nil {
		response.FailWithHttpStatus(c, 403, err.Error())
		return
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
	//todo 从header获取软件元数据

	run := runner.NewRunner(&runnerMeta)
	req.RequestJsonPath = req.GetRequestFilePath(run.GetInstallPath())

	err = jsonx.SaveFile(req.RequestJsonPath, req) //todo 存储请求参数
	if err != nil {
		response.FailWithHttpStatus(c, 500, xerrors.Wrapf(err, "req.RequestJsonPath %s faild", req.RequestJsonPath).Error())
		return
	}
	//todo 请求参数需要删除
	//defer os.Remove(req.RequestJsonPath)
	getCall, err := r.Engine.GetCall(&req, &runnerMeta)

	//call, err := run.Call(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	GetSoftLogs(getCall) //记录用户日志
	if getCall.StatusCode != 200 {
		c.JSON(getCall.StatusCode, gin.H{
			"msg": getCall.Data,
		})
		return
	}

	if getCall.ContentType == "file" {
		if getCall.HasFile {
			//if call.HasFile
			fileName := filepath.Base(getCall.FilePath) // 获取文件名
			if getCall.DeleteFile {
				defer os.Remove(getCall.FilePath)
			}
			// 如果请求中有自定义文件名，则使用自定义文件名
			if customFileName := c.Query("filename"); customFileName != "" {
				fileName = customFileName
			}
			c.Writer.Header().Add("Content-Disposition", "attachment; filename="+fileName)
			c.File(getCall.FilePath)
			return
		}
	}

	if getCall.ContentType == "json" {
		c.Data(200, "application/json; charset=utf-8", []byte(jsonx.JSONString(getCall.Data)))
	}
	if getCall.ContentType == "text" {
		c.Data(200, "text/plain; charset=utf-8", []byte(fmt.Sprintf("%v", getCall.Data)))
	}

}
