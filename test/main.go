package main

import (
	"fmt"
	"github.com/liu-cn/runbox/model"
	"github.com/liu-cn/runbox/pkg/store"
	"github.com/liu-cn/runbox/runner"
)

func main() {
	fileStore := store.NewDefaultQiNiu()
	runnerEl := model.Runner{
		AppCode:    "timer",
		ToolType:   "可执行程序(linux)",
		Version:    "v1",
		TenantUser: "beiluo",
		OssPath:    "runner/beiluo/timer/v1/timer.zip",
	}
	cmd := runner.NewCmd(&runnerEl)
	install, err := cmd.Install(fileStore)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("%+v", install))
}
