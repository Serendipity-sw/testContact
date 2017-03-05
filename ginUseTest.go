package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
)

func main() {
	r := gin.New()
	f:=r.Group("/f")
	{
		f.Use(Logger())
		f.GET("/test", func(c *gin.Context) {
			//example := c.MustGet("example").(string)
			return
			// it would print: "12345"
			fmt.Println(2)
		})
	}
	r.GET("/test", func(c *gin.Context) {
		//example := c.MustGet("example").(string)

		// it would print: "12345"
		fmt.Println(3)
	})


	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(1)
		fmt.Println(c.Request.Host)
		c.Redirect(http.StatusFound,"https://www.baidu.com")

	}
}
