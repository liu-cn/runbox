package main

import (
	"github.com/liu-cn/runbox/runner_code/apphub/biz/array"
	"github.com/liu-cn/runbox/runner_code/apphub/biz/images"
	"github.com/liu-cn/runbox/runner_code/apphub/biz/strings"
	"github.com/liu-cn/runbox/runner_code/apphub/biz/table"
	"github.com/liu-cn/runbox/sdk/runner"
)

func main() {
	r := runner.New()
	r.Get("hello", func(ctx *runner.Context) {
		ctx.ResponseOkWithJSON(map[string]string{"hello": "world"})
	})
	r.Get("assets", images.View)
	r.Post("array/Diff", array.Diff)
	r.Post("array/ComputeIntersection", array.ComputeIntersection)
	r.Post("array/Split", array.Split)
	r.Post("strings/Split", strings.Split)
	r.Post("strings/ReplaceAll", strings.ReplaceAll)
	r.Post("table/Demo", table.Demo)
	r.Post("array/SplitJoin", array.SplitJoin)
	r.Run()
}
