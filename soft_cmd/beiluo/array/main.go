package main

import (
	"github.com/liu-cn/runbox/sdk/runner"
	"github.com/liu-cn/runbox/soft_cmd/beiluo/array/biz"
)

func main() {
	r := runner.New()
	r.Post("diff", biz.Diff)
	r.Post("Split", biz.Split)
	r.Post("ComputeIntersection", biz.ComputeIntersection)
	r.Run()
}
