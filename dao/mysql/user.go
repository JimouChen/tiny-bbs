package mysql

import (
	"tiny-bbs/models"
)

func CheckUserIsExist(username string) (bool, error) {
	sql := "select count(user_id) from user where username = ?;"
	var cnt int
	if err := db.Get(&cnt, sql, username); err != nil {
		return false, err
	}
	return cnt > 0, nil
}

func InsertUser(user *models.User) (err error) {
	//加密密码

	//写入数据库
	sql := "insert user(user_id, username, password) value(?, ?, ?);"
	_, err = db.Exec(sql, user.UserId, user.Username, user.Password)
	if err != nil {
		return err
	}
	return
}
