package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/liu-cn/runbox/api/v1"
)

func Register(r *gin.Engine) {
	box := v1.NewDefaultApi()
	runnerGroup := r.Group("/runner")
	runnerGroup.GET("/run/:user/:soft/*command", box.Run)  //run
	runnerGroup.POST("/run/:user/:soft/*command", box.Run) //run
	runnerGroup.POST("/install", box.Install)              //安装
	runnerGroup.POST("/updateVersion", box.UpdateVersion)  //版本更新
}
