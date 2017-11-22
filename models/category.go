package models

import (
	"github.com/astaxie/beego/logs"
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
		logs.Error("get Category error", err)
	}
	return category
}

func SaveOrUpdateCategory(category *VegeCategory) (id int, err error) {
	o := orm.NewOrm()
	newCategory, err := GetByNameAndCompanyId(category.CategoryName, category.SysCompanyId)
	if err != nil {
		category.CreateDate = time.Now()
		// fmt.Println("insert:", category)
		id, err := o.Insert(category)
		return int(id), err
	}
	newCategory.ModifyDate = time.Now()
	_, err = o.Update(&newCategory)
	return newCategory.Id, err

}

func GetByNameAndCompanyId(categoryName string, sysCompanyId int) (category VegeCategory, err error) {
	o := orm.NewOrm()
	category = VegeCategory{CategoryName: categoryName, SysCompanyId: sysCompanyId}
	err = o.Read(&category, "CategoryName", "SysCompanyId")
	return category, err
}
