package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

/******sql******
CREATE TABLE `product_category` (
  `category_id` smallint(5) unsigned NOT NULL AUTO_INCREMENT COMMENT '分类ID',
  `category_name` varchar(10) NOT NULL COMMENT '分类名称',
  `category_code` varchar(10) NOT NULL COMMENT '分类编码',
  `parent_id` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '父分类ID',
  `category_level` tinyint(4) NOT NULL DEFAULT '1' COMMENT '分类层级',
  `category_status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '分类状态',
  `modified_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  PRIMARY KEY (`category_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='商品分类表'
******sql******/
// ProductCategory 商品分类表
type ProductCategory struct {
	CategoryID     int16              `gorm:"primary_key;column:category_id;type:smallint(5) unsigned;not null"`                         // 分类ID
	CategoryName   string             `form:"category_id" binding:"required" gorm:"column:category_name;type:varchar(10);not null"`      // 分类名称
	CategoryCode   string             `form:"category_id" binding:"required" gorm:"column:category_code;type:varchar(10);not null"`      // 分类编码
	ParentID       int16              `gorm:"column:parent_id;type:smallint(5) unsigned;not null"`                                       // 父分类ID
	CategoryLevel  int8               `gorm:"column:category_level;type:tinyint(4);not null"`                                            // 分类层级
	CategoryStatus int8               `form:"category_status" binding:"required" gorm:"column:category_status;type:tinyint(4);not null"` // 分类状态
	ModifiedTime   time.Time          `gorm:"column:modified_time;type:timestamp;not null"`                                              // 最后修改时间
	Children       []*ProductCategory `json:"list,omitempty"`
}

/**
更新时间
*/
func (*ProductCategory) BeforeCreate(scope *gorm.Scope) error {

	scope.SetColumn("ModifiedTime", time.Now())

	return nil
}
func (productCategory *ProductCategory) BeforeUpdate(scope *gorm.Scope) error {

	scope.SetColumn("ModifiedTime", time.Now())

	return nil
}

/**
数据组装
*/
//func (*ProductCategory) BuildData(list []*ProductCategory) []*ProductCategory {
//
//	var data = make(map[int]map[int]*ProductCategory)
//
//	for _, v := range list {
//
//		CategoryID := int(v.CategoryID)
//
//		ParentID :=int(v.ParentID)
//
//		if _, ok := data[ParentID]; !ok {
//
//			data[ParentID] = make(map[int]*ProductCategory)
//
//		}
//
//		data[ParentID][CategoryID] = v
//	}
//	result := ProductCategory.MakeTreeCore(0, data)
//	return result
//}
/**
数据排序
*/
//func (productCategory *ProductCategory) MakeTreeCore(index int, data map[int]map[int]*ProductCategory) []*ProductCategory {
//
//	tmp := make([]*ProductCategory, 0)
//
//	for id, item := range data[index] {
//
//		if data[id] != nil {
//
//			item.Children = productCategory.MakeTreeCore(id, data)
//
//		}
//
//		tmp = append(tmp, item)
//	}
//
//	return tmp
//}
