// Package classification loveauth API.
//
// Auth server for World of Love.
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http, https
//     BasePath: /api/v1
//     Version: v1.0.64 -- develop(0efd925)
//     License: Apache License 2.0 http://www.apache.org/licenses/LICENSE-2.0.html
//     Contact: xy.kong<xy.kong@gmail.com>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//
// swagger:meta

package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var logo = `
    BasePath: /api/v1
    Version: v1.0.64 -- develop(0efd925)
    Contact: 1455975151@qq.com <1455975151@qq.com>
`

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2

func Run() {
	logrus.Info(logo)
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// use ginSwagger middleware to
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	r.Run() // listen and serve on 0.0.0.0:8080
}