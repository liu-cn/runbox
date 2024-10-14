package runner

import (
	"encoding/json"
	"fmt"
	"github.com/liu-cn/runbox/pkg/logger"
	logger2 "github.com/liu-cn/runbox/sdk/runner/logger"
	"github.com/liu-cn/runbox/sdk/runner/request"
	"os"
)

func init() {
	logger.Setup()
}

type Context struct {
	IsDebug  bool
	Cmd      string   `json:"cmd"`
	WorkPath string   `json:"work_path"` //工作目录
	Request  string   `json:"request"`
	FileList []string `json:"file_list"`

	Req    *request.Call
	runner *Runner `json:"-"`
}

func (c *Context) GetLogger() *logger2.Logger {

	mp := make(map[string]interface{})
	mp["a_tenant"] = c.Req.User
	mp["a_soft"] = c.Req.Soft
	mp["a_command"] = c.Req.Command
	if c.runner != nil {
		if c.runner.Version != "" {
			mp["a_version"] = c.runner.Version
		}
	}

	mp["a_soft_classify"] = fmt.Sprintf("/%s/%s", c.Req.User, c.Req.Soft)
	return &logger2.Logger{
		DataMap: mp,
	}
}

func (r *Runner) Notfound(fn func(ctx *Context)) {
	r.NotFound = fn
}

func (c *Context) DebugPrintf(format string, args ...interface{}) {
	if c.IsDebug {
		fmt.Printf(format, args...)
	}
}

func (r *Runner) DebugPrintf(format string, args ...interface{}) {
	if r.IsDebug {
		fmt.Printf(format, args...)
	}
}

//func bind(ctx *Context) error {
//	var ca request.Call
//	err := jsonx.UnmarshalFromFile(ctx.Request, &ca)
//	ctx.DebugPrintf("bind err:%v", err)
//	if err != nil {
//		return err
//	}
//	ctx.Req = &ca
//	return nil
//}

func (c *Context) ShouldBindJSON(jsonBody interface{}) error {
	if c.Req != nil {
		marshal, err := json.Marshal(c.Req.Body)
		if err != nil {
			return err
		}
		err = json.Unmarshal(marshal, jsonBody)
		if err != nil {
			return err
		}
	}
	return nil
}
func (c *Context) ReqMap() map[string]interface{} {
	if c.Req != nil {
		return c.Req.Body
	}
	return nil
}

func New() *Runner {
	r := &Runner{
		CmdMapHandel: make(map[string]*Worker),
	}
	for path, fn := range GetMap {
		r.Get(path, fn)
	}
	return r
}

type Worker struct {
	Handel []func(*Context)
	Path   string
	Method string
	Config *Config
}

type Runner struct {
	IsDebug bool
	Version string

	About        bool
	CmdMapHandel map[string]*Worker
	NotFound     func(ctx *Context)
}

func (r *Runner) SetVersion(version string) {
	r.Version = version
}
func (r *Runner) Post(commandName string, handelFunc func(ctx *Context), opts ...Option) {
	_, ok := r.CmdMapHandel[commandName]
	if !ok {
		worker := &Worker{
			Handel: []func(*Context){handelFunc},
			Method: "POST",
			Path:   commandName,
			Config: &Config{},
		}
		if len(opts) > 0 {
			for _, opt := range opts {
				opt(worker.Config)
			}
		}
		r.CmdMapHandel[commandName+".POST"] = worker
	} else {
		r.CmdMapHandel[commandName].Handel = append(r.CmdMapHandel[commandName].Handel, handelFunc)
	}

}
func (r *Runner) Get(commandName string, handelFunc func(ctx *Context), opts ...Option) {
	_, ok := r.CmdMapHandel[commandName]
	if !ok {
		worker := &Worker{
			Handel: []func(*Context){handelFunc},
			Method: "GET",
			Path:   commandName,
			Config: &Config{},
		}
		if len(opts) > 0 {
			for _, opt := range opts {
				opt(worker.Config)
			}
		}
		r.CmdMapHandel[commandName+".GET"] = worker
	} else {
		r.CmdMapHandel[commandName].Handel = append(r.CmdMapHandel[commandName].Handel, handelFunc)
	}
}

func (r *Runner) Run() {
	command := os.Args[1]
	jsonFileName := os.Args[2]
	//workPath := os.Args[3]
	if len(os.Args) > 3 {
		r.IsDebug = true
	}
	r.DebugPrintf("run ....")

	context := &Context{Request: jsonFileName, IsDebug: r.IsDebug, Cmd: command}
	Handel(context, r)
	if r.About {
		return
	}

}
