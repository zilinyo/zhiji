package models

import (
	"github.com/jinzhu/gorm"
	"time"
	db "zhiji/api/database"
)

/******sql******
CREATE TABLE `product_pic_info` (
  `product_pic_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '商品图片ID',
  `product_id` int(10) unsigned NOT NULL COMMENT '商品ID',
  `specifications_name` varchar(50) NOT NULL COMMENT '规格名称',
  `specifications_value` varchar(200) NOT NULL DEFAULT '0' COMMENT '规格值',
  `specifications_price` decimal(8,2) NOT NULL DEFAULT '0.00' COMMENT '规格价格',
  `specifications_count` int(10) NOT NULL DEFAULT '1' COMMENT '规格数量',
  `modified_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  `specifications_url` varchar(255) NOT NULL,
  PRIMARY KEY (`product_pic_id`),
  KEY `FK_i1q2cf5pxfr8r69cfci3ssss` (`product_id`),
  CONSTRAINT `FK_i1q2cf5pxfr8r69cfci3ssss` FOREIGN KEY (`product_id`) REFERENCES `product_info` (`product_id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8 COMMENT='商品图片信息表';
******sql******/
// ProductPicInfo 商品图片信息表
type ProductPicInfo struct {
	ProductPicID        int     `gorm:"primary_key;column:product_pic_id;type:int(10) unsigned;not null"`                                      // 商品图片ID
	ProductID           int     `form:"product_id" binding:"required" gorm:"column:product_id;type:int(10) unsigned;not null"`                 // 商品ID
	SpecificationsName  string  `form:"specifications_name" binding:"required" gorm:"column:specifications_name;type:varchar(50)"`             // 规格名称
	SpecificationsValue string  `form:"specifications_value" binding:"required" gorm:"column:specifications_value;type:varchar(200);not null"` // 规格值
	SpecificationsPrice float64 `form:"specifications_price" binding:"required" gorm:"column:specifications_price;type:decimal(8,2);not null"` // 规格价格
	SpecificationsCount int     `form:"specifications_count" binding:"required" gorm:"column:specifications_count;type:int(10);not null"`      // 规格数量
	SpecificationsUrl   string  `form:"specifications_url" binding:"required" gorm:"column:specifications_url;type:varchar(200);not null"`     // url

	ModifiedTime time.Time `form:"modified_time" gorm:"column:modified_time;type:timestamp;not null"` // 最后修改时间
}

func (p *ProductPicInfo) BeforeCreate(scope *gorm.Scope) error {

	scope.SetColumn("ModifiedTime", time.Now())

	return nil
}

//添加
func (p ProductPicInfo) Insert(id int) (err error) {
	p.ProductID = id
	result := db.Eloquent.Create(&p)
	if result.Error != nil {
		return
	}
	return
}
