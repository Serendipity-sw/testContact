package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
	. "github.com/CodyGuo/win"
)

var (
	rt            *gin.Engine
	rootPrefix    string
)

func main() {
	rt = gin.Default()
	router(rt)

	go rt.Run(":8945")

	c := make(chan os.Signal, 1)

	// 信号处理
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	// 等待信号
	<-c


	os.Exit(0)
}


func router(r *gin.Engine) {
	g := &r.RouterGroup
	if rootPrefix != "" {
		g = r.Group(rootPrefix)
	}

	{
		g.GET("/", func(c *gin.Context) { c.String(200, "ok") })
		g.GET("/ShutDown", ShutDown)
	}
}

func ShutDown(c *gin.Context) {
	getPrivileges()
	ExitWindowsEx(EWX_SHUTDOWN, 0)
}

func getPrivileges() {
	var hToken HANDLE
	var tkp TOKEN_PRIVILEGES

	OpenProcessToken(GetCurrentProcess(), TOKEN_ADJUST_PRIVILEGES|TOKEN_QUERY, &hToken)
	LookupPrivilegeValueA(nil, StringToBytePtr(SE_SHUTDOWN_NAME), &tkp.Privileges[0].Luid)
	tkp.PrivilegeCount = 1
	tkp.Privileges[0].Attributes = SE_PRIVILEGE_ENABLED
	AdjustTokenPrivileges(hToken, false, &tkp, 0, nil, nil)
}