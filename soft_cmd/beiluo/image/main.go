package main

import (
	"fmt"
	"github.com/liu-cn/runbox/sdk/runner"
)

func main() {

	r := runner.New()
	r.Get("views", func(ctx *runner.Context) {
		reqMap := ctx.ReqMap()
		imageName := reqMap["image_name"].(string)
		err := ctx.Response(runner.Response{
			HttpStatusCode: 200,
			FilePath:       "./assets/" + imageName,
			DeleteFileTime: -1,
			Body:           nil,
		})
		if err != nil {
			fmt.Println(err)
		}
	})

	r.Get("assets", func(ctx *runner.Context) {
		reqMap := ctx.ReqMap()
		name := reqMap["name"].(string)
		err := ctx.Response(runner.Response{
			HttpStatusCode: 200,
			FilePath:       "./assets/" + name,
			DeleteFileTime: -1,
			Body:           nil,
		})
		if err != nil {
			fmt.Println(err)
		}
	})

	r.Run()
}
