package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
)

func CreatePost(p *models.Post) (err error) {
	//1.为帖子生成postid
	p.ID = int64(snowflake.GenID())
	//2.保存到数据库
	return mysql.CreatePost(p)
}
