package mysql

import (
	"database/sql"
	"go.uber.org/zap"
	"web_app/models"
)

func GetCommunityList() (communityList []*models.Community,err error){
	sqlStr := "select community_id,community_name from community"

	if err := db.Select(&communityList,sqlStr); err != nil{
		if err == sql.ErrNoRows{
			zap.L().Warn("there is not community in db")
			err = nil
		}
	}
	return
}

func GetCommunityDetailByID(id int64)(community *models.CommunityDetail,err error){
	community = new(models.CommunityDetail)
	sqlStr := `select community_id, community_name, introduction, create_time
	from community
	where community_id = ?`
	if err=db.Get(community, sqlStr, id); err != nil{
		if err == sql.ErrNoRows{
			err = ErrorInvalidID
		}
		return
	}
	return
}
