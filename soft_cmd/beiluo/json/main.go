package main

import (
	"github.com/liu-cn/runbox/pkg/jsonx"
	"github.com/liu-cn/runbox/sdk/runner"
	"github.com/sirupsen/logrus"
)

func main() {
	r := runner.New()
	r.SetVersion("v1")
	r.Post("echo", func(ctx *runner.Context) {
		s := struct {
			Label string                 `json:"label"`
			Data  map[string]interface{} `json:"data"`
		}{
			Label: "json",
			Data:  ctx.ReqMap(),
		}
		ctx.ResponseOkWithJSON(s)

	})

	r.Get("hi", func(ctx *runner.Context) {
		ctx.GetLogger().Info("hello info log")
		ctx.ResponseOkWithText("hello info res")
	})

	r.Get("file", func(ctx *runner.Context) {
		logrus.Infof("file" + jsonx.String(ctx.ReqMap()))
		jsonx.SaveFile("./request.json", ctx.ReqMap())
		ctx.ResponseOkWithFile("./request.json", true)
	})

	r.Get("helloWorld", func(ctx *runner.Context) {
		logrus.Infof("helloWorld" + jsonx.String(ctx.ReqMap()))
		ctx.ResponseOkWithText("hello world")
	})
	//todo 最后必须执行run方法
	r.Run()
}
