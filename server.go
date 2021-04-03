package main

import (
	"github.com/onurrsalmann/api_log/helper"
	"github.com/onurrsalmann/api_log/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.Use(gin.Recovery(), middlewares.Logger())
	r.GET("/get", func (context *gin.Context) {
		res := helper.BuildResponse(true, "OK!", "")
		context.JSON(http.StatusOK, res)
	})
	r.POST("/post", func (context *gin.Context) {
		res := helper.BuildResponse(true, "OK!", "")
		context.JSON(http.StatusOK, res)
	})
	r.PUT("/put", func (context *gin.Context) {
		res := helper.BuildResponse(true, "OK!", "")
		context.JSON(http.StatusOK, res)
	})
	r.DELETE("/delete", func(context *gin.Context) {
		res := helper.BuildResponse(true, "OK!", "")
		context.JSON(http.StatusOK, res)
	})

	r.Run()

}