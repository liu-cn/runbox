package testarray

import (
	"github.com/liu-cn/runbox/model"
	"github.com/liu-cn/runbox/pkg/jsonx"
)

func TestArray() {
	runnerEl := model.Runner{
		AppCode:    "array",
		ToolType:   "windows",
		Version:    "v1",
		TenantUser: "beiluo",
		OssPath:    "runner/beiluo/timer/v1/array.zip",
	}
	jsonx.MustPrintJSON(runnerEl)
}
