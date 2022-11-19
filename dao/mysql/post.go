package mysql

import "tiny-bbs/models"

func CreatePost(param *models.PostParam) (err error) {
	sql := `insert into post(
			post_id, title, content, author_id, community_id) 
			value(?,?,?,?,?);`
	_, err = db.Exec(sql, param.ID, param.Title, param.Content, param.AuthorID, param.CommunityID)
	if err != nil {
		return err
	}
	return
}

func GetPostMsgById(id int64) (data *models.PostParam, err error) {
	data = new(models.PostParam)
	sql := `select
			post_id, title, content, author_id, community_id, create_time
			from post where post_id = ?;`
	err = db.Get(data, sql, id)
	if err != nil {
		return nil, ErrServerBusy
	}
	return
}
