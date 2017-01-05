package main

import (
	"github.com/guotie/deferinit"
	"github.com/gin-gonic/gin"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)
var (
	rt                                   *gin.Engine
	rootPrefix=""
)

func main() {
	deferinit.RunRoutines()
	rt = gin.Default()
	router(rt)
	go rt.Run(fmt.Sprintf(":%d", 8754))

	c := make(chan os.Signal, 1)
	// 信号处理
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	// 等待信号
	<-c
}

func router(r *gin.Engine) {

	g := &r.RouterGroup
	if rootPrefix != "" {
		g = r.Group(rootPrefix)
	}
	{
		g.GET("/", func(c *gin.Context) { c.String(200, "ok") })
		g.POST("/nihao",nihao)
	}
}

func nihao(c *gin.Context) {
	fmt.Println(c.Request.Host)
}