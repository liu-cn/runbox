package engine

import (
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
