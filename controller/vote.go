package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strconv"
	"web_app/dao/redis"
	"web_app/models"
)


func POSTVoteController(c *gin.Context){
	/* 参数校验
	userid  [在请求头中]
	PostID int64 (帖子id)
	Direction int (赞成还是反对票)
	 */
	p := new( models.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil{
		errs,ok := err.(validator.ValidationErrors)
		if !ok{  //不是因为bind产生的错误
			ResponseError(c,CodeServerBusy )
			return
		}else{
			errData := removeTopStruct( errs.Translate(trans) )
			ResponseErrorWithMsg(c,CodeInvalidParam,errData)
			return
		}
	}
	//获取当前请求的用户id
	userID,err := GetCurrentUser(c)
	if err != nil{
		ResponseError( c, CodeNeedLogin )
		return
	}
	useridstring := strconv.FormatInt(userID,10)
	redis.PostVote(useridstring,p)
	ResponseSuccess(c,nil)
}