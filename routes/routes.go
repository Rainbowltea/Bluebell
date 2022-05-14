package routes

import (
	"bluebell/controllers"
	"bluebell/logger"
	"bluebell/middlerwares"
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

	r.GET("/ping", middlerwares.JWTAuthMiddleware(), func(ctx *gin.Context) {
		// 判断是否是已登录的用户，判断请求头中是否有 有效的JWT
		ctx.String(http.StatusOK, "pong")
	})
	return r
}
