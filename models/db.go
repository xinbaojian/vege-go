package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	fmt.Println("init database.....")
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:dove1989@tcp(39.106.47.1:3306)/crm_test?charset=utf8")
	// 需要在init中注册定义的model
	orm.RegisterModel(new(VegeProduct), new(VegeCategory))
}
