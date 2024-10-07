package images

import (
	"github.com/liu-cn/runbox/sdk/runner"
	"github.com/liu-cn/runbox/soft_cmd/beiluo/apphub/model/dto"
)

func View(ctx *runner.Context) {

	var r dto.ViewReq

	err := ctx.ShouldBindJSON(&r)
	if err != nil {
		ctx.ResponseFailDefaultJSONWithMsg("参数错误")
		return
	}
	//ctx.ResponseOkWithFile()
}
