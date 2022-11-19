package service

import (
	"go.uber.org/zap"
	"tiny-bbs/dao/mysql"
	"tiny-bbs/models"
	"tiny-bbs/pkg/snowflake"
)

func CreatePost(p *models.PostParam) (err error) {
	// 生成id
	p.ID = snowflake.GenID()
	// 保存到数据库,返回
	return mysql.CreatePost(p)
}

func GetPostMsgById(id int64) (data *models.PostApiDetail, err error) {
	data = new(models.PostApiDetail)
	// 获取帖子信息
	postMsg, err := mysql.GetPostMsgById(id)
	if err != nil {
		zap.L().Error("get post msg failed",
			zap.Int64("post id", id),
			zap.Error(err))
		return
	}
	// 获取作者名字
	authorId := postMsg.AuthorID
	authorMsg, err := mysql.GetUserById(authorId)
	if err != nil {
		zap.L().Error("get author name failed.",
			zap.Int64("author id", authorId),
			zap.Error(err))
		return
	}
	// 获取社区信息
	communityMsg, err := GetIntroductionById(postMsg.CommunityID)
	if err != nil {
		zap.L().Error("grt community msg failed",
			zap.Int64("community id", postMsg.CommunityID),
			zap.Error(err))
		return
	}
	//data = &models.PostApiDetail{
	//	AuthorName:      authorMsg.Username,
	//	PostParam:       postMsg,
	//	CommunityDetail: communityMsg,
	//}
	data.AuthorName = authorMsg.Username
	data.PostParam = postMsg
	data.CommunityDetail = communityMsg
	return
}

func GetPostMsgList(page, size int64) (data []*models.PostApiDetail, err error) {
	posts, err := mysql.GetPostMsgList(page, size)
	if err != nil {
		zap.L().Error("mysql.GetPostMsgList() failed", zap.Error(err))
		return nil, mysql.ErrServerBusy
	}
	data = make([]*models.PostApiDetail, 0, len(posts))

	for _, post := range posts {
		// 获取帖子信息
		postMsg, err := mysql.GetPostMsgById(post.ID)
		if err != nil {
			zap.L().Error("get post msg failed",
				zap.Int64("post id", post.ID),
				zap.Error(err))
			continue
		}
		// 获取作者名字
		authorId := postMsg.AuthorID
		authorMsg, err := mysql.GetUserById(authorId)
		if err != nil {
			zap.L().Error("get author name failed.",
				zap.Int64("author id", authorId),
				zap.Error(err))
			continue
		}
		// 获取社区信息
		communityMsg, err := GetIntroductionById(postMsg.CommunityID)
		if err != nil {
			zap.L().Error("grt community msg failed",
				zap.Int64("community id", postMsg.CommunityID),
				zap.Error(err))
			continue
		}
		d := &models.PostApiDetail{
			AuthorName:      authorMsg.Username,
			PostParam:       postMsg,
			CommunityDetail: communityMsg,
		}
		data = append(data, d)
	}

	return
}
