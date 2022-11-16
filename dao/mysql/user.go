package mysql

import (
	"crypto/md5"
	sql2 "database/sql"
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

func Login(user *models.User) (err error) {
	sqlStr := "select user_id, username, password from user where username = ?;"
	res := models.User{}
	err = db.Get(&res, sqlStr, user.Username)
	if err == sql2.ErrNoRows {
		return errors.New("用户不存在")
	}
	if err != nil {
		return errors.New("数据库查询失败")
	}
	if res.Password != Md5Psw(user.Password) {
		return errors.New("用户密码错误")
	}
	return

}
