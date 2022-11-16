package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"tiny-bbs/models"
)

const secretBase = "neayacom"

func CheckUserIsExist(username string) (err error) {
	sql := "select count(user_id) from user where username = ?;"
	var cnt int
	if err := db.Get(&cnt, sql, username); err != nil {
		return err
	}
	if cnt > 0 {
		return errors.New("用户已经存在！")
	}
	return
}

func InsertUser(user *models.User) (err error) {
	//加密密码
	user.Password = Md5Psw(user.Password)
	//写入数据库
	sql := "insert user(user_id, username, password) value(?, ?, ?);"
	_, err = db.Exec(sql, user.UserId, user.Username, user.Password)
	if err != nil {
		return err
	}
	return
}

func Md5Psw(password string) string {
	h := md5.New()
	h.Write([]byte(secretBase))
	return hex.EncodeToString(h.Sum([]byte(password)))
}
