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
  `pic_desc` varchar(50) DEFAULT NULL COMMENT '图片描述',
  `pic_url` varchar(200) NOT NULL COMMENT '图片URL',
  `is_master` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否主图：0.非主图1.主图',
  `pic_order` tinyint(4) NOT NULL DEFAULT '0' COMMENT '图片排序',
  `pic_status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '图片是否有效：0无效 1有效',
  `modified_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  PRIMARY KEY (`product_pic_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商品图片信息表'
******sql******/
// ProductPicInfo 商品图片信息表
type ProductPicInfo struct {
	ProductPicID int       `gorm:"primary_key;column:product_pic_id;type:int(10) unsigned;not null"`                      // 商品图片ID
	ProductID    int       `form:"product_id" binding:"required" gorm:"column:product_id;type:int(10) unsigned;not null"` // 商品ID
	PicDesc      string    `form:"pic_desc" binding:"required" gorm:"column:pic_desc;type:varchar(50)"`                   // 图片描述
	PicURL       string    `form:"pic_url" binding:"required" gorm:"column:pic_url;type:varchar(200);not null"`           // 图片URL
	IsMaster     int8      `form:"is_master" binding:"required" gorm:"column:is_master;type:tinyint(4);not null"`         // 是否主图：0.非主图1.主图
	PicOrder     int8      `form:"pic_order" binding:"required" gorm:"column:pic_order;type:tinyint(4);not null"`         // 图片排序
	PicStatus    int8      ` form:"pic_status" binding:"required" gorm:"column:pic_status;type:tinyint(4);not null"`      // 图片是否有效：0无效 1有效
	ModifiedTime time.Time `form:"modified_time" gorm:"column:modified_time;type:timestamp;not null"`                     // 最后修改时间
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
