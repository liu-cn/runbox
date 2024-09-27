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

// Install 安装软件
func (b *RunBox) Install(runnerConf *model.Runner) (*runner.InstallInfo, error) {
	newRunner := runner.NewRunner(runnerConf)
	info, err := newRunner.Install(b.FileStore)
	if err != nil {
		return nil, err
	}
	return info, nil
}

// UpdateVersion 更新软件
func (b *RunBox) UpdateVersion(updateRunner *model.UpdateVersion) (*runner.UpdateVersion, error) {
	newRunner := runner.NewRunner(updateRunner.RunnerConf)
	info, err := newRunner.UpdateVersion(updateRunner, b.FileStore)
	if err != nil {
		return nil, err
	}
	return info, nil
}
