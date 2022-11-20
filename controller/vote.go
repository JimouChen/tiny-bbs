package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"tiny-bbs/models"
	"tiny-bbs/service"
)

func PostVoteController(ctx *gin.Context) {
	// 参数校验
	p := new(models.ParamVoteData)
	if err := ctx.ShouldBindJSON(p); err != nil {
		e, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseErr(ctx, CodeInvalidParams)
			return
		}
		zap.L().Error("err.(validator.ValidationErrors) ", zap.Error(e))
		ResponseErr(ctx, CodeServerBusy)
		return
	}
	uid, err := GetCurrentUserId(ctx)
	if err != nil {
		zap.L().Error("GetCurrentUserId(ctx) failed", zap.Error(err))
		ResponseErr(ctx, CodeUserNotLogin)
		return
	}
	if err := service.VoteForPost(uid, p); err != nil {
		zap.L().Error("service.VoteForPost(uid, p) failed.", zap.Error(err))
		ResponseErr(ctx, CodeServerBusy)
		return
	}
	ResponseSuccess(ctx, nil)
}
