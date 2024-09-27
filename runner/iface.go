package runner

import (
	"github.com/liu-cn/runbox/model"
	"github.com/liu-cn/runbox/model/request"
	"github.com/liu-cn/runbox/model/response"
	"github.com/liu-cn/runbox/pkg/store"
)

func NewRunner(runner *model.Runner) Runner {
	switch runner.ToolType {
	case "docker":
		//	todo 待实现
	case "python":
		//todo 待实现
	case "windows", "linux", "macos":
		return NewCmd(runner)
	}
	return NewCmd(runner)
}

// Runner RunBox 引擎可以调度任何实现Runner 接口的程序（可执行程序|批处理文件|python代码|lua|ruby|nodejs|docker镜像）
type Runner interface {
	Install(store store.FileStore) (installInfo *InstallInfo, err error)                      //安装程序
	GetInstallPath() (path string)                                                            //获取安装路径
	UnInstall() (unInstallInfo *UnInstallInfo, err error)                                     //卸载
	UpdateVersion(up *model.UpdateVersion, fileStore store.FileStore) (*UpdateVersion, error) //更新版本
	Run(req *request.Run) (*response.Run, error)                                              //运行程序
}
