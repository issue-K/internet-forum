package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/jwt"
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

func Login(p *models.ParamLogin)(token string,err error){
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	//传递的是指针
	if err := mysql.Login(user); err != nil{
		return "",err
	}

	// 于是可以拿到user.UserID去生成jwt了
	return jwt.GenToken(user.UserID,user.Username)
}