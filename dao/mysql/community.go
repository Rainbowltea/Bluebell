package mysql

import (
	"bluebell/models"
	"database/sql"

	"go.uber.org/zap"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlSt := `select community_id,community_name from community`
	if err := db.Select(&communityList, sqlSt); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}
	return communityList, err
}

func GetCommunityDetailByID(id int64) (community *models.CommunityDetail, err error) {
	//分配内存
	community = new(models.CommunityDetail)
	sqlSt := `select 
	community_id, community_name, introduction, create_time
	from community 
	where community_id = ?`
	if err := db.Get(community, sqlSt, id); err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInvalidID
		}
	}
	return
}
