package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
	_ "errors"
	_ "go/token"
)

func SignUp(p *models.ParamSignUp) (err error) {

	//1.判断用户存不存在
	//2.若不存在则生成id
	//3.保存id进数据库

	//1判断
	if err := mysql.CheckUserExist(p.Username); err != nil {
		//数据库查询出错
		return err
	}
	//2生成Uid
	userID := snowflake.GenID()
	//构造一个User实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	//保存进数据库
	return mysql.InsertUser(user)
}
func Login(p *models.ParamLogin) (token string, err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	//  mysql.Login 此处传入user指针，在指针查询时有对user进行一个“完整”赋值，
	//  即可以获得user的ID
	if err = mysql.Login(user); err != nil {
		return "", err
	}
	//生成JWT
	token, err = jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return
	}
	user.Token = token
	return
}
