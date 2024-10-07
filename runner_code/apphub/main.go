package main

import (
	"github.com/liu-cn/runbox/runner_code/apphub/biz/images"
	"github.com/liu-cn/runbox/sdk/runner"
)

func main() {
	r := runner.New()
	r.Get("hello", func(ctx *runner.Context) {
		ctx.ResponseOkWithJSON(map[string]string{"hello": "world"})
	})
	r.Get("view", images.View)
	r.Run()

}
