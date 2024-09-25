package runner

import (
	"github.com/liu-cn/runbox/model"
	"github.com/liu-cn/runbox/model/request"
	"github.com/liu-cn/runbox/model/response"
	"github.com/liu-cn/runbox/pkg/store"
)

func NewRunner(runner *model.Runner) Runner {
	return NewCmd(runner)
}

type Runner interface {
	Install(store store.FileStore) (installInfo *InstallInfo, err error)
	GetInstallPath() (path string)
	UnInstall() (unInstallInfo *UnInstallInfo, err error)
	Call(req *request.Run) (*response.Run, error)
}
