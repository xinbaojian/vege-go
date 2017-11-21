package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type VegeCategory struct {
	Id           int
	CreateDate   time.Time
	ModifyDate   time.Time
	Creator      string
	Modifier     string
	ParentId     int
	CategoryName string
	isValid      int
	SysCompanyId int
	index        int
}

/**
 * 根据分类名称和菜场标识获取分类
 * @param {[type]} categoryName string [description]
 * @param {[type]} sysCompanyId int)   (category     VegeCategory [description]
 */
func GetCategoryByNameAndCompanyId(categoryName string, sysCompanyId int) (category VegeCategory) {
	o := orm.NewOrm()
	// category := new(VegeCategory)
	qs := o.QueryTable(category)
	err := qs.Filter("CategoryName", categoryName).Filter("SysCompanyId", sysCompanyId).One(&category)
	if err != nil {
		fmt.Println("get Category error", err)
	}
	fmt.Println(category)
	return category
}
