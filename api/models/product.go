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
  `one_category_id` smallint(5) unsigned NOT NULL COMMENT '一级分类ID',
  `two_category_id` smallint(5) unsigned NOT NULL COMMENT '二级分类ID',
  `three_category_id` smallint(5) unsigned NOT NULL COMMENT '三级分类ID',
  `supplier_id` int(10) unsigned NOT NULL COMMENT '商品的供应商ID',
  `price` decimal(8,2) NOT NULL COMMENT '商品销售价格',
  `average_cost` decimal(18,2) NOT NULL COMMENT '商品加权平均成本',
  `publish_status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '上下架状态：0下架1上架',
  `audit_status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '审核状态：0未审核，1已审核',
  `weight` float DEFAULT NULL COMMENT '商品重量',
  `length` float DEFAULT NULL COMMENT '商品长度',
  `height` float DEFAULT NULL COMMENT '商品高度',
  `width` float DEFAULT NULL COMMENT '商品宽度',
  `color_type` enum('红','黄','蓝','黑') DEFAULT NULL,
  `production_date` datetime NOT NULL COMMENT '生产日期',
  `shelf_life` int(11) NOT NULL COMMENT '商品有效期',
  `descript` text NOT NULL COMMENT '商品描述',
  `indate` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '商品录入时间',
  `modified_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  PRIMARY KEY (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商品信息表'
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
	ProductID       int              ` gorm:"primary_key;column:product_id;type:int(10) unsigned;not null"`                                       // 商品ID
	ProductCore     string           `form:"product_core" binding:"required" gorm:"column:product_core;type:char(16);not null"`                   // 商品编码
	ProductName     string           `form:"product_name" binding:"required" gorm:"column:product_name;type:varchar(20);not null"`                // 商品名称
	BarCode         string           `form:"bar_code" binding:"required" gorm:"column:bar_code;type:varchar(50);not null"`                        // 国条码
	BrandID         int              `form:"brand_id" binding:"required" gorm:"column:brand_id;type:int(10) unsigned;not null"`                   // 品牌表的ID
	OneCategoryID   int16            `form:"one_category_id" binding:"required" gorm:"column:one_category_id;type:smallint(5) unsigned;not null"` // 一级分类ID
	TwoCategoryID   int16            `gorm:"column:two_category_id;type:smallint(5) unsigned;not null"`                                           // 二级分类ID
	ThreeCategoryID int16            `gorm:"column:three_category_id;type:smallint(5) unsigned;not null"`                                         // 三级分类ID
	Price           float64          `form:"price" binding:"required" gorm:"column:price;type:decimal(8,2);not null"`                             // 商品销售价格
	AverageCost     float64          `gorm:"column:average_cost;type:decimal(18,2);not null"`                                                     // 商品加权平均成本
	PublishStatus   int8             `form:"publish_status" binding:"required" gorm:"column:publish_status;type:tinyint(4);not null"`             // 上下架状态：0下架1上架
	AuditStatus     int8             `form:"audit_status" binding:"required" gorm:"column:audit_status;type:tinyint(4);not null"`                 // 审核状态：0未审核，1已审核
	Weight          float32          `form:"weight" gorm:"column:weight;type:float"`                                                              // 商品重量
	Length          float32          `form:"length" gorm:"column:length;type:float"`                                                              // 商品长度
	Height          float32          `form:"height" gorm:"column:height;type:float"`                                                              // 商品高度
	Width           float32          `form:"width" gorm:"column:width;type:float"`                                                                // 商品宽度
	ColorType       string           `form:"color_type" gorm:"column:color_type;type:enum('红','黄','蓝','黑')"`
	ProductionDate  string           `form:"production_date" gorm:"column:production_date;type:datetime;not null"` // 生产日期
	Descript        string           `form:"descript" gorm:"column:descript;type:text;not null"`                   // 商品描述
	Indate          time.Time        `form:"indate" gorm:"column:indate;type:timestamp;not null"`                  // 商品录入时间
	ModifiedTime    time.Time        `form:"modified_time" gorm:"column:modified_time;type:timestamp;not null"`    // 最后修改时间
	PicInFo         []ProductPicInfo `gorm:"FOREIGNKEY:ProductID;ASSOCIATION_FOREIGNKEY:ProductID"`
}

func (p *ProductInfo) BeforeCreate(scope *gorm.Scope) error {

	scope.SetColumn("Indate", time.Now())
	scope.SetColumn("ModifiedTime", time.Now())

	return nil
}

//添加
func (p ProductInfo) Insert() (ProductID int, err error) {
	tx := db.Eloquent.Begin()
	result := tx.Create(&p)
	ProductID = p.ProductID
	if result.Error != nil {
		tx.Rollback()
		return 0, result.Error
	}
	tx.Commit()
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
