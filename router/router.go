package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/liu-cn/runbox/api/v1"
)

func Register(r *gin.Engine) {
	box := v1.NewDefaultApi()
	runnerGroup := r.Group("/runner")
	runnerGroup.GET("/run/:user/:soft/*command", box.Run)
	runnerGroup.POST("/run/:user/:soft/*command", box.Run)
}
