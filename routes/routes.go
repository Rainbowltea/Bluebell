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

	v1 := r.Group("/api/v1")

	// 注册
	v1.POST("/signup", controllers.SignUpHandler)
	// 登录
	v1.POST("/login", controllers.LoginHandler)

	v1.Use(middlerwares.JWTAuthMiddleware()) // 应用JWT认证中间件

	{
		v1.GET("/community", controllers.CommunityHandler)
		v1.GET("/communtiy/:id", controllers.CommunityDetailHandler)
		v1.POST("/post", controllers.CreatePostHandler)
		v1.GET("/post/:id", controllers.PostDetailHandler)
	}
	return r
}
