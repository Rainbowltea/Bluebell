package controllers

import (
	"bluebell/dao/mysql"
	"bluebell/logic"
	"bluebell/models"
	"errors"
	_ "go/token"
	"net/http"

	"github.com/go-playground/validator/v10"

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
		//为了让前端更加清楚出现的错误，使用翻译器
		// 判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			//现开发阶段打印完整报错有助于快速定位错误
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})

			/*
				自定义错误（便于前端开发人员识别）
				ResponseError(c, CodeInvalidParam)
			*/

			return
		}
		/*
			附带额外信息返回
			RespoonseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		*/

		c.JSON(http.StatusOK, gin.H{
			//添加一个函数来去除前端中结构体名称
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}
	//一般使用validator库来对参数进行校验
	//shouldbindjson只能检验参数格式是否正确
	//当要求输入不能为空等情况需手动加入判断
	// if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.Password != p.RePassword {
	// 	zap.L().Error("SignUp with invalid param")
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"msg": "请求参数有误",
	// 	})
	// 	return
	// }

	//2
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		/*
			若不是“用户已存在”的错误，就可能是数据库服务繁忙
			ResponseError(c,CodeServerBusy)
		*/
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		//函数返回
		return
	}

	//3
	// c.JSON(http.StatusOK, "ok")
	//返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
	/*
		便于前端开发人员识别
		ResponseSuccess(c, nil)
	*/
}

func LoginHandler(c *gin.Context) {
	//1.参数处理
	//2.业务逻辑处理
	//3.返回响应
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		//判断err是不是属于validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			/*
				ResponseError(c,CodeInvalidParam)
			*/
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		/*
			RespoonseErrorWithMsg(c,CodeInvalidParam,removeTopStruct(errs.Translate(trans)))
		*/
		return
	}
	//业务逻辑处理
	token, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login falied", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": "用户名或者密码错误",
		})
		/*
			ResponseError(c,CodeInvalidPassword)
		*/
	}
	//3.返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "登录成功",
	})
	ResponseSuccess(c, token)
}
