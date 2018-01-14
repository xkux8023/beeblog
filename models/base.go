package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func RegisterDB() {
	//获取服务器数据库相关信息
	dbhost := beego.AppConfig.String("db_host")
	dbport := beego.AppConfig.String("db_port")
	dbuser := beego.AppConfig.String("db_user")
	dbpassword := beego.AppConfig.String("db_pass")
	dbname := beego.AppConfig.String("db_name")
	dbIdle, _ := beego.AppConfig.Int("db_max_idle_conn")
	dbConn, _ := beego.AppConfig.Int("db_max_open_conn")

	dburl := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"

	orm.RegisterModel(new(Category), new(Topic), new(Comment))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dburl, dbIdle, dbConn)

}
