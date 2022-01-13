package mysql

import (
	"log"
	"web_app/models"
)

func InsertPost(commit *models.ParamPostCommit) (err error){
	sqlStr := `insert into post (post_id,title,content,author_id,community_id) values(?,?,?,?,?)`
	log.Println("post_id = ",commit.PostID)
	_,err = db.Exec(sqlStr,commit.PostID,commit.Title,commit.Content,commit.UserID,commit.Community_id)
	return err
}

func GetPostByID(pid int64) (post *models.Post,err error){
	post = new(models.Post)
	sqlStr := `select post_id,title,content,author_id,community_id from post where post_id = ?`
	err = db.Get(post,sqlStr,pid )
	return
}
func GetPostList(start ,limit int64,order string)( ans []*models.ParamPostCommit,err error){
	if order == "time" {
		order = "create_time"
	}else{
		order = "id"
	}


	ans = make( []*models.ParamPostCommit,0,limit+1 )
	sqlStr := `select post_id,title,content from post order by ? limit ?,?`
	err = db.Select(&ans,sqlStr,order,start,limit )
	return
}