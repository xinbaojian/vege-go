package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strconv"
	"time"
	"vege-go/models"
	_ "vege-go/routers"
)

func main() {
	SyscData()
	go Timer()
	beego.Run()
}

func Timer() {
	duration, err := beego.AppConfig.Int64("sysc_duration")
	if err != nil {
		logs.Error("f", err)
		duration = 3 * 60 * 60
	}
	ticker := time.NewTicker(time.Duration(duration) * time.Second)
	for _ = range ticker.C {
		SyscData()
	}
}

func SyscData() {
	logs.Info("开始同步数据.............")

	sysCompanyIdStr := beego.AppConfig.String("sys_company_id")
	sysCompanyId, err := strconv.Atoi(sysCompanyIdStr)
	if err != nil {
		logs.Error("菜场标识不能为空")
		panic("菜场标识不能为空")
	}
	defer func() {
		if err := recover(); err != nil {
			models.UpdateStatus(sysCompanyId, 1)
		}
	}()
	models.UpdateStatus(sysCompanyId, 0)
	pojos := models.ParseJson()
	for _, pojo := range pojos {
		//categroy
		var category models.VegeCategory
		category.CategoryName = pojo.CategoryName
		category.SysCompanyId = sysCompanyId
		id, err := models.SaveOrUpdateCategory(&category)
		logs.Info("保存或更新分类信息;", category, err)
		if err == nil {
			//product
			product := models.ConvertType(pojo)
			product.SysCompanyId = sysCompanyId
			product.CategoryId = id
			product.Status = 1
			models.SaveOrUpdateProduct(&product)
			logs.Info("保存或更新菜品数据", product)
		} else {
			logs.Error("保存分类信息出错了", err)
		}
	}
	logs.Info("同步数据结束.....")
}
