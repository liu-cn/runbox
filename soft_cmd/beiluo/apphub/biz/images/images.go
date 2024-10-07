package images

import (
	"fmt"
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
	err = ctx.Response(runner.Response{
		FilePath:       "./assets/" + r.FilePath,
		DeleteFileTime: -1,
	})
	if err != nil {
		fmt.Println(err)
	}
}
