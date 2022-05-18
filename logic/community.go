package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
)

// GetCommunityList 获取所有社区的集合
func GetCommunityList() ([]*models.Community, error) {
	return mysql.GetCommunityList()
}

// GetCommunityDetail 返回某 一个社区的详细信息
func GetCommunityDetail(id int64) (community *models.CommunityDetail, err error) {
	return mysql.GetCommunityDetailByID(id)
}
