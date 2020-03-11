package models

import (
	"time"
	db "zhiji/api/database"
)

/******sql******
CREATE TABLE `order_master` (
  `order_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '订单ID',
  `order_sn` bigint(20) unsigned NOT NULL COMMENT '订单编号 yyyymmddnnnnnnnn',
  `customer_id` int(10) unsigned NOT NULL COMMENT '下单人ID',
  `shipping_user` varchar(10) NOT NULL COMMENT '收货人姓名',
  `province` smallint(6) NOT NULL COMMENT '省',
  `city` smallint(6) NOT NULL COMMENT '市',
  `district` smallint(6) NOT NULL COMMENT '区',
  `address` varchar(100) NOT NULL COMMENT '地址',
  `payment_method` tinyint(4) NOT NULL COMMENT '支付方式：1现金，2余额，3网银，4支付宝，5微信',
  `order_money` decimal(8,2) NOT NULL COMMENT '订单金额',
  `district_money` decimal(8,2) NOT NULL DEFAULT '0.00' COMMENT '优惠金额',
  `shipping_money` decimal(8,2) NOT NULL DEFAULT '0.00' COMMENT '运费金额',
  `payment_money` decimal(8,2) NOT NULL DEFAULT '0.00' COMMENT '支付金额',
  `shipping_comp_name` varchar(10) DEFAULT NULL COMMENT '快递公司名称',
  `shipping_sn` varchar(50) DEFAULT NULL COMMENT '快递单号',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '下单时间',
  `shipping_time` datetime DEFAULT NULL COMMENT '发货时间',
  `pay_time` datetime DEFAULT NULL COMMENT '支付时间',
  `receive_time` datetime DEFAULT NULL COMMENT '收货时间',
  `order_status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '订单状态',
  `order_point` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '订单积分',
  `invoice_time` varchar(100) DEFAULT NULL COMMENT '发票抬头',
  `modified_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  PRIMARY KEY (`order_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='订单主表'
******sql******/
// OrderMaster 订单主表
type OrderMaster struct {
	OrderID          int       `gorm:"primary_key;column:order_id;type:int(10) unsigned;not null"`                             // 订单ID
	OrderSn          int64     `form:"order_sn" binding:"required" gorm:"column:order_sn;type:bigint(20) unsigned;not null"`   // 订单编号 yyyymmddnnnnnnnn
	CustomerID       int       `form:"customer_id" binding:"required"gorm:"column:customer_id;type:int(10) unsigned;not null"` // 下单人ID
	ShippingUser     string    `form:"shipping_user" binding:"required"gorm:"column:shipping_user;type:varchar(10);not null"`  // 收货人姓名
	Province         int16     `form:"province" binding:"required"gorm:"column:province;type:smallint(6);not null"`            // 省
	City             int16     `form:"city" binding:"required"gorm:"column:city;type:smallint(6);not null"`                    // 市
	District         int16     `form:"district" binding:"required"gorm:"column:district;type:smallint(6);not null"`            // 区
	Address          string    `form:"address" binding:"required"gorm:"column:address;type:varchar(100);not null"`             // 地址
	PaymentMethod    int8      `form:"payment_method" gorm:"column:payment_method;type:tinyint(4);not null"`                   // 支付方式：1现金，2余额，3网银，4支付宝，5微信
	OrderMoney       float64   `form:"order_money" gorm:"column:order_money;type:decimal(8,2);not null"`                       // 订单金额
	DistrictMoney    float64   `form:"district_money" gorm:"column:district_money;type:decimal(8,2);not null"`                 // 优惠金额
	ShippingMoney    float64   `form:"shipping_money" gorm:"column:shipping_money;type:decimal(8,2);not null"`                 // 运费金额
	PaymentMoney     float64   `form:"payment_money" gorm:"column:payment_money;type:decimal(8,2);not null"`                   // 支付金额
	ShippingCompName string    `form:"shipping_comp_name" gorm:"column:shipping_comp_name;type:varchar(10)"`                   // 快递公司名称
	ShippingSn       string    `form:"shipping_sn" gorm:"column:shipping_sn;type:varchar(50)"`                                 // 快递单号
	CreateTime       string    `form:"create_time"gorm:"column:create_time;type:timestamp;not null"`                           // 下单时间
	ShippingTime     string    `form:"shipping_time" gorm:"column:shipping_time;type:datetime"`                                // 发货时间
	PayTime          time.Time `form:"pay_time" orm:"column:pay_time;type:datetime"`                                           // 支付时间
	ReceiveTime      string    `form:"receive_time" binding:"required"gorm:"column:receive_time;type:datetime"`                // 收货时间
	OrderStatus      int8      `form:"order_status" binding:"required"gorm:"column:order_status;type:tinyint(4);not null"`     // 订单状态
	OrderPoint       int       `form:"order_point" binding:"required"gorm:"column:order_point;type:int(10) unsigned;not null"` // 订单积分
	InvoiceTime      string    `form:"invoice_time" binding:"required"gorm:"column:invoice_time;type:varchar(100)"`            // 发票抬头
	ModifiedTime     string    `form:"modified_time" binding:"required"gorm:"column:modified_time;type:timestamp;not null"`    // 最后修改时间
}

/******sql******
CREATE TABLE `order_detail` (
  `order_detail_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '订单详情表ID',
  `order_id` int(10) unsigned NOT NULL COMMENT '订单表ID',
  `product_id` int(10) unsigned NOT NULL COMMENT '订单商品ID',
  `product_name` varchar(50) NOT NULL COMMENT '商品名称',
  `product_cnt` int(11) NOT NULL DEFAULT '1' COMMENT '购买商品数量',
  `product_price` decimal(8,2) NOT NULL COMMENT '购买商品单价',
  `average_cost` decimal(8,2) NOT NULL COMMENT '平均成本价格',
  `weight` float DEFAULT NULL COMMENT '商品重量',
  `fee_money` decimal(8,2) NOT NULL DEFAULT '0.00' COMMENT '优惠分摊金额',
  `w_id` int(10) unsigned NOT NULL COMMENT '仓库ID',
  `modified_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  PRIMARY KEY (`order_detail_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='订单详情表'
******sql******/
// OrderDetail 订单详情表
type OrderDetail struct {
	OrderDetailID int     ` gorm:"primary_key;column:order_detail_id;type:int(10) unsigned;not null"`                      // 订单详情表ID
	OrderID       int     ` gorm:"column:order_id;type:int(10) unsigned;not null"`                                         // 订单表ID
	ProductID     int     `form:"product_id" binding:"required" gorm:"column:product_id;type:int(10) unsigned;not null"`   // 订单商品ID
	ProductName   string  `form:"product_name" binding:"required" gorm:"column:product_name;type:varchar(50);not null"`    // 商品名称
	ProductCnt    int     `form:"product_cnt" binding:"required" gorm:"column:product_cnt;type:int(11);not null"`          // 购买商品数量
	ProductPrice  float64 `form:"product_price" binding:"required" gorm:"column:product_price;type:decimal(8,2);not null"` // 购买商品单价
	AverageCost   float64 ` gorm:"column:average_cost;type:decimal(8,2);not null"`                                         // 平均成本价格
	Weight        float32 ` gorm:"column:weight;type:float"`                                                               // 商品重量
	FeeMoney      float64 `gorm:"column:fee_money;type:decimal(8,2);not null"`                                             // 优惠分摊金额
	ModifiedTime  string  `form:"order_sn" binding:"required" gorm:"column:modified_time;type:timestamp;not null"`         // 最后修改时间
}

//添加
func (o OrderMaster) Insert() (OrderId int, err error) {
	tx := db.Eloquent.Begin()
	result := tx.Create(&o)
	OrderId = o.OrderID
	if result.Error != nil {
		tx.Rollback()
		return 0, result.Error
	}
	tx.Commit()
	return OrderId, err
}
