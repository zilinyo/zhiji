package apis

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	db "zhiji/api/database"
	models "zhiji/api/models"
)

const SecretKey = "hello world"
const UserCollection = "user"

func Test(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"code": "200", "data": "success"})

}

func Login(c *gin.Context) {

	var Check map[string]interface{}

	var Custom = models.CustomerLogin{}

	body, _ := ioutil.ReadAll(c.Request.Body)

	json.Unmarshal(body, &Check)

	password := Check["Password"]

	passwords := models.CheckPassword(password.(string))

	err := db.Eloquent.Where(map[string]interface{}{"login_name": Check["LoginName"], "password": passwords}).Find(&Custom).Error

	if err != nil {

		c.JSON(http.StatusOK, gin.H{"status": -1, "msg": "账号密码不对", "err": err})
		return

	}

	c.JSON(http.StatusOK, gin.H{"token": Custom.GetToken()})

}
func Register(c *gin.Context) {

	var register = new(models.CustomerLogin)

	if err := c.BindJSON(register); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"status": -1, "err": err.Error()})
		return

	}

	err := register.Register()

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return

	}

	c.JSON(http.StatusOK, gin.H{"code": "200", "data": "success"})

}

var counter int = 0

func Lock(c *gin.Context) {
	//lock := &sync.Mutex{}
	//for i := 0; i < 10; i++ {
	//	go count(lock)
	//}
	//for {
	//	lock.Lock()
	//	a := counter
	//	lock.Unlock()
	//	runtime.Gosched()
	//	if a >= 10 {
	//		fmt.Println(a)
	//
	//		break
	//	}
	//}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "sssssss",
	})
}

//type address struct {
//	Recipients string `json:"recipients" binding:"required"`
//	Province int `json:"province" binding:"required"`
//	City string `json:"city" binding:"required"`
//	Area string `json:"area" binding:"required"`
//	Address string `json:"address" binding:"required"`
//	Postcodes string `json:"postcodes" binding:"required"`
//	Tel string `json:"tel" binding:"required"`
//}
//type additional struct {
//	Address address
//}
//type param struct {
//
//	StoreCode string `json:"store_code" binding:"required"`
//	Uid int `json:"uid" binding:"required"`
//	Additional additional
//	WhTradeCode string `json:"wh_trade_code" binding:"required"`
//	Str string `json:"str" binding:"required"`
//}

type arr struct {
	pro_code                 string
	unit_price               float64
	original_price           float64
	merchant_pro_standard_id int
	sku                      string
	plat_pid                 int
	ship_id                  int
	types                    int8
	cost_type                int8
	weight                   float64
	volume                   float64
	global_send_type         int8
	m_batch_price            float64
	property_1               string
	property_2               string
	pic                      string
	tax_rate                 float64
	spu_id                   int
	goods_source             int8
	supplier_type            int8
	wh_price                 string
	is_special               string
	count                    string
	merchant_pro_code        string
}

func Routine(ctx *gin.Context) {

	go func() {
		time.Sleep(1 * time.Hour)
	}()
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
	for i := range c {

		fmt.Println(i)

	}
	fmt.Println("finish")

	//time.Sleep(time.Second)

}

//订单配合Redis秒杀

func Saveorder(c *gin.Context) {

	err := db.RedisClient.Exists("count").Err()
	if err != nil {
		//设置库存
		err := db.RedisClient.Set("count", 10000, 0).Err()
		if err != nil {
			panic("something was wrong" + err.Error())
		}
	}
	//自增
	incre := db.RedisClient.Incr("lin").Val()
	result := db.RedisClient.Get("count").Val()
	data, err := strconv.ParseInt(result, 10, 64)
	if err != nil {
		panic("wrong")
	}

	if incre < data {
		//存入list todo not finish
		err := db.RedisClient.RPush("Mylist", "goodid", []string{
			"incrementid", "store_id",
		}).Err()
		if err != nil {
			panic("mylist save error")
		}
		//插入订单
		//db.Eloquent.
		//减库存
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": "抢购成功 ",
		})
	} else {
		//计算抢购失败用户
		db.RedisClient.Incr("fail")
		c.JSON(http.StatusOK, gin.H{
			"code": 303,
			"data": "抢购失败",
		})
	}

}

/**
设置库存
*/
func Setcount(c *gin.Context) {
	val, err := db.RedisClient.Get("setCount").Result()
	if err == nil {
		panic("setCount is exists")
	}
	val, seterr := db.RedisClient.Set("setCount", 100, 0).Result()
	if seterr != nil {
		panic("设置库存失败")
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": val,
	})
}

//获取库存
func Getcount(c *gin.Context) {
	val, err := db.RedisClient.Get("setCount").Result()
	if err != nil {
		panic("setCount is not exists")
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": val,
	})
}

//param:=param{}
//	c.BindJSON(&param)
//	arr:=strings.Split(param.Str,"|")
//	product := make(map[string]string)
//	increment_id:=UniqueId()
//for _,value:=range arr{
//		tempProduct:=strings.Split(value,",")
//		product[tempProduct[0]] = tempProduct[1]
//		arrs:=orderMessage(tempProduct)
//		saveItem(arrs,increment_id)
//	}

//fmt.Println(UniqueId())
//	c.JSON(http.StatusOK, gin.H{
//	"code": 200,
//	//"data":arrs,
//})

//}
//保存item表
func saveItem(product arr, increment_id string) {
	item := models.GbOrderItem{
		IncrementID:     increment_id,
		Sku:             product.sku,
		UnitPrice:       product.unit_price,
		UnitFreight:     0,
		ShipID:          product.ship_id,
		Count:           product.count,
		Property1:       product.property_1,
		Property2:       product.property_2,
		WhPrice:         product.wh_price,
		MBatchPrice:     product.m_batch_price,
		Weight:          product.weight,
		Volume:          product.volume,
		MerchantProCode: product.merchant_pro_code,
		Pic:             product.pic,
	}
	db.Eloquent.Create(&item)
}
func saveChidren(product arr, increment_id string) {

}

/**
获取下单产品相关信息
*/
func orderMessage(tempProduct []string) arr {
	arr := arr{}
	var GbMerchantProStandard = models.GbMerchantProStandard{}

	var gbstoreproduct = models.GbStoreProduct{}

	var GbProStandard = models.GbProStandard{}

	db.Eloquent.Find(&GbMerchantProStandard, tempProduct[0])

	db.Eloquent.Where("mp_code=?", GbMerchantProStandard.GbMerchantProduct.Code).First(&gbstoreproduct)

	db.Eloquent.Where(map[string]interface{}{"code": GbMerchantProStandard.GbMerchantProduct.PlatProCode, "sku": GbMerchantProStandard.Sku}).Find(&GbProStandard)

	arr.pro_code = GbMerchantProStandard.PlatProCode
	arr.unit_price = GbMerchantProStandard.Price
	arr.original_price = GbMerchantProStandard.Price
	arr.merchant_pro_standard_id = GbMerchantProStandard.ID
	arr.sku = GbMerchantProStandard.Sku
	arr.plat_pid = GbMerchantProStandard.ID
	arr.ship_id = gbstoreproduct.GbProduct.ShipTemplate
	arr.types = gbstoreproduct.GbProduct.Type
	arr.cost_type = gbstoreproduct.GbProduct.CostType
	arr.weight = gbstoreproduct.GbProduct.Weight
	arr.volume = gbstoreproduct.GbProduct.Volume
	arr.global_send_type = gbstoreproduct.GbProduct.GlobalSendType
	arr.m_batch_price = GbProStandard.MBatchPrice
	arr.property_1 = GbProStandard.Property1
	arr.property_2 = GbProStandard.Property2
	arr.pic = GbProStandard.Pic
	arr.tax_rate = GbProStandard.TaxRate
	arr.spu_id = gbstoreproduct.ID
	arr.goods_source = gbstoreproduct.GoodsSource
	arr.supplier_type = gbstoreproduct.SupplierType
	arr.wh_price = tempProduct[2]
	arr.is_special = tempProduct[3]
	arr.count = tempProduct[1]
	return arr

}
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func UniqueId() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}

//func GetProInfo(tempProduct []string )map[string]interface{}{
//	arr := make(map[string]interface{})
//	var	gbmerchantpro=models.GbMerchantProStandard{}
//	var	gbmerchantproduct=models.GbMerchantProduct{}
//	var	gbstoreproduct=models.GbStoreProduct{}
//	var	gbproduct=models.GbProduct{}
//	var	prostan=models.GbProStandard{}
//
//	arr["pro_code"]=gbmerchantpro.PlatProCode
//	arr["unit_price"]=gbmerchantpro.Price
//	arr["original_price"]=gbmerchantpro.Price
//	arr["merchant_pro_standard_id"]=gbmerchantpro.ID
//	arr["sku"]=gbmerchantpro.Sku
//	arr["plat_pid"]=gbproduct.ID
//	arr["ship_id"]=gbproduct.ShipTemplate
//	arr["type"]=gbproduct.Type
//	arr["cost_type"]=gbproduct.CostType
//	arr["weight"]=gbproduct.Weight
//	arr["volume"]=gbproduct.Volume
//	arr["global_send_type"]=gbproduct.GlobalSendType
//	arr["m_batch_price"]=prostan.MBatchPrice
//	arr["property_1"]=prostan.Property1
//	arr["property_2"]=prostan.Property2
//	arr["pic"]=prostan.Pic
//	arr["tax_rate"]=prostan.TaxRate
//	arr["spu_id"]=gbstoreproduct.ID
//	arr["goods_source"]=gbstoreproduct.GoodsSource
//	arr["supplier_type"]=gbstoreproduct.SupplierType
//	arr["wh_price"]=(tempProduct[2])
//	arr["is_special"]=(tempProduct[3])
//
//
//
//	meerr:=db.Eloquent.Where("id=?",tempProduct[0]).First(&gbmerchantpro).Error
//	if meerr!=nil {
//		errors.New("GbMerchantProStandard 不能为空")
//	}
//	merproerr:=db.Eloquent.Where("id=?",gbmerchantpro.ProID).First(&gbmerchantproduct).Error
//	if merproerr!=nil{
//		errors.New("gbmerchantproduct 不能为空")
//	}
//	stoerr:=db.Eloquent.Where("mp_code=?",gbmerchantproduct.Code).First(&gbstoreproduct).Error
//	if  stoerr!=nil{
//		errors.New("gbstoreproduct 不能为空")
//	}
//
//	proerr:=db.Eloquent.Where("code=?",gbstoreproduct.PlatProCode).First(&gbproduct).Error
//	if  proerr!=nil{
//		errors.New("gbstoreproduct 不能为空")
//	}
//	prostanerr:=db.Eloquent.Where(map[string]interface{}{"code": gbmerchantpro.PlatProCode, "sku": "11111"}).Find(&prostan).Error
//	if  prostanerr!=nil{
//	}
//
//	arr["pro_code"]=gbmerchantpro.PlatProCode
//	arr["unit_price"]=gbmerchantpro.Price
//	arr["original_price"]=gbmerchantpro.Price
//	arr["merchant_pro_standard_id"]=gbmerchantpro.ID
//	arr["sku"]=gbmerchantpro.Sku
//	arr["plat_pid"]=gbproduct.ID
//	arr["ship_id"]=gbproduct.ShipTemplate
//	arr["type"]=gbproduct.Type
//	arr["cost_type"]=gbproduct.CostType
//	arr["weight"]=gbproduct.Weight
//	arr["volume"]=gbproduct.Volume
//	arr["global_send_type"]=gbproduct.GlobalSendType
//	arr["m_batch_price"]=prostan.MBatchPrice
//	arr["property_1"]=prostan.Property1
//	arr["property_2"]=prostan.Property2
//	arr["pic"]=prostan.Pic
//	arr["tax_rate"]=prostan.TaxRate
//	arr["spu_id"]=gbstoreproduct.ID
//	arr["goods_source"]=gbstoreproduct.GoodsSource
//	arr["supplier_type"]=gbstoreproduct.SupplierType
//	arr["wh_price"]=(tempProduct[2])
//	arr["is_special"]=(tempProduct[3])
//	return arr
//
//}
//func order(p param){
//
//}
//func  item(p param)  {
//
//}
//func orderChidren(p param)  {
//
//}
//func addre(p param){
//
//}

// 所有用户
func List(c *gin.Context) {
	//list := []string{}
	//
	////db.Eloquent.Quer
	//var order models.GbOrder
	//err:=db.Eloquent.Preload("GbOrderChildren").Where("increment_id=191120004233025").Find(&order).Error
	//	message:="-----"
	////for _,value:=range order.GbOrderChildren{
	////	//fmt.Println(value)
	////	list=append(list,value.ProName)
	//
	//}
	//fmt.Println(list)
	//fmt.Println(order.ID)
	//if err !=nil{
	//	return
	//}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		//"data": order,
		//"message": message,
	})
}
