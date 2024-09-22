package router

import "github.com/gin-gonic/gin"

func Init(r *gin.Engine) {
	r.GET("/call", func(c *gin.Context) {})
}
