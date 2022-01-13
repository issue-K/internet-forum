package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"web_app/logic"
)

func CommunityHandler(c *gin.Context){
	//logic层查询所有的社区(community_id,community_name)
	data, err := logic.GetCommunityList()
	if err != nil{
		zap.L().Error("logic.GetCommunityList() failed",zap.Error(err) )
		ResponseError(c,CodeServerBusy)
		return
	}
	ResponseSuccess(c,data)
}

func CommunityDetailHandler(c *gin.Context){
	communityID := c.Param("id")
	id,err := strconv.ParseInt( communityID,10,64 )
	if err != nil{
		ResponseError(c,CodeInvalidParam)
		return
	}
	//logic层查询所有的社区(community_id,community_name)
	data, err := logic.GetCommunityDetail(id)
	if err != nil{
		zap.L().Error("logic.GetCommunityList() failed",zap.Error(err) )
		ResponseError(c,CodeServerBusy)
		return
	}
	ResponseSuccess(c,data)
}