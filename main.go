package main

import (
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/runbox/conf"
	"github.com/liu-cn/runbox/pkg/logger"
	"github.com/liu-cn/runbox/pkg/natsx"
	"github.com/liu-cn/runbox/router"
	"github.com/liu-cn/runbox/service"
	"log"
)

func init() {
	logger.Setup()
	conf.Setup(&conf.LocalConfig{FilePath: "./config.yaml"})
}

func main() {
	v1 := gin.Default()
	natsx.Setup(conf.Config.Nats.Url)
	defer natsx.Nats.Close()
	box := service.NewDefaultRunBox()
	err := box.HandelMsg()
	if err != nil {
		panic(err)
	}
	defer box.Close()
	router.Register(v1)
	log.Fatal(v1.Run(":17777"))
}
