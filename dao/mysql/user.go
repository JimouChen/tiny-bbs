package mysql

import (
	"crypto/md5"
	sql2 "database/sql"
	"encoding/hex"
	"errors"
	"tiny-bbs/models"
)

const secretBase = "neayacom"

var (
	ErrUserNotExist    = errors.New("用户不存在")
	ErrUserExist       = errors.New("用户已经存在")
	ErrPswUName        = errors.New("用户名或密码输入错误")
	ErrInvalidPswUName = errors.New("用户名或密码不合法")
	ErrServerBusy      = errors.New("服务器繁忙")
)

func CheckUserIsExist(username string) (err error) {
	sql := "select count(user_id) from user where username = ?;"
	var cnt int
	if err := db.Get(&cnt, sql, username); err != nil {
		return ErrServerBusy
	}
	if cnt > 0 {
		return ErrUserExist
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
		return ErrServerBusy
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
	originPsw := user.Password
	err = db.Get(user, sqlStr, user.Username)
	if err == sql2.ErrNoRows {
		return ErrUserNotExist
	}
	if err != nil {
		return ErrServerBusy
	}
	if user.Password != Md5Psw(originPsw) {
		return ErrPswUName
	}
	return

}
