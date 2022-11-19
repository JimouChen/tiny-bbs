package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"tiny-bbs/models"
	"tiny-bbs/service"
)

func PostMsgController(ctx *gin.Context) {
	// 参数获取和校验
	p := new(models.PostParam)
	if err := ctx.ShouldBindJSON(p); err != nil {
		ResponseErr(ctx, CodeInvalidParams)
		return
	}
	authorID, err := GetCurrentUserId(ctx)
	if err != nil {
		ResponseErr(ctx, CodeUserNotLogin)
		return
	}
	p.AuthorID = authorID

	// 发布/创建帖子
	if err := service.CreatePost(p); err != nil {
		zap.L().Error("service.CreatePost(p) failed.", zap.Error(err))
		ResponseErr(ctx, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(ctx, nil)
}

func GetPostMsgByIdController(ctx *gin.Context) {
	postId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		zap.L().Error("param error", zap.Error(err))
		ResponseErr(ctx, CodeParamTypeErr)
	}
	data, err := service.GetPostMsgById(postId)
	if err != nil {
		zap.L().Error("service.GetPostMsgById err...", zap.Error(err))
		ResponseErr(ctx, CodeServerBusy)
		return
	}
	ResponseSuccess(ctx, data)
}

func GetPostMsgListController(ctx *gin.Context) {
	page, err := strconv.ParseInt(ctx.Query("page"), 10, 64)
	if err != nil {
		page = 1
		zap.L().Error("get page failed", zap.Error(err))
	}
	size, err := strconv.ParseInt(ctx.Query("size"), 10, 64)
	if err != nil {
		size = 10
		zap.L().Error("get page size failed", zap.Error(err))
	}
	data, err := service.GetPostMsgList(page, size)
	if err != nil {
		ResponseErr(ctx, CodeServerBusy)
		return
	}
	ResponseSuccess(ctx, data)
}
