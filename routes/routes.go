package routes

import (
	"bluebell/controllers"
	"bluebell/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUp() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello脚手架")
	})

	//注册业务路由
	r.POST("/signup", controllers.SignUpHandler)
	return r
}
