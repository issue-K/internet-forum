package logic

import (
	"go.uber.org/zap"
	"web_app/dao/mysql"
	"web_app/models"
)

func PostCommit(commit *models.ParamPostCommit) (err error){
	return mysql.InsertPost(commit)
}

func GetPostById(pid int64)(data *models.ApiPostDetail,err error){
	data = new( models.ApiPostDetail )
	post,err := mysql.GetPostByID( pid )
	if err != nil{
		zap.L().Error("mysql.GetPostById.GetPostByID failed",zap.Error(err) )
		return
	}
	user,err := mysql.GetUserById( post.AuthorId )
	if err != nil{
		zap.L().Error("mysql.GetPostById.GetUserById failed",zap.Error(err) )
		return
	}
	community,err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetPostById.GetUserById failed",zap.Error(err) )
		return
	}
	data.AuthorName = user.Username
	data.CommunityName = community.CommunityName
	data.Post = post
	return
}

func GetPostList(pageNum int64,order string )( []*models.ParamPostCommit, error ){
	//
	var limit,start int64
	limit = 8
	start = limit*(pageNum-1)
	return mysql.GetPostList(start,limit,order)
}