package main

import (
	"github.com/liu-cn/runbox/sdk/runner"
	"github.com/liu-cn/runbox/soft_cmd/beiluo/apphub/biz/array"
	"github.com/liu-cn/runbox/soft_cmd/beiluo/apphub/biz/images"
	"github.com/liu-cn/runbox/soft_cmd/beiluo/apphub/biz/strings"
	"github.com/liu-cn/runbox/soft_cmd/beiluo/apphub/biz/table"
)

func main() {
	r := runner.New()
	r.Get("hello", func(ctx *runner.Context) {
		ctx.ResponseOkWithJSON(map[string]string{"hello": "world"})
	}, runner.WithPublicApi())
	r.Get("assets", images.View, runner.WithPublicApi())
	r.Post("array/Diff", array.Diff, runner.WithPublicApi())
	r.Post("array/ComputeIntersection", array.ComputeIntersection, runner.WithPublicApi())
	r.Post("array/Split", array.Split, runner.WithPublicApi())
	r.Post("strings/Split", strings.Split, runner.WithPublicApi())
	r.Post("strings/ReplaceAll", strings.ReplaceAll, runner.WithPublicApi())
	r.Post("table/Demo", table.Demo, runner.WithPublicApi())
	r.Post("array/SplitJoin", array.SplitJoin, array.WithSplitJoinOpt())
	r.Post("strings/StatisticsText", strings.StatisticsText, strings.WithStatisticsTextOpt())
	r.Run()
}
