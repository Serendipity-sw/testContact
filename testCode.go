package main

import (
	"github.com/swgloomy/gutil"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/dchest/captcha"
	"time"
	"fmt"
	"github.com/hanguofeng/gocaptcha"
)

func main() {
	gutil.GinInit(true, 8077, "",SetGinRouter )
	time.Sleep(30*time.Hour)
}

func SetGinRouter(r *gin.Engine, rootPrefix string)  {
	g := &r.RouterGroup
	if rootPrefix != "" {
		g = r.Group(rootPrefix)
	}
	{
		g.GET("/", func(c *gin.Context) { c.String(http.StatusOK, "ok") }) //确认接口服务程序是否健在
		g.GET("/unitImageCode",unitImageCode) //验证码
	}
}

var (
	ccaptcha   *gocaptcha.Captcha
)

//图片验证码
//create by gloomy 2017-10-9 17:51:57
func unitImageCode(c *gin.Context) {
	messageCode:=captcha.New()
	fmt.Println(messageCode)
	captcha.WriteImage(c.Writer,messageCode,100,40)
}