package models

import (
	"time"
)

/******sql******
CREATE TABLE `customer_login_log` (
  `login_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '登陆日志ID',
  `customer_id` int(10) unsigned NOT NULL COMMENT '登陆用户ID',
  `login_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '用户登陆时间',
  `login_ip` int(10) unsigned NOT NULL COMMENT '登陆IP',
  `login_type` tinyint(4) NOT NULL COMMENT '登陆类型：0未成功，1成功',
  PRIMARY KEY (`login_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户登陆日志表'
******sql******/
// CustomerLoginLog 用户登陆日志表
type CustomerLoginLog struct {
	LoginID    int       `gorm:"primary_key;column:login_id;type:int(10) unsigned;not null"` // 登陆日志ID
	CustomerID int       `gorm:"column:customer_id;type:int(10) unsigned;not null"`          // 登陆用户ID
	LoginTime  time.Time `gorm:"column:login_time;type:timestamp;not null"`                  // 用户登陆时间
	LoginIP    int       `gorm:"column:login_ip;type:int(10) unsigned;not null"`             // 登陆IP
	LoginType  int8      `gorm:"column:login_type;type:tinyint(4);not null"`                 // 登陆类型：0未成功，1成功
}

/******sql******
CREATE TABLE `supplier_info` (
  `supplier_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '供应商ID',
  `supplier_code` char(8) NOT NULL COMMENT '供应商编码',
  `supplier_name` char(50) NOT NULL COMMENT '供应商名称',
  `supplier_type` tinyint(4) NOT NULL COMMENT '供应商类型：1.自营，2.平台',
  `link_man` varchar(10) NOT NULL COMMENT '供应商联系人',
  `phone_number` varchar(50) NOT NULL COMMENT '联系电话',
  `bank_name` varchar(50) NOT NULL COMMENT '供应商开户银行名称',
  `bank_account` varchar(50) NOT NULL COMMENT '银行账号',
  `address` varchar(200) NOT NULL COMMENT '供应商地址',
  `supplier_status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态：0禁止，1启用',
  `modified_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  PRIMARY KEY (`supplier_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='供应商信息表'
******sql******/
// SupplierInfo 供应商信息表
type SupplierInfo struct {
	SupplierID     int       `gorm:"primary_key;column:supplier_id;type:int(10) unsigned;not null"` // 供应商ID
	SupplierCode   string    `gorm:"column:supplier_code;type:char(8);not null"`                    // 供应商编码
	SupplierName   string    `gorm:"column:supplier_name;type:char(50);not null"`                   // 供应商名称
	SupplierType   int8      `gorm:"column:supplier_type;type:tinyint(4);not null"`                 // 供应商类型：1.自营，2.平台
	LinkMan        string    `gorm:"column:link_man;type:varchar(10);not null"`                     // 供应商联系人
	PhoneNumber    string    `gorm:"column:phone_number;type:varchar(50);not null"`                 // 联系电话
	BankName       string    `gorm:"column:bank_name;type:varchar(50);not null"`                    // 供应商开户银行名称
	BankAccount    string    `gorm:"column:bank_account;type:varchar(50);not null"`                 // 银行账号
	Address        string    `gorm:"column:address;type:varchar(200);not null"`                     // 供应商地址
	SupplierStatus int8      `gorm:"column:supplier_status;type:tinyint(4);not null"`               // 状态：0禁止，1启用
	ModifiedTime   time.Time `gorm:"column:modified_time;type:timestamp;not null"`                  // 最后修改时间
}

/******sql******
CREATE TABLE `customer_addr` (
  `customer_addr_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键ID',
  `customer_id` int(10) unsigned NOT NULL COMMENT 'customer_login表的自增ID',
  `zip` smallint(6) NOT NULL COMMENT '邮编',
  `province` smallint(6) NOT NULL COMMENT '地区表中省份的ID',
  `city` smallint(6) NOT NULL COMMENT '地区表中城市的ID',
  `district` smallint(6) NOT NULL COMMENT '地区表中的区ID',
  `address` varchar(200) NOT NULL COMMENT '具体的地址门牌号',
  `is_default` tinyint(4) NOT NULL COMMENT '是否默认',
  `modified_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  PRIMARY KEY (`customer_addr_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户地址表'
******sql******/
// CustomerAddr 用户地址表
type CustomerAddr struct {
	CustomerAddrID int       `gorm:"primary_key;column:customer_addr_id;type:int(10) unsigned;not null"` // 自增主键ID
	CustomerID     int       `gorm:"column:customer_id;type:int(10) unsigned;not null"`                  // customer_login表的自增ID
	Zip            int16     `gorm:"column:zip;type:smallint(6);not null"`                               // 邮编
	Province       int16     `gorm:"column:province;type:smallint(6);not null"`                          // 地区表中省份的ID
	City           int16     `gorm:"column:city;type:smallint(6);not null"`                              // 地区表中城市的ID
	District       int16     `gorm:"column:district;type:smallint(6);not null"`                          // 地区表中的区ID
	Address        string    `gorm:"column:address;type:varchar(200);not null"`                          // 具体的地址门牌号
	IsDefault      int8      `gorm:"column:is_default;type:tinyint(4);not null"`                         // 是否默认
	ModifiedTime   time.Time `gorm:"column:modified_time;type:timestamp;not null"`                       // 最后修改时间
}

/******sql******
CREATE TABLE `customer_level_inf` (
  `customer_level` tinyint(4) NOT NULL AUTO_INCREMENT COMMENT '会员级别ID',
  `level_name` varchar(10) NOT NULL COMMENT '会员级别名称',
  `min_point` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '该级别最低积分',
  `max_point` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '该级别最高积分',
  `modified_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  PRIMARY KEY (`customer_level`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户级别信息表'
******sql******/
// CustomerLevelInf 用户级别信息表
type CustomerLevelInf struct {
	CustomerLevel int8      `gorm:"primary_key;column:customer_level;type:tinyint(4);not null"` // 会员级别ID
	LevelName     string    `gorm:"column:level_name;type:varchar(10);not null"`                // 会员级别名称
	MinPoint      int       `gorm:"column:min_point;type:int(10) unsigned;not null"`            // 该级别最低积分
	MaxPoint      int       `gorm:"column:max_point;type:int(10) unsigned;not null"`            // 该级别最高积分
	ModifiedTime  time.Time `gorm:"column:modified_time;type:timestamp;not null"`               // 最后修改时间
}

/******sql******
CREATE TABLE `product_comment` (
  `comment_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '评论ID',
  `product_id` int(10) unsigned NOT NULL COMMENT '商品ID',
  `order_id` bigint(20) unsigned NOT NULL COMMENT '订单ID',
  `customer_id` int(10) unsigned NOT NULL COMMENT '用户ID',
  `title` varchar(50) NOT NULL COMMENT '评论标题',
  `content` varchar(300) NOT NULL COMMENT '评论内容',
  `audit_status` tinyint(4) NOT NULL COMMENT '审核状态：0未审核，1已审核',
  `audit_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '评论时间',
  `modified_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  PRIMARY KEY (`comment_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商品评论表'
******sql******/
// ProductComment 商品评论表
type ProductComment struct {
	CommentID    int       `gorm:"primary_key;column:comment_id;type:int(10) unsigned;not null"` // 评论ID
	ProductID    int       `gorm:"column:product_id;type:int(10) unsigned;not null"`             // 商品ID
	OrderID      int64     `gorm:"column:order_id;type:bigint(20) unsigned;not null"`            // 订单ID
	CustomerID   int       `gorm:"column:customer_id;type:int(10) unsigned;not null"`            // 用户ID
	Title        string    `gorm:"column:title;type:varchar(50);not null"`                       // 评论标题
	Content      string    `gorm:"column:content;type:varchar(300);not null"`                    // 评论内容
	AuditStatus  int8      `gorm:"column:audit_status;type:tinyint(4);not null"`                 // 审核状态：0未审核，1已审核
	AuditTime    time.Time `gorm:"column:audit_time;type:timestamp;not null"`                    // 评论时间
	ModifiedTime time.Time `gorm:"column:modified_time;type:timestamp;not null"`                 // 最后修改时间
}

/******sql******
CREATE TABLE `warehouse_product` (
  `wp_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '商品库存ID',
  `product_id` int(10) unsigned NOT NULL COMMENT '商品ID',
  `w_id` smallint(5) unsigned NOT NULL COMMENT '仓库ID',
  `current_cnt` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '当前商品数量',
  `lock_cnt` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '当前占用数据',
  `in_transit_cnt` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '在途数据',
  `average_cost` decimal(8,2) NOT NULL DEFAULT '0.00' COMMENT '移动加权成本',
  `modified_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  PRIMARY KEY (`wp_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商品库存表'
******sql******/
// WarehouseProduct 商品库存表
type WarehouseProduct struct {
	WpID         int       `gorm:"primary_key;column:wp_id;type:int(10) unsigned;not null"` // 商品库存ID
	ProductID    int       `gorm:"column:product_id;type:int(10) unsigned;not null"`        // 商品ID
	WID          int16     `gorm:"column:w_id;type:smallint(5) unsigned;not null"`          // 仓库ID
	CurrentCnt   int       `gorm:"column:current_cnt;type:int(10) unsigned;not null"`       // 当前商品数量
	LockCnt      int       `gorm:"column:lock_cnt;type:int(10) unsigned;not null"`          // 当前占用数据
	InTransitCnt int       `gorm:"column:in_transit_cnt;type:int(10) unsigned;not null"`    // 在途数据
	AverageCost  float64   `gorm:"column:average_cost;type:decimal(8,2);not null"`          // 移动加权成本
	ModifiedTime time.Time `gorm:"column:modified_time;type:timestamp;not null"`            // 最后修改时间
}

/******sql******
CREATE TABLE `customer_balance_log` (
  `balance_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '余额日志ID',
  `customer_id` int(10) unsigned NOT NULL COMMENT '用户ID',
  `source` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '记录来源：1订单，2退货单',
  `source_sn` int(10) unsigned NOT NULL COMMENT '相关单据ID',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录生成时间',
  `amount` decimal(8,2) NOT NULL DEFAULT '0.00' COMMENT '变动金额',
  PRIMARY KEY (`balance_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户余额变动表'
******sql******/
// CustomerBalanceLog 用户余额变动表
type CustomerBalanceLog struct {
	BalanceID  int       `gorm:"primary_key;column:balance_id;type:int(10) unsigned;not null"` // 余额日志ID
	CustomerID int       `gorm:"column:customer_id;type:int(10) unsigned;not null"`            // 用户ID
	Source     int8      `gorm:"column:source;type:tinyint(3) unsigned;not null"`              // 记录来源：1订单，2退货单
	SourceSn   int       `gorm:"column:source_sn;type:int(10) unsigned;not null"`              // 相关单据ID
	CreateTime time.Time `gorm:"column:create_time;type:timestamp;not null"`                   // 记录生成时间
	Amount     float64   `gorm:"column:amount;type:decimal(8,2);not null"`                     // 变动金额
}

/******sql******
CREATE TABLE `customer_point_log` (
  `point_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '积分日志ID',
  `customer_id` int(10) unsigned NOT NULL COMMENT '用户ID',
  `source` tinyint(3) unsigned NOT NULL COMMENT '积分来源：0订单，1登陆，2活动',
  `refer_number` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '积分来源相关编号',
  `change_point` smallint(6) NOT NULL DEFAULT '0' COMMENT '变更积分数',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '积分日志生成时间',
  PRIMARY KEY (`point_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户积分日志表'
******sql******/
// CustomerPointLog 用户积分日志表
type CustomerPointLog struct {
	PointID     int       `gorm:"primary_key;column:point_id;type:int(10) unsigned;not null"` // 积分日志ID
	CustomerID  int       `gorm:"column:customer_id;type:int(10) unsigned;not null"`          // 用户ID
	Source      int8      `gorm:"column:source;type:tinyint(3) unsigned;not null"`            // 积分来源：0订单，1登陆，2活动
	ReferNumber int       `gorm:"column:refer_number;type:int(10) unsigned;not null"`         // 积分来源相关编号
	ChangePoint int16     `gorm:"column:change_point;type:smallint(6);not null"`              // 变更积分数
	CreateTime  time.Time `gorm:"column:create_time;type:timestamp;not null"`                 // 积分日志生成时间
}

/******sql******
CREATE TABLE `order_cart` (
  `cart_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '购物车ID',
  `customer_id` int(10) unsigned NOT NULL COMMENT '用户ID',
  `product_id` int(10) unsigned NOT NULL COMMENT '商品ID',
  `product_amount` int(11) NOT NULL COMMENT '加入购物车商品数量',
  `price` decimal(8,2) NOT NULL COMMENT '商品价格',
  `add_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '加入购物车时间',
  `modified_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  PRIMARY KEY (`cart_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='购物车表'
******sql******/
// OrderCart 购物车表
type OrderCart struct {
	CartID        int       `gorm:"primary_key;column:cart_id;type:int(10) unsigned;not null"` // 购物车ID
	CustomerID    int       `gorm:"column:customer_id;type:int(10) unsigned;not null"`         // 用户ID
	ProductID     int       `gorm:"column:product_id;type:int(10) unsigned;not null"`          // 商品ID
	ProductAmount int       `gorm:"column:product_amount;type:int(11);not null"`               // 加入购物车商品数量
	Price         float64   `gorm:"column:price;type:decimal(8,2);not null"`                   // 商品价格
	AddTime       time.Time `gorm:"column:add_time;type:timestamp;not null"`                   // 加入购物车时间
	ModifiedTime  time.Time `gorm:"column:modified_time;type:timestamp;not null"`              // 最后修改时间
}

/******sql******
CREATE TABLE `shipping_info` (
  `ship_id` tinyint(3) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `ship_name` varchar(20) NOT NULL COMMENT '物流公司名称',
  `ship_contact` varchar(20) NOT NULL COMMENT '物流公司联系人',
  `telephone` varchar(20) NOT NULL COMMENT '物流公司联系电话',
  `price` decimal(8,2) NOT NULL DEFAULT '0.00' COMMENT '配送价格',
  `modified_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  PRIMARY KEY (`ship_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='物流公司信息表'
******sql******/
// ShippingInfo 物流公司信息表
type ShippingInfo struct {
	ShipID       int8      `gorm:"primary_key;column:ship_id;type:tinyint(3) unsigned;not null"` // 主键ID
	ShipName     string    `gorm:"column:ship_name;type:varchar(20);not null"`                   // 物流公司名称
	ShipContact  string    `gorm:"column:ship_contact;type:varchar(20);not null"`                // 物流公司联系人
	Telephone    string    `gorm:"column:telephone;type:varchar(20);not null"`                   // 物流公司联系电话
	Price        float64   `gorm:"column:price;type:decimal(8,2);not null"`                      // 配送价格
	ModifiedTime time.Time `gorm:"column:modified_time;type:timestamp;not null"`                 // 最后修改时间
}

/******sql******
CREATE TABLE `warehouse_info` (
  `w_id` smallint(5) unsigned NOT NULL AUTO_INCREMENT COMMENT '仓库ID',
  `warehouse_sn` char(5) NOT NULL COMMENT '仓库编码',
  `warehoust_name` varchar(10) NOT NULL COMMENT '仓库名称',
  `warehouse_phone` varchar(20) NOT NULL COMMENT '仓库电话',
  `contact` varchar(10) NOT NULL COMMENT '仓库联系人',
  `province` smallint(6) NOT NULL COMMENT '省',
  `city` smallint(6) NOT NULL COMMENT '市',
  `distrct` smallint(6) NOT NULL COMMENT '区',
  `address` varchar(100) NOT NULL COMMENT '仓库地址',
  `warehouse_status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '仓库状态：0禁用，1启用',
  `modified_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  PRIMARY KEY (`w_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='仓库信息表'
******sql******/
// WarehouseInfo 仓库信息表
type WarehouseInfo struct {
	WID             int16     `gorm:"primary_key;column:w_id;type:smallint(5) unsigned;not null"` // 仓库ID
	WarehouseSn     string    `gorm:"column:warehouse_sn;type:char(5);not null"`                  // 仓库编码
	WarehoustName   string    `gorm:"column:warehoust_name;type:varchar(10);not null"`            // 仓库名称
	WarehousePhone  string    `gorm:"column:warehouse_phone;type:varchar(20);not null"`           // 仓库电话
	Contact         string    `gorm:"column:contact;type:varchar(10);not null"`                   // 仓库联系人
	Province        int16     `gorm:"column:province;type:smallint(6);not null"`                  // 省
	City            int16     `gorm:"column:city;type:smallint(6);not null"`                      // 市
	Distrct         int16     `gorm:"column:distrct;type:smallint(6);not null"`                   // 区
	Address         string    `gorm:"column:address;type:varchar(100);not null"`                  // 仓库地址
	WarehouseStatus int8      `gorm:"column:warehouse_status;type:tinyint(4);not null"`           // 仓库状态：0禁用，1启用
	ModifiedTime    time.Time `gorm:"column:modified_time;type:timestamp;not null"`               // 最后修改时间
}

/******sql******
CREATE TABLE `customer_inf` (
  `customer_inf_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键ID',
  `customer_id` int(10) unsigned NOT NULL COMMENT 'customer_login表的自增ID',
  `customer_name` varchar(20) NOT NULL COMMENT '用户真实姓名',
  `identity_card_type` tinyint(4) NOT NULL DEFAULT '1' COMMENT '证件类型：1 身份证，2 军官证，3 护照',
  `identity_card_no` varchar(20) DEFAULT NULL COMMENT '证件号码',
  `mobile_phone` int(10) unsigned DEFAULT NULL COMMENT '手机号',
  `customer_email` varchar(50) DEFAULT NULL COMMENT '邮箱',
  `gender` char(1) DEFAULT NULL COMMENT '性别',
  `user_point` int(11) NOT NULL DEFAULT '0' COMMENT '用户积分',
  `register_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '注册时间',
  `birthday` datetime DEFAULT NULL COMMENT '会员生日',
  `customer_level` tinyint(4) NOT NULL DEFAULT '1' COMMENT '会员级别：1 普通会员，2 青铜，3白银，4黄金，5钻石',
  `user_money` decimal(8,2) NOT NULL DEFAULT '0.00' COMMENT '用户余额',
  `modified_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  PRIMARY KEY (`customer_inf_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户信息表'
******sql******/
// CustomerInf 用户信息表
type CustomerInf struct {
	CustomerInfID    int       `gorm:"primary_key;column:customer_inf_id;type:int(10) unsigned;not null"` // 自增主键ID
	CustomerID       int       `gorm:"column:customer_id;type:int(10) unsigned;not null"`                 // customer_login表的自增ID
	CustomerName     string    `gorm:"column:customer_name;type:varchar(20);not null"`                    // 用户真实姓名
	IDentityCardType int8      `gorm:"column:identity_card_type;type:tinyint(4);not null"`                // 证件类型：1 身份证，2 军官证，3 护照
	IDentityCardNo   string    `gorm:"column:identity_card_no;type:varchar(20)"`                          // 证件号码
	MobilePhone      int       `gorm:"column:mobile_phone;type:int(10) unsigned"`                         // 手机号
	CustomerEmail    string    `gorm:"column:customer_email;type:varchar(50)"`                            // 邮箱
	Gender           string    `gorm:"column:gender;type:char(1)"`                                        // 性别
	UserPoint        int       `gorm:"column:user_point;type:int(11);not null"`                           // 用户积分
	RegisterTime     time.Time `gorm:"column:register_time;type:timestamp;not null"`                      // 注册时间
	Birthday         time.Time `gorm:"column:birthday;type:datetime"`                                     // 会员生日
	CustomerLevel    int8      `gorm:"column:customer_level;type:tinyint(4);not null"`                    // 会员级别：1 普通会员，2 青铜，3白银，4黄金，5钻石
	UserMoney        float64   `gorm:"column:user_money;type:decimal(8,2);not null"`                      // 用户余额
	ModifiedTime     time.Time `gorm:"column:modified_time;type:timestamp;not null"`                      // 最后修改时间
}
