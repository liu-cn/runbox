package main

import (
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/runbox/cmd/runbelet/proxy"
	"log"
)

func main() {
	r := gin.Default()
	newProxy, err := proxy.NewHttpProxy("http://127.0.0.1:17777")
	if err != nil {
		panic(err)
	}
	//代理
	//前端访问：http://127.0.0.1:16666/api/runner/run/beiluo/image/views/test?image_name=2.png
	//请求会代理转发到下游地址：http://127.0.0.1:17777/runner/run/beiluo/image/views/test?image_name=2.png
	//其中NewProxy 有对请求做了一下操作，比如添加元数据等
	r.Any("/api/*path", func(c *gin.Context) {
		// 修改请求路径，去掉/proxy前缀
		c.Request.URL.Path = c.Param("path")
		newProxy.ServeHTTP(c.Writer, c.Request)
	})

	log.Println("Starting server on :16666")
	if err := r.Run(":16666"); err != nil {
		log.Fatal(err)
	}

}
