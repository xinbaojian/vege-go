package models

import (
	"github.com/astaxie/beego/orm"
	"time"
	"fmt"
)

type VegeProduct struct {
	Id int
	categoryId int
	categoryName string
	number string
	name string
	barcode string
	unit string
	price float64
	allowWeigh string
	shopId int
	shopName string
	sellerId string
	sellerName string
	createDate time.Time
	updateDate time.Time
	placeOfOrigin string
	specification string
	grade string
	promotionFlag string
	promotionPrice float64
	promoTimeFrom time.Time
	promoTimeTo time.Time
	traceCode string
	photo string
	mobileDesc string
	pcDesc string
	sysCompanyId int

}

func init() {
	fmt.Println("init database.....")
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:root@tcp(39.106.47.1:3306)/crm_test?charset=utf8")
	// 需要在init中注册定义的model
	orm.RegisterModel(new(VegeProduct))
}
