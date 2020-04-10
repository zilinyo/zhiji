package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

/*
@Time : 2020/4/9 1:48 PM
@Author : apple
@File : BrandInfo
@Software: GoLand
*/
/******sql******
CREATE TABLE `brand_info` (
  `brand_id` smallint(5) unsigned NOT NULL AUTO_INCREMENT COMMENT '品牌ID',
  `brand_name` varchar(50) NOT NULL COMMENT '品牌名称',
  `telephone` varchar(50) NOT NULL COMMENT '联系电话',
  `brand_web` varchar(100) DEFAULT NULL COMMENT '品牌网络',
  `brand_logo` varchar(100) DEFAULT NULL COMMENT '品牌logo URL',
  `brand_desc` varchar(150) DEFAULT NULL COMMENT '品牌描述',
  `brand_status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '品牌状态,0禁用,1启用',
  `brand_order` tinyint(4) NOT NULL DEFAULT '0' COMMENT '排序',
  `modified_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  PRIMARY KEY (`brand_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='品牌信息表'
******sql******/
// BrandInfo 品牌信息表
type BrandInfo struct {
	BrandID      int16     `gorm:"primary_key;column:brand_id;type:smallint(5) unsigned;not null"`                      // 品牌ID
	BrandName    string    `form:"brand_name" binding:"required" gorm:"column:brand_name;type:varchar(50);not null"`    // 品牌名称
	Telephone    string    `form:"telephone" binding:"required" gorm:"column:telephone;type:varchar(50);not null"`      // 联系电话
	BrandWeb     string    `form:"brand_web" binding:"required" gorm:"column:brand_web;type:varchar(100)"`              // 品牌网络
	BrandLogo    string    `form:"brand_logo" binding:"required" gorm:"column:brand_logo;type:varchar(100)"`            // 品牌logo URL
	BrandDesc    string    `form:"brand_desc" binding:"required" gorm:"column:brand_desc;type:varchar(150)"`            // 品牌描述
	BrandStatus  int8      `form:"brand_status" binding:"required" gorm:"column:brand_status;type:tinyint(4);not null"` // 品牌状态,0禁用,1启用
	BrandOrder   int8      `form:"brand_order" binding:"required" gorm:"column:brand_order;type:tinyint(4);not null"`   // 排序
	ModifiedTime time.Time `gorm:"column:modified_time;type:timestamp;not null"`                                        // 最后修改时间
}

/**
更新时间
*/
func (brandInfo *BrandInfo) BeforeCreate(scope *gorm.Scope) error {

	scope.SetColumn("ModifiedTime", time.Now())

	return nil
}
func (brandInfo *BrandInfo) BeforeUpdate(scope *gorm.Scope) error {

	scope.SetColumn("ModifiedTime", time.Now())

	return nil
}
