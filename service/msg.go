package service

import (
	"encoding/json"
	"fmt"
	"github.com/liu-cn/runbox/engine"
	"github.com/liu-cn/runbox/model"
	"github.com/liu-cn/runbox/model/request"
	"github.com/liu-cn/runbox/pkg/natsx"
	"github.com/liu-cn/runbox/pkg/store"
	"github.com/liu-cn/runbox/runner"
	"github.com/nats-io/nats.go"
	"os"
	"strconv"
	"strings"
)

type RunBox struct {
	Nats *nats.Conn
	Eng  *engine.RunBox
	sub  *nats.Subscription
}

func (r *RunBox) Close() {
	r.sub.Unsubscribe()
	r.Nats.Close()
}

func NewRunBox(conn *nats.Conn) *RunBox {
	return &RunBox{
		Nats: conn,
		Eng:  engine.NewRunBox(store.NewDefaultQiNiu()),
	}
}
func NewDefaultRunBox() *RunBox {
	return &RunBox{
		Nats: natsx.Nats,
		Eng:  engine.NewRunBox(store.NewDefaultQiNiu()),
	}
}
func (r *RunBox) HandelMsg() error {
	sub, err := r.Nats.Subscribe("runner.run.*.*.*", func(msg *nats.Msg) {
		var (
			req request.Run
			err error
		)
		subs := strings.Split(msg.Subject, ".")
		user := subs[2]
		soft := subs[3]
		cmd := subs[4]
		runnerMeta := &model.Runner{
			AppCode:    soft,
			TenantUser: user,
			OssPath:    msg.Header.Get("fs_path"),
			ToolType:   msg.Header.Get("type"),
			Version:    msg.Header.Get("version"),
		}
		req.User = user
		req.Soft = soft
		req.Command = cmd
		req.Command = strings.TrimPrefix(req.Command, "/")
		req.Method = msg.Header.Get("method")
		rspMsp := nats.NewMsg(msg.Subject)

		if msg.Data != nil {
			err = json.Unmarshal(msg.Data, &req.Body)
			if err != nil {
				rspMsp.Header.Set("error", "参数错误")
				msg.RespondMsg(rspMsp)
				return
			}
		}

		run := runner.NewRunner(runnerMeta)

		req.RequestJsonPath = strings.ReplaceAll(req.GetRequestFilePath(run.GetInstallPath()), "\\", "/")
		defer os.Remove(req.RequestJsonPath)
		rsp, err := r.Eng.Run(&req, runnerMeta)
		if err != nil {
			rspMsp.Header.Set("error", err.Error())
			msg.RespondMsg(rspMsp)
			return
		}
		rsp.GetResponseMetaData().Print()

		rspMsp.Header.Set("Status-Code", strconv.Itoa(rsp.StatusCode))
		rspMsp.Header.Set("Content-Type", rsp.GetContentType())
		rspMsp.Header.Set("Cost", strconv.FormatInt(rsp.CallCostTime.Nanoseconds(), 10))
		data, err := json.Marshal(rsp.Body)
		if err != nil {
			fmt.Println(err)
		}
		//cost := rsp.CallCostTime
		//timex.Println(cost)
		//执行引擎发起调用到程序执行结束的总耗时

		rspMsp.Data = data
		err = msg.RespondMsg(rspMsp)
		if err != nil {
			fmt.Println(err)
		}

		//name := msg.Subject[len("greet."):]
		//msg.Respond([]byte("hello, " + name + string(msg.Data)))

	})
	if err != nil {
		return err
	}
	r.sub = sub
	return nil
}
