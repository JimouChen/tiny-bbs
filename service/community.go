package service

import (
	"tiny-bbs/dao/mysql"
	"tiny-bbs/models"
)

func GetCommunityList() ([]*models.CommunityList, error) {
	return mysql.GetCommunityList()
}

func GetIntroductionById(id int64) (introduction *models.CommunityDetail, err error) {
	return mysql.GetIntroductionById(id)
}
