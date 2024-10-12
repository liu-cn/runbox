package runner

import (
	"fmt"
	"github.com/liu-cn/runbox/pkg/jsonx"
	//"github.com/liu-cn/runbox/runner"
	"github.com/liu-cn/runbox/sdk/runner/request"
	"net/http"
	"strings"
	"time"
)

type Param struct {
	Name  string
	Field string
	Type  string
}

type CloudFunc struct {
	Name     string
	Path     string
	Method   string
	Request  []*Param
	Response []*Param
}

func bind(ctx *Context) error {
	var ca request.Call
	err := jsonx.UnmarshalFromFile(ctx.Request, &ca)
	ctx.WorkPath = ca.MetaData.WorkPath
	ctx.DebugPrintf("bind err:%v", err)
	if err != nil {
		return err
	}
	ctx.Req = &ca
	return nil
}
func Handel(context *Context, runner *Runner) {
	err := bind(context)
	if err != nil {
		context.ResponseFailJSONWithCode(http.StatusBadRequest, map[string]interface{}{
			"msg": "参数解析失败: " + err.Error(),
		})
		runner.About = true
		return
	}

	method := strings.ToUpper(context.Req.Method)
	//fmt.Println(command)
	//todo
	if context.Cmd == "_docs_info_text" && method == "GET" { //获取接口文档
		var s []string
		for _, worker := range runner.CmdMapHandel {
			if worker.Config == nil {
				continue
			}
			if !worker.Config.IsPublicApi {
				continue
			}
			s = append(s, fmt.Sprintf("%s\t %s \t %s", worker.Path, worker.Method, worker.Config.ApiDesc))
		}
		//res := append([]string{"请求地址 \t 请求方式 \t 接口描述"}, s...)
		context.ResponseOkWithText(strings.Join(s, "\n"))
		runner.About = true
		return
	}

	if context.Cmd == "_cloud_func" && method == "GET" {
		var s []string
		for _, worker := range runner.CmdMapHandel {
			if worker.Config == nil {
				continue
			}
			if !worker.Config.IsPublicApi {
				continue
			}
			s = append(s, fmt.Sprintf("%s\t %s \t %s", worker.Path, worker.Method, worker.Config.ApiDesc))
		}
		//res := append([]string{"请求地址 \t 请求方式 \t 接口描述"}, s...)
		context.ResponseOkWithText(strings.Join(s, "\n"))
		runner.About = true
		return
	}

	worker, ok := runner.CmdMapHandel[context.Cmd+"."+method]
	if ok {
		if context.Req == nil {
			panic("context.Req == nil")
		}
		handelList := worker.Handel
		//if context.Req.Method == "GET" {
		//	handelList = worker.Handel
		//}
		if len(handelList) == 0 {
			context.ResponseFailTextWithCode(http.StatusBadRequest, "bad request: method not handel")
			return
		}
		for _, fn := range handelList {
			now := time.Now()
			fn(context)
			t := time.Since(now)
			fmt.Println(fmt.Sprintf("<UserCost>%s</UserCost>", t.String()))
		}

	} else { //not found
		if runner.NotFound != nil {
			runner.NotFound(context)
		} else {
			context.ResponseFailTextWithCode(http.StatusNotFound, "command not found")
		}
	}

}
