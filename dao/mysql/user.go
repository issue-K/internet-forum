package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"web_app/models"
)
const secret = "cl"

func CheckUserExist(username string) (error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	if err := db.Get(&count,sqlStr,username ); err!=nil{
		return err
	}
	if count>0{
		return errors.New("用户已存在")
	}
	return nil
}

func InsertUser(user *models.User) (err error ) {
	//对密码进行加密
	user.Password = encryptPassword( user.Password )
	//执行SQL语句入库
	sqlStr := `insert into user(user_id,username,password) values(?,?,?)`
	_,err = db.Exec(sqlStr,user.UserID,user.Username,user.Password)
	return
}

func encryptPassword(opassword string) string{
	h := md5.New()
	h.Write([]byte(secret) )
	return hex.EncodeToString( h.Sum( []byte(opassword) ) )
}

func Login(user *models.User)( err error ){
	password := encryptPassword( user.Password )
	sqlStr := `select user_id,username,password from user where username=?`
	err = db.Get(user,sqlStr,user.Username)
	if err == sql.ErrNoRows{
		return errors.New("用户不存在")
	}
	if err != nil{ //查询数据库失败
		return err
	}

	if password != user.Password{ //密码错误
		return errors.New("密码错误")
	}

	return
}

func GetUserById(id int64)(user *models.User,err error){
	user = new(models.User)
	sqlStr := `select user_id,username from user where user_id = ?`
	err = db.Get(user,sqlStr,id)
	return user,err
}