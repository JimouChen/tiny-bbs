package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"tiny-bbs/service"
)

// CommunityController 获取社区id和社区name
func CommunityController(ctx *gin.Context) {
	data, err := service.GetCommunityList()
	if err != nil {
		zap.L().Error("service.GetCommunityList() get community list failed...", zap.Error(err))
		ResponseErr(ctx, CodeServerBusy)
		return
	}
	ResponseSuccess(ctx, data)
}

// CommIntroByIdController 通过community_id获取详情
func CommIntroByIdController(ctx *gin.Context) {
	communityId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		zap.L().Error("param error", zap.Error(err))
		ResponseErr(ctx, CodeParamTypeErr)
		return
	}
	introduction, err := service.GetIntroductionById(communityId)
	if err != nil {
		zap.L().Error("service.CommIntroByIdController() get community detail failed...", zap.Error(err))
		ResponseErr(ctx, CodeServerBusy)
		return
	}
	ResponseSuccess(ctx, introduction)
}
