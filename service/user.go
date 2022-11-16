package service

import (
	"tiny-bbs/dao/mysql"
	"tiny-bbs/models"
	"tiny-bbs/pkg/snowflake"
)

func SignUp(user *models.ParmaRegister) (err error) {
	//判断用户存不存在
	err = mysql.CheckUserIsExist(user.Username)
	if err != nil {
		return err
	}
	//生成uid
	userId := snowflake.GenID()
	u := &models.User{
		UserId:   userId,
		Username: user.Username,
		Password: user.Password,
	}
	//保存入库
	return mysql.InsertUser(u)
}
