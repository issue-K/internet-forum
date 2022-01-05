package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/snowflake"
)

//
func SignUp(p *models.ParamSignUp) error {
	//判断用户不存在
	if err := mysql.CheckUserExist(p.Username);err != nil{
		return err
	}
	//构造user对象
	userID := snowflake.GenID()
	u := &models.User{
		UserID: userID,
		Username: p.Username,
		Password: p.Password,
	}
	//保存进数据库
	return mysql.InsertUser( u )
}

func Login(p *models.ParamLogin)(err error){
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	return mysql.Login(user)
}