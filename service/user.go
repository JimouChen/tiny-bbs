package service

import (
	"fmt"
	"tiny-bbs/dao/mysql"
	"tiny-bbs/models"
	"tiny-bbs/pkg/jwt"
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

func Login(user *models.ParmaLogin) (token string, uid string, err error) {
	u := &models.User{
		Username: user.Username,
		Password: user.Password,
	}
	if err := mysql.Login(u); err != nil {
		return "", "", err
	}
	token, err = jwt.GenToken(u.UserId, u.Username)
	if err != nil {
		return "", fmt.Sprintf("%d", u.UserId), err
	}
	return token, fmt.Sprintf("%d", u.UserId), err
}
