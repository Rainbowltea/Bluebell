package controllers

import (
	"bluebell/logic"
	"bluebell/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context) {

	//1.参数校验
	//2.业务处理
	//3.返回响应
	// var p models.ParamSignUp
	p := new(models.ParamSignUp)
	//1
	if err := c.ShouldBindJSON(p); err != nil {
		//请求参数有误
		//日志记录错误的格式：类型+详细信息
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg": "请求参数有误",
		})
		return
	}
	//shouldbindjson只能检验参数格式是否正确
	//当要求输入不能为空等情况需手动加入判断
	if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.Password != p.RePassword {
		zap.L().Error("SignUp with invalid param")
		c.JSON(http.StatusOK, gin.H{
			"msg": "请求参数有误",
		})
		return
	}

	//2
	logic.SignUp(p)

	//3
	// c.JSON(http.StatusOK, "ok")
	//返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})

}
