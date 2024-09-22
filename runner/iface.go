package runner

import (
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/runbox/model/request"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/runbox/model/response"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/runbox/pkg/store"
)

type Runner interface {
	Install(store store.FileStore) (installInfo *InstallInfo, err error)
	UnInstall() (unInstallInfo *UnInstallInfo, err error)
	Call(req *request.Call) (*response.Call, error)
}
