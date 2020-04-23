package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

/*
@Time : 2020/4/20 2:33 PM
@Author : apple
@File : productParam
@Software: GoLand
*/
/**
CREATE TABLE `product_param` (
  `product_param_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '商品参数ID',
  `product_id` int(10) unsigned NOT NULL COMMENT '商品ID',
  `param_name` varchar(50) DEFAULT NULL COMMENT '参数名称',
  `param_value` varchar(200) NOT NULL COMMENT '参数值',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `modified_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  PRIMARY KEY (`product_param_id`),
  KEY `FK_i1q2cf5pxfr8r69cfci3yyari` (`product_id`),
  CONSTRAINT `FK_i1q2cf5pxfr8r69cfci3yyari` FOREIGN KEY (`product_id`) REFERENCES `product_info` (`product_id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8 COMMENT='商品产品规格表';
*/
type ProductParam struct {
	ProductParamId int       `gorm:"primary_key;column:product_param_id;type:int(10) unsigned;not null"`                    // 商品图片ID
	ProductID      int       `form:"product_id" binding:"required" gorm:"column:product_id;type:int(10) unsigned;not null"` // 商品ID
	ParamName      string    `form:"param_name" binding:"required" gorm:"column:param_name;type:varchar(50);not null"`      // 参数名称
	ParamValue     string    `form:"param_value" binding:"required" gorm:"column:param_value;type:varchar(200);not null"`   // 参数值
	CreateTime     time.Time `form:"create_time" gorm:"column:create_time;type:timestamp;not null"`                         // 创建
	ModifiedTime   time.Time `form:"modified_time" gorm:"column:modified_time;type:timestamp;not null"`                     // 最后修改时间
}

func (pa *ProductParam) BeforeCreate(scope *gorm.Scope) error {

	scope.SetColumn("ModifiedTime,CreateTime", time.Now())

	return nil
}
