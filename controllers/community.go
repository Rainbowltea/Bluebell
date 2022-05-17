package controllers

import (
	"bluebell/logic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CommunityHandler(c *gin.Context) {
	//1参数处理
	//逻辑处理
	//返回参数

	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunicaty failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
	}
	ResponseSuccess(c, data)
}
