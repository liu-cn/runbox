package engine

import (
	"github.com/liu-cn/runbox/model"
	"github.com/liu-cn/runbox/model/request"
	"github.com/liu-cn/runbox/model/response"
	"github.com/liu-cn/runbox/pkg/store"
	"github.com/liu-cn/runbox/runner"
)

type RunBox struct {
	FileStore store.FileStore
}

func NewRunBox(fileStore store.FileStore) *RunBox {
	return &RunBox{FileStore: fileStore}
}

//func (b *RunBox) Call() {
//
//}

func (b *RunBox) CallRunner(call request.Run) *response.Run {

	return nil

}

//
//// GetCall 执行get请求
//func (b *RunBox) GetCall(call *request.Call, runnerConf *model.Runner) (*response.Call, error) {
//	newRunner := runner.NewRunner(runnerConf)
//	rspCall, err := newRunner.Call(call)
//	if err != nil {
//		return nil, err
//	}
//	return rspCall, nil
//}

// Run 执行请求
func (b *RunBox) Run(call *request.Run, runnerConf *model.Runner) (*response.Run, error) {
	newRunner := runner.NewRunner(runnerConf)
	rspCall, err := newRunner.Call(call)
	if err != nil {
		return nil, err
	}
	return rspCall, nil
}
