package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var ErrorUserNotLogin = errors.New("用户未登录")

const CtxUserIDKey = "userID"

// Go中不支持包的互相调用，此处调用middlerwares中的变量，而middlerwares中又调用controlllers中的response
func getCurrentUserID(c *gin.Context) (userId int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
	}
	//考虑转化失败的情况
	userId, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
