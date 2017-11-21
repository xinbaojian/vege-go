package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"vege-go/models"
	_ "vege-go/routers"
)

func main() {
	fmt.Println("解析商品数据")
	models.ParseJson()
	models.GetCategoryByNameAndCompanyId("海鲜", 2)
	beego.Run()
}
