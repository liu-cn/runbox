package main

import (
	"github.com/liu-cn/runbox/sdk/runner"
	"github.com/liu-cn/runbox/soft_cmd/beiluo/strings/biz"
)

func main() {

	r := runner.New()
	r.Post("split", biz.Split)
	r.Post("ReplaceAll", biz.ReplaceAll)
	r.Run()
}
