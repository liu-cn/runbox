package main

import (
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/runbox/pkg/logger"
	"github.com/liu-cn/runbox/router"
	"log"
)

func init() {
	logger.Setup()
}

func main() {
	v1 := gin.Default()
	router.Register(v1)
	log.Fatal(v1.Run(":17777"))
}
