package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)
var(
	db *sqlx.DB
)
func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"))
	db = sqlx.MustConnect("mysql", dsn)  //若出现错误,会panic

	db.SetMaxOpenConns(viper.GetInt("mysql.max_open_conns"))  //设置最大连接数
	db.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns")) //最大闲置连接数
	return
}
func Close(){
	db.Close()
}