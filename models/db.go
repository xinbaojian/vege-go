package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {

	db_type := beego.AppConfig.String("db_type")
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_pass")
	db_name := beego.AppConfig.String("db_name")

	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", db_type, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", db_user, db_pass, db_host, db_port, db_name))
	// 需要在init中注册定义的model
	orm.RegisterModel(new(VegeProduct), new(VegeCategory))
	orm.Debug = true

	//setting logs
	logs.SetLogger(logs.AdapterFile, `{"filename":"vege-go.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
}
