package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
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
	PromotionFlag  string    //`json:promotion_flag`
	PromotionPrice float64   `json:promotion_price`
	PromoTimeFrom  time.Time //`json:promo_time_from`
	PromoTimeTo    time.Time //`json:promo_time_to`
	TraceCode      string    `json:trace_code`
	Photo          string
	MobileDesc     string
	PcDesc         string
	SysCompanyId   int
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

/**
 * 根据商品名称和菜场标识获取商品
 * @param {[type]} name         string [description]
 * @param {[type]} sysCompanyId int)   (product      VegeProduct [description]
 */
func GetVegeProduct(name string, sysCompanyId int) (product VegeProduct) {
	o := orm.NewOrm()
	// product := new(VegeProduct)
	err := o.QueryTable(product).Filter("Name", name).Filter("SysCompanyId", sysCompanyId).One(&product)
	if err != nil {
		fmt.Println("get product info error", err)
	}
	return product
}

func GetRemoteData() (result string) {
	resp, err := http.Get("http://at.itsmore.com/goods/")
	if err != nil {
		fmt.Println("获取商品信息出错了")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("商品信息出错了", err)
	}
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
	return string(body)
}

func ParseJson() (products []VegeProduct) {
	content := GetRemoteData()
	// fmt.Println(content)
	err := json.Unmarshal([]byte(content), &products)
	if err != nil {
		fmt.Println("error", err)
	}
	var product VegeProduct
	for i := 0; i < len(products); i++ {
		product = products[i]
		product.SellerName = product.ShopName
		// fmt.Println(product)
	}
	// fmt.Printf("%+v", products)
	return products
}
