package mysql

import (
	"bluebell/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	_ "errors"
	_ "fmt"
)

const secret = "codecffee.xyz"

// InsertUser 向数据库中插入一条新的用户信息
func InsertUser(user *models.User) (err error) {
	//对用户密码进行一个加密
	user.Password = encryptPassword(user.Password)
	//fmt.Println(user.Password)
	sqlStr := `insert into user(user_id, username, password) values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)

	return
}

// CheckUserExist
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int64
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

// encryptPassword 密码加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))

	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

//Login 用户登录
func Login(user *models.User) (err error) {
	oPassword := user.Password
	sqlStr := `select user_id ,username, password from user where username=?`
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		return err
	}
	if user.Password != encryptPassword(oPassword) {
		return ErrorInvalidPassword
	}
	return
}
