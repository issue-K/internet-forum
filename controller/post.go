package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"strconv"
	"web_app/logic"
	"web_app/models"
	"web_app/pkg/snowflake"
)

// PostCommitHandler 创建帖子的处理函数
func PostCommitHandler(c *gin.Context){
	//获取参数
	postCommit := new( models.ParamPostCommit )
	if err := c.ShouldBindJSON(postCommit); err != nil{
		zap.L().Error("post.PostCommitHandler() failed",zap.Error(err) )
		ResponseErrorWithMsg(c,CodeInvalidParam,CodeInvalidParam.Msg())
		return
	}

	userid,_ := c.Get("userID")
	postCommit.UserID = userid.(int64)
	postCommit.PostID = snowflake.GenID()
	//业务层处理
	if err := logic.PostCommit( postCommit ); err != nil{
		zap.L().Error("post.PostCommitHandler() failed",zap.Error(err) )
		ResponseError(c,CodeServerBusy  )
	}
	//返回数据
	ResponseSuccess(c,nil )
}

// GetPostDetailHandler 获取帖子详情的处理函数
func GetPostDetailHandler(c *gin.Context){
	//1.获取参数(得到帖子的id)
	pidstr := c.Param("id")
	pid,err := strconv.ParseInt(pidstr,10,64)
	if err != nil{
		zap.L().Error("GetPostDetailHandler error,参数异常无法转型",zap.Error(err) )
		ResponseError(c,CodeInvalidParam)
		return
	}

	//2.logic层,根据id取出帖子
	data,err := logic.GetPostById(pid)
	if err != nil{
		zap.L().Error("GetPostDetailHandler logic error,",zap.Error(err))
		ResponseError(c,CodeInvalidParam)
		return
	}
	//3.返回结果
	ResponseSuccess(c,data)
}

func PostListHandler(c *gin.Context){
	order, _ := c.GetQuery("order")
	pageStr,ok := c.GetQuery("page")
	if !ok{
		pageStr = "1"
	}
	pageNum, err := strconv.ParseInt(pageStr,10,64)
	if err != nil{
		pageNum = 1
	}

	log.Printf("pageNum = %v order = %v\n",pageNum,order )
	posts,err := logic.GetPostList(pageNum,order )
	if err != nil{
		zap.L().Error("PostListHandler failed in logic.GetPostList",zap.Error(err) )
		ResponseError(c,CodeServerBusy)
		return
	}
	ResponseSuccess(c,posts)
}