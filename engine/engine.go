package engine

import (
	"github.com/liu-cn/runbox/model"
	"github.com/liu-cn/runbox/model/request"
	"github.com/liu-cn/runbox/model/response"
	"github.com/liu-cn/runbox/pkg/store"
)

type RunBox struct {
	FileStore store.FileStore
}

func NewRunBox(fileStore store.FileStore) *RunBox {
	return &RunBox{FileStore: fileStore}
}

func (b *RunBox) Call() {

}

func (b *RunBox) CallRunner(call request.Call) *response.Call {

	return nil

}

// GetCall 执行get请求
func (b *RunBox) GetCall(call *request.Call, runner *model.Runner) (*response.Call, error) {
	return nil, nil
}
