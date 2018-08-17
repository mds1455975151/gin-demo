package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var logo = `
    BasePath: /api/v1
    Version: v1.0.64 -- develop(0efd925)
    Contact: 1455975151@qq.com <1455975151@qq.com>
`

func Run() {
	logrus.Info(logo)
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}