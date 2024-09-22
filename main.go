package main

import (
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/runbox/pkg/logger"
	"log"
)

func init() {
	logger.Setup()
}

func main() {
	v1 := gin.Default()
	log.Fatal(v1.Run(":17777"))
}
