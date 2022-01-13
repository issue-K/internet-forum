package models

//定义请求的参数结构体

type ParamSignUp struct{
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type ParamLogin struct{
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamPostCommit struct{
	Title string `json:"title" binding:"required" db:"title"`
	Content string `json:"content" binding:"required" db:"content"`
	Community_id int64 `json:"community_id" binding:"required" db:"community_id"`
	UserID int64
	PostID int64 `db:"post_id"`
}

type ParamVoteData struct{
	PostID string `json:"post_id,string" binding:"required"`
	Direction int `json:"direction,string" binding:"required"`
}