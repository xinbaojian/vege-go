package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"io/ioutil"
	"net/http"
	"time"
)

type VegeProductPojo struct {
	Id             int
	CategoryId     int       `json:"category_id"`
	CategoryName   string    `json:"caregory_name"`
	Number         string    `json:"Number"`
	Name           string    `json:"Name"`
	Barcode        string    `json:"barcode"`
	Unit           string    `json:"Unit"`
	Expr           float64   `json:"Expr1"`
	Price          float64   `json:"price"`
	AllowWeighBool bool      `json:"allow_weigh"`
	ShopId         int       `json:"shop_id"`
	ShopName       string    `json:"shop_name"`
	SellerId       int       `json:"seller_id"`
	SellerName     string    `json:"seller_name"`
	CreateDate     string    `json:"creation_time"`
	UpdateDate     string    `json:"update_time"`
	PlaceOfOrigin  string    `json:"place_of_origin"`
	Specification  string    `json:"specification"`
	Grade          string    `json:"grade"`
	PromotionFlag  string    //`json:"promotion_flag"`
	PromotionPrice string    `json:"promotion_price"`
	PromoTimeFrom  time.Time //`json:"promo_time_from"`
	PromoTimeTo    time.Time //`json:"promo_time_to"`
	TraceCode      string    `json:"trace_code"`
	Photo          string
	MobileDesc     string
	PcDesc         string
	SysCompanyId   int
}

type VegeProduct struct {
	Id             int
	CategoryId     int
	CategoryName   string
	Number         string
	Name           string
	Barcode        string
	Unit           string
	Price          float64
	AllowWeigh     string
	ShopId         int
	ShopName       string
	SellerId       int
	SellerName     string
	CreateDate     time.Time
	UpdateDate     time.Time
	PlaceOfOrigin  string
	Specification  string
	Grade          string
	PromotionFlag  string
	PromotionPrice string
	PromoTimeFrom  time.Time
	PromoTimeTo    time.Time
	TraceCode      string
	Photo          string
	MobileDesc     string
	PcDesc         string
	SysCompanyId   int
	Status         int
}

func GetVegeProductList() {
	o := orm.NewOrm()
	var products []*VegeProduct
	o.QueryTable("vege_product").All(&products)
	for i := 0; i < len(products); i++ {
		fmt.Println(products[i].Id)
	}
}

/**
 * 根据商品名称和菜场标识获取商品
 * @param {[type]} name         string [description]
 * @param {[type]} sysCompanyId int)   (product      VegeProduct [description]
 */
func GetVegeProduct(name string, sysCompanyId int) (product VegeProduct, err error) {
	o := orm.NewOrm()
	// product := new(VegeProduct)
	err = o.QueryTable(product).Filter("Name", name).Filter("SysCompanyId", sysCompanyId).One(&product)
	if err != nil {
		logs.Error("get product info error", err)
	}
	return product, err
}

func GetRemoteData() (result string) {
	resp, err := http.Get(beego.AppConfig.String("sysc_goods_url"))
	if err != nil {
		logs.Error("获取商品信息出错了", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error("商品信息出错了", err)
	}
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
	logs.Info("json data:%s", string(body))
	return string(body)
}

func ParseJson() (products []VegeProductPojo) {
	content := GetRemoteData()
	err := json.Unmarshal([]byte(content), &products)
	if err != nil {
		logs.Error("Parse json Error", err)
	}
	return products
}

func SaveOrUpdateProduct(product *VegeProduct) {
	o := orm.NewOrm()
	dbProduct, err := GetVegeProduct(product.Name, product.SysCompanyId)
	if err == orm.ErrNoRows {
		product.CreateDate = time.Now()
		id, err := o.Insert(product)
		logs.Info("Insert Product", "Id:", id, "Error:", err, "Product:", product)
	} else if err == nil {
		product.Id = dbProduct.Id
		product.PromotionFlag = dbProduct.PromotionFlag
		product.PromotionPrice = dbProduct.PromotionPrice
		product.PromoTimeFrom = dbProduct.PromoTimeFrom
		product.PromoTimeTo = dbProduct.PromoTimeTo
		product.Photo = dbProduct.Photo
		o.Update(product)
		logs.Info("Update Product", product)
	}
	fmt.Println(dbProduct)

}

func ConvertType(pojo VegeProductPojo) (product VegeProduct) {
	product.CategoryName = pojo.CategoryName
	product.Number = pojo.Number
	product.Name = pojo.Name
	product.Barcode = pojo.Barcode
	product.Unit = pojo.Unit
	product.Price = pojo.Expr
	product.Price = pojo.Price
	if pojo.AllowWeighBool {
		product.AllowWeigh = "true"
	} else {
		product.AllowWeigh = "false"
	}
	product.ShopId = pojo.ShopId
	product.ShopName = pojo.ShopName
	product.SellerId = pojo.SellerId
	var sellName = beego.AppConfig.String("seller_name")
	if sellName != nil && sellName != "" {
		product.SellerName = sellName
	} else {
		product.SellerName = pojo.ShopName
	}

	product.CreateDate = time.Now()
	product.UpdateDate = time.Now()
	product.PlaceOfOrigin = pojo.PlaceOfOrigin
	product.Specification = pojo.Specification
	product.Grade = pojo.Grade
	// product.PromotionFlag = pojo.PromotionFlag
	// product.PromotionPrice = pojo.PromotionPrice
	// product.PromoTimeFrom = pojo.PromoTimeFrom
	// product.PromoTimeTo = pojo.PromoTimeTo
	product.TraceCode = pojo.TraceCode
	return product
}

func UpdateStatus(sysCompanyId int, status int) {
	o := orm.NewOrm()
	res, err := o.Raw("UPDATE vege_product SET status = ? WHERE sys_company_id  = ?", status, sysCompanyId).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		logs.Info("update product status succefully;mysql row affected nums: ", num)
	}

}
