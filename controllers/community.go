package controllers

import (
	"bluebell/logic"
	"strconv"

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
		return
	}
	ResponseSuccess(c, data)
}

//  CommunityDetailHandler 社区分类详情
func CommunityDetailHandler(c *gin.Context) {
	//1参数处理
	//逻辑处理
	//返回参数

	//获取社区id
	idstr := c.Param("id")
	CommunityId, err := strconv.ParseInt(idstr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	//调用loigc处理事务：查询
	data, err := logic.GetCommunityDetail(CommunityId)
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(c, data)
}
