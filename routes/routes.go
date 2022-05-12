package routes

import (
	"bluebell/controllers"
	"bluebell/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUp(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) //将gin设置成发布模式：即不在控制台打印
	}
	r := gin.New()
	//使用自定义日志打印
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello脚手架")
	})

	//注册业务路由
	r.POST("/signup", controllers.SignUpHandler)

	//登录业务路由
	r.POST("/login", controllers.LoginHandler)
	return r
}
