package controllers

import "bluebell/models"

//用来存放接口文档用到的model
type _ResponsePostList struct {
	Code    ResCode                 `json:"code"`
	Message string                  `json:"message"`
	Data    []*models.ApiPostDetail `json:"data"` //数据
}
