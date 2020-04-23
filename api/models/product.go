package models

import (
	"github.com/jinzhu/gorm"
	"time"
	db "zhiji/api/database"
)

/******sql******
CREATE TABLE `product_info` (
  `product_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '商品ID',
  `product_core` char(16) NOT NULL COMMENT '商品编码',
  `product_name` varchar(20) NOT NULL COMMENT '商品名称',
  `bar_code` varchar(50) NOT NULL COMMENT '国条码',
  `brand_id` int(10) unsigned NOT NULL COMMENT '品牌表的ID',
  `category_id` smallint(5) unsigned NOT NULL COMMENT '一级分类ID',
  `main_url` varchar(255) NOT NULL COMMENT '主图',
  `gallery` text COMMENT '画廊',
  `price` decimal(8,2) NOT NULL COMMENT '商品销售价格',
  `average_cost` decimal(18,2) NOT NULL COMMENT '商品加权平均成本',
  `publish_status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '上下架状态：0下架1上架',
  `is_new` tinyint(4) DEFAULT '1' COMMENT '1表示是普通，2表示新品',
  `is_hot` tinyint(4) NOT NULL DEFAULT '1' COMMENT '1表示普通，2表示热卖',
  `audit_status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '审核状态：0未审核，1已审核',
  `weight` varchar(20) DEFAULT NULL COMMENT '商品重量',
  `length` varchar(20) DEFAULT NULL COMMENT '商品长度',
  `height` varchar(20) DEFAULT NULL COMMENT '商品高度',
  `width` varchar(20) DEFAULT NULL COMMENT '商品宽度',
  `color_type` varchar(20) DEFAULT NULL,
  `descript` text NOT NULL COMMENT '商品描述',
  `indate` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '商品录入时间',
  `modified_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  `is_banner` tinyint(2) NOT NULL DEFAULT '2' COMMENT '是否为banner图 取前五张 1表示是 2表示否 ',
  PRIMARY KEY (`product_id`,`is_banner`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8 COMMENT='商品信息表';
******sql******/

//Weight	float32	`form:"weight"`	// 排序
//Length	float32	`form:"length"`	// 排序
//Height	float32	`form:"height"`
//Width	float32	`form:"width"`
//ColorType	string	`form:"color_type"`
//ProductionDate	string	`form:"production_date"`
//ShelfLife	int	`form:"shelf_life"`
//Descript	string	`form:"descript"`
//Indate	string	`form:"indate" `
//ModifiedTime	string	`form:"modified_time"`
//PicInFo			[]*PPCvidate
type ProductInfo struct {
	ProductID     int       ` gorm:"primary_key;column:product_id;type:int(10) unsigned;not null"`                           // 商品ID
	ProductCore   string    `form:"product_core" binding:"required" gorm:"column:product_core;type:char(16);not null"`       // 商品编码
	ProductName   string    `form:"product_name" binding:"required" gorm:"column:product_name;type:varchar(20);not null"`    // 商品名称
	BarCode       string    `form:"bar_code" binding:"required" gorm:"column:bar_code;type:varchar(50);not null"`            // 国条码
	BrandID       int       `form:"brand_id" binding:"required" gorm:"column:brand_id;type:int(10) unsigned;not null"`       // 品牌表的ID
	CategoryID    int16     `form:"category_id" binding:"required" gorm:"column:category_id;type:int(5) unsigned;not null"`  // 一级分类ID
	MainURL       string    `form:"main_url" binding:"required" gorm:"column:main_url;type:varchar(200);not null"`           // 图片URL
	Price         float64   `form:"price" binding:"required" gorm:"column:price;type:decimal(8,2);not null"`                 // 商品销售价格
	AverageCost   float64   `gorm:"column:average_cost;type:decimal(18,2);not null"`                                         // 商品加权平均成本
	PublishStatus int8      `form:"publish_status" binding:"required" gorm:"column:publish_status;type:tinyint(4);not null"` // 上下架状态：0下架1上架
	AuditStatus   int8      `form:"audit_status" binding:"required" gorm:"column:audit_status;type:tinyint(4);not null"`     // 审核状态：0未审核，1已审核
	Weight        string    ` gorm:"column:weight;type:string"`                                                              // 商品重量
	Length        string    ` gorm:"column:length;type:string"`                                                              // 商品长度
	Height        string    ` gorm:"column:height;type:string"`                                                              // 商品高度
	Width         string    `gorm:"column:width;type:string"`                                                                // 商品宽度
	ColorType     string    `form:"color_type" gorm:"column:color_type;varchar(200)"`
	Descript      string    `form:"descript" gorm:"column:descript;type:text;not null"`                // 商品描述
	Indate        time.Time `form:"indate" gorm:"column:indate;type:timestamp;not null"`               // 商品录入时间
	ModifiedTime  time.Time `form:"modified_time" gorm:"column:modified_time;type:timestamp;not null"` // 最后修改时间
	Gallery       string    `form:"gallery" gorm:"column:gallery;type:text;not null"`
	IsBanner      int8      `gorm:"column:is_banner;type:tinyint(2)"` //是否为banner图 取前五张
	IsNew         int8      `gorm:"column:is_new;type:tinyint(2)"`    //是否为新品
	IsHot         int8      `gorm:"column:is_hot;type:tinyint(2)"`    //是否为热图

	//Gallery   		[]struct{
	//	PicURL string
	//}            	 `gorm:"column:gallery;type:text"`                               //是否为banner图 取前五张

	PicInFo      []ProductPicInfo `gorm:"FOREIGNKEY:ProductID;ASSOCIATION_FOREIGNKEY:ProductID"` //画廊
	ProductParam []ProductParam   `gorm:"FOREIGNKEY:ProductID;ASSOCIATION_FOREIGNKEY:ProductID"` //参数图
}

func (p *ProductInfo) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("Indate,ModifiedTime", time.Now())
	return nil
}

//添加
func (p ProductInfo) Insert() (ProductID int, err error) {

	tx := db.Eloquent.Begin() //开启事务

	result := tx.Create(&p)

	ProductID = p.ProductID

	if result.Error != nil {

		tx.Rollback() //回滚

		return 0, result.Error
	}
	tx.Commit() //提交

	return ProductID, err
}

//列表
func (p ProductInfo) Users(data map[string]int) (P []ProductInfo, err error) {
	var total int = 0
	db.Eloquent.Find(&p).Count(&total)
	db.Eloquent.Limit(data["pageSize"]).Offset((data["page"] - 1) * data["pageSize"]).Find(&p)
	if err = db.Eloquent.Find(&p).Error; err != nil {
		return
	}
	return
}

//新品推荐
func (p ProductInfo) NewGoodsList() (P []ProductInfo) {

	var productInfo []ProductInfo

	err := db.Eloquent.Preload("PicInFo", func(query *gorm.DB) *gorm.DB {

		return query.Where("is_master=1")

	}).Where("is_banner=1").Order("modified_time").Limit(6).Find(&productInfo).Error

	if err != nil {
		return
	}
	return productInfo
}

//首页banner图
func (p ProductInfo) Banner() (P []ProductInfo) {

	var productInfo []ProductInfo

	err := db.Eloquent.Preload("PicInFo", func(query *gorm.DB) *gorm.DB {

		return query.Where("is_master=1")

	}).Where("is_banner=1").Limit(5).Find(&productInfo).Error

	if err != nil {
		return
	}
	return productInfo
}

//分类
func (p ProductInfo) Channel() (C []ProductCategory) {

	var ProductCategory []ProductCategory
	err := db.Eloquent.Find(&ProductCategory).Error

	if err != nil {
		return
	}
	return ProductCategory
}
