package controllers

import (
	"bluebell/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func PostVoteHandler(c *gin.Context) {
	//参数校验
	//业务处理
	//返回数据

	vote1 := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(vote1); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("PostVote with invaliedparam", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors) // 类型断言
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans)) // 翻译并去除掉错误提示中的结构体标识
		RespoonseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}
}
