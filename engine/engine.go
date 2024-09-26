package engine

import (
	"github.com/liu-cn/runbox/model"
	"github.com/liu-cn/runbox/model/request"
	"github.com/liu-cn/runbox/model/response"
	"github.com/liu-cn/runbox/pkg/store"
	"github.com/liu-cn/runbox/runner"
)

// RunBox 引擎负责调度，管理和执行 各种Runner
type RunBox struct {
	FileStore store.FileStore
}

func NewRunBox(fileStore store.FileStore) *RunBox {
	return &RunBox{FileStore: fileStore}
}

// Run 执行请求
func (b *RunBox) Run(call *request.Run, runnerConf *model.Runner) (*response.Run, error) {
	newRunner := runner.NewRunner(runnerConf)
	rspCall, err := newRunner.Run(call)
	if err != nil {
		return nil, err
	}
	return rspCall, nil
}
