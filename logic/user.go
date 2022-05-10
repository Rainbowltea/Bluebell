package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) {

	//1.判断用户存不存在
	//2.若不存在则生成id
	//2.1保存id进数据库
	//3.若存在则与数据库中已经有的数据进行比对
	//3.1相等或不相等

	mysql.QueryUserByUsername()

	snowflake.GenID()

	mysql.InsertUser()
}
