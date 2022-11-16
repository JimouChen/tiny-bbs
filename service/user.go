package service

import (
	"errors"
	"tiny-bbs/dao/mysql"
	"tiny-bbs/models"
	"tiny-bbs/pkg/snowflake"
)

func SignUp(user *models.ParmaRegister) (err error) {
	//判断用户存不存在
	isExist, err := mysql.CheckUserIsExist(user.Username)
	if err != nil {
		return err
	}
	if isExist {
		return errors.New("注册失败，用户已经存在")
	}
	//生成uid
	userId := snowflake.GenID()
	u := &models.User{
		UserId:   userId,
		Username: user.Username,
		Password: user.Password,
	}
	//保存入库
	mysql.InsertUser(u)
	return nil
}
