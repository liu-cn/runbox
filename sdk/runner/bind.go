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

type Api struct {
	Path   string `json:"path"`
	Method string `json:"method"`
	Config
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

	if context.Cmd == "_func_all" && method == "GET" { //获取接口文档
		var apis []Api
		for _, worker := range runner.CmdMapHandel {
			//if worker.Config == nil {
			//	continue
			//}
			cfg := worker.Config
			params, err := cfg.GetParams()
			if err != nil {
				context.ResponseFailJSONWithCode(http.StatusBadRequest, map[string]interface{}{
					"msg": "参数解析失败: " + err.Error(),
				})
				runner.About = true
			}
			cfg.Params = params
			apis = append(apis, Api{
				Path:   worker.Path,
				Method: worker.Method,
				Config: *cfg,
			})

			//s = append(s, fmt.Sprintf("%s\t %s \t %s", worker.Path, worker.Method, worker.Config.ApiDesc))
		}
		//res := append([]string{"请求地址 \t 请求方式 \t 接口描述"}, s...)
		//context.ResponseOkWithText(strings.Join(s, "\n"))
		context.ResponseOkWithJSON(map[string]interface{}{
			"data": apis,
		})
		runner.About = true
		return
	}

	//todo
	if context.Cmd == "_docs_info_text" && method == "GET" { //获取接口文档
		var s []string
		for _, worker := range runner.CmdMapHandel {
			if worker.Config == nil {
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
		api := context.ReqMap()["api"]
		methodApi := context.ReqMap()["method"]
		if api == nil || methodApi == nil {
			context.ResponseFailJSONWithCode(http.StatusBadRequest, map[string]interface{}{
				"msg": "请填写api",
			})
			runner.About = true
			return
		}
		apiPath := fmt.Sprintf("%v", api)
		methodApiStr := fmt.Sprintf("%v", methodApi)
		//var s []string
		key := fmt.Sprintf("%v.%v", apiPath, methodApiStr)
		worker := runner.CmdMapHandel[key]
		if worker == nil {
			context.ResponseFailJSONWithCode(http.StatusBadRequest, map[string]interface{}{
				"msg": "api不存在",
			})
			runner.About = true
			return
		}
		params, err := worker.Config.GetParams()
		if err != nil {
			context.ResponseFailJSONWithCode(http.StatusBadRequest, map[string]interface{}{
				"msg": err.Error(),
			})
			runner.About = true
			return
		}
		worker.Config.Params = params
		context.ResponseFailJSONWithCode(http.StatusOK, map[string]interface{}{
			"code": 0,
			"msg":  "ok",
			"data": worker.Config,
		})
		runner.About = true
		return
		//for _, worker := range runner.CmdMapHandel {
		//	if worker.Config == nil {
		//		continue
		//	}
		//	if !worker.Config.IsPublicApi {
		//		continue
		//	}
		//	s = append(s, fmt.Sprintf("%s\t %s \t %s", worker.Path, worker.Method, worker.Config.ApiDesc))
		//}
		//res := append([]string{"请求地址 \t 请求方式 \t 接口描述"}, s...)
		//context.ResponseOkWithText(strings.Join(s, "\n"))
		//runner.About = true
		//return
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
