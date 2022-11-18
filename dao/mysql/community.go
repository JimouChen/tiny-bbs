package mysql

import (
	sql2 "database/sql"
	"go.uber.org/zap"
	"tiny-bbs/models"
)

func GetCommunityList() (datalist []*models.CommunityList, err error) {
	sql := "select community_id, community_name from community;"
	if err = db.Select(&datalist, sql); err != nil {
		if err == sql2.ErrNoRows {
			zap.L().Warn("community list is empty")
			err = nil
			return nil, ErrUserNotExist
		} else {
			err = ErrServerBusy
			return
		}
	}
	return
}

func GetIntroductionById(id int64) (*models.CommunityDetail, error) {
	sql := "select community_id, community_name, introduction, create_time from community where community_id = ?;"
	introduction := new(models.CommunityDetail)
	err := db.Get(introduction, sql, id)
	if err != nil {
		zap.L().Warn("community list is empty")
		err = ErrServerBusy
		return nil, err
	}
	return introduction, err
}
