package models

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
	"time"
)

type VegeProduct struct {
	Id             int
	CategoryId     int       `json:category_id`
	CategoryName   string    `json:caregory_name`
	Number         string    `json:Number`
	Name           string    `json:Name`
	Barcode        string    `json:barcode`
	Unit           string    `json:Unit`
	Price          float64   `json:Expr1`
	AllowWeigh     string    `json:allow_weigh`
	ShopId         int       `json:shop_id`
	ShopName       string    `json:shop_name`
	SellerId       string    `json:seller_id`
	SellerName     string    `json:seller_name`
	CreateDate     time.Time `json:creation_time`
	UpdateDate     time.Time `json:update_time`
	PlaceOfOrigin  string    `json:place_of_origin`
	Specification  string    `json:specification`
	Grade          string    `json:grade`
	PromotionFlag  string    `json:promotion_flag`
	PromotionPrice float64   `json:promotion_price`
	PromoTimeFrom  time.Time `json:promo_time_from`
	PromoTimeTo    time.Time `json:promo_time_to`
	TraceCode      string    `json:trace_code`
	Photo          string
	MobileDesc     string
	PcDesc         string
	SysCompanyId   int
}

func init() {
	fmt.Println("init database.....")
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:dove1989@tcp(39.106.47.1:3306)/crm_test?charset=utf8")
	// 需要在init中注册定义的model
	orm.RegisterModel(new(VegeProduct))
}

func GetVegeProductList() {
	o := orm.NewOrm()
	var products []*VegeProduct
	num, err := o.QueryTable("vege_product").All(&products)
	fmt.Println(num, err)
	for i := 0; i < len(products); i++ {
		fmt.Println(products[i].Id)
	}

}

func GetRemoteData() (result string) {
	resp, err := http.Get("http://at.itsmore.com/goods/")
	if err != nil {
		fmt.Println("获取商品信息出错了")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("获取商品信息出错了")
	}
	return string(body)
}

func ParseJson() (products []*VegeProduct) {
	content := GetRemoteData()
	json.Unmarshal([]byte(content), products)
	return products
}
