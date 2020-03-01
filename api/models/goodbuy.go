package models

import (
	"time"
)

// GbAllAccountDetails 平台账户流水
type GbAllAccountDetails struct {
	ID            int     `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	WorksheetCode string  `gorm:"column:worksheet_code;type:char(32);not null" json:"worksheet_code"`      // 工单编码
	BalanceAction float64 `gorm:"column:balance_action;type:decimal(10,2);not null" json:"balance_action"` // 操作金额动作
	ActionTepy    int8    `gorm:"column:action_tepy;type:tinyint(2);not null" json:"action_tepy"`          // 操作类型：1,购买（采购）；2，提现，3，退款，4，冻结,5，解冻，6，充值，7，退款临时存储进账，8，退款临时存储出账
	Status        int8    `gorm:"column:status;type:tinyint(2);not null" json:"status"`                    // 状态：1,正常；2，关闭
	Creattime     int     `gorm:"column:creattime;type:int(10) unsigned;not null" json:"creattime"`        // 创建时间
	Updatatime    int     `gorm:"column:updatatime;type:int(10)" json:"updatatime"`                        // 更新时间
}

type GbMerchantYigroupConfig struct {
	ID            int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	AppName       string    `gorm:"column:app_name;type:varchar(50)" json:"app_name"`                 // 小程序名称
	ShareImage    string    `gorm:"column:share_image;type:varchar(250);not null" json:"share_image"` // 小程序分享图片
	SpringImage   string    `gorm:"column:spring_image;type:varchar(250)" json:"spring_image"`        // 春-头部背景
	SummerImage   string    `gorm:"column:summer_image;type:varchar(250)" json:"summer_image"`        // 夏-头部背景
	AutumnImage   string    `gorm:"column:autumn_image;type:varchar(250)" json:"autumn_image"`        // 秋-头部背景
	WinterImage   string    `gorm:"column:winter_image;type:varchar(250)" json:"winter_image"`        // 冬-头部背景
	FestivalDate  time.Time `gorm:"column:festival_date;type:datetime" json:"festival_date"`          // 节日日期
	FestivalImage string    `gorm:"column:festival_image;type:varchar(250)" json:"festival_image"`    // 节日-头部背景
	ShareTemplate string    `gorm:"column:share_template;type:varchar(250)" json:"share_template"`    // 商品分享图模板
	CashMinMoney  float64   `gorm:"column:cash_min_money;type:decimal(10,0)" json:"cash_min_money"`   // 最小提现金额
	CashDayLimit  int       `gorm:"column:cash_day_limit;type:int(10)" json:"cash_day_limit"`         // 每日允许提现次数
	CreatedAt     int       `gorm:"column:created_at;type:int(10)" json:"created_at"`                 // 创建时间
	UpdatedAt     int       `gorm:"column:updated_at;type:int(10)" json:"updated_at"`                 // 修改时间
}

type GbStoreKeyword struct {
	ID         int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	StoreID    int       `gorm:"column:store_id;type:int(11);not null" json:"store_id"`   // 关键词所属店铺id
	Keyword    string    `gorm:"column:keyword;type:varchar(32);not null" json:"keyword"` // 关键词内容
	Type       int8      `gorm:"column:type;type:tinyint(4);not null" json:"type"`        // 关键词类型：0.普通用户热门关键词1.导购用户热门关键词
	Status     int8      `gorm:"column:status;type:tinyint(4);not null" json:"status"`    // 关键词状态0.显示1.不显示
	IsDel      int8      `gorm:"column:is_del;type:tinyint(2) unsigned;not null" json:"is_del"`
	Createtime time.Time `gorm:"column:createtime;type:timestamp;not null" json:"createtime"`
}

type GbWhUserQualification struct {
	ID           int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	WhUserID     int    `gorm:"column:wh_user_id;type:int(11) unsigned;not null" json:"wh_user_id"`      // 网红用户ID
	Nickname     string `gorm:"column:nickname;type:varchar(600);not null" json:"nickname"`              // 网红昵称
	FansNumber   int64  `gorm:"column:fans_number;type:bigint(20) unsigned;not null" json:"fans_number"` // 粉丝数量
	ScreenShot   string `gorm:"column:screen_shot;type:varchar(800);not null" json:"screen_shot"`        // 粉丝数截图
	ContactName  string `gorm:"column:contact_name;type:varchar(255)" json:"contact_name"`               // 联系人名称
	ContactPhone string `gorm:"column:contact_phone;type:varchar(20)" json:"contact_phone"`              // 联系手机号
	Status       int8   `gorm:"column:status;type:tinyint(2);not null" json:"status"`                    // 资质状态 0 待审核 1审核通过 2审核未通过
	CreatedAt    int    `gorm:"column:created_at;type:int(11);not null" json:"created_at"`
	UpdatedAt    int    `gorm:"column:updated_at;type:int(11);not null" json:"updated_at"`
}

type GbLogOrderChangeBk struct {
	ID          int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	UId         int    `gorm:"column:uid;type:int(11) unsigned;not null" json:"uid"`                 // 操作者ID
	IncrementID string `gorm:"index;column:increment_id;type:char(20);not null" json:"increment_id"` // 订单流水号
	Remarks     string `gorm:"column:remarks;type:varchar(255)" json:"remarks"`                      // 备注
	AfterStatus int    `gorm:"column:after_status;type:int(11);not null" json:"after_status"`        // 处理后状态
	Creattime   int    `gorm:"column:creattime;type:int(10)" json:"creattime"`                       // 创建时间
}

// GbRetailPro 零售商-产品库表
type GbRetailPro struct {
	ID       int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	RetailID int    `gorm:"column:retail_id;type:int(11);not null" json:"retail_id"` // 零售商ID
	ProID    int    `gorm:"index;column:pro_id;type:int(11);not null" json:"pro_id"` // 产品ID(来自商铺产品库  gd_store_pro表)
	Sku      string `gorm:"column:sku;type:char(40);not null" json:"sku"`            // 产品SKU
	Num      int    `gorm:"column:num;type:int(11);not null" json:"num"`             // 销量
	Ctime    int    `gorm:"column:ctime;type:int(11);not null" json:"ctime"`         // 创建时间
	Mtime    int    `gorm:"column:mtime;type:int(11);not null" json:"mtime"`         // 更新时间
	Status   int8   `gorm:"column:status;type:tinyint(1);not null" json:"status"`    // 状态(1:上架    0:下架)
}

type GbStoreBrandCateRelation struct {
	ID        int       `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	Cid       int       `gorm:"index;column:cid;type:int(11) unsigned;not null" json:"cid"` // 品牌分类id
	Bid       int       `gorm:"index;column:bid;type:int(11) unsigned;not null" json:"bid"` // 店铺品牌ID
	Sort      int       `gorm:"column:sort;type:int(11);not null" json:"sort"`              // 分类排序
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;not null" json:"updated_at"`
}

type GbSupLogisticsOrder struct {
	ID              int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	Esn             string `gorm:"column:esn;type:varchar(255);not null" json:"esn"`                  // 电子面单号
	FacilitatorID   int    `gorm:"column:facilitator_id;type:int(11);not null" json:"facilitator_id"` // 发货网点
	Time            string `gorm:"column:time;type:varchar(30);not null" json:"time"`                 // 获取时间
	Order           string `gorm:"column:order;type:varchar(255);not null" json:"order"`              // 订单号
	EsnStatus       int    `gorm:"column:esn_status;type:int(1);not null" json:"esn_status"`          // 单号状态 1:有效 2：已回收3：取消
	Start           int    `gorm:"column:start;type:int(1);not null" json:"start"`                    // 运输状态1-未打印
	SupID           int    `gorm:"column:sup_id;type:int(11);not null" json:"sup_id"`                 // 用户id
	OrderID         int    `gorm:"column:order_id;type:int(11)" json:"order_id"`                      // 订单id
	ZiOrder         string `gorm:"column:zi_order;type:varchar(255)" json:"zi_order"`                 // 子订单号
	PrintData       string `gorm:"column:print_data;type:text" json:"print_data"`                     // 面单打印数据
	CpID            int    `gorm:"column:cp_id;type:int(11)" json:"cp_id"`                            // 快递公司id
	CustomPrintData string `gorm:"column:custom_print_data;type:text" json:"custom_print_data"`       // 自定义打印数据
}

// GbMerchantGroup 商户拼团列表
type GbMerchantGroup struct {
	ID                int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Title             string    `gorm:"column:title;type:varchar(128);not null" json:"title"`                            // 活动名称
	ActivitySellpoint string    `gorm:"column:activity_sellpoint;type:varchar(1000);not null" json:"activity_sellpoint"` // 活动卖点
	Starttime         time.Time `gorm:"column:starttime;type:datetime" json:"starttime"`                                 // 活动开始时间
	Endtime           time.Time `gorm:"column:endtime;type:datetime" json:"endtime"`                                     // 活动结束时间
	SellerID          int       `gorm:"column:seller_id;type:int(11)" json:"seller_id"`
	StoreID           int       `gorm:"column:store_id;type:int(11);not null" json:"store_id"`       // 拼团所属店铺
	Status            int       `gorm:"column:status;type:int(11)" json:"status"`                    // 拼团状态：1.进行中2.已结束3.未开始4.已成团
	IsDel             int       `gorm:"column:is_del;type:int(11)" json:"is_del"`                    // 逻辑删除状态
	ActivityMod       int8      `gorm:"column:activity_mod;type:tinyint(4)" json:"activity_mod"`     // 活动执行板块
	Createtime        time.Time `gorm:"column:createtime;type:timestamp;not null" json:"createtime"` // 拼团创建时间
}

type GbOrderMergeInvoice struct {
	ID                  int     `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`                       // 主键ID
	IncrementID         string  `gorm:"index;column:increment_id;type:char(20);not null" json:"increment_id"`                // 总订单流水号:123456789451245
	PayJifen            float64 `gorm:"column:pay_jifen;type:decimal(10,2);not null" json:"pay_jifen"`                       // 用户花费积分
	PayMoney            float64 `gorm:"column:pay_money;type:decimal(10,2);not null" json:"pay_money"`                       // 用户实际支付金额
	Commision           float64 `gorm:"column:commision;type:decimal(10,2);not null" json:"commision"`                       // 用户优惠券金额
	Freight             float64 `gorm:"column:freight;type:decimal(10,2);not null" json:"freight"`                           // 用户订单商品运费总价
	Price               float64 `gorm:"column:price;type:decimal(10,2);not null" json:"price"`                               // 商品总价，不含邮费
	CountPrice          float64 `gorm:"column:count_price;type:decimal(10,2);not null" json:"count_price"`                   // 用户订单商品总价
	BatchValue          int8    `gorm:"column:batch_value;type:smallint(2) unsigned" json:"batch_value"`                     // 1表示可以开发票，0表示不可以开发票
	PurchaseIncrementID string  `gorm:"column:purchase_increment_id;type:char(20);not null" json:"purchase_increment_id"`    // 平台采购订单号
	PurchasePrice       float64 `gorm:"column:purchase_price;type:decimal(10,2);not null" json:"purchase_price"`             // 采购总价，不含邮费
	PurchaseFreight     float64 `gorm:"column:purchase_freight;type:decimal(10,2);not null" json:"purchase_freight"`         // 采购运费
	CountPurchasePrice  float64 `gorm:"column:count_purchase_price;type:decimal(10,2);not null" json:"count_purchase_price"` // 采购订单商品总价，包含邮费
	OrderQuantity       int     `gorm:"column:order_quantity;type:int(11) unsigned;not null" json:"order_quantity"`          // 订单数量
	GoodsQuantity       int     `gorm:"column:goods_quantity;type:int(11) unsigned;not null" json:"goods_quantity"`          // 商品数量
	Creattime           int     `gorm:"index;column:creattime;type:int(10) unsigned;not null" json:"creattime"`              // 创建时间
	Updatetime          int     `gorm:"column:updatetime;type:int(10) unsigned;not null" json:"updatetime"`                  // 更新时间
}

// GbProComment 商品评论
type GbProComment struct {
	ID            int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	ProID         int    `gorm:"column:pro_id;type:int(11)" json:"pro_id"`                         // 平台商品ID
	ProBarCode    string `gorm:"column:pro_bar_code;type:varchar(20)" json:"pro_bar_code"`         // 平台商品编码 ENA13码
	ProStandardID int    `gorm:"column:pro_standard_id;type:int(11)" json:"pro_standard_id"`       // 平台商品SKUID
	UserID        int    `gorm:"column:user_id;type:int(11)" json:"user_id"`                       // 评论人ID
	UserNickname  string `gorm:"column:user_nickname;type:varchar(50)" json:"user_nickname"`       // 用户昵称
	UserIP        string `gorm:"column:user_ip;type:varchar(20)" json:"user_ip"`                   // 评论提交IP地址
	Content       string `gorm:"column:content;type:varchar(250)" json:"content"`                  // 评论内容
	StarNum       int8   `gorm:"column:star_num;type:tinyint(1) unsigned" json:"star_num"`         // 用户评论等级，10非常好 7很好 5一般 1差 0 糟透了
	IsAnon        int8   `gorm:"column:is_anon;type:smallint(2) unsigned;not null" json:"is_anon"` // 是否匿名评论
	IsPic         int8   `gorm:"column:is_pic;type:smallint(2) unsigned" json:"is_pic"`            // 评论如果上传了图片请在这里标记 0无图 1有图
	Ctime         int    `gorm:"column:ctime;type:int(11)" json:"ctime"`                           // 评论时间
	FromType      int8   `gorm:"column:from_type;type:tinyint(1)" json:"from_type"`                // 评论来源  0 本站 ，1 天猫  2京东
	Status        int8   `gorm:"column:status;type:tinyint(2)" json:"status"`                      // 显示状态 1 显示 0 隐藏
	IsDel         int8   `gorm:"column:is_del;type:tinyint(1)" json:"is_del"`
}

type GbStoreBrandCate struct {
	ID        int       `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	Name      string    `gorm:"column:name;type:varchar(255);not null" json:"name"`                   // 分类名称
	StoreID   int       `gorm:"index;column:store_id;type:int(11) unsigned;not null" json:"store_id"` // 店铺ID
	IsDel     int8      `gorm:"column:is_del;type:tinyint(2);not null" json:"is_del"`                 // 删除状态 0 为删除 1已删除
	Sort      int       `gorm:"column:sort;type:int(11);not null" json:"sort"`                        // 排序 升序
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;not null" json:"updated_at"`
}

type GbSupAccountDetailsLog struct {
	ID       int     `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"` // 编号
	Name     int     `gorm:"column:name;type:int(2)" json:"name"`                           // 名称 1表示前款 2表示尾款
	OrderID  string  `gorm:"column:order_id;type:varchar(32);not null" json:"order_id"`     // 订单编号
	Type     int8    `gorm:"column:type;type:tinyint(3);not null" json:"type"`              // 明细类型（1 收入 2 退款）
	CreateAt int     `gorm:"column:create_at;type:int(10);not null" json:"create_at"`       // 创建时间
	Money    float64 `gorm:"column:money;type:decimal(19,4);not null" json:"money"`         // 金额
	IsDel    int8    `gorm:"column:is_del;type:tinyint(1) unsigned" json:"is_del"`          // 是否删除 预留误删 陈磊
	SupUId   int     `gorm:"column:sup_uid;type:int(10);not null" json:"sup_uid"`           // 供应商编号
}

type GbAdminLog struct {
	ID          int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	AdminID     int64  `gorm:"index:idx_user_id;column:admin_id;type:bigint(20);not null" json:"admin_id"` // 用户ID
	Module      int8   `gorm:"index:idx_user_id;column:module;type:tinyint(1);not null" json:"module"`     // 后台模块 1平台端 2商家端 3店铺端 4供应商
	URL         string `gorm:"column:url;type:varchar(255);not null" json:"url"`                           // 访问的url
	Route       string `gorm:"index:idx_user_id;column:route;type:varchar(100)" json:"route"`              // 请求跟由
	QueryParams string `gorm:"column:query_params;type:longtext" json:"query_params"`                      // get和post参数
	UserAgent   string `gorm:"column:user_agent;type:varchar(255);not null" json:"user_agent"`             // 访问User-Agent
	IP          string `gorm:"column:ip;type:varchar(32);not null" json:"ip"`                              // 访问ip
	Method      string `gorm:"column:method;type:varchar(20)" json:"method"`                               // 提交方式
	Note        string `gorm:"column:note;type:varchar(1000);not null" json:"note"`                        // 备注
	CreatedAt   int    `gorm:"column:created_at;type:int(10) unsigned;not null" json:"created_at"`         // 创建时间
	UpdatedAt   int    `gorm:"column:updated_at;type:int(10) unsigned;not null" json:"updated_at"`         // 更新时间
}

type GbMerchantBargainRate struct {
	ID        int  `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	BargainID int  `gorm:"column:bargain_id;type:int(11)" json:"bargain_id"`
	RateNum   int8 `gorm:"column:rate_num;type:tinyint(2)" json:"rate_num"`
}

// GbStoreProductCate 商品分类
type GbStoreProductCate struct {
	ID         int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	StoreID    int       `gorm:"column:store_id;type:int(11);not null" json:"store_id"`  // 店铺ID
	Pid        int       `gorm:"index;column:pid;type:int(11);not null" json:"pid"`      // 上级ID
	Name       string    `gorm:"column:name;type:varchar(30);not null" json:"name"`      // 类别名称
	Logo       string    `gorm:"column:logo;type:varchar(200)" json:"logo"`              // 图片
	Sort       int       `gorm:"column:sort;type:int(11);not null" json:"sort"`          // 排序
	Display    int8      `gorm:"column:display;type:tinyint(1);not null" json:"display"` // 是否默认展示：0否，1是
	Status     int8      `gorm:"column:status;type:tinyint(1);not null" json:"status"`   // 状态：1可用，0不可用，2特价
	Type       int8      `gorm:"column:type;type:tinyint(1);not null" json:"type"`       // 分类类型：1手动定义类型，2系统自动，3商家追加
	CreateDate time.Time `gorm:"column:create_date;type:datetime" json:"create_date"`    // 创建时间
	IsDel      int8      `gorm:"column:is_del;type:tinyint(1);not null" json:"is_del"`   // 是否删除  1：是   0：否
	TempOID    int       `gorm:"column:temp_o_id;type:int(11)" json:"temp_o_id"`         // 导入之前的记录的ID号
}

type GbWhStore struct {
	StoreID          int       `gorm:"primary_key;column:store_id;type:int(11) unsigned;not null" json:"store_id"`
	StoreName        string    `gorm:"column:store_name;type:char(30);not null" json:"store_name"`              // 店铺名称
	UId              int64     `gorm:"column:uid;type:bigint(20) unsigned;not null" json:"uid"`                 // 用户ID
	PlatformType     string    `gorm:"column:platform_type;type:varchar(12);not null" json:"platform_type"`     // 所属平台编号：taobao淘宝 jingdong京东
	PlatformUserID   string    `gorm:"column:platform_user_id;type:varchar(100)" json:"platform_user_id"`       // 平台店铺ID
	PlatformUserName string    `gorm:"column:platform_user_name;type:varchar(255)" json:"platform_user_name"`   // 平台店铺用户名
	CreateTime       int       `gorm:"column:create_time;type:int(10) unsigned;not null" json:"create_time"`    // 创建时间
	UpdateTime       int       `gorm:"column:update_time;type:int(10) unsigned" json:"update_time"`             // 修改时间
	DeadTime         int       `gorm:"column:dead_time;type:int(10) unsigned" json:"dead_time"`                 // 到期时间
	Status           int       `gorm:"column:status;type:int(1) unsigned" json:"status"`                        // 店铺状态：0关闭 1正常 2授权完成
	AccessToken      string    `gorm:"column:access_token;type:varchar(96)" json:"access_token"`                // 商店的sessionKey或Access token
	RefreshToken     string    `gorm:"column:refresh_token;type:varchar(100)" json:"refresh_token"`             // Refresh token，可用来刷新access_token
	ExpiresIn        int       `gorm:"column:expires_in;type:int(10) unsigned" json:"expires_in"`               // Access token过期时间
	ReExpiresIn      int       `gorm:"column:re_expires_in;type:int(10) unsigned" json:"re_expires_in"`         // Refresh token过期时间
	R1ExpiresIn      int       `gorm:"column:r1_expires_in;type:int(10) unsigned" json:"r1_expires_in"`         // R1级别API或字段的访问过期时间
	R2ExpiresIn      int       `gorm:"column:r2_expires_in;type:int(10) unsigned" json:"r2_expires_in"`         // R2级别API或字段的访问过期时间
	W1ExpiresIn      int       `gorm:"column:w1_expires_in;type:int(10) unsigned" json:"w1_expires_in"`         // W1级别API或字段的访问过期时间
	W2ExpiresIn      int       `gorm:"column:w2_expires_in;type:int(10) unsigned" json:"w2_expires_in"`         // W2级别API或字段的访问过期时间
	LastModifyTime   time.Time `gorm:"column:last_modify_time;type:timestamp;not null" json:"last_modify_time"` // 记录最后更新时间
}

type GbSupProductEditLog struct {
	ID         int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	BeforeData string    `gorm:"column:before_data;type:text;not null" json:"before_data"`     // 变更前json数据
	SupProid   int       `gorm:"column:sup_proid;type:int(11);not null" json:"sup_proid"`      // 变更供应商产品id
	AfterData  string    `gorm:"column:after_data;type:text;not null" json:"after_data"`       // 变更后json数据
	SubmitTime time.Time `gorm:"column:submit_time;type:datetime;not null" json:"submit_time"` // 变更审核时间
}

// GbAkcActivity 爱库存活动表
type GbAkcActivity struct {
	ID             int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Code           string    `gorm:"column:code;type:varchar(150);not null" json:"code"`               // 爱库存活动编码
	StartDate      time.Time `gorm:"column:startDate;type:datetime;not null" json:"startDate"`         // 开始时间
	EndDate        time.Time `gorm:"index;column:endDate;type:datetime;not null" json:"endDate"`       // 结束时间
	Name           string    `gorm:"column:name;type:varchar(100);not null" json:"name"`               // 活动名称
	Description    string    `gorm:"column:description;type:varchar(100);not null" json:"description"` // 活动描述
	Brand          string    `gorm:"column:brand;type:varchar(20);not null" json:"brand"`              // 活动品牌
	Category       string    `gorm:"column:category;type:varchar(20)" json:"category"`                 // 分类
	StatementByDay int       `gorm:"column:statementByDay;type:int(1);not null" json:"statementByDay"` // 结单标识
	Creattime      time.Time `gorm:"column:creattime;type:datetime;not null" json:"creattime"`         // 创建时间
	IsDel          int       `gorm:"index;column:is_del;type:int(1)" json:"is_del"`                    // 1代表没有过期，0代表已经过期
	Banner         string    `gorm:"column:banner;type:varchar(100)" json:"banner"`                    // 活动图片
	BrandLogoURL   string    `gorm:"column:brandLogoUrl;type:varchar(100)" json:"brandLogoUrl"`        // 品牌图标地址
	DeliverTime    time.Time `gorm:"column:deliverTime;type:datetime" json:"deliverTime"`              // 预计发货时间
	InvalidAreas   string    `gorm:"column:invalidAreas;type:text" json:"invalidAreas"`                // 偏远不发货区域列表
}

type GbErpSetting struct {
	ID       int       `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	Key      string    `gorm:"column:key;type:varchar(512);not null" json:"key"`              // 设置KEY
	Value    string    `gorm:"column:value;type:varchar(1024);not null" json:"value"`         // 设置value
	IsDel    int8      `gorm:"column:is_del;type:tinyint(2) unsigned;not null" json:"is_del"` // 是否禁用 0未禁用 1禁用
	UpdateAt time.Time `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	CreateAt time.Time `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
}

type GbOrderMergeInvoiceIndex struct {
	ID             int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`          // 主键ID
	MergeInvoiceID int    `gorm:"column:merge_invoice_id;type:int(11);not null" json:"merge_invoice_id"`  // 表order_merge_invoice的id
	ChildrenID     int    `gorm:"index;column:children_id;type:int(11);not null" json:"children_id"`      // 表order_children表的id
	InvoiceType    int8   `gorm:"column:invoice_type;type:tinyint(4);not null" json:"invoice_type"`       // 发票号类型，1表示普票，2表示增值票
	InvoiceNumber  string `gorm:"column:invoice_number;type:varchar(256);not null" json:"invoice_number"` // 发票号，由财务系统导入
	Creattime      int    `gorm:"index;column:creattime;type:int(10) unsigned;not null" json:"creattime"` // 创建时间
	Updatetime     int    `gorm:"column:updatetime;type:int(10) unsigned;not null" json:"updatetime"`     // 更新时间
}

// GbOrderSetting 订单限时触发配置
type GbOrderSetting struct {
	ID                     int  `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	SeckillOrderEtime      int8 `gorm:"column:seckill_order_etime;type:smallint(2) unsigned;not null" json:"seckill_order_etime"`             // 秒杀订单限时设置（分钟）
	GeneralOrderEtime      int8 `gorm:"column:general_order_etime;type:smallint(2) unsigned;not null" json:"general_order_etime"`             // 正常订单限时设置（分钟）
	WxappOrderEtime        int8 `gorm:"column:wxapp_order_etime;type:smallint(2) unsigned;not null" json:"wxapp_order_etime"`                 // 小程序订单限时设置（分钟）
	PackageOrderEtime      int8 `gorm:"column:package_order_etime;type:smallint(2) unsigned;not null" json:"package_order_etime"`             // 已发货订单限时设置（天） 超过则订单自动完成
	CloseOrderEtime        int8 `gorm:"column:close_order_etime;type:smallint(2) unsigned;not null" json:"close_order_etime"`                 // 订单完成（天数）超过该设置天数则不能申请售后
	CloseOrderCommentEtime int8 `gorm:"column:close_order_comment_etime;type:smallint(2) unsigned;not null" json:"close_order_comment_etime"` // 订单完成(天数)，超过该设置天数则自动五星好评
	AutoPurchase           int8 `gorm:"column:auto_purchase;type:tinyint(1)" json:"auto_purchase"`                                            // ERP平台端订单自动采购，0为手动采购，1为自动采购
}

type GbStoreApply struct {
	ID             int       `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	CheckUId       int       `gorm:"column:check_uid;type:int(11) unsigned;not null" json:"check_uid"`      // 审核人
	CheckType      int8      `gorm:"column:check_type;type:tinyint(2) unsigned;not null" json:"check_type"` // 0 入驻 1变更
	StoreID        int       `gorm:"column:store_id;type:int(11) unsigned;not null" json:"store_id"`
	CheckNote      string    `gorm:"column:check_note;type:varchar(512);not null" json:"check_note"`            // 审核备注
	CheckStatus    int8      `gorm:"column:check_status;type:tinyint(2) unsigned;not null" json:"check_status"` // 1通过 2不通过
	ApplyContent   string    `gorm:"column:apply_content;type:text" json:"apply_content"`                       // 申请内容 serialize
	CreateTime     time.Time `gorm:"column:create_time;type:datetime;not null" json:"create_time"`
	UpdateTime     time.Time `gorm:"column:update_time;type:datetime;not null" json:"update_time"`
	IsDel          int8      `gorm:"column:is_del;type:tinyint(2) unsigned;not null" json:"is_del"` // 软删除
	ChangeUserInfo string    `gorm:"column:change_user_info;type:text" json:"change_user_info"`     // 申请变更人
}

type GbSupLogisticsTemplate struct {
	ID                  int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	Name                string `gorm:"column:name;type:varchar(255);not null" json:"name"`                                   // 模板名称
	FacilitatorID       int    `gorm:"column:facilitator_id;type:int(10);not null" json:"facilitator_id"`                    // 快递公司id
	TemplateType        int    `gorm:"column:template_type;type:int(1)" json:"template_type"`                                // 模板规格id
	Type                int    `gorm:"column:type;type:int(1);not null" json:"type"`                                         // 模板类型1-标准 2-自定义
	StandardTemplateURL string `gorm:"column:standard_template_url;type:varchar(100);not null" json:"standard_template_url"` // 标准模板url
	CustomMappingID     int    `gorm:"column:custom_mapping_id;type:int(10)" json:"custom_mapping_id"`                       // 自定义区映射id,用来获取自定义区url
	Text                string `gorm:"column:text;type:text" json:"text"`                                                    // 备注
	Time                string `gorm:"column:time;type:varchar(26)" json:"time"`                                             // 创建时间
	CustomKey           string `gorm:"column:custom_key;type:varchar(20)" json:"custom_key"`                                 // 模板自定义字段
	CustomURL           string `gorm:"column:custom_url;type:varchar(100)" json:"custom_url"`                                // 自定义区url
	StandardType        int    `gorm:"column:standard_type;type:int(10)" json:"standard_type"`                               // 菜鸟返回模板类型id,方便后续添加别的标准模板
	SupID               int    `gorm:"column:sup_id;type:int(11)" json:"sup_id"`                                             // 供应商id,标准模板属于共享的0表示
	UpdateTime          string `gorm:"column:update_time;type:varchar(26)" json:"update_time"`                               // 更新时间
}

// GbMerchantShareModelItem 分成模式的具体分成结构
type GbMerchantShareModelItem struct {
	ID          int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Pid         int    `gorm:"column:pid;type:int(11) unsigned" json:"pid"`  // 父类ID
	ModelID     int    `gorm:"column:model_id;type:int(11)" json:"model_id"` // 所属的分成模式ID
	DimensionID int    `gorm:"column:dimension_id;type:int(11)" json:"dimension_id"`
	Name        string `gorm:"column:name;type:varchar(20)" json:"name"`    // 分成名称
	Ctime       int    `gorm:"column:ctime;type:int(11)" json:"ctime"`      // 创建时间
	Mtime       int    `gorm:"column:mtime;type:int(11)" json:"mtime"`      // 修改时间
	IsDel       int8   `gorm:"column:is_del;type:tinyint(1)" json:"is_del"` // 删除状态 0未删除，1已删除
	Path        string `gorm:"column:path;type:varchar(150)" json:"path"`   // 路径
}

type GbPayment struct {
	ID        int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Name      string `gorm:"column:name;type:varchar(250);not null" json:"name"`                 // 支付方式
	Code      string `gorm:"column:code;type:varchar(20);not null" json:"code"`                  // 支付编码（维一） 如 wxapp
	MchID     string `gorm:"column:mch_id;type:varchar(250);not null" json:"mch_id"`             // 商户号
	AppKey    string `gorm:"column:app_key;type:varchar(250);not null" json:"app_key"`           // APPKEY
	AppID     string `gorm:"column:app_id;type:varchar(250);not null" json:"app_id"`             // APPID
	AppSecret string `gorm:"column:app_secret;type:varchar(250);not null" json:"app_secret"`     // 公众帐号secert
	Sort      int    `gorm:"column:sort;type:int(11)" json:"sort"`                               // 排序，倒序
	Status    int8   `gorm:"column:status;type:tinyint(1)" json:"status"`                        // 预留，暂未启用 是否启用 0为未启用 1为启用
	Image     string `gorm:"column:image;type:varchar(250);not null" json:"image"`               // 图标
	CreatedAt int    `gorm:"column:created_at;type:int(10) unsigned;not null" json:"created_at"` // 创建时间
	UpdatedAt int    `gorm:"column:updated_at;type:int(10) unsigned;not null" json:"updated_at"` // 更新时间
}

// GbRetailWithdraw 提现记录表
type GbRetailWithdraw struct {
	ID       int     `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	StoreID  int     `gorm:"index;column:store_id;type:int(11);not null" json:"store_id"`   // 店铺ID
	RetailID int     `gorm:"column:retail_id;type:int(11);not null" json:"retail_id"`       // 零售商ID
	Code     string  `gorm:"column:code;type:varchar(60);not null" json:"code"`             // 提现编号
	TotalFee float64 `gorm:"column:total_fee;type:decimal(10,2);not null" json:"total_fee"` // 申请提现金额
	Mode     int8    `gorm:"column:mode;type:smallint(1);not null" json:"mode"`             // 结算方式(1:银行卡   2:支付宝   3:微信)
	Account  string  `gorm:"column:account;type:varchar(60);not null" json:"account"`       // 结算账户
	Ctime    int64   `gorm:"column:ctime;type:bigint(20)" json:"ctime"`                     // 申请时间
	Status   int8    `gorm:"column:status;type:smallint(6)" json:"status"`                  // 申请状态   0:申请中待审核   1:申请通过   2:申请驳回
	Etime    int64   `gorm:"column:etime;type:bigint(20)" json:"etime"`                     // 审核时间
	Ereason  string  `gorm:"column:ereason;type:varchar(255)" json:"ereason"`               // 驳回原因
	EadminID int     `gorm:"index;column:eadmin_id;type:int(11)" json:"eadmin_id"`          // 审核人ID
}

// GbSupAccount 供应商账户资金流水
type GbSupAccount struct {
	ID                  int     `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	Code                string  `gorm:"column:code;type:char(20)" json:"code"`                                      // 流水编号（YYMMDD+）
	SupplierCode        string  `gorm:"column:supplier_code;type:char(20);not null" json:"supplier_code"`           // 供应商编码
	SupplierID          string  `gorm:"column:supplier_id;type:char(20);not null" json:"supplier_id"`               // 供应商ID
	IncrementID         string  `gorm:"column:increment_id;type:char(15)" json:"increment_id"`                      // 订单流水号
	ChildrenIncrementID string  `gorm:"column:children_increment_id;type:char(20)" json:"children_increment_id"`    // 子订单流水号
	GoodsPrice          float64 `gorm:"column:goods_price;type:decimal(10,2) unsigned;not null" json:"goods_price"` // 商品总价
	Freight             float64 `gorm:"column:freight;type:decimal(10,2) unsigned;not null" json:"freight"`         // 运费总价
	TotalPrice          float64 `gorm:"column:total_price;type:decimal(10,2) unsigned;not null" json:"total_price"` // 订单总价（商品总价+运费总价）
	Action              string  `gorm:"column:action;type:char(20)" json:"action"`                                  // 流水动作（+订单总价、—订单总价）
	Remark              string  `gorm:"column:remark;type:varchar(250)" json:"remark"`                              // 备注
	Creattime           int     `gorm:"column:creattime;type:int(10) unsigned;not null" json:"creattime"`           // 创建时间
}

type GbSupLogisticsAddress struct {
	ID          int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	Country     int    `gorm:"column:country;type:int(10)" json:"country"`                 // 国
	Province    int    `gorm:"column:province;type:int(10)" json:"province"`               // 对应的省id-不需要
	City        int    `gorm:"column:city;type:int(10)" json:"city"`                       // 对应的市id-不需要
	Area        int    `gorm:"column:area;type:int(10)" json:"area"`                       // 对应的区id-不需要
	Detailed    string `gorm:"column:detailed;type:varchar(255);not null" json:"detailed"` // 地址详情
	Time        string `gorm:"column:time;type:varchar(20);not null" json:"time"`          // 创建时间
	DetailedCon string `gorm:"column:detailed_con;type:varchar(255)" json:"detailed_con"`  // 拼接后的数据用于显示
	Status      int8   `gorm:"column:status;type:tinyint(1)" json:"status"`                // 1-开通 0-审核
}

type GbAdminRole struct {
	ID        int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`      // ID
	Name      string `gorm:"column:name;type:varchar(50);not null" json:"name"`                  // 名称
	Module    int8   `gorm:"column:module;type:tinyint(1);not null" json:"module"`               // 后台模块 1平台端 2商家端 3店铺端 4供应商 5官网
	Status    int8   `gorm:"column:status;type:tinyint(1);not null" json:"status"`               // 状态：1有效、0无效
	CreatedAt int    `gorm:"column:created_at;type:int(10) unsigned;not null" json:"created_at"` // 创建时间
	UpdatedAt int    `gorm:"column:updated_at;type:int(10) unsigned;not null" json:"updated_at"` // 更新时间
}

// GbAreaProvince 省份分区表
type GbAreaProvince struct {
	ID   int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Name string `gorm:"column:name;type:varchar(20);not null" json:"name"`
}

type GbMerchantShareModelTemplateItem struct {
	ID               int `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	TemplateID       int `gorm:"column:template_id;type:int(11)" json:"template_id"`                 // 所属模板ID
	ShareModelItemID int `gorm:"column:share_model_item_id;type:int(11)" json:"share_model_item_id"` // 分成模式维度ID，对应 gb_merchant_share_model_item.id
	VN               int `gorm:"column:v_n;type:int(11)" json:"v_n"`                                 // 分成比例数值，1-100
}

// GbStore 店铺表
type GbStore struct {
	ID                      int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	SellerID                int    `gorm:"column:seller_id;type:int(11);not null" json:"seller_id"`              // 商家ID
	Password                string `gorm:"column:password;type:varchar(100);not null" json:"password"`           // 登录密码  sha加密
	Code                    string `gorm:"index;column:code;type:char(80)" json:"code"`                          // 加密串
	StoreCode               string `gorm:"column:store_code;type:varchar(20)" json:"store_code"`                 // 小程序店铺识别码
	AdminMobile             string `gorm:"column:admin_mobile;type:varchar(50);not null" json:"admin_mobile"`    // 店长账号  手机号
	StoreName               string `gorm:"column:store_name;type:varchar(100);not null" json:"store_name"`       // 店铺名称
	StoreType               int8   `gorm:"column:store_type;type:tinyint(2);not null" json:"store_type"`         // 店铺类型(1:自营店  2:专营店  3:品牌店   4:连锁店)
	DeviceDesc              string `gorm:"column:device_desc;type:varchar(255);not null" json:"device_desc"`     // 终端描述
	LogoURL                 string `gorm:"column:logo_url;type:varchar(255);not null" json:"logo_url"`           // 店铺logo地址
	BisnessScope            string `gorm:"column:bisness_scope;type:varchar(250);not null" json:"bisness_scope"` // 主营范围
	Tag                     string `gorm:"column:tag;type:varchar(255);not null" json:"tag"`                     // 店铺标签
	ShareModelID            int    `gorm:"column:share_model_id;type:int(11)" json:"share_model_id"`             // 分成模式ID
	ShareTemplateID         int    `gorm:"column:share_template_id;type:int(11)" json:"share_template_id"`       // 分成模板ID
	StartTime               string `gorm:"column:start_time;type:varchar(50);not null" json:"start_time"`        // 合约开始时间
	EndTime                 string `gorm:"column:end_time;type:varchar(50);not null" json:"end_time"`            // 合约结束时间
	IsDel                   int8   `gorm:"column:is_del;type:tinyint(1) unsigned;not null" json:"is_del"`        // 删除状态 1删除，0未删除
	SectionID               int    `gorm:"column:section_id;type:int(11)" json:"section_id"`
	Status                  int8   `gorm:"column:status;type:tinyint(1);not null" json:"status"`                                           // 店铺状态 1 开启，0 关闭
	Ctime                   int    `gorm:"column:ctime;type:int(11);not null" json:"ctime"`                                                // 创建时间
	UpdatedAt               int    `gorm:"column:updated_at;type:int(11)" json:"updated_at"`                                               // 更新时间
	PartnerName             string `gorm:"column:partner_name;type:varchar(255);not null" json:"partner_name"`                             // 合伙人姓名
	PartnerContactsMethod   string `gorm:"column:partner_contacts_method;type:varchar(500);not null" json:"partner_contacts_method"`       // 合伙人联系方式
	PartnerActivePlatform   int8   `gorm:"column:partner_active_platform;type:tinyint(2);not null" json:"partner_active_platform"`         // 合伙人活跃平台
	PartnerActivePlatformID string `gorm:"column:partner_active_platform_id;type:varchar(255);not null" json:"partner_active_platform_id"` // 合伙人活跃平台账号id
	PartnerFansNumber       int    `gorm:"column:partner_fans_number;type:int(11) unsigned;not null" json:"partner_fans_number"`           // 合伙人粉丝数量
	PartnerField            int8   `gorm:"column:partner_field;type:tinyint(2);not null" json:"partner_field"`                             // 合伙人所属领域
	AddCheck                int8   `gorm:"column:add_check;type:tinyint(2);not null" json:"add_check"`                                     // 入驻审核状态 0 未审核 1 通过 2不通过
	RegPhone                string `gorm:"column:reg_phone;type:varchar(20);not null" json:"reg_phone"`                                    // 注册手机号
	AccessToken             string `gorm:"column:access_token;type:varchar(255);not null" json:"access_token"`                             // 接口身份凭证
	AccessTokenExpire       int    `gorm:"column:access_token_expire;type:int(11) unsigned;not null" json:"access_token_expire"`           // access_token 失效时间
}

type GbUserBindRecord struct {
	ID           int       `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	UserID       int       `gorm:"column:user_id;type:int(11) unsigned;not null" json:"user_id"` // 本系统用户ID
	Phone        string    `gorm:"column:phone;type:varchar(11);not null" json:"phone"`
	SourceUserID int       `gorm:"column:source_user_id;type:int(11) unsigned;not null" json:"source_user_id"` // 外部系统UID
	Source       string    `gorm:"column:source;type:varchar(255);not null" json:"source"`                     // 绑定来源
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime;not null" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;type:datetime;not null" json:"updated_at"`
}

type GbCollectArticle struct {
	ID        int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	SiteID    int    `gorm:"column:site_id;type:int(11)" json:"site_id"`       // 站点ID
	KeywordID int    `gorm:"column:keyword_id;type:int(11)" json:"keyword_id"` // 关键词ID
	Title     string `gorm:"column:title;type:varchar(255);not null" json:"title"`
	Img       string `gorm:"column:img;type:varchar(255)" json:"img"`              // 列表主图
	Status    int8   `gorm:"column:status;type:tinyint(1);not null" json:"status"` // 0未发布  1已发布
	IsDel     int8   `gorm:"column:is_del;type:tinyint(1);not null" json:"is_del"`
	CreatedAt int    `gorm:"column:created_at;type:int(11);not null" json:"created_at"`
	UpdatedAt int    `gorm:"column:updated_at;type:int(11);not null" json:"updated_at"`
	PublishAt int    `gorm:"column:publish_at;type:int(11);not null" json:"publish_at"` // 发布时间
	IsImg     int8   `gorm:"column:is_img;type:tinyint(1);not null" json:"is_img"`      // 是否更新图片  0未替换  1已替换
}

// GbMerchantFrontendBargainRecord 砍价专区-帮砍价人记录表
type GbMerchantFrontendBargainRecord struct {
	ID                int     `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	UUID              int     `gorm:"column:uuid;type:int(11) unsigned;not null" json:"uuid"`                               // 砍价人UID
	FrontendBargainID int     `gorm:"column:frontend_bargain_id;type:int(11) unsigned;not null" json:"frontend_bargain_id"` // H5端发起砍价表主键ID
	KillRatioID       int     `gorm:"column:kill_ratio_id;type:int(11) unsigned;not null" json:"kill_ratio_id"`             // 砍价比率ID
	KillPercent       int     `gorm:"column:kill_percent;type:int(2) unsigned;not null" json:"kill_percent"`                // 砍价的百分比
	KillPrice         float64 `gorm:"column:kill_price;type:decimal(11,2) unsigned;not null" json:"kill_price"`             // 砍掉的价格
	KillPosition      int8    `gorm:"column:kill_position;type:tinyint(2) unsigned;not null" json:"kill_position"`          // 玩家当前是第几个砍的
	KillText          string  `gorm:"column:kill_text;type:varchar(255);not null" json:"kill_text"`
	GotProfit         float64 `gorm:"column:got_profit;type:decimal(10,2);not null" json:"got_profit"`         // 用户砍到的红包金额
	StoreID           int     `gorm:"column:store_id;type:int(11) unsigned;not null" json:"store_id"`          // 当前砍价所在的店铺ID
	IsGiveOut         int8    `gorm:"column:is_give_out;type:tinyint(2) unsigned;not null" json:"is_give_out"` // 是否已发放
	IsShare           int8    `gorm:"column:is_share;type:tinyint(2) unsigned;not null" json:"is_share"`       // 是否分享
	Ctime             int     `gorm:"column:ctime;type:int(10) unsigned;not null" json:"ctime"`
	ActID             int     `gorm:"column:act_id;type:int(11) unsigned;not null" json:"act_id"` // 商家活动id
}

type GbPostageNumsection struct {
	ID         int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`       // id
	PostageID  int       `gorm:"column:postage_id;type:int(11);not null" json:"postage_id"`  // 所属运费模板
	MinVal     float32   `gorm:"column:min_val;type:float;not null" json:"min_val"`          // 区间最小值
	MaxVal     float32   `gorm:"column:max_val;type:float;not null" json:"max_val"`          // 区间最大值
	Price      float64   `gorm:"column:price;type:decimal(10,2)" json:"price"`               // 运费
	Unit       int8      `gorm:"column:unit;type:smallint(6);not null" json:"unit"`          // 区间计量单位:0.单价1.数量
	Area       string    `gorm:"column:area;type:varchar(255);not null" json:"area"`         // 运费规则应用区域
	Type       int       `gorm:"column:type;type:int(11);not null" json:"type"`              // 区间类型：0小于等于区间;1.范围区间;2.大于等于区间
	Createtime time.Time `gorm:"column:createtime;type:datetime;not null" json:"createtime"` // 创建时间
}

type GbWebsiteMould struct {
	ID          int       `gorm:"primary_key;column:id;type:int(10);not null" json:"-"`    // 模板管理自增ID
	WebsiteType int8      `gorm:"column:website_type;type:tinyint(1)" json:"website_type"` // 官网类型 1神雀 2鼎翰
	MouldType   int8      `gorm:"column:mould_type;type:tinyint(1)" json:"mould_type"`     // 模板类型 1页首 2导航 3banner 4页尾
	URL         string    `gorm:"column:url;type:varchar(255)" json:"url"`                 // 模板展示图图片地址
	CreateTime  time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`     // 创建时间
	UpdateTime  time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`     // 更新时间
}

// GbArea 省|市|区、县市、县|街道、乡镇|四级行政区域信息表
type GbArea struct {
	ID           int64  `gorm:"primary_key;column:id;type:bigint(11) unsigned;not null" json:"-"`
	Name         string `gorm:"column:name;type:char(30);not null" json:"name"`          // 区域名称
	ParentID     int    `gorm:"column:parent_id;type:int(11) unsigned" json:"parent_id"` // 父ID
	Level        int8   `gorm:"column:level;type:tinyint(2) unsigned;not null" json:"level"`
	Code         string `gorm:"column:code;type:char(6)" json:"code"`                   // 邮编号
	ProvinceArea int    `gorm:"column:province_area;type:int(11)" json:"province_area"` // 省份所属区域
}

type GbCollectKeyword struct {
	ID           int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	SiteID       int    `gorm:"index;column:site_id;type:int(11);not null" json:"site_id"`         // 站点ID
	Keyword      string `gorm:"column:keyword;type:varchar(255);not null" json:"keyword"`          // 关键词
	FirstCateID  int    `gorm:"column:first_cate_id;type:int(11);not null" json:"first_cate_id"`   // 一级分类
	SecondCateID int    `gorm:"column:second_cate_id;type:int(11);not null" json:"second_cate_id"` // 二级分类
	ThirdCateID  int    `gorm:"column:third_cate_id;type:int(11);not null" json:"third_cate_id"`   // 三级分类
	Status       int8   `gorm:"column:status;type:tinyint(1);not null" json:"status"`              // 1启用  0禁用
	IsDel        int8   `gorm:"column:is_del;type:tinyint(1);not null" json:"is_del"`              // 0未删除  1已删除
	CreatedAt    int    `gorm:"column:created_at;type:int(11);not null" json:"created_at"`
	UpdatedAt    int    `gorm:"column:updated_at;type:int(11);not null" json:"updated_at"`
}

// GbMerchantSeckillStore 执行秒杀活动的店铺
type GbMerchantSeckillStore struct {
	ID        int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	SeckillID int    `gorm:"column:seckill_id;type:int(11)" json:"seckill_id"`     // 秒杀活动ID
	StoreID   int    `gorm:"column:store_id;type:int(11)" json:"store_id"`         // 执行秒杀活动的店铺ID
	StoreCode string `gorm:"column:store_code;type:varchar(50)" json:"store_code"` // 店铺编码
}

type GbUserJifenAccount struct {
	ID                    int     `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`                               // 用户积分账户表
	StoreID               int     `gorm:"column:store_id;type:int(11) unsigned" json:"store_id"`                                       // 店铺ID
	UId                   int     `gorm:"column:uid;type:int(11) unsigned" json:"uid"`                                                 // 用户ID
	CardCouponBalance     float64 `gorm:"column:card_coupon_balance;type:decimal(10,2);not null" json:"card_coupon_balance"`           // 用户可用卡券积分余额
	CardCouponNotActivate float64 `gorm:"column:card_coupon_not_activate;type:decimal(10,2);not null" json:"card_coupon_not_activate"` // 用户待激活卡券积分余额
	IsDel                 int8    `gorm:"column:is_del;type:tinyint(2) unsigned;not null" json:"is_del"`                               // 是否禁用 0未禁用 1禁用
	Creattime             int     `gorm:"column:creattime;type:int(10) unsigned;not null" json:"creattime"`                            // 创建时间
	Updatatime            int     `gorm:"column:updatatime;type:int(10) unsigned;not null" json:"updatatime"`                          // 更新时间
	LockJifen             float64 `gorm:"column:lock_jifen;type:decimal(10,2)" json:"lock_jifen"`                                      // 下单时候锁定用户的积分
}

// GbActivityLockQty 活动产品库存锁定表
type GbActivityLockQty struct {
	ID         int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	SellerID   int    `gorm:"column:seller_id;type:int(11) unsigned;not null" json:"seller_id"`        // 商家ID
	ActType    int8   `gorm:"column:act_type;type:tinyint(2) unsigned;not null" json:"act_type"`       // 活动类型 1普通产品 2拼团 3秒杀 4砍价
	ActID      int    `gorm:"column:act_id;type:int(11) unsigned;not null" json:"act_id"`              // 活动ID，默认0为普通产品
	ActStime   int    `gorm:"column:act_stime;type:int(10) unsigned" json:"act_stime"`                 // 活动开始时间
	ActEtime   int    `gorm:"column:act_etime;type:int(10) unsigned" json:"act_etime"`                 // 活动结束时间
	ActName    string `gorm:"column:act_name;type:varchar(255)" json:"act_name"`                       // 活动名称
	StandID    int    `gorm:"column:stand_id;type:int(11) unsigned;not null" json:"stand_id"`          // 平台库SKU属性ID
	LockQty    int    `gorm:"column:lock_qty;type:int(11) unsigned;not null" json:"lock_qty"`          // 锁定库存数量
	UseQty     int    `gorm:"column:use_qty;type:int(11) unsigned" json:"use_qty"`                     // 已使用库存
	Ctime      int    `gorm:"column:ctime;type:int(10) unsigned;not null" json:"ctime"`                // 创建时间
	IsDel      int8   `gorm:"column:is_del;type:tinyint(2) unsigned;not null" json:"is_del"`           // 1为已删除 0未删除
	RollStatus int8   `gorm:"column:roll_status;type:tinyint(2) unsigned;not null" json:"roll_status"` // 库存回滚状态 0 未处理 1已回滚
}

type GbCrmBearer struct {
	ID             int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	BearerLogID    int    `gorm:"column:bearer_log_id;type:int(11)" json:"bearer_log_id"`
	UserID         int    `gorm:"column:user_id;type:int(11);not null" json:"user_id"`
	UserType       int8   `gorm:"column:user_type;type:tinyint(1);not null" json:"user_type"` // 1 H5用户  2商家  3供应商
	Nickname       string `gorm:"column:nickname;type:varchar(255);not null" json:"nickname"`
	CreateTime     int    `gorm:"column:create_time;type:int(11);not null" json:"create_time"`
	CustomerID     int    `gorm:"column:customer_id;type:int(11);not null" json:"customer_id"`
	Username       string `gorm:"column:username;type:varchar(255)" json:"username"`    // 客服昵称
	Status         int8   `gorm:"column:status;type:tinyint(1);not null" json:"status"` // 1会话中  2已离线
	LastUserFd     int    `gorm:"column:last_user_fd;type:int(10)" json:"last_user_fd"`
	LastCustomerFd int    `gorm:"column:last_customer_fd;type:int(10)" json:"last_customer_fd"`
	CompanyID      int    `gorm:"column:company_id;type:int(11)" json:"company_id"`
	StoreCode      string `gorm:"column:store_code;type:varchar(16)" json:"store_code"`
}

type GbMerchantSetting struct {
	ID       int       `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	Key      string    `gorm:"column:key;type:varchar(512);not null" json:"key"`                 // 设置KEY
	Value    string    `gorm:"column:value;type:varchar(1024);not null" json:"value"`            // 设置value
	SellerID int       `gorm:"column:seller_id;type:int(11) unsigned;not null" json:"seller_id"` // 商家ID
	IsDel    int8      `gorm:"column:is_del;type:tinyint(2) unsigned;not null" json:"is_del"`    // 是否禁用 0未禁用 1禁用
	UpdateAt time.Time `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	CreateAt time.Time `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
}

type GbMessageProcess struct {
	ID            int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	SkuID         string    `gorm:"column:skuId;type:varchar(30)" json:"skuId"`                     // skuId
	Price         float64   `gorm:"column:price;type:decimal(10,2)" json:"price"`                   // 批发价
	MerchantPrice float64   `gorm:"column:merchant_price;type:decimal(10,2)" json:"merchant_price"` // 商家卖价
	Ctime         int       `gorm:"column:ctime;type:int(11);not null" json:"ctime"`                // 创建时间
	Time          time.Time `gorm:"column:time;type:datetime" json:"time"`
	Type          int8      `gorm:"column:type;type:tinyint(1)" json:"type"` // 1代表定时任务变更，2代表用户手动提交变更
}

// GbProStandardRelaSup 平台商品规格属性 和供应商的关联关系
type GbProStandardRelaSup struct {
	ID               int     `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	IsTemp           int8    `gorm:"column:is_temp;type:tinyint(1)" json:"is_temp"`                      // 是否生效，0 未生效  1已生效（在没有确定导入之前，所有合并和导入都是未生效 的）
	ProStandardID    int     `gorm:"index;column:pro_standard_id;type:int(11)" json:"pro_standard_id"`   // 平台商品规格ID,对应gb_pro_standard的id
	ProID            int     `gorm:"index;column:pro_id;type:int(11)" json:"pro_id"`                     // 商品ID 对应 gb_product.id
	Sku              string  `gorm:"column:sku;type:varchar(60)" json:"sku"`                             // 对应gb_pro_standard的SKU
	SupplierID       int     `gorm:"column:supplier_id;type:int(11)" json:"supplier_id"`                 // 供应商id
	SupplierName     string  `gorm:"column:supplier_name;type:varchar(50)" json:"supplier_name"`         // 供应商名称
	SupProStandardID int     `gorm:"column:sup_pro_standard_id;type:int(11)" json:"sup_pro_standard_id"` // 供应商 商品规格ID 对应 gb_sup_pro_standard 的ID
	MarketPrice      float64 `gorm:"column:market_price;type:decimal(10,2)" json:"market_price"`         // 市场价
	MinPrice         float64 `gorm:"column:min_price;type:decimal(10,2)" json:"min_price"`               // 最低价
	MinStatus        int8    `gorm:"column:min_status;type:tinyint(1)" json:"min_status"`                // 是否可低限价出售 1代表是
	BatchPrice       float64 `gorm:"column:batch_price;type:decimal(10,2)" json:"batch_price"`           // 批发价
	Qty              int     `gorm:"column:qty;type:int(11)" json:"qty"`                                 // 库存数
	MinQty           int     `gorm:"column:min_qty;type:int(11)" json:"min_qty"`                         // 库存预警
	IsDel            int8    `gorm:"column:is_del;type:tinyint(1)" json:"is_del"`                        // 删除状态 0未删除，1已经删除
	Mark             string  `gorm:"column:mark;type:varchar(200)" json:"mark"`
	ActType          string  `gorm:"column:act_type;type:varchar(10)" json:"act_type"` // 操作类型，merge 合并，add 新增
}

type GbWebsite struct {
	ID            int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	SectionID     int    `gorm:"column:section_id;type:int(11) unsigned;not null" json:"section_id"` // 角色ID
	Username      string `gorm:"index;column:username;type:varchar(20);not null" json:"username"`    // 用户名
	Password      string `gorm:"column:password;type:char(40);not null" json:"password"`             // 密码
	Name          string `gorm:"column:name;type:varchar(255);not null" json:"name"`                 // 姓名
	Mail          string `gorm:"column:mail;type:varchar(255)" json:"mail"`                          // 邮箱
	Salt          string `gorm:"column:salt;type:varchar(10);not null" json:"salt"`                  // 混淆码
	Telphone      string `gorm:"column:telphone;type:varchar(20)" json:"telphone"`                   // 联系电话
	Mobile        string `gorm:"column:mobile;type:char(11)" json:"mobile"`                          // 手机号
	Sex           int8   `gorm:"column:sex;type:tinyint(1)" json:"sex"`                              // 1、男性   2、女性
	LastLoginTime int    `gorm:"column:last_login_time;type:int(11)" json:"last_login_time"`         // 上次登陆时间
	LastLoginIP   string `gorm:"column:last_login_ip;type:varchar(15)" json:"last_login_ip"`         // 上次登陆IP
	Ctime         int    `gorm:"column:ctime;type:int(11)" json:"ctime"`                             // 创建时间
	Mtime         int    `gorm:"column:mtime;type:int(11)" json:"mtime"`                             // 更新时间
	Status        int8   `gorm:"column:status;type:tinyint(1) unsigned" json:"status"`               // 状态  0：不可用  1：可用
	IsDel         int8   `gorm:"column:is_del;type:smallint(2) unsigned;not null" json:"is_del"`     // 删除状态   1表示已删除
	Remark        string `gorm:"column:remark;type:text" json:"remark"`                              // 备注信息
}

type GbCrmChatCounts struct {
	ID          int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	UserID      int    `gorm:"unique_index:idx_0;column:user_id;type:int(11);not null" json:"user_id"`
	FormatDate  string `gorm:"unique_index:idx_0;column:format_date;type:char(8);not null" json:"format_date"`
	FormatMonth string `gorm:"column:format_month;type:char(6);not null" json:"format_month"`
	Counts      int    `gorm:"column:counts;type:int(255);not null" json:"counts"`
}

type GbProfitLog struct {
	ID             int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	IncrementID    string `gorm:"column:increment_id;type:varchar(15);not null" json:"increment_id"` // 订单号
	UId            int    `gorm:"column:uid;type:int(11);not null" json:"uid"`                       // 用户ID
	PushContent    string `gorm:"column:push_content;type:text" json:"push_content"`                 // 推送内容
	ReceiveContent string `gorm:"column:receive_content;type:text" json:"receive_content"`           // 接收内容
	IsReceive      int8   `gorm:"column:is_receive;type:tinyint(1)" json:"is_receive"`               // 0未收到回复1为收到回复
	CreatedAt      int    `gorm:"column:created_at;type:int(10)" json:"created_at"`                  // 创建时间
	UpdatedAt      int    `gorm:"column:updated_at;type:int(10)" json:"updated_at"`                  // 创建时间
}

// GbStoreTag 店铺标签表
type GbStoreTag struct {
	ID      int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	StoreID int    `gorm:"column:store_id;type:int(11) unsigned;not null" json:"store_id"` // 店铺ID
	Name    string `gorm:"column:name;type:varchar(24);not null" json:"name"`              // 标签名称
	Sort    int    `gorm:"column:sort;type:int(11);not null" json:"sort"`                  // 排序
	IsDel   int8   `gorm:"column:is_del;type:tinyint(1);not null" json:"is_del"`           // 1删除，0未删除
	Status  int8   `gorm:"column:status;type:tinyint(1);not null" json:"status"`           // 状态 1 开启，0 关闭
	Ctime   int    `gorm:"column:ctime;type:int(11);not null" json:"ctime"`                // 创建时间
}

// GbSupBankAccount 供应商银行账户信息
type GbSupBankAccount struct {
	ID       int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	Account  string `gorm:"column:account;type:varchar(50);not null" json:"account"`      // 账户名
	RealName string `gorm:"column:real_name;type:varchar(255);not null" json:"real_name"` // 真实姓名
	Type     int8   `gorm:"column:type;type:tinyint(1);not null" json:"type"`             // 账户类型 0 支付宝
	SupUId   int    `gorm:"column:sup_uid;type:int(10);not null" json:"sup_uid"`          // 供应商编号
	CreateAt int    `gorm:"column:create_at;type:int(10);not null" json:"create_at"`      // 创建时间
	UpdateAt int    `gorm:"column:update_at;type:int(10);not null" json:"update_at"`      // 更新时间
	NickName string `gorm:"column:nick_name;type:varchar(50);not null" json:"nick_name"`  // 昵称
}

type GbSupplierManagement struct {
	ID                 int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	ManagementName     string `gorm:"column:management_name;type:varchar(50);not null" json:"management_name"`       // 供货商名称
	ManagementContacts string `gorm:"column:management_contacts;type:varchar(15)" json:"management_contacts"`        // 供货商联系人
	ManagementTel      string `gorm:"column:management_tel;type:varchar(15)" json:"management_tel"`                  // 供货商电话
	ManagementQq       string `gorm:"column:management_qq;type:varchar(30)" json:"management_qq"`                    // 供货商QQ号
	ManagementWx       string `gorm:"column:management_wx;type:varchar(30)" json:"management_wx"`                    // 供货商微信号
	ManagementWw       string `gorm:"column:management_ww;type:varchar(30)" json:"management_ww"`                    // 供货商旺旺号
	ProductDeveloper   string `gorm:"column:product_developer;type:varchar(30);not null" json:"product_developer"`   // 商品开发人
	DocumentaryMan     string `gorm:"column:documentary_man;type:varchar(30);not null" json:"documentary_man"`       // 商品跟单人
	PaymentMethod      int8   `gorm:"column:payment_method;type:tinyint(1);not null" json:"payment_method"`          // 收款方式(1代表微信支付，2代表支付宝支付，3代表银行转账，4代表平台支付)
	PayableOrCompany   string `gorm:"column:payable_or_company;type:varchar(50);not null" json:"payable_or_company"` // 收款人/单位
	ReceivableAccount  string `gorm:"column:receivable_account;type:varchar(60);not null" json:"receivable_account"` // 收款账号
	OpeningBank        string `gorm:"column:opening_bank;type:varchar(60)" json:"opening_bank"`                      // 开户行
	Remarks            string `gorm:"column:remarks;type:varchar(60)" json:"remarks"`                                // 备注
	SupID              int    `gorm:"column:sup_id;type:int(11)" json:"sup_id"`                                      // 供应商ID
	IsDel              int8   `gorm:"column:is_del;type:tinyint(1)" json:"is_del"`                                   // 是否删除（0是存在 1是删除）
	CreateAt           int    `gorm:"column:create_at;type:int(11);not null" json:"create_at"`
	UpdateAt           int    `gorm:"column:update_at;type:int(11)" json:"update_at"` // 更新时间
	Type               int8   `gorm:"column:type;type:tinyint(1)" json:"type"`        // 1实物 2虚拟 3服务商品
	Code               int    `gorm:"index;column:code;type:int(11)" json:"code"`     // 供货商编码
}

type GbCrmBearerLog struct {
	ID          int `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	UserID      int `gorm:"column:user_id;type:int(11);not null" json:"user_id"`
	CompanyID   int `gorm:"column:company_id;type:int(11);not null" json:"company_id"`
	CreateAt    int `gorm:"column:create_at;type:int(11);not null" json:"create_at"`
	LastVisitAt int `gorm:"column:last_visit_at;type:int(11)" json:"last_visit_at"`
}

// GbOrderReason 订单原因设置
type GbOrderReason struct {
	ID       int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	Name     string `gorm:"column:name;type:varchar(150);not null" json:"name"`   // 原因类型
	Sort     int8   `gorm:"column:sort;type:smallint(255);not null" json:"sort"`  // 排序
	Status   int8   `gorm:"column:status;type:tinyint(1);not null" json:"status"` // 状态 1启用 0不启用
	IsDel    int8   `gorm:"column:is_del;type:tinyint(1);not null" json:"is_del"`
	CreateAt int    `gorm:"column:create_at;type:int(11);not null" json:"create_at"`
	UpdateAt int    `gorm:"column:update_at;type:int(11);not null" json:"update_at"`
}

type GbQuexinSample struct {
	ID         int       `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	UserID     int       `gorm:"column:user_id;type:int(11) unsigned;not null" json:"user_id"`           // 用户id
	StoreProID int       `gorm:"column:store_pro_id;type:int(11) unsigned;not null" json:"store_pro_id"` // 店铺产品id
	Status     int8      `gorm:"column:status;type:tinyint(1) unsigned" json:"status"`                   // 申请状态：1申请中，2已确定
	Remark     string    `gorm:"column:remark;type:varchar(255)" json:"remark"`                          // 补充信息
	CreatedAt  time.Time `gorm:"column:created_at;type:timestamp;not null" json:"created_at"`            // 创建时间
	UpdatedAt  time.Time `gorm:"column:updated_at;type:timestamp" json:"updated_at"`                     // 更新时间
	Pic        string    `gorm:"column:pic;type:varchar(255)" json:"pic"`                                // 商品图片
	ProName    string    `gorm:"column:pro_name;type:varchar(200)" json:"pro_name"`                      // 商品名
	BarCode    string    `gorm:"column:bar_code;type:varchar(20)" json:"bar_code"`                       // 商品编码
	Profit     float64   `gorm:"column:profit;type:decimal(10,2)" json:"profit"`                         // 预计利润
	SellPrice  float64   `gorm:"column:sell_price;type:decimal(10,2) unsigned" json:"sell_price"`        // 售价
}

// GbRetailIncome 零售店-收益明细表
type GbRetailIncome struct {
	ID       int     `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	StoreID  int     `gorm:"column:store_id;type:int(11) unsigned;not null" json:"store_id"`       // 商铺ID
	UserID   int     `gorm:"column:user_id;type:int(11);not null" json:"user_id"`                  // 收益人ID (user表)
	RetailID int     `gorm:"column:retail_id;type:int(11) unsigned;not null" json:"retail_id"`     // 零售商ID
	OrderID  int     `gorm:"index;column:order_id;type:int(11) unsigned;not null" json:"order_id"` // 订单ID
	Sku      string  `gorm:"column:sku;type:char(40);not null" json:"sku"`                         // 产品
	Count    int     `gorm:"column:count;type:int(11) unsigned;not null" json:"count"`             // 交易数量
	Income   float64 `gorm:"column:income;type:decimal(10,2);not null" json:"income"`              // 收益
	Status   int8    `gorm:"column:status;type:smallint(1);not null" json:"status"`                // 状态： 0:待结算  1:取消(用户确认收货7天内点击退换货)  2:完成(已入账 可提现)
	Ctime    int     `gorm:"column:ctime;type:int(11);not null" json:"ctime"`                      // 创建时间
}

type GbWhProStandard struct {
	ID           int     `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	WhProductID  int     `gorm:"column:wh_product_id;type:int(11);not null" json:"wh_product_id"`     // 网红产品id
	SupProID     int     `gorm:"index;column:sup_pro_id;type:int(11);not null" json:"sup_pro_id"`     // 供应商产品ID
	StoProID     int     `gorm:"index;column:sto_pro_id;type:int(11);not null" json:"sto_pro_id"`     // 店铺商品id
	Sku          string  `gorm:"column:sku;type:varchar(30)" json:"sku"`                              // 供应商sku
	Price        float64 `gorm:"column:price;type:decimal(10,2)" json:"price"`                        // 网红售价
	PlatProCode  string  `gorm:"column:plat_pro_code;type:varchar(40);not null" json:"plat_pro_code"` // 编号(产品平台编号)
	Properties   string  `gorm:"column:properties;type:text" json:"properties"`                       // 商品sku属性，用于sku后续操作
	SkuProps     string  `gorm:"column:sku_props;type:text" json:"sku_props"`                         // 销售属性
	LastSkuProps string  `gorm:"column:last_sku_props;type:text" json:"last_sku_props"`               // 上一次的销售属性，记录用于第三方sku下架
	IsUpdate     int8    `gorm:"column:is_update;type:tinyint(4);not null" json:"is_update"`          // 是否修改了类目属性数据 0否 1是
}

type GbLogOrderChangeCopy1 struct {
	ID          int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	UId         int    `gorm:"column:uid;type:int(11) unsigned;not null" json:"uid"`                 // 操作者ID
	IncrementID string `gorm:"index;column:increment_id;type:char(20);not null" json:"increment_id"` // 订单流水号
	Remarks     string `gorm:"column:remarks;type:varchar(255)" json:"remarks"`                      // 备注
	AfterStatus int    `gorm:"column:after_status;type:int(11);not null" json:"after_status"`        // 处理后状态
	Creattime   int    `gorm:"column:creattime;type:int(10)" json:"creattime"`                       // 创建时间
}

// GbTaobaoCates 淘宝类目表
type GbTaobaoCates struct {
	ID              int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Level           int8   `gorm:"column:level;type:tinyint(1);not null" json:"level"`                           // 1一级分类  2二级分类 3三级分类 4三级分类
	Cid             int    `gorm:"index;column:cid;type:int(11);not null" json:"cid"`                            // 类目ID
	Name            string `gorm:"column:name;type:varchar(20);not null" json:"name"`                            // 类目名称
	ParentCid       int    `gorm:"column:parent_cid;type:int(11);not null" json:"parent_cid"`                    // 父类目ID
	IsParent        int8   `gorm:"column:is_parent;type:tinyint(1);not null" json:"is_parent"`                   // 是否为父级分类
	LocalCid        int    `gorm:"index;column:local_cid;type:int(11);not null" json:"local_cid"`                // 匹配的本地分类ID
	IsFixedLocalCid int8   `gorm:"column:is_fixed_local_cid;type:tinyint(1);not null" json:"is_fixed_local_cid"` // 是否固定本地分类
	IsFinished      int8   `gorm:"column:is_finished;type:tinyint(1);not null" json:"is_finished"`               // 是否完成属性匹配
}

type GbCash struct {
	ID            int     `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	UserID        int     `gorm:"column:user_id;type:int(10) unsigned;not null" json:"user_id"`          // 用户ID
	UserType      int8    `gorm:"column:user_type;type:tinyint(3) unsigned;not null" json:"user_type"`   // 1 : user表的用户
	OutTradeNo    string  `gorm:"column:out_trade_no;type:varchar(64);not null" json:"out_trade_no"`     // 商户订单号
	TransactionNo string  `gorm:"column:transaction_no;type:varchar(255)" json:"transaction_no"`         // 第三方订单号
	Money         float64 `gorm:"column:money;type:decimal(10,2) unsigned;not null" json:"money"`        // 实发金额
	OriginMoney   float64 `gorm:"column:origin_money;type:decimal(10,2);not null" json:"origin_money"`   // 应发金额
	Tax           float64 `gorm:"column:tax;type:decimal(10,2);not null" json:"tax"`                     // 个税
	ServiceMoney  float64 `gorm:"column:service_money;type:decimal(10,2);not null" json:"service_money"` // 手续费
	CashType      int8    `gorm:"column:cash_type;type:tinyint(2);not null" json:"cash_type"`            // 提现类型：1微信  2支付宝  3银行卡
	CashAccount   string  `gorm:"column:cash_account;type:varchar(255);not null" json:"cash_account"`    // 提现 账号
	Status        int8    `gorm:"column:status;type:tinyint(2) unsigned;not null" json:"status"`         // 0 待提现   1提现中  2提现成功  3提现失败
	AdminID       int     `gorm:"column:admin_id;type:int(11);not null" json:"admin_id"`                 // 操作员ID
	CreatedAt     int     `gorm:"column:created_at;type:int(11);not null" json:"created_at"`
	UpdatedAt     int     `gorm:"column:updated_at;type:int(11);not null" json:"updated_at"`
	Remark        string  `gorm:"column:remark;type:text" json:"remark"`                                   // 备注
	CashSource    int8    `gorm:"column:cash_source;type:tinyint(3) unsigned;not null" json:"cash_source"` // 提现来源：1财务系统  2线下转账 3壹团
	Count         int     `gorm:"column:count;type:int(10) unsigned;not null" json:"count"`                // 订单笔数
	CommissionIDs string  `gorm:"column:commission_ids;type:text" json:"commission_ids"`
	IsCal         int8    `gorm:"column:is_cal;type:tinyint(1);not null" json:"is_cal"` // 0资金未扣税，不可发放  1可发放
	Year          string  `gorm:"column:year;type:char(4);not null" json:"year"`
}

type GbWhUser struct {
	ID            int     `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	Username      string  `gorm:"column:username;type:varchar(255);not null" json:"username"`          // 账户名
	Password      string  `gorm:"column:password;type:varchar(255);not null" json:"password"`          // 密码
	AccessToken   string  `gorm:"column:access_token;type:varchar(255)" json:"access_token"`           // 用户身份凭证
	LastLoginIP   string  `gorm:"column:last_login_ip;type:varchar(255)" json:"last_login_ip"`         // 最后一次登录IP
	LastLoginTime int     `gorm:"column:last_login_time;type:int(11) unsigned" json:"last_login_time"` // 最后一次登录时间
	Avatar        string  `gorm:"column:avatar;type:varchar(255)" json:"avatar"`                       // 微信头像地址
	Nickname      string  `gorm:"column:nickname;type:varchar(255)" json:"nickname"`                   // 微信昵称
	Openid        string  `gorm:"column:openid;type:varchar(255)" json:"openid"`                       // 微信openid
	CreatedAt     int     `gorm:"column:created_at;type:int(11) unsigned;not null" json:"created_at"`
	UpdatedAt     int     `gorm:"column:updated_at;type:int(11);not null" json:"updated_at"`
	PayAmount     float64 `gorm:"column:pay_amount;type:decimal(10,2);not null" json:"pay_amount"` // 网红进货总额
}

type GbCrmKnowledgebaseCate struct {
	ID         int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Name       string    `gorm:"column:name;type:varchar(200);not null" json:"name"`         // 分类名称
	Companyid  int       `gorm:"column:companyid;type:int(11)" json:"companyid"`             // 商户id
	Parentid   int       `gorm:"column:parentid;type:int(11);not null" json:"parentid"`      // 所属上级分类
	Status     int8      `gorm:"column:status;type:smallint(6);not null" json:"status"`      // 状态：0已禁用1已启用
	Createtime time.Time `gorm:"column:createtime;type:datetime;not null" json:"createtime"` // 创建时间
}

// GbMenu 后台系统菜单
type GbMenu struct {
	ID           int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	Level        int8   `gorm:"column:level;type:tinyint(2) unsigned;not null" json:"level"`        // 0为大类菜单
	AccessPath   string `gorm:"column:accessPath;type:varchar(255)" json:"accessPath"`              // URL路径
	IconPath     string `gorm:"column:iconPath;type:varchar(255)" json:"iconPath"`                  // 图标，当为大类菜单时，需要记录此值
	ParentID     int    `gorm:"column:parentID;type:int(11) unsigned;not null" json:"parentID"`     // 父级ID
	ResourceName string `gorm:"column:resourceName;type:varchar(255);not null" json:"resourceName"` // 菜单名称
	Sort         int8   `gorm:"column:sort;type:tinyint(2) unsigned;not null" json:"sort"`          // 排序
	Status       int8   `gorm:"column:status;type:smallint(2) unsigned;not null" json:"status"`     // 状态 0为禁用
	Module       string `gorm:"column:module;type:varchar(255)" json:"module"`                      // 模块名称
}

type GbStoreBrand struct {
	ID        int       `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	StoreID   int       `gorm:"index;column:store_id;type:int(11) unsigned;not null" json:"store_id"` // 店铺ID
	Name      string    `gorm:"column:name;type:varchar(256);not null" json:"name"`                   // 品牌名称
	Logo      string    `gorm:"column:logo;type:varchar(1200)" json:"logo"`                           // 品牌LOGO
	Banner    string    `gorm:"column:banner;type:varchar(1200);not null" json:"banner"`              // 广告图
	Sort      int       `gorm:"column:sort;type:int(11);not null" json:"sort"`                        // 排序
	History   string    `gorm:"column:history;type:varchar(1200);not null" json:"history"`            // 品牌故事
	IsDel     int8      `gorm:"column:is_del;type:tinyint(2) unsigned;not null" json:"is_del"`        // 是否删除 0 否 1 删除
	CreatedAt time.Time `gorm:"column:created_at;type:datetime;not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime;not null" json:"updated_at"`
}

// GbTaobaoPropValues 淘宝类目属性表
type GbTaobaoPropValues struct {
	ID   int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Cid  int    `gorm:"index:ids_cid_pid;index:ids_cid_vid;column:cid;type:int(11);not null" json:"cid"` // 类目ID
	Pid  int    `gorm:"index:ids_cid_pid;column:pid;type:int(11);not null" json:"pid"`                   // 属性ID
	Vid  int    `gorm:"index;index:ids_cid_vid;column:vid;type:int(11);not null" json:"vid"`             // 属性值ID
	Name string `gorm:"column:name;type:varchar(50);not null" json:"name"`                               // 属性值名称
}

// GbTaobaoPropsOriginal 淘宝类目属性表
type GbTaobaoPropsOriginal struct {
	ID            int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Cid           int    `gorm:"index;column:cid;type:int(11);not null" json:"cid"`                    // 类目ID
	Pid           int    `gorm:"index;column:pid;type:int(11);not null" json:"pid"`                    // 属性ID
	ParentPid     int    `gorm:"column:parent_pid;type:int(11);not null" json:"parent_pid"`            // 上级属性ID
	ParentVid     int    `gorm:"column:parent_vid;type:int(11);not null" json:"parent_vid"`            // 上级属性值ID
	Name          string `gorm:"column:name;type:varchar(50);not null" json:"name"`                    // 属性名称
	IsKeyProp     int8   `gorm:"column:is_key_prop;type:tinyint(1);not null" json:"is_key_prop"`       // 是否关键属性：1为是
	IsSaleProp    int8   `gorm:"column:is_sale_prop;type:tinyint(1);not null" json:"is_sale_prop"`     // 是否销售属性：1为是
	IsColorProp   int8   `gorm:"column:is_color_prop;type:tinyint(1);not null" json:"is_color_prop"`   // 是否颜色属性：1为是
	IsEnumProp    int8   `gorm:"column:is_enum_prop;type:tinyint(1);not null" json:"is_enum_prop"`     // 是否是可枚举属性：1为是
	IsInputProp   int8   `gorm:"column:is_input_prop;type:tinyint(1);not null" json:"is_input_prop"`   // 是否是卖家可以自行输入的属性：1为是
	IsItemProp    int8   `gorm:"column:is_item_prop;type:tinyint(1);not null" json:"is_item_prop"`     // 是否商品属性：1为是
	ChildTemplate string `gorm:"column:child_template;type:varchar(50)" json:"child_template"`         // 子属性的模板（卖家自行输入属性时需要用到）
	IsAllowAlias  int8   `gorm:"column:is_allow_alias;type:tinyint(1);not null" json:"is_allow_alias"` // 是否允许别名：1为是
	IsTaosir      int8   `gorm:"column:is_taosir;type:tinyint(1);not null" json:"is_taosir"`           // 是否度量衡属性项：1为是
	TaosirDo      string `gorm:"column:taosir_do;type:text" json:"taosir_do"`                          // 度量衡相关信息集合
	IsMaterial    int8   `gorm:"column:is_material;type:tinyint(1);not null" json:"is_material"`       // 是否是材质属性项：1为是
	MaterialDo    string `gorm:"column:material_do;type:text" json:"material_do"`                      // 材质属性信息集合
	Multi         int8   `gorm:"column:multi;type:tinyint(1);not null" json:"multi"`                   // 发布产品或商品时是否可以多选
	Must          int8   `gorm:"column:must;type:tinyint(1);not null" json:"must"`                     // 发布产品或商品时是否为必选属性
	Value         string `gorm:"column:value;type:varchar(50)" json:"value"`                           // 属性所匹配过的值
	IsFixedValue  int8   `gorm:"column:is_fixed_value;type:tinyint(1);not null" json:"is_fixed_value"` // 是否固定属性值
}

type GbProfitRatio struct {
	ID         int     `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	TypeID     int     `gorm:"column:type_id;type:int(11) unsigned;not null" json:"type_id"`              // 利润分配业务类型 对应 gb_profit_type
	Name       string  `gorm:"column:name;type:char(60);not null" json:"name"`                            // 方案名称
	OwnRatio   float64 `gorm:"column:own_ratio;type:float(6,2) unsigned;not null" json:"own_ratio"`       // 自留比例
	OwnARatio  float64 `gorm:"column:own_a_ratio;type:float(6,2) unsigned;not null" json:"own_a_ratio"`   // 自留比便-内部分配比例A
	OwnBRatio  float64 `gorm:"column:own_b_ratio;type:float(6,2) unsigned;not null" json:"own_b_ratio"`   // 自留比便-内部分配比例A
	ExtRatio   float64 `gorm:"column:ext_ratio;type:float(6,2) unsigned;not null" json:"ext_ratio"`       // 外部分配比例
	CostARatio float64 `gorm:"column:cost_a_ratio;type:float(6,2) unsigned;not null" json:"cost_a_ratio"` // 硬件成本比例
	CostBRatio float64 `gorm:"column:cost_b_ratio;type:float(6,2) unsigned;not null" json:"cost_b_ratio"` // 人工成本比例
	IsTemplate int8    `gorm:"column:is_template;type:tinyint(2) unsigned;not null" json:"is_template"`   // 是否设置为模板
	Status     int8    `gorm:"column:status;type:tinyint(2) unsigned;not null" json:"status"`             // 默认1打开 0关闭
	IsDel      int8    `gorm:"column:is_del;type:tinyint(2) unsigned;not null" json:"is_del"`             // 默认0未删除 1已删除
	CreateUId  int     `gorm:"column:create_uid;type:int(11) unsigned;not null" json:"create_uid"`        // 创建人
	Ctime      int     `gorm:"column:ctime;type:int(10) unsigned;not null" json:"ctime"`                  // 创建时间
}

// GbDatabaseBackup 数据库备份记录表
type GbDatabaseBackup struct {
	ID       int     `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`  // 数据库备份记录
	Version  string  `gorm:"column:version;type:char(40);not null" json:"version"`           // 数据库版本号
	FileName string  `gorm:"column:file_name;type:varchar(255);not null" json:"file_name"`   // 备份名称
	Size     float64 `gorm:"column:size;type:decimal(11,2) unsigned;not null" json:"size"`   // 文件大小
	IsDel    int8    `gorm:"column:is_del;type:smallint(2) unsigned;not null" json:"is_del"` // 删除状态
	Ctime    int     `gorm:"column:ctime;type:int(10) unsigned;not null" json:"ctime"`       // 创建时间
}

// GbOrderChildren 拆单信息表
type GbOrderChildren struct {
	ID                  int     `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	ProName             string  `gorm:"column:pro_name;type:varchar(100);not null" json:"pro_name"`                    // 商品名称
	Code                string  `gorm:"column:code;type:char(24);not null" json:"code"`                                // 拆单号
	SupplierID          int     `gorm:"column:supplier_id;type:int(10);not null" json:"supplier_id"`                   // 供应商id
	ChildrenIncrementID string  `gorm:"index;column:children_increment_id;type:char(20)" json:"children_increment_id"` // 订单商品表唯一编号
	SupplierCode        string  `gorm:"column:supplier_code;type:char(20);not null" json:"supplier_code"`              // 供应商CODE
	Freight             float64 `gorm:"column:freight;type:decimal(10,2) unsigned;not null" json:"freight"`            // 运费总价
	MBatchPrice         float64 `gorm:"column:m_batch_price;type:decimal(10,4);not null" json:"m_batch_price"`         // 平台批发价
	Count               int     `gorm:"column:count;type:int(10);not null" json:"count"`                               // 数量
	GoodsPrice          float64 `gorm:"column:goods_price;type:decimal(10,4) unsigned;not null" json:"goods_price"`    // 商品总价(平台批发价*数量)；如果是网红订单的话；按照网红的商品价格来算的，因为有公开价和特供价（卢子林和陈磊）。
	TotalPrice          float64 `gorm:"column:total_price;type:decimal(10,4) unsigned;not null" json:"total_price"`    // 子单总价（商品总价+运费总价）;如果是网红订单的话，该值等于good_price + flight（卢子林和陈磊）。
	Status              int8    `gorm:"index;column:status;type:tinyint(2);not null" json:"status"`                    // 状态 -1已取消，1新子订单 2待发货 3待收货 5完成（用户收货） 9已完成 10售后中
	IsConfirm           int8    `gorm:"column:is_confirm;type:tinyint(2) unsigned" json:"is_confirm"`                  // 客服确认 0未确认 1确认
	Creattime           int     `gorm:"column:creattime;type:int(10) unsigned;not null" json:"creattime"`              // 创建时间
	IncrementID         string  `gorm:"index;column:increment_id;type:char(15)" json:"increment_id"`                   // 订单流水号（YYMMDD+）
	Updatetime          int     `gorm:"column:updatetime;type:int(10)" json:"updatetime"`                              // 更新时间
	IsDel               int8    `gorm:"column:is_del;type:smallint(2) unsigned" json:"is_del"`
	PurchaseIncrementID string  `gorm:"column:purchase_increment_id;type:char(20)" json:"purchase_increment_id"`     // 平台采购流水号
	PurchaseID          int     `gorm:"column:purchase_id;type:int(10)" json:"purchase_id"`                          // 平台采购ID
	BatchPrice          float64 `gorm:"column:batch_price;type:decimal(10,2)" json:"batch_price"`                    // 供应商供货单价(平台采购价)
	SupUnitFreight      float64 `gorm:"column:sup_unit_freight;type:decimal(10,2)" json:"sup_unit_freight"`          // 供应商运费单价(平台采购运费单件价价格)
	SupProStandardID    int     `gorm:"column:sup_pro_standard_id;type:int(10)" json:"sup_pro_standard_id"`          // 供应商商品规格ID
	ShipmentID          int     `gorm:"column:shipment_id;type:int(10)" json:"shipment_id"`                          // 快递公司ID
	FreightNumber       string  `gorm:"column:freight_number;type:char(40)" json:"freight_number"`                   // 物流单号
	FreightTime         int     `gorm:"column:freight_time;type:int(10)" json:"freight_time"`                        // 发货时间
	Remark              string  `gorm:"column:remark;type:text" json:"remark"`                                       // 子订单备注
	SupProStandardSku   string  `gorm:"column:sup_pro_standard_sku;type:char(100)" json:"sup_pro_standard_sku"`      // 供应商产品规格SKU（同步第三方SKU）
	ThirdOrderID        string  `gorm:"column:third_order_id;type:char(32)" json:"third_order_id"`                   // 第三方平台订单号
	ThirdOrderChildID   string  `gorm:"column:third_order_child_id;type:char(32)" json:"third_order_child_id"`       // 第三方平台子订单号
	RefundTepy          int8    `gorm:"index;column:refund_tepy;type:tinyint(2)" json:"refund_tepy"`                 // 售后类型  1为退款，2退换货
	MergeOrderID        int     `gorm:"column:merge_order_id;type:int(10)" json:"merge_order_id"`                    // 供应商订单ID，OrderMergeChildrenModelID
	ThirdPlatformName   string  `gorm:"column:third_platform_name;type:char(40)" json:"third_platform_name"`         // 第三方平台名称
	ThirdOrderStatus    int8    `gorm:"column:third_order_status;type:tinyint(2)" json:"third_order_status"`         // 第三方平台拆单状态 1为未拆单  2为已拆单
	PurchaseFreight     float64 `gorm:"column:purchase_freight;type:decimal(10,2);not null" json:"purchase_freight"` // 该笔子订单的运费
	ManagementName      string  `gorm:"column:management_name;type:varchar(50)" json:"management_name"`              // 商品供货商（陈伟强加）
	ProductDeveloper    string  `gorm:"column:product_developer;type:varchar(10)" json:"product_developer"`          // 商品开发人（陈伟强加）
	DocumentaryMan      string  `gorm:"column:documentary_man;type:varchar(10)" json:"documentary_man"`              // 商品跟单人（陈伟强加）
	AkcActivityCode     string  `gorm:"column:akc_activity_code;type:varchar(100)" json:"akc_activity_code"`         // 第三方活动ID(liveId)
	AkcProductID        string  `gorm:"column:akc_product_id;type:varchar(100)" json:"akc_product_id"`               // 第三方产品ID(spuId)
	TaxDate             int     `gorm:"column:tax_date;type:int(11)" json:"tax_date"`                                // 税率
	FreightType         int     `gorm:"column:freight_type;type:int(2)" json:"freight_type"`                         // 物流信息来源，1：待打印 3：已打印；5：已完成；4：未定义来源
	ProCode             string  `gorm:"column:pro_code;type:varchar(60);not null" json:"pro_code"`                   // 商品编号
	ReceiveTime         int     `gorm:"column:receive_time;type:int(10);not null" json:"receive_time"`               // 收货时间，网红7天之后方便计算网红供应商可提现金额（陈磊加）
}

type GbOrderChildrenPayDetails struct {
	ID                      int     `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`                                   // 主键ID
	ChildrenID              int     `gorm:"column:children_id;type:int(11) unsigned;not null" json:"children_id"`                            // gb_order_children表的id
	IncrementID             string  `gorm:"index;column:increment_id;type:char(20);not null" json:"increment_id"`                            // 总订单流水号:123456789451245
	PayJifen                float64 `gorm:"column:pay_jifen;type:decimal(10,2);not null" json:"pay_jifen"`                                   // 花费积分
	PayMoney                float64 `gorm:"column:pay_money;type:decimal(10,2);not null" json:"pay_money"`                                   // 实际支付金额
	Commision               float64 `gorm:"column:commision;type:decimal(10,2);not null" json:"commision"`                                   // 优惠券金额
	Price                   float64 `gorm:"column:price;type:decimal(10,2);not null" json:"price"`                                           // 商品总价，不含邮费
	Count                   int     `gorm:"column:count;type:int(11) unsigned;not null" json:"count"`                                        // 购买数量
	OrderChildrenCountPrice float64 `gorm:"column:order_children_count_price;type:decimal(10,2);not null" json:"order_children_count_price"` // 用户支付的商品总价
	BatchType               int8    `gorm:"column:batch_type;type:smallint(2) unsigned" json:"batch_type"`                                   // 含税类型 1未税 2含增值税 3含普票
	ItemID                  int     `gorm:"column:item_id;type:int(11) unsigned;not null" json:"item_id"`                                    // 订单商品ID,表order_item的id
	Creattime               int     `gorm:"index;column:creattime;type:int(10) unsigned;not null" json:"creattime"`                          // 创建时间
	Updatetime              int     `gorm:"column:updatetime;type:int(10) unsigned;not null" json:"updatetime"`                              // 更新时间
}

type GbCrmWorkOrderCopy1 struct {
	ID                   int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	JobNo                string `gorm:"column:job_no;type:varchar(64);not null" json:"job_no"`                             // 工单号
	Title                string `gorm:"column:title;type:varchar(255);not null" json:"title"`                              // 标题
	Content              string `gorm:"column:content;type:text;not null" json:"content"`                                  // 内容
	Enclosure            string `gorm:"column:enclosure;type:text" json:"enclosure"`                                       // 附件
	Level                int8   `gorm:"column:level;type:tinyint(1);not null" json:"level"`                                // 1非常紧急  2紧急  3一般  4低
	AcceptanceSectionID  int    `gorm:"column:acceptance_section_id;type:int(11);not null" json:"acceptance_section_id"`   // 受理部门
	AcceptancePassportID int    `gorm:"column:acceptance_passport_id;type:int(11);not null" json:"acceptance_passport_id"` // 受理人
	CateID               int8   `gorm:"column:cate_id;type:tinyint(2);not null" json:"cate_id"`                            // 工单分类
	Username             string `gorm:"column:username;type:varchar(64);not null" json:"username"`                         // 用户姓名
	Mobile               string `gorm:"column:mobile;type:char(11);not null" json:"mobile"`                                // 用户手机
	Status               int8   `gorm:"column:status;type:tinyint(1);not null" json:"status"`                              // 0 受理中  1已处理  2已完结
	CreateAt             int    `gorm:"column:create_at;type:int(11);not null" json:"create_at"`
	IsDel                int8   `gorm:"column:is_del;type:tinyint(1);not null" json:"is_del"`
	CustomerID           int    `gorm:"column:customer_id;type:int(11);not null" json:"customer_id"` // 发起人ID
	UserID               int    `gorm:"column:user_id;type:int(11);not null" json:"user_id"`         // 用户ID
}

type GbDyyImportLog struct {
	ID  int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Log string `gorm:"column:log;type:text" json:"log"`
}

type GbLogisticsRs struct {
	ID            int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	ChildrenID    int    `gorm:"column:children_id;type:int(10);not null" json:"children_id"`      // 子订单ID
	IncrementID   string `gorm:"column:increment_id;type:char(15)" json:"increment_id"`            // 订单流水号
	ShipmentID    int    `gorm:"column:shipment_id;type:int(10)" json:"shipment_id"`               // 快递公司ID
	FreightNumber string `gorm:"column:freight_number;type:char(40)" json:"freight_number"`        // 物流单号
	Creattime     int    `gorm:"column:creattime;type:int(10) unsigned;not null" json:"creattime"` // 创建时间
}

type GbRetailRemittance struct {
	ID           int     `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`   // 分销用户转账记录表
	WithdrawalID int     `gorm:"column:withdrawal_id;type:int(11)" json:"withdrawal_id"` // 提现申请表id
	DealerID     int     `gorm:"column:dealer_id;type:int(11)" json:"dealer_id"`         // 汇款的分销商id
	SetMode      int8    `gorm:"column:set_mode;type:tinyint(1)" json:"set_mode"`        // 结算方式(1:微信 2:支付宝)
	SetAccount   string  `gorm:"column:set_account;type:varchar(32)" json:"set_account"` // 结算账户
	TotalFee     float64 `gorm:"column:total_fee;type:decimal(10,2)" json:"total_fee"`   // 汇款金额
	Status       int8    `gorm:"column:status;type:tinyint(1)" json:"status"`            // 0汇款失败 1汇款成功
	CTime        int     `gorm:"column:c_time;type:int(11)" json:"c_time"`               // 汇款时间
	Remark       string  `gorm:"column:remark;type:varchar(1024)" json:"remark"`         // 汇款备注
	AdminID      int     `gorm:"column:admin_id;type:int(11)" json:"admin_id"`           // 管理员id
}

type GbWhGoodsMatching struct {
	ID           int     `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	Sid          int     `gorm:"unique_index:idx_uid_sid_oid_osid;column:sid;type:int(20) unsigned;not null" json:"sid"`       // 店铺ID
	UId          int64   `gorm:"unique_index:idx_uid_sid_oid_osid;column:uid;type:bigint(20) unsigned;not null" json:"uid"`    // 用户ID
	OuterID      string  `gorm:"column:outer_id;type:varchar(60)" json:"outer_id"`                                             // 本地商品编码
	ProID        int     `gorm:"unique_index:idx_uid_sid_oid_osid;column:pro_id;type:int(11) unsigned;not null" json:"pro_id"` // 本地匹配的商品ID
	OpNumIid     string  `gorm:"column:op_num_iid;type:varchar(50);not null" json:"op_num_iid"`                                // 平台商品ID
	GoodsStatus  int8    `gorm:"column:goods_status;type:tinyint(2);not null" json:"goods_status"`                             // 平台商品状态 0-可用 1-上架 2-下架 3-删除
	Price        float64 `gorm:"column:price;type:decimal(10,2)" json:"price"`                                                 // 平台商品售价
	ItemStatus   int8    `gorm:"column:item_status;type:tinyint(1);not null" json:"item_status"`                               // 商品铺货状态 -1铺货失败 1铺货中 2铺货成功
	ImageStatus  int8    `gorm:"column:image_status;type:tinyint(1);not null" json:"image_status"`                             // 图片上传状态  -1处理失败 1处理中 2处理完成
	PlatformType string  `gorm:"column:platform_type;type:varchar(12);not null" json:"platform_type"`                          // 平台类型 淘宝 taobao 京东 jingdong
}

type GbCrmWorkOrderCateCopy1 struct {
	ID          int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Name        string    `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Description string    `gorm:"column:description;type:varchar(255);not null" json:"description"`
	Status      int8      `gorm:"column:status;type:tinyint(4);not null" json:"status"`
	Createtime  time.Time `gorm:"column:createtime;type:datetime;not null" json:"createtime"`
	Companyid   int       `gorm:"column:companyid;type:int(11);not null" json:"companyid"`
}

// GbOrderInvoice 订单发票信息表
type GbOrderInvoice struct {
	ID          int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"` // 用户发票信息存根
	UId         int    `gorm:"column:uid;type:int(11) unsigned;not null" json:"uid"`
	OrderNum    string `gorm:"column:order_num;type:char(40);not null" json:"order_num"`
	Type        int8   `gorm:"column:type;type:tinyint(2) unsigned;not null" json:"type"`                 // 发票类型 1普通发票 2电子普通发票  3增值税发票
	InvoiceType int8   `gorm:"column:invoice_type;type:tinyint(2) unsigned;not null" json:"invoice_type"` // 发票抬头 1个人 2单位
	Phone       string `gorm:"column:phone;type:char(20)" json:"phone"`                                   // 收票人手机
	Email       string `gorm:"column:email;type:varchar(255)" json:"email"`                               // 邮箱
	Name        string `gorm:"column:name;type:varchar(255)" json:"name"`                                 // 发票抬头名称
	ShopType    int8   `gorm:"column:shop_type;type:tinyint(2) unsigned" json:"shop_type"`                // 1商品明细，2商品类别
	TaxNum      string `gorm:"column:tax_num;type:char(30)" json:"tax_num"`                               // 税号识别码
	WriteStatus int8   `gorm:"column:write_status;type:tinyint(2) unsigned" json:"write_status"`          // 填写状态 1开票内容已完整填写，2开票内容未填写完整
	Ctime       int    `gorm:"column:ctime;type:int(10) unsigned;not null" json:"ctime"`
}

type GbRechargeLog struct {
	ID         int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	OrderNo    string `gorm:"column:order_no;type:varchar(50)" json:"order_no"`               // 订单号
	NotifyType int8   `gorm:"column:notify_type;type:tinyint(2);not null" json:"notify_type"` // 通知方式 1为同步通知2为异步通知
	Source     int8   `gorm:"column:source;type:tinyint(1)" json:"source"`                    // 来源
	URL        string `gorm:"column:url;type:varchar(255);not null" json:"url"`               // 通知网址
	Content    string `gorm:"column:content;type:text;not null" json:"content"`               // 通知内容
	CreatedAt  int    `gorm:"column:created_at;type:int(13);not null" json:"created_at"`      // 创建时间
	UpdatedAt  int    `gorm:"column:updated_at;type:int(13);not null" json:"updated_at"`      // 更新时间
	GameState  int    `gorm:"column:game_state;type:int(2)" json:"game_state"`                // 充值状态（1成功，9失败）
}

type GbStoreCardCreateTaskRecord struct {
	ID         int       `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	StoreID    int       `gorm:"column:store_id;type:int(11) unsigned;not null" json:"store_id"`       // 店铺ID
	Name       string    `gorm:"column:name;type:varchar(512);not null" json:"name"`                   // 任务名称
	NotifyMsg  string    `gorm:"column:notify_msg;type:varchar(512);not null" json:"notify_msg"`       // 提示信息
	MakeNum    int       `gorm:"column:make_num;type:int(11);not null" json:"make_num"`                // 预计生成数量
	MakeValue  int       `gorm:"column:make_value;type:int(11) unsigned;not null" json:"make_value"`   // 生成面值
	JobID      int       `gorm:"column:job_id;type:int(11) unsigned;not null" json:"job_id"`           // 任务ID
	NowCreated int       `gorm:"column:now_created;type:int(11) unsigned;not null" json:"now_created"` // 当前已创建条数
	Status     int8      `gorm:"column:status;type:tinyint(2) unsigned;not null" json:"status"`        // 任务状态0等待中1运行中2已取消3已完成4创建异常
	CreatedAt  time.Time `gorm:"column:created_at;type:datetime;not null" json:"created_at"`           // 创建时间
	UpdatedAt  time.Time `gorm:"column:updated_at;type:datetime;not null" json:"updated_at"`           // 更新时间
}

type GbStoreTrendCate struct {
	ID      int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Title   string `gorm:"column:title;type:varchar(250);not null" json:"title"`            // 类名
	Image   string `gorm:"column:image;type:varchar(50);not null" json:"image"`             // 图片地址
	IsTop   int8   `gorm:"column:is_top;type:tinyint(2) unsigned;not null" json:"is_top"`   // 是否置顶  1置顶 0未置顶
	Display int8   `gorm:"column:display;type:tinyint(2) unsigned;not null" json:"display"` // 展示状态  0不展示  1展示
	Sort    int    `gorm:"column:sort;type:int(11);not null" json:"sort"`                   // 排序
	CTime   int    `gorm:"column:c_time;type:int(10) unsigned;not null" json:"c_time"`
}

type GbUserRelation struct {
	ID         int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	UId        int    `gorm:"column:uid;type:int(11)" json:"uid"`                          // 用户ID
	StoreID    int    `gorm:"column:store_id;type:int(11) unsigned" json:"store_id"`       // 店铺ID
	MerchantID int    `gorm:"column:merchant_id;type:int(11) unsigned" json:"merchant_id"` // 商家ID
	ShareUId   int    `gorm:"column:share_uid;type:int(11)" json:"share_uid"`              // 分享用户ID
	ShareDepth string `gorm:"column:share_depth;type:varchar(255)" json:"share_depth"`     // 分享级别目录
	Type       int    `gorm:"column:type;type:int(2)" json:"type"`                         // 注册来源类型 1为H5 2为小程序
	CreatedAt  int    `gorm:"column:created_at;type:int(10)" json:"created_at"`            // 注册时间
	UpdateAt   int    `gorm:"column:update_at;type:int(10)" json:"update_at"`              // 更新时间
}

type GbCatchSQL struct {
	ID        int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	Category  string `gorm:"column:category;type:varchar(255)" json:"category"`                  // 分类
	SQL       string `gorm:"column:sql;type:text" json:"sql"`                                    // sql命令
	CreatedAt int    `gorm:"column:created_at;type:int(10) unsigned;not null" json:"created_at"` // 创建时间
	UpdatedAt int    `gorm:"column:updated_at;type:int(10)" json:"updated_at"`                   // 更新时间
}

type GbOrderPurchaseMergeInvoiceIndex struct {
	ID             int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`          // 主键ID
	MergeInvoiceID int    `gorm:"column:merge_invoice_id;type:int(11);not null" json:"merge_invoice_id"`  // 表order_merge_invoice的id
	ChildrenID     int    `gorm:"index;column:children_id;type:int(11);not null" json:"children_id"`      // 表order_children表的id
	InvoiceType    int8   `gorm:"column:invoice_type;type:tinyint(4);not null" json:"invoice_type"`       // 发票号类型，1表示普票，2表示增值票
	InvoiceNumber  string `gorm:"column:invoice_number;type:varchar(256);not null" json:"invoice_number"` // 发票号，由财务系统导入
	Creattime      int    `gorm:"index;column:creattime;type:int(10) unsigned;not null" json:"creattime"` // 创建时间
	Updatetime     int    `gorm:"column:updatetime;type:int(10) unsigned;not null" json:"updatetime"`     // 更新时间
}

type GbProfitPassport struct {
	ID       int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	Username string `gorm:"column:username;type:char(60);not null" json:"username"`        // 登陆用户名
	Password string `gorm:"column:password;type:char(60);not null" json:"password"`        // 登陆密码
	Status   int8   `gorm:"column:status;type:tinyint(2) unsigned;not null" json:"status"` // 状态1为开启 0为停用
	IsDel    int8   `gorm:"column:is_del;type:tinyint(2) unsigned;not null" json:"is_del"` // 删除状态 0为正常 1为删除
	Menu     string `gorm:"column:menu;type:text;not null" json:"menu"`                    // 菜单权限
	Ctime    int    `gorm:"column:ctime;type:int(10) unsigned;not null" json:"ctime"`      // 创建时间
}

type GbProfitRule struct {
	ID         int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	CreateUId  int    `gorm:"column:create_uid;type:int(11) unsigned;not null" json:"create_uid"`      // 创建人
	TypeID     int    `gorm:"column:type_id;type:int(11) unsigned;not null" json:"type_id"`            // 业务类型 对应gb_profit_type表
	RatioID    int    `gorm:"column:ratio_id;type:int(11) unsigned;not null" json:"ratio_id"`          // 内部分配方案ID，对应 profit_ratio表
	Name       string `gorm:"column:name;type:char(60);not null" json:"name"`                          // 外部分配方案名称
	IsTemplate int8   `gorm:"column:is_template;type:tinyint(2) unsigned;not null" json:"is_template"` // 是否保存为模板 1是 0否
	Status     int8   `gorm:"column:status;type:tinyint(2) unsigned;not null" json:"status"`           // 开启状态 1开 0关
	IsDel      int8   `gorm:"column:is_del;type:tinyint(2) unsigned;not null" json:"is_del"`           // 删除状态 1是 0否
	Ctime      int    `gorm:"column:ctime;type:int(10) unsigned;not null" json:"ctime"`                // 创建时间
}

// GbViewAppProfit VIEW
type GbViewAppProfit struct {
	ID                 int       `gorm:"column:id;type:int(11) unsigned" json:"-"`
	AppID              int       `gorm:"column:app_id;type:int(11)" json:"app_id"`                      // 终端类型 0 h5 5 wxapp
	UserTag            int8      `gorm:"column:user_tag;type:tinyint(2)" json:"user_tag"`               // 用户类型 0 消费者 1BD 2代理商
	Commision          float64   `gorm:"column:commision;type:decimal(26,6) unsigned" json:"commision"` // 佣金
	SkuID              int       `gorm:"column:sku_id;type:int(11) unsigned" json:"sku_id"`             // 商家产品规格id
	SpuID              int       `gorm:"column:spu_id;type:int(11) unsigned" json:"spu_id"`             // 商家产品id spu
	SellerID           int       `gorm:"column:seller_id;type:int(11) unsigned" json:"seller_id"`       // 商家ID
	CreatedAt          time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt          time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`
	SecretKey          string    `gorm:"column:secret_key;type:char(16)" json:"secret_key"` // 密钥
	Sku                string    `gorm:"column:sku;type:varchar(60)" json:"sku"`            // 产品SKU
	GmpsID             int       `gorm:"column:gmps_id;type:int(11)" json:"gmps_id"`
	PlatProCode        string    `gorm:"column:plat_pro_code;type:varchar(100)" json:"plat_pro_code"` // 编号(产品平台编号)
	PriceAfterDiscount float64   `gorm:"column:price_after_discount;type:decimal(27,6)" json:"price_after_discount"`
	CommisionPercent   float64   `gorm:"column:commision_percent;type:decimal(36,10)" json:"commision_percent"`
}

// GbMerchantBargain 砍价活动信息表
type GbMerchantBargain struct {
	ID            int     `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Seller        int     `gorm:"column:seller;type:int(11)" json:"seller"`                       // 商家ID
	Name          string  `gorm:"column:name;type:varchar(100)" json:"name"`                      // 砍价活动名称
	StartTime     int     `gorm:"column:start_time;type:int(11)" json:"start_time"`               // 活动开始时间
	EndTime       int     `gorm:"column:end_time;type:int(11)" json:"end_time"`                   // 结束时间
	Status        int     `gorm:"column:status;type:int(2)" json:"status"`                        // 砍价活动状态 ，10 未开始，20 进行中，30 已结束
	MerchantSkuID int     `gorm:"column:merchant_sku_id;type:int(11)" json:"merchant_sku_id"`     // 商家SKUID
	PlatSkuID     int     `gorm:"column:plat_sku_id;type:int(11)" json:"plat_sku_id"`             // 平台SKU  ID 号
	MinPrice      float64 `gorm:"column:min_price;type:decimal(10,2)" json:"min_price"`           // 底价，砍到最后的销售价格
	MinCommission float64 `gorm:"column:min_commission;type:decimal(10,2)" json:"min_commission"` // 保底分成金额
	TemplateID    int     `gorm:"column:template_id;type:int(11)" json:"template_id"`             // 砍价模板ID，对应表gb_merchant_bargain_template.id
	IsDel         int8    `gorm:"column:is_del;type:tinyint(1)" json:"is_del"`
	TextConfigID  int     `gorm:"column:text_config_id;type:int(11)" json:"text_config_id"`               // 关联的文案配置ID  gb_merchant_bargain_text_config.id
	StartJobID    int     `gorm:"column:start_job_id;type:int(11) unsigned;not null" json:"start_job_id"` // 自动开启活动任务id
	CloseJobID    int     `gorm:"column:close_job_id;type:int(11) unsigned;not null" json:"close_job_id"` // 自动关闭活动任务id
}

// GbProExaminLog 平台商品审核记录
type GbProExaminLog struct {
	ID       int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	ProCode  string `gorm:"column:pro_code;type:varchar(20)" json:"pro_code"`  // 商品ENA13编码
	UId      int    `gorm:"column:uid;type:int(11)" json:"uid"`                // 审核人员ID
	ExResult int8   `gorm:"column:ex_result;type:tinyint(1)" json:"ex_result"` // 审核结果，0 审核不通过，1 审核通过
	ExMsg    string `gorm:"column:ex_msg;type:varchar(250)" json:"ex_msg"`     // 审核原因，不通过的原因
	Ctime    int    `gorm:"column:ctime;type:int(11)" json:"ctime"`            // 创建时间
	IsDel    int8   `gorm:"column:is_del;type:tinyint(1)" json:"is_del"`       // 删除状态 0 未删除 ，1 已删除
}

// GbMerchantProTag 商家-产品标签表
type GbMerchantProTag struct {
	ID          int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	SellerID    int    `gorm:"column:seller_id;type:int(11) unsigned;not null" json:"seller_id"` // 商家ID
	Name        string `gorm:"column:name;type:varchar(24);not null" json:"name"`                // 店铺名称
	Pic         string `gorm:"column:pic;type:varchar(100)" json:"pic"`
	Content     string `gorm:"column:content;type:varchar(255);not null" json:"content"`
	Sort        int    `gorm:"column:sort;type:int(11);not null" json:"sort"`          // 排序
	IsDel       int8   `gorm:"column:is_del;type:tinyint(1);not null" json:"is_del"`   // 删除状态 1删除，0未删除
	Ctime       int    `gorm:"column:ctime;type:int(11);not null" json:"ctime"`        // 创建时间
	Mtime       int    `gorm:"column:mtime;type:int(11);not null" json:"mtime"`        // 更新时间
	SourceTagID int    `gorm:"column:source_tag_id;type:int(11)" json:"source_tag_id"` // 来源标签ID
}

type GbStoreSetting struct {
	ID       int       `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	Key      string    `gorm:"column:key;type:varchar(512);not null" json:"key"`               // 设置KEY
	Value    string    `gorm:"column:value;type:varchar(1024);not null" json:"value"`          // 设置value
	StoreID  int       `gorm:"column:store_id;type:int(11) unsigned;not null" json:"store_id"` // 店铺ID
	IsDel    int8      `gorm:"column:is_del;type:tinyint(2) unsigned;not null" json:"is_del"`  // 是否禁用 0未禁用 1禁用
	UpdateAt time.Time `gorm:"column:update_at;type:datetime;not null" json:"update_at"`
	CreateAt time.Time `gorm:"column:create_at;type:datetime;not null" json:"create_at"`
}

// GbStoreUserFavorite 用户商品收藏表
type GbStoreUserFavorite struct {
	ID       int     `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	UId      int     `gorm:"column:uid;type:int(11);not null" json:"uid"`                  // 用户ID
	StoreID  int     `gorm:"column:store_id;type:int(11);not null" json:"store_id"`        // 店铺ID
	ProCode  string  `gorm:"column:pro_code;type:varchar(50);not null" json:"pro_code"`    // 产品编号
	ProName  string  `gorm:"column:pro_name;type:varchar(255);not null" json:"pro_name"`   // 商品名
	CateID   int     `gorm:"column:cate_id;type:int(11);not null" json:"cate_id"`          // 分类ID
	CateName string  `gorm:"column:cate_name;type:varchar(255);not null" json:"cate_name"` // 分类名
	Price    float64 `gorm:"column:price;type:decimal(20,6);not null" json:"price"`        // 售价
	Ctime    int     `gorm:"column:ctime;type:int(11);not null" json:"ctime"`              // 创建日期
	IsDel    int8    `gorm:"column:is_del;type:tinyint(1);not null" json:"is_del"`         // 是否删除
}

type GbSupProductOpLog struct {
	ID            int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	UId           int       `gorm:"column:uid;type:int(11)" json:"uid"`                            // 平台操作员id
	ProID         int       `gorm:"column:pro_id;type:int(11)" json:"pro_id"`                      // 产品id
	SuppID        int       `gorm:"column:supp_id;type:int(11);not null" json:"supp_id"`           // 供应商id
	BeforeData    string    `gorm:"column:before_data;type:text" json:"before_data"`               // 操作之前数据
	AfterData     string    `gorm:"column:after_data;type:text" json:"after_data"`                 // 操作之后数据
	ExamineRemark string    `gorm:"column:examine_remark;type:varchar(200)" json:"examine_remark"` // 审核说明
	ExamineStatus int       `gorm:"column:examine_status;type:int(11)" json:"examine_status"`      // 审核状态
	IsDel         int       `gorm:"column:is_del;type:int(11)" json:"is_del"`                      // 删除状态
	Type          int       `gorm:"column:type;type:int(11);not null" json:"type"`                 // 操作类型：1.新增2.更新.3审核4.下架5.下架审核
	CreateTime    time.Time `gorm:"column:create_time;type:timestamp" json:"create_time"`          // 操作时间
}

type GbVirtualOrderLog struct {
	ID            int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	OrderID       int    `gorm:"column:order_id;type:int(11) unsigned;not null" json:"order_id"`     // 订单id
	IncrementID   string `gorm:"column:increment_id;type:varchar(255);not null" json:"increment_id"` // 订单流水号id
	Code          string `gorm:"column:code;type:char(100);not null" json:"code"`                    // 订单号
	Creattime     int    `gorm:"column:creattime;type:int(10) unsigned;not null" json:"creattime"`   // 创建时间
	Type          int8   `gorm:"column:type;type:tinyint(1);not null" json:"type"`                   // 订单类型  1：已提交   2：返回成功 3：返回失败
	ReturnMessage string `gorm:"column:return_message;type:text" json:"return_message"`              // 第三方返回示例消息(默认json)
}

type GbAPIPermission struct {
	ID     int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	UserID int    `gorm:"unique_index:unique_idx_0;column:user_id;type:int(11)" json:"user_id"`
	Name   string `gorm:"unique_index:unique_idx_0;column:name;type:varchar(64)" json:"name"`
	IsDel  int8   `gorm:"unique_index:unique_idx_0;column:is_del;type:tinyint(1)" json:"is_del"`
}

type GbMerchantProRelaShareTemplate struct {
	ID            int     `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	MerchantID    int     `gorm:"column:merchant_id;type:int(11)" json:"merchant_id"`         // 商家ID
	MerchantProID int     `gorm:"column:merchant_pro_id;type:int(11)" json:"merchant_pro_id"` // 商家产品ID  Merchant_product.id
	MerSkuID      int     `gorm:"column:mer_sku_id;type:int(11)" json:"mer_sku_id"`           // 商家商品规格
	TemplateID    int     `gorm:"column:template_id;type:int(11)" json:"template_id"`         // 分成模板ID
	IsDel         int8    `gorm:"column:is_del;type:tinyint(1)" json:"is_del"`
	PartolNum     float64 `gorm:"column:partol_num;type:float(10,2)" json:"partol_num"` // 分成金额
}

// GbOrderMergeChildren 供应商合并订单表
type GbOrderMergeChildren struct {
	ID                  int     `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	SupplierID          int     `gorm:"column:supplier_id;type:int(10);not null" json:"supplier_id"`                // 供应商id
	SupplierCode        string  `gorm:"column:supplier_code;type:char(20);not null" json:"supplier_code"`           // 供应商CODE
	IncrementID         string  `gorm:"column:increment_id;type:char(15)" json:"increment_id"`                      // 订单流水号
	OrderChildrenIDs    string  `gorm:"column:order_children_ids;type:char(100)" json:"order_children_ids"`         // 子订单ID号，各ID号之间用,隔开
	Freight             float64 `gorm:"column:freight;type:decimal(10,2) unsigned;not null" json:"freight"`         // 运费总价
	GoodsPrice          float64 `gorm:"column:goods_price;type:decimal(10,2) unsigned;not null" json:"goods_price"` // 商品总价(供应商供货价*数量)
	TotalPrice          float64 `gorm:"column:total_price;type:decimal(10,2) unsigned;not null" json:"total_price"` // 总价（商品总价+运费总价）
	IsThird             int8    `gorm:"column:is_third;type:smallint(2) unsigned" json:"is_third"`                  // 是否为第三方平台 0不是，1是
	Status              int8    `gorm:"column:status;type:tinyint(2);not null" json:"status"`                       // 状态 0未支付，1已提交(待付款)  2待发货 3待收货 5:已完成 9已关闭  -1已取消
	Creattime           int     `gorm:"column:creattime;type:int(10) unsigned;not null" json:"creattime"`           // 创建时间
	Updatetime          int     `gorm:"column:updatetime;type:int(10)" json:"updatetime"`                           // 更新时间
	IsDel               int8    `gorm:"column:is_del;type:smallint(2) unsigned" json:"is_del"`
	PurchaseIncrementID string  `gorm:"column:purchase_increment_id;type:char(20)" json:"purchase_increment_id"` // 平台采购流水号
	PurchaseID          int     `gorm:"column:purchase_id;type:int(10)" json:"purchase_id"`                      // 平台采购ID
	ThirdOrderID        string  `gorm:"column:third_order_id;type:char(32)" json:"third_order_id"`               // 第三方平台订单号
	ShipmentID          int8    `gorm:"column:shipment_id;type:smallint(3) unsigned" json:"shipment_id"`         // 快递公司代码  common/status/Logistics
	FreightNumber       string  `gorm:"column:freight_number;type:char(40)" json:"freight_number"`               // 运单号
	FreightTime         int     `gorm:"column:freight_time;type:int(10) unsigned" json:"freight_time"`           // 发货时间
	Remark              string  `gorm:"column:remark;type:text" json:"remark"`                                   // 订单备注
	PaymentNo           string  `gorm:"column:paymentNo;type:char(30)" json:"paymentNo"`                         // 第三方订单支付单号
	ThirdPayStatus      int     `gorm:"column:third_pay_status;type:int(1)" json:"third_pay_status"`             // 第三方订单是否支付（1代表第三方已支付等待回调）
}

type GbPaymentNotifyLog struct {
	ID         int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	OrderNo    string `gorm:"column:order_no;type:varchar(50);not null" json:"order_no"`          // 订单号
	PaymentID  int    `gorm:"column:payment_id;type:int(11);not null" json:"payment_id"`          // 支付方式
	NotifyType string `gorm:"column:notify_type;type:varchar(10);not null" json:"notify_type"`    // 通知方式 1为同步通知2为异步通知3微信退款
	Source     int8   `gorm:"column:source;type:tinyint(1)" json:"source"`                        // 来源 1为H5 2为小程序
	URL        string `gorm:"column:url;type:varchar(250)" json:"url"`                            // 通知网址
	Content    string `gorm:"column:content;type:text" json:"content"`                            // 通知内容
	CreatedAt  int    `gorm:"column:created_at;type:int(10) unsigned;not null" json:"created_at"` // 创建时间
	UpdatedAt  int    `gorm:"column:updated_at;type:int(10) unsigned;not null" json:"updated_at"` // 更新时间
}

type GbPushLog struct {
	ID             int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	PushType       int    `gorm:"index:id;column:push_type;type:int(2)" json:"push_type"`                     // 类别：1计算订单可分配利润接口,2订单取消佣金回退接口
	IncrementID    string `gorm:"index:id;column:increment_id;type:varchar(15);not null" json:"increment_id"` // 订单号
	UId            int    `gorm:"index:id;column:uid;type:int(11);not null" json:"uid"`                       // 用户ID
	PushContent    string `gorm:"column:push_content;type:text" json:"push_content"`                          // 推送内容
	ReceiveContent string `gorm:"column:receive_content;type:text" json:"receive_content"`                    // 接收内容
	IsReceive      int8   `gorm:"index:id;column:is_receive;type:tinyint(1)" json:"is_receive"`               // 0未收到回复1为收到回复
	CreatedAt      int    `gorm:"column:created_at;type:int(10)" json:"created_at"`                           // 创建时间
	UpdatedAt      int    `gorm:"column:updated_at;type:int(10)" json:"updated_at"`                           // 创建时间
}

// GbSupplier 供应商
type GbSupplier struct {
	ID                int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Code              string `gorm:"index;column:code;type:varchar(20);not null" json:"code"`               // 供应商编号
	Applicant         string `gorm:"column:applicant;type:varchar(20);not null" json:"applicant"`           // 联系人
	BackAccount       string `gorm:"column:back_account;type:varchar(100)" json:"back_account"`             // 供应商登陆帐号
	Scope             string `gorm:"column:scope;type:varchar(100)" json:"scope"`                           // 经营项目
	Area              string `gorm:"column:area;type:varchar(100)" json:"area"`                             // 所在地区
	Address           string `gorm:"column:address;type:varchar(100)" json:"address"`                       // 详细地址
	Telphone          string `gorm:"column:telphone;type:varchar(20);not null" json:"telphone"`             // 联系人手机
	Company           string `gorm:"column:company;type:varchar(50);not null" json:"company"`               // 公司名称
	Nature            int8   `gorm:"column:nature;type:tinyint(1) unsigned;not null" json:"nature"`         // 1、普通 2、一般纳税人  3、个人
	RegisteredAddress string `gorm:"column:registered_address;type:varchar(150)" json:"registered_address"` // 注册地址
	OfficePlace       string `gorm:"column:office_place;type:varchar(150)" json:"office_place"`             // 办公地址
	Producer          string `gorm:"column:producer;type:varchar(150)" json:"producer"`                     // 生产厂址
	PassportType      int8   `gorm:"column:passport_type;type:smallint(2) unsigned" json:"passport_type"`   // 供应商类型 0平台供应商 1第三方供应商
	Postcode          int    `gorm:"column:postcode;type:int(6);not null" json:"postcode"`                  // 邮编
	CreateTime        int    `gorm:"column:create_time;type:int(10);not null" json:"create_time"`           // 创建时间
	UpdatedAt         int    `gorm:"column:updated_at;type:int(10)" json:"updated_at"`                      // 更新时间
	Status            int8   `gorm:"column:status;type:tinyint(2);not null" json:"status"`                  // 0、待审核    1、审核通过
	BackPassword      string `gorm:"column:back_password;type:varchar(100)" json:"back_password"`           // 登录密码
	LastLoginTime     int    `gorm:"column:last_login_time;type:int(11)" json:"last_login_time"`
	LastLoginIP       string `gorm:"column:last_login_ip;type:varchar(100)" json:"last_login_ip"`
	TaxNumber         string `gorm:"column:tax_number;type:varchar(30)" json:"tax_number"`                       // 税号
	Payment           int8   `gorm:"column:payment;type:tinyint(1)" json:"payment"`                              // 付款方式(1:月结30天,2:半月结,3:二个月结,4:发货前付清,5:货到付款,6:三个月结,7:周结,8:款到发货,9:货到货运站付款,10:支付宝,11:预付全款,12:预付X%,余下款到发货,13:预付X%(X>=30%),余下货到付款,14:预付X%(X<30%),余下货到付款,15:预付X%(X<30%),余下货到货运站付款)
	Bank              string `gorm:"column:bank;type:varchar(50);not null" json:"bank"`                          // 开户行
	Account           string `gorm:"column:account;type:varchar(20);not null" json:"account"`                    // 账号
	AccountName       string `gorm:"column:account_name;type:varchar(20);not null" json:"account_name"`          // 户名
	CapacityTotal     string `gorm:"column:capacity_total;type:varchar(30)" json:"capacity_total"`               // 总产能
	Remark            string `gorm:"column:remark;type:varchar(250)" json:"remark"`                              // 备注
	BusinessLicence   string `gorm:"column:business_licence;type:varchar(200);not null" json:"business_licence"` // 营业执照
	FoodPermitImg     string `gorm:"column:food_permit_img;type:varchar(200)" json:"food_permit_img"`            // 食品流通许可证
	Mcode             string `gorm:"index;column:mcode;type:varchar(10)" json:"mcode"`                           // 创建人
	Verify            int8   `gorm:"column:verify;type:tinyint(4)" json:"verify"`                                // 审核 0:待供应链审核 1:待财务审核 2:审核未通过-1:存草稿3:财务审核通过
	Reason            string `gorm:"column:reason;type:varchar(255)" json:"reason"`                              // 理由
	Tax               string `gorm:"column:tax;type:varchar(20)" json:"tax"`                                     // 税率
	AgreementImg      string `gorm:"column:agreement_img;type:varchar(255)" json:"agreement_img"`                // 合作协议
	Disabled          int8   `gorm:"column:disabled;type:tinyint(4);not null" json:"disabled"`                   // 0启用1禁用
	AccountImg        string `gorm:"column:account_img;type:varchar(255)" json:"account_img"`                    // 开户许可证
	BillImg           string `gorm:"column:bill_img;type:varchar(255)" json:"bill_img"`                          // 开票资料
	LastConnTime      int    `gorm:"column:last_conn_time;type:int(11);not null" json:"last_conn_time"`          // 最后联系时间
	Source            int8   `gorm:"column:source;type:tinyint(1);not null" json:"source"`                       // 0、前台    1、后台
	GoodsSource       int8   `gorm:"column:goods_source;type:tinyint(1);not null" json:"goods_source"`           // 商品来源：0、境内    1、境外  @author 万朝
	IsDel             int8   `gorm:"column:is_del;type:tinyint(1);not null" json:"is_del"`                       // 0、正常   1、删除
	UserID            int    `gorm:"column:user_id;type:int(11)" json:"user_id"`
	IsSubmit          int8   `gorm:"column:is_submit;type:tinyint(1);not null" json:"is_submit"`         // 0、未提交    1、已提交
	ManagementScope   string `gorm:"column:management_scope;type:text" json:"management_scope"`          // 经营范围
	Profile           string `gorm:"column:profile;type:text" json:"profile"`                            // 简介
	Email             string `gorm:"column:email;type:varchar(50)" json:"email"`                         // 邮箱
	HeadName          string `gorm:"column:head_name;type:varchar(10)" json:"head_name"`                 // 负责人姓名
	HeadTel           string `gorm:"column:head_tel;type:varchar(11)" json:"head_tel"`                   // 负责人手机
	FixedTelephone    string `gorm:"column:fixed_telephone;type:varchar(15)" json:"fixed_telephone"`     // 固定电话
	Fax               string `gorm:"column:fax;type:varchar(15)" json:"fax"`                             // 传真
	OrganizationCode  string `gorm:"column:organization_code;type:varchar(30)" json:"organization_code"` // 组织机构代码
	Province          string `gorm:"column:province;type:varchar(20)" json:"province"`                   // 省
	City              string `gorm:"column:city;type:varchar(20)" json:"city"`                           // 市
	AccountAddress    string `gorm:"column:account_address;type:varchar(20)" json:"account_address"`     // 开户所在地
	Type              int8   `gorm:"column:type;type:tinyint(1)" json:"type"`                            // 1代表京东 2代表爱库存 3代表全球采销联盟
	ShipModelID       int8   `gorm:"column:ship_model_id;type:tinyint(2) unsigned" json:"ship_model_id"` // 1 单品 2是运单
	CnToken           string `gorm:"column:cn_token;type:varchar(255)" json:"cn_token"`                  // 物流token
}

// GbCashLog 提现日志
type GbCashLog struct {
	ID             int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	CashType       int8   `gorm:"column:cash_type;type:tinyint(1);not null" json:"cash_type"`         // 提现类型 1微信
	OutTradeNo     string `gorm:"column:out_trade_no;type:varchar(250);not null" json:"out_trade_no"` // 提现订单号
	PostURL        string `gorm:"column:post_url;type:varchar(255)" json:"post_url"`                  // 提交地址
	PostContent    string `gorm:"column:post_content;type:text" json:"post_content"`                  // 提交内容
	ReceiveContent string `gorm:"column:receive_content;type:text" json:"receive_content"`            // 接收内容
	Status         int8   `gorm:"column:status;type:tinyint(3) unsigned;not null" json:"status"`      // 状态  0：失败 1：成功
	Remark         string `gorm:"column:remark;type:varchar(250)" json:"remark"`                      // 备注
	CreatedAt      int    `gorm:"column:created_at;type:int(10) unsigned;not null" json:"created_at"` // 创建时间
	UpdatedAt      int    `gorm:"column:updated_at;type:int(10) unsigned;not null" json:"updated_at"` // 更新时间
}

type GbOrderAddressReplace struct {
	ID              int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Type            int8   `gorm:"column:type;type:tinyint(2);not null" json:"type"`                           // 来源类型：1,微信小程序 2,网红(陈磊)
	Find            string `gorm:"column:find;type:varchar(250);not null" json:"find"`                         // 查找项
	ReplaceProvince string `gorm:"column:replace_province;type:varchar(100);not null" json:"replace_province"` // 替换省
	ReplaceCity     string `gorm:"column:replace_city;type:varchar(100);not null" json:"replace_city"`         // 替换市
	ReplaceArea     string `gorm:"column:replace_area;type:varchar(100);not null" json:"replace_area"`         // 替换区县
	ReplaceVillage  string `gorm:"column:replace_village;type:varchar(100)" json:"replace_village"`
	Status          int8   `gorm:"column:status;type:tinyint(2);not null" json:"status"`               // 状态：1,正常；2，关闭
	CreatedAt       int    `gorm:"column:created_at;type:int(10) unsigned;not null" json:"created_at"` // 创建时间
	UpdatedAt       int    `gorm:"column:updated_at;type:int(10)" json:"updated_at"`                   // 更新时间
}

type GbMerchantYigroupPro struct {
	ShareImg      string  `gorm:"column:share_img;type:varchar(200)" json:"share_img"` // 小程序分享图片
	ID            int     `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	GroupID       int     `gorm:"column:group_id;type:int(11)" json:"group_id"`               // 团购ID
	MerchantSpuID int     `gorm:"column:merchant_spu_id;type:int(11)" json:"merchant_spu_id"` // 商家商品SPU ID
	MerchantSkuID int     `gorm:"column:merchant_sku_id;type:int(11)" json:"merchant_sku_id"`
	GroupPrice    float64 `gorm:"column:group_price;type:decimal(10,2)" json:"group_price"` // 团购价格
	IsDel         int8    `gorm:"column:is_del;type:tinyint(1)" json:"is_del"`
	BaseShareImg  string  `gorm:"column:base_share_img;type:varchar(250)" json:"base_share_img"`
}

type GbOrderNetRed struct {
	ID                        int     `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	ChildrenIncrementID       string  `gorm:"column:children_increment_id;type:char(20);not null" json:"children_increment_id"`                    // 订单流水号id
	SerialID                  string  `gorm:"column:serial_id;type:char(16)" json:"serial_id"`                                                     // 流水号
	CollectionAmount          float64 `gorm:"column:collection_amount;type:decimal(11,4) unsigned" json:"collection_amount"`                       // 代收额
	Taxes                     float64 `gorm:"column:taxes;type:decimal(11,4) unsigned" json:"taxes"`                                               // 含税平台费
	Proportion                string  `gorm:"column:proportion;type:varchar(255)" json:"proportion"`                                               // 比例
	PaidAmount                float64 `gorm:"primary_key;column:paid_amount;type:decimal(10,4) unsigned;not null" json:"paid_amount"`              // 代付额（原始可体现的金额）
	PayoutBalance             float64 `gorm:"column:payout_balance;type:decimal(10,4)" json:"payout_balance"`                                      // 代付余额
	WithdrawPaidAmount        float64 `gorm:"column:withdraw_paid_amount;type:decimal(10,4) unsigned" json:"withdraw_paid_amount"`                 // 剩余可提现金额
	TaxRate                   float64 `gorm:"column:tax_rate;type:decimal(11,4) unsigned" json:"tax_rate"`                                         // 税率
	UntaxedPlatformServiceFee float64 `gorm:"column:untaxed_platform_service_fee;type:decimal(11,4) unsigned" json:"untaxed_platform_service_fee"` // 未税平台服务费
	Percentage                string  `gorm:"column:percentage;type:varchar(255)" json:"percentage"`                                               // 平台费用百分之八十和百分之二十预存
	Tax                       float64 `gorm:"column:tax;type:decimal(11,4) unsigned" json:"tax"`                                                   // 税额
	CreateTime                int     `gorm:"column:create_time;type:int(11) unsigned" json:"create_time"`                                         // 创建时间
	UpdateTime                int     `gorm:"column:update_time;type:int(11) unsigned zerofill" json:"update_time"`                                // 更新时间
	Type                      int     `gorm:"column:type;type:int(2)" json:"type"`                                                                 // 1表示网红 0表示其他
}

type GbQuexinCode struct {
	ID          int       `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`   // id
	AdminID     int       `gorm:"column:admin_id;type:int(11) unsigned;not null" json:"admin_id"`  // 用户id
	UserID      int       `gorm:"column:user_id;type:int(11) unsigned;not null" json:"user_id"`    // 申请id
	Code        string    `gorm:"unique;column:code;type:char(8);not null" json:"code"`            // 邀请码
	Application string    `gorm:"column:application;type:varchar(20);not null" json:"application"` // 申请人
	Status      int8      `gorm:"column:status;type:tinyint(1);not null" json:"status"`            // 激活状态:0未激活，1已激活
	ActiveAt    time.Time `gorm:"column:active_at;type:timestamp" json:"active_at"`                // 激活时间
	ExpiredAt   time.Time `gorm:"column:expired_at;type:timestamp" json:"expired_at"`              // 过期时间
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamp;not null" json:"created_at"`     // 创建时间
}

type GbWhSupplierProWhUserRel struct {
	ID        int  `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	SupSpuID  int  `gorm:"column:sup_spu_id;type:int(11) unsigned;not null" json:"sup_spu_id"` // 供应商产品SPU id
	WhUserID  int  `gorm:"column:wh_user_id;type:int(11) unsigned;not null" json:"wh_user_id"` // 网红用户ID
	IsDel     int8 `gorm:"column:is_del;type:tinyint(2) unsigned;not null" json:"is_del"`      // 删除状态
	CreatedAt int  `gorm:"column:created_at;type:int(11) unsigned" json:"created_at"`          // 创建时间
}

// GbAccountDetails 平台账户流水
type GbAccountDetails struct {
	ID            int     `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	WorksheetCode string  `gorm:"column:worksheet_code;type:char(32);not null" json:"worksheet_code"`      // 工单编码
	BalanceAction float64 `gorm:"column:balance_action;type:decimal(10,2);not null" json:"balance_action"` // 操作金额动作
	ActionTepy    int8    `gorm:"column:action_tepy;type:tinyint(2);not null" json:"action_tepy"`          // 操作类型：1,购买（采购）；2，提现，3，退款，4，冻结,5，解冻，7，退款临时存储进账，8，退款临时存储出账
	Status        int8    `gorm:"column:status;type:tinyint(2);not null" json:"status"`                    // 状态：1,正常；2，关闭
	Creattime     int     `gorm:"column:creattime;type:int(10) unsigned;not null" json:"creattime"`        // 创建时间
	Updatatime    int     `gorm:"column:updatatime;type:int(10)" json:"updatatime"`                        // 更新时间
}

type GbAdminPermission struct {
	ID        int  `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`      // ID
	RouteID   int  `gorm:"column:route_id;type:int(11);not null" json:"route_id"`              // 路由ID
	RoleID    int  `gorm:"column:role_id;type:int(11);not null" json:"role_id"`                // 角色ID
	Module    int8 `gorm:"column:module;type:tinyint(1);not null" json:"module"`               // 后台模块 1平台端 2商家端 3店铺端 4供应商
	Status    int8 `gorm:"column:status;type:tinyint(1);not null" json:"status"`               // 状态：1有效、0无效
	CreatedAt int  `gorm:"column:created_at;type:int(10) unsigned;not null" json:"created_at"` // 创建时间
	UpdatedAt int  `gorm:"column:updated_at;type:int(10) unsigned;not null" json:"updated_at"` // 更新时间
}

type GbCollectSites struct {
	ID        int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	Name      string `gorm:"column:name;type:varchar(64);not null" json:"name"`          // 名称
	BaseURL   string `gorm:"column:base_url;type:varchar(255);not null" json:"base_url"` // url
	Sort      int8   `gorm:"column:sort;type:tinyint(3)" json:"sort"`                    // 排序
	IsDel     int8   `gorm:"column:is_del;type:tinyint(1)" json:"is_del"`
	Status    int8   `gorm:"column:status;type:tinyint(1)" json:"status"`      // 状态  1启用  0禁用
	CreatedAt int    `gorm:"column:created_at;type:int(11)" json:"created_at"` // 创建时间
	UpdatedAt int    `gorm:"column:updated_at;type:int(11)" json:"updated_at"` // 更新时间
}

type GbCrmWorkOrderCate struct {
	ID          int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Name        string    `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Description string    `gorm:"column:description;type:varchar(255);not null" json:"description"`
	Status      int8      `gorm:"column:status;type:tinyint(4);not null" json:"status"`
	Createtime  time.Time `gorm:"column:createtime;type:datetime;not null" json:"createtime"`
	Companyid   int       `gorm:"column:companyid;type:int(11);not null" json:"companyid"`
}

type GbMerchantYigroup struct {
	ID          int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Name        string `gorm:"column:name;type:varchar(50)" json:"name"`            // 名称
	Description string `gorm:"column:description;type:tinytext" json:"description"` // 描述
	Ctime       int    `gorm:"column:ctime;type:int(11)" json:"ctime"`              // 开始时间
	Stime       int    `gorm:"column:stime;type:int(11)" json:"stime"`              // 开始时间
	Etime       int    `gorm:"column:etime;type:int(11)" json:"etime"`              // 结束时间
	GoodsNum    int    `gorm:"column:goods_num;type:int(11)" json:"goods_num"`      // 商品数量
	SellerID    int    `gorm:"column:seller_id;type:int(11)" json:"seller_id"`      // 商家ID
	Status      int8   `gorm:"column:status;type:tinyint(1)" json:"status"`         // 状态  1 活动进行中，2 活动结束
	IsDel       int8   `gorm:"column:is_del;type:tinyint(1)" json:"is_del"`
}

type GbRefundGoodsStatus struct {
	ID            int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	From          string `gorm:"column:from;type:varchar(255)" json:"from"`                          // 来自于哪一个平台
	RefundGoodsID int    `gorm:"column:refund_goods_id;type:int(11)" json:"refund_goods_id"`         // 退货ID
	Worksheet     string `gorm:"column:worksheet;type:char(32)" json:"worksheet"`                    // 退货工单
	Status        int8   `gorm:"column:status;type:tinyint(2);not null" json:"status"`               // 当前状态
	UserName      string `gorm:"column:user_name;type:varchar(255)" json:"user_name"`                // 操作人
	Creattime     int    `gorm:"column:creattime;type:int(10) unsigned;not null" json:"creattime"`   // 创建时间
	Updatetime    int    `gorm:"column:updatetime;type:int(10) unsigned;not null" json:"updatetime"` // 修改时间
	BeforeStatus  int8   `gorm:"column:before_status;type:tinyint(2)" json:"before_status"`          // 修改前状态
	Remarks       string `gorm:"column:remarks;type:varchar(255)" json:"remarks"`                    // 备注
	Operating     int8   `gorm:"column:operating;type:tinyint(2)" json:"operating"`                  // 操作事项
}

// GbMerchantGroupRelaPro 产品参与拼团关联表
type GbMerchantGroupRelaPro struct {
	ID          int     `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	GroupID     int     `gorm:"column:group_id;type:int(11);not null" json:"group_id"`   // 参团id
	Peoplecount int     `gorm:"column:peoplecount;type:int(11)" json:"peoplecount"`      // 参团人数限定
	Discount    float32 `gorm:"column:discount;type:float" json:"discount"`              // 拼团优惠价格
	MpProID     int     `gorm:"column:mp_pro_id;type:int(11);not null" json:"mp_pro_id"` // 商户产品id
	ProID       int     `gorm:"column:pro_id;type:int(11)" json:"pro_id"`                // 店铺产品id
	IsDel       int8    `gorm:"column:is_del;type:tinyint(4);not null" json:"is_del"`    // 逻辑删除状态0.未删除1.已删除
}

// GbMerchantSeckill 秒杀活动表
type GbMerchantSeckill struct {
	ID         int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	MerchantID int    `gorm:"column:merchant_id;type:int(11)" json:"merchant_id"` // 创建商家ID
	Name       string `gorm:"column:name;type:varchar(50)" json:"name"`           // 秒杀活动名称
	StartTime  int    `gorm:"column:start_time;type:int(11)" json:"start_time"`   // 开始时间
	EndTime    int    `gorm:"column:end_time;type:int(11)" json:"end_time"`       // 结束时间
	Status     int8   `gorm:"column:status;type:tinyint(1)" json:"status"`        // 状态 1 上线  0下线
	IsDel      int8   `gorm:"column:is_del;type:tinyint(1)" json:"is_del"`
	TopSkuNum  int8   `gorm:"column:top_sku_num;type:tinyint(1)" json:"top_sku_num"` // 置顶SKU数量
}

type GbOrderItemPayDetails struct {
	ID                  int     `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`                           // 订单商品支付详情表
	ItemID              int     `gorm:"index;column:item_id;type:int(11) unsigned;not null" json:"item_id"`                      // 订单商品ID
	IncrementID         string  `gorm:"index;column:increment_id;type:char(20);not null" json:"increment_id"`                    // 总订单流水号:123456789451245
	PayJifen            float64 `gorm:"column:pay_jifen;type:decimal(10,2);not null" json:"pay_jifen"`                           // 花费积分
	PayMoney            float64 `gorm:"column:pay_money;type:decimal(10,2);not null" json:"pay_money"`                           // 实际支付金额
	Creattime           int     `gorm:"column:creattime;type:int(10) unsigned;not null" json:"creattime"`                        // 创建时间
	Updatatime          int     `gorm:"column:updatatime;type:int(10) unsigned;not null" json:"updatatime"`                      // 更新时间
	Commision           float64 `gorm:"column:commision;type:decimal(10,2);not null" json:"commision"`                           // 优惠券金额
	OrderItemCountPrice float64 `gorm:"column:order_item_count_price;type:decimal(10,2);not null" json:"order_item_count_price"` // 商品总价
	Freight             float64 `gorm:"column:freight;type:decimal(10,2);not null" json:"freight"`                               // 商品运费总价
	UnitCommission      float64 `gorm:"column:unit_commission;type:decimal(10,2)" json:"unit_commission"`                        // 单个商品佣金金额
	CountCommission     float64 `gorm:"column:count_commission;type:decimal(10,2)" json:"count_commission"`                      // 商品总佣金金额
	CrossBorderTax      float64 `gorm:"column:cross_border_tax;type:decimal(10,2);not null" json:"cross_border_tax"`             // 跨境综合税额度总额度
}

type GbVirtualCare struct {
	ID          int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	BatchNumber string `gorm:"column:batch_number;type:varchar(30);not null" json:"batch_number"` // 批次号
	VirtualID   int    `gorm:"column:virtual_id;type:int(11);not null" json:"virtual_id"`         // 关联virtual表id
	CareNumber  string `gorm:"column:care_number;type:varchar(100);not null" json:"care_number"`  // 卡号或兑换码
	CarePass    string `gorm:"column:care_pass;type:varchar(10)" json:"care_pass"`                // 卡密
	EndTime     int    `gorm:"column:end_time;type:int(11)" json:"end_time"`                      // 过期时间
	EndNum      int    `gorm:"column:end_num;type:int(2)" json:"end_num"`                         // 过期天数
	Status      int8   `gorm:"column:status;type:tinyint(1)" json:"status"`                       // 使用状态 1：未使用 2：已使用 3：已过期
	CreateTime  int    `gorm:"column:create_time;type:int(11)" json:"create_time"`                // 创建时间
	OrderID     int    `gorm:"column:order_id;type:int(11)" json:"order_id"`                      // 订单id
	Number      int    `gorm:"column:number;type:int(4)" json:"number"`                           // 批次号尾部四位数
	IncermentID string `gorm:"column:incerment_id;type:varchar(30)" json:"incerment_id"`          // 订单流水号
}

// GbMerchantShareModelTemplate 分成模板，确定店铺获取商品的边界
type GbMerchantShareModelTemplate struct {
	ID             int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	ModelID        int    `gorm:"column:model_id;type:int(11)" json:"model_id"`                 // 所属的模式ID
	Name           string `gorm:"column:name;type:varchar(255)" json:"name"`                    // 模板名称
	Description    string `gorm:"column:description;type:varchar(255)" json:"description"`      // 模板描述
	SellerRate     int8   `gorm:"column:seller_rate;type:tinyint(3)" json:"seller_rate"`        // 商家分成比例1-100
	StoreRate      int8   `gorm:"column:store_rate;type:tinyint(3)" json:"store_rate"`          // 店铺分成比例1-100
	GUIDeRate      int8   `gorm:"column:guide_rate;type:tinyint(3)" json:"guide_rate"`          // 导购分成比例1-100
	PromotionRate  int8   `gorm:"column:promotion_rate;type:tinyint(3)" json:"promotion_rate"`  // 推广分成占比
	RelateStoreNum int    `gorm:"column:relate_store_num;type:int(11)" json:"relate_store_num"` // 关联的店铺数量
	Ctime          int    `gorm:"column:ctime;type:int(10) unsigned" json:"ctime"`              // 创建时间
	IsDel          int8   `gorm:"column:is_del;type:tinyint(1)" json:"is_del"`                  // 删除状态 0 未删除 1 已删除
}

type GbMerchantYigroupMaster struct {
	ID               int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Serial           string `gorm:"column:serial;type:varchar(20)" json:"serial"`                      // 壹团番号
	MasterMobile     string `gorm:"column:master_mobile;type:varchar(20)" json:"master_mobile"`        // 团长手机号
	PromoMobile      string `gorm:"column:promo_mobile;type:varchar(20)" json:"promo_mobile"`          // 推荐人手机号
	ApplyTime        int    `gorm:"column:apply_time;type:int(11)" json:"apply_time"`                  // 团长申请时间
	GroupImage       string `gorm:"column:group_image;type:varchar(200)" json:"group_image"`           // 团购群截图
	UId              int    `gorm:"column:uid;type:int(11)" json:"uid"`                                // 用户ID
	Fullname         string `gorm:"column:fullname;type:varchar(20)" json:"fullname"`                  // 姓名
	IDcard           string `gorm:"column:idcard;type:varchar(20)" json:"idcard"`                      // 身份证号
	IDcardFront      string `gorm:"column:idcard_front;type:varchar(255)" json:"idcard_front"`         // 身份证正面
	IDcardBankend    string `gorm:"column:idcard_bankend;type:varchar(255)" json:"idcard_bankend"`     // 身份证反面
	ExStatus         int8   `gorm:"column:ex_status;type:tinyint(1)" json:"ex_status"`                 // 审核状态  0未审核 1已审核  2,审核未通过
	IsIDentification int8   `gorm:"column:is_identification;type:tinyint(1)" json:"is_identification"` // 是否实名认证 0为待认证 1为认证通过 2为认证失败
	IsDel            int8   `gorm:"column:is_del;type:tinyint(1)" json:"is_del"`                       // 删除状态 0 未删除 1 已经删除
	WxAccount        string `gorm:"column:wx_account;type:varchar(30)" json:"wx_account"`              // wx_name
}

// GbPurchase 平台采购订单信息表
type GbPurchase struct {
	ID                  int     `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	PurchaseIncrementID string  `gorm:"column:purchase_increment_id;type:char(26)" json:"purchase_increment_id"`    // 平台采购流水号-分号
	SupplierCode        string  `gorm:"column:supplier_code;type:char(20);not null" json:"supplier_code"`           // 供应商编码
	SupplierID          string  `gorm:"column:supplier_id;type:char(20);not null" json:"supplier_id"`               // 供应商编码
	Freight             float64 `gorm:"column:freight;type:decimal(10,2) unsigned;not null" json:"freight"`         // 运费总价
	GoodsPrice          float64 `gorm:"column:goods_price;type:decimal(10,2) unsigned;not null" json:"goods_price"` // 商品总价
	TotalPrice          float64 `gorm:"column:total_price;type:decimal(10,2) unsigned;not null" json:"total_price"` // 订单总价（商品总价+运费总价）
	Status              int8    `gorm:"column:status;type:tinyint(2);not null" json:"status"`                       // 订单状态： 待确认 1； 已确认  2；
	IsConfirm           int8    `gorm:"column:is_confirm;type:tinyint(2) unsigned" json:"is_confirm"`               // 客服确认 0未确认 1确认
	Creattime           int     `gorm:"column:creattime;type:int(10) unsigned;not null" json:"creattime"`           // 创建时间
	Updatetime          int     `gorm:"column:updatetime;type:int(10) unsigned;not null" json:"updatetime"`         // 修改时间
	CountOrders         int     `gorm:"column:count_orders;type:int(10)" json:"count_orders"`                       // 订单数量
	CountChildrenOrders int     `gorm:"column:count_children_orders;type:int(10)" json:"count_children_orders"`     // 子订单数
	CountGoodsTepy      int     `gorm:"column:count_goods_tepy;type:int(10)" json:"count_goods_tepy"`               // 商品种类
	CountGoods          int     `gorm:"column:count_goods;type:int(10)" json:"count_goods"`                         // 商品数量
	AdminID             int     `gorm:"column:admin_id;type:int(10)" json:"admin_id"`                               // 操作人（平台用户ID）
	IndexPurchaseID     int     `gorm:"column:index_purchase_id;type:int(10)" json:"index_purchase_id"`             // 采购主表ID（关联字段）
}

type GbWebsiteContent struct {
	ID            int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`       // 内容ID
	WebsiteType   int8      `gorm:"column:website_type;type:tinyint(1)" json:"website_type"`    // 官网类型 1神雀 2鼎翰
	Founder       string    `gorm:"column:founder;type:varchar(255)" json:"founder"`            // 创建人
	Name          string    `gorm:"column:name;type:varchar(255)" json:"name"`                  // 页面名称
	Banner        string    `gorm:"column:banner;type:varchar(255)" json:"banner"`              // banner图标
	PublishType   int8      `gorm:"column:publish_type;type:tinyint(1)" json:"publish_type"`    // 发布类型 1新闻 2活动 3合作案例
	ChannelTypeID int       `gorm:"column:channel_type_id;type:int(10)" json:"channel_type_id"` // 发布渠道ID
	NewsTypeID    int       `gorm:"column:news_type_id;type:int(10)" json:"news_type_id"`       // 对应的分类ID
	Title         string    `gorm:"column:title;type:varchar(255)" json:"title"`                // 标题
	ViceTitle     string    `gorm:"column:vice_title;type:varchar(255)" json:"vice_title"`      // 副标题
	Author        string    `gorm:"column:author;type:varchar(255)" json:"author"`              // 作者
	Info          string    `gorm:"column:info;type:text" json:"info"`                          // 简介
	Content       string    `gorm:"column:content;type:text" json:"content"`                    // 正文
	Tag           string    `gorm:"column:tag;type:text" json:"tag"`                            // 新闻-标签 合作案例-案例 json格式
	Top           int8      `gorm:"column:top;type:tinyint(1)" json:"top"`                      // 置顶 0未置顶 1置顶
	Shelf         int8      `gorm:"column:shelf;type:tinyint(1)" json:"shelf"`                  // 上、下架状态 1上架 2下架
	Status        int8      `gorm:"column:status;type:tinyint(1)" json:"status"`                // 状态 0删除 1正常
	CreateTime    time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`        // 创建时间
	UpdateTime    time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`        // 更新时间
}

type GbLogOrderChange struct {
	ID          int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	UId         int    `gorm:"column:uid;type:int(11) unsigned;not null" json:"uid"`                 // 操作者ID
	IncrementID string `gorm:"index;column:increment_id;type:char(20);not null" json:"increment_id"` // 订单流水号
	Remarks     string `gorm:"column:remarks;type:varchar(255)" json:"remarks"`                      // 备注
	AfterStatus int    `gorm:"column:after_status;type:int(11);not null" json:"after_status"`        // 处理后状态
	Creattime   int    `gorm:"column:creattime;type:int(10)" json:"creattime"`                       // 创建时间
}

type GbOrderPayDetails struct {
	ID                  int     `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`                           // 订单支付详情表
	OrderID             int     `gorm:"index;column:order_id;type:int(11) unsigned;not null" json:"order_id"`                    // 订单商品ID
	IncrementID         string  `gorm:"index;column:increment_id;type:char(20);not null" json:"increment_id"`                    // 总订单流水号:123456789451245
	OrderCountPrice     float64 `gorm:"column:order_count_price;type:decimal(10,2)" json:"order_count_price"`                    // 订单理论总金额
	Freight             float64 `gorm:"column:freight;type:decimal(10,2)" json:"freight"`                                        // 运费
	Commision           float64 `gorm:"column:commision;type:decimal(10,2)" json:"commision"`                                    // 优惠金额
	PayJifen            float64 `gorm:"column:pay_jifen;type:decimal(10,2);not null" json:"pay_jifen"`                           // 花费积分
	PayMoney            float64 `gorm:"column:pay_money;type:decimal(10,2);not null" json:"pay_money"`                           // 花费金额
	Creattime           int     `gorm:"column:creattime;type:int(10) unsigned;not null" json:"creattime"`                        // 创建时间
	Updatatime          int     `gorm:"column:updatatime;type:int(10) unsigned;not null" json:"updatatime"`                      // 更新时间
	LaterAdd            int8    `gorm:"column:later_add;type:tinyint(4);not null" json:"later_add"`                              // 0表示是之前有的数据，1表示是后面手工加的
	CountCommission     float64 `gorm:"column:count_commission;type:decimal(10,2);not null" json:"count_commission"`             // 订单总佣金金额
	TotalCrossBorderTax float64 `gorm:"column:total_cross_border_tax;type:decimal(10,2);not null" json:"total_cross_border_tax"` // 该订单总跨境综合税额度额度
}

// GbProCommentLiked 评论点赞记录
type GbProCommentLiked struct {
	ID         int  `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	UId        int  `gorm:"column:uid;type:int(11) unsigned;not null" json:"uid"`                   // 用户ID
	ProID      int  `gorm:"column:pro_id;type:int(11) unsigned;not null" json:"pro_id"`             // 产品ID
	ProStandID int  `gorm:"column:pro_stand_id;type:int(11) unsigned;not null" json:"pro_stand_id"` // 产品规格ID
	CommentID  int  `gorm:"column:comment_id;type:int(11) unsigned;not null" json:"comment_id"`     // 评论表ID
	IsDel      int8 `gorm:"column:is_del;type:smallint(2) unsigned;not null" json:"is_del"`         // 0已点赞 1取消点赞
}

type GbStoreColumnNameBind struct {
	ID          int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	StoreID     int    `gorm:"index;column:store_id;type:int(11) unsigned;not null" json:"store_id"`     // 店铺ID
	ColumnIndex string `gorm:"index;column:column_index;type:varchar(255);not null" json:"column_index"` // 模块索引
	ColumnName  string `gorm:"column:column_name;type:varchar(255);not null" json:"column_name"`         // 栏目定制名称
	CreatedAt   int    `gorm:"column:created_at;type:int(11) unsigned;not null" json:"created_at"`
	UpdatedAt   int    `gorm:"column:updated_at;type:int(11) unsigned;not null" json:"updated_at"`
}

type GbStoreHotsaleCate struct {
	ID        int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	StoreID   int    `gorm:"column:store_id;type:int(11);not null" json:"store_id"`           // 店铺ID
	Title     string `gorm:"column:title;type:varchar(250);not null" json:"title"`            // 类名
	Image     string `gorm:"column:image;type:varchar(250);not null" json:"image"`            // 图片地址
	IsTop     int8   `gorm:"column:is_top;type:tinyint(2) unsigned;not null" json:"is_top"`   // 是否置顶  1置顶 0未置顶
	Display   int8   `gorm:"column:display;type:tinyint(2) unsigned;not null" json:"display"` // 展示状态  0不展示  1展示
	Sort      int    `gorm:"column:sort;type:int(11);not null" json:"sort"`                   // 排序
	CreatecAt int    `gorm:"column:createc_at;type:int(10) unsigned" json:"createc_at"`       // 创建时间
	UpdatedAt int    `gorm:"column:updated_at;type:int(10)" json:"updated_at"`                // 更新时间
}

// GbUserAccountDetails 用户账户流水
type GbUserAccountDetails struct {
	ID                int     `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	AccountID         int     `gorm:"column:account_id;type:int(11) unsigned;not null" json:"account_id"`         // 账户ID
	UserID            int     `gorm:"column:user_id;type:int(11) unsigned;not null" json:"user_id"`               // 用户ID
	OrderIncrementID  string  `gorm:"column:order_increment_id;type:char(32);not null" json:"order_increment_id"` // 订单流水号ID
	WorksheetCode     string  `gorm:"column:worksheet_code;type:char(32)" json:"worksheet_code"`                  // 工单编码
	Commission        float64 `gorm:"column:commission;type:decimal(10,6) unsigned;not null" json:"commission"`   // 返利佣金
	SurplusCommission float64 `gorm:"column:surplus_commission;type:decimal(10,6)" json:"surplus_commission"`     // 部分提现后剩余金额
	CashStatus        int8    `gorm:"column:cash_status;type:tinyint(1)" json:"cash_status"`                      // 提现状态 0未提现 1已提现 2部分提现
	ActionTepy        int8    `gorm:"column:action_tepy;type:tinyint(2) unsigned;not null" json:"action_tepy"`    // 操作类型：1进帐(冻结)  2提现(解冻)  3佣金已退，4第三方用户购买，5第三方用户退款
	Status            int8    `gorm:"column:status;type:tinyint(2) unsigned;not null" json:"status"`              // 状态：1,正常(未领冻结佣金)；2，已关闭(已领冻结佣金)
	SecretKey         string  `gorm:"column:secret_key;type:char(20);not null" json:"secret_key"`                 // 保存佣金写进来时的 分配规则编码 gb_app_profit
	Creattime         int     `gorm:"column:creattime;type:int(10) unsigned;not null" json:"creattime"`           // 创建时间
	Updatatime        int     `gorm:"column:updatatime;type:int(10) unsigned" json:"updatatime"`                  // 更新时间
}

type GbAdminRoute struct {
	ID        int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`      // ID
	Name      string `gorm:"column:name;type:varchar(50);not null" json:"name"`                  // 名称
	Route     string `gorm:"column:route;type:varchar(50);not null" json:"route"`                // 路由
	ParentID  int    `gorm:"column:parent_id;type:int(11);not null" json:"parent_id"`            // 父ID
	IsMenu    int8   `gorm:"column:is_menu;type:tinyint(1);not null" json:"is_menu"`             // 是否菜单0否1是
	Module    int8   `gorm:"column:module;type:tinyint(1);not null" json:"module"`               // 后台模块 1平台端 2商家端 3店铺端 4供应商
	Status    int8   `gorm:"column:status;type:tinyint(1);not null" json:"status"`               // 状态：1有效、0无效
	Icon      string `gorm:"column:icon;type:varchar(30);not null" json:"icon"`                  // 图标
	Sort      int8   `gorm:"column:sort;type:tinyint(2);not null" json:"sort"`                   // 排序
	CreatedAt int    `gorm:"column:created_at;type:int(10) unsigned;not null" json:"created_at"` // 创建时间
	UpdatedAt int    `gorm:"column:updated_at;type:int(10) unsigned;not null" json:"updated_at"` // 更新时间
}

// GbSection 云平台 部门|角色表
type GbSection struct {
	ID     int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"` // 部门表
	Name   string `gorm:"column:name;type:varchar(255);not null" json:"name"`            // 部门名称
	Desc   string `gorm:"column:desc;type:text;not null" json:"desc"`                    // 部门描述
	Enable int8   `gorm:"column:enable;type:tinyint(2) unsigned;not null" json:"enable"` // 启用状态 1为启用
	IsDel  int8   `gorm:"column:is_del;type:tinyint(2) unsigned;not null" json:"is_del"` // 删除状态 1为已删除
	Menu   string `gorm:"column:menu;type:text" json:"menu"`                             // 能见菜单
	Ctime  int    `gorm:"column:ctime;type:int(10) unsigned;not null" json:"ctime"`
}

// GbCrmPosition CRM职位表
type GbCrmPosition struct {
	ID        int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`      // 职位ID
	CompanyID int    `gorm:"column:company_id;type:int(11);not null" json:"company_id"` // 关联的公司ID
	SectionID int    `gorm:"column:section_id;type:int(11);not null" json:"section_id"` // 关联的部门ID
	Name      string `gorm:"column:name;type:varchar(255);not null" json:"name"`        // 职位名称
	Desc      string `gorm:"column:desc;type:text;not null" json:"desc"`                // 职位描述
	Enable    int8   `gorm:"column:enable;type:tinyint(2)" json:"enable"`               // 启用状态  1为启用
	IsDel     int8   `gorm:"column:is_del;type:tinyint(2)" json:"is_del"`               // 删除状态  1位删除
	Ctime     int    `gorm:"column:ctime;type:int(10);not null" json:"ctime"`           // 创建时间
}

// GbSupAccountDetails 供应商账户流水
type GbSupAccountDetails struct {
	ID            int     `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	AccountID     int     `gorm:"column:account_id;type:int(11) unsigned;not null" json:"account_id"`      // 账户ID
	WorksheetCode string  `gorm:"column:worksheet_code;type:char(32);not null" json:"worksheet_code"`      // 工单编码
	BalanceAction float64 `gorm:"column:balance_action;type:decimal(10,2);not null" json:"balance_action"` // 操作金额动作
	ActionTepy    int8    `gorm:"column:action_tepy;type:tinyint(2);not null" json:"action_tepy"`          // 操作类型：1,收单；2，提现，3，退款，4，冻,5，解冻
	Status        int8    `gorm:"column:status;type:tinyint(2);not null" json:"status"`                    // 状态：1,正常；2，关闭
	Creattime     int     `gorm:"column:creattime;type:int(10) unsigned;not null" json:"creattime"`        // 创建时间
	Updatatime    int     `gorm:"column:updatatime;type:int(10)" json:"updatatime"`                        // 更新时间
}

// GbTaobaoProps 淘宝类目属性表
type GbTaobaoProps struct {
	ID            int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Cid           int    `gorm:"index;column:cid;type:int(11);not null" json:"cid"`                    // 类目ID
	Pid           int    `gorm:"index;column:pid;type:int(11);not null" json:"pid"`                    // 属性ID
	ParentPid     int    `gorm:"column:parent_pid;type:int(11);not null" json:"parent_pid"`            // 上级属性ID
	ParentVid     int    `gorm:"column:parent_vid;type:int(11);not null" json:"parent_vid"`            // 上级属性值ID
	Name          string `gorm:"column:name;type:varchar(50);not null" json:"name"`                    // 属性名称
	IsKeyProp     int8   `gorm:"column:is_key_prop;type:tinyint(1);not null" json:"is_key_prop"`       // 是否关键属性：1为是
	IsSaleProp    int8   `gorm:"column:is_sale_prop;type:tinyint(1);not null" json:"is_sale_prop"`     // 是否销售属性：1为是
	IsColorProp   int8   `gorm:"column:is_color_prop;type:tinyint(1);not null" json:"is_color_prop"`   // 是否颜色属性：1为是
	IsEnumProp    int8   `gorm:"column:is_enum_prop;type:tinyint(1);not null" json:"is_enum_prop"`     // 是否是可枚举属性：1为是
	IsInputProp   int8   `gorm:"column:is_input_prop;type:tinyint(1);not null" json:"is_input_prop"`   // 是否是卖家可以自行输入的属性：1为是
	IsItemProp    int8   `gorm:"column:is_item_prop;type:tinyint(1);not null" json:"is_item_prop"`     // 是否商品属性：1为是
	ChildTemplate string `gorm:"column:child_template;type:varchar(50)" json:"child_template"`         // 子属性的模板（卖家自行输入属性时需要用到）
	IsAllowAlias  int8   `gorm:"column:is_allow_alias;type:tinyint(1);not null" json:"is_allow_alias"` // 是否允许别名：1为是
	IsTaosir      int8   `gorm:"column:is_taosir;type:tinyint(1);not null" json:"is_taosir"`           // 是否度量衡属性项：1为是
	TaosirDo      string `gorm:"column:taosir_do;type:text" json:"taosir_do"`                          // 度量衡相关信息集合
	IsMaterial    int8   `gorm:"column:is_material;type:tinyint(1);not null" json:"is_material"`       // 是否是材质属性项：1为是
	MaterialDo    string `gorm:"column:material_do;type:text" json:"material_do"`                      // 材质属性信息集合
	Multi         int8   `gorm:"column:multi;type:tinyint(1);not null" json:"multi"`                   // 发布产品或商品时是否可以多选
	Must          int8   `gorm:"column:must;type:tinyint(1);not null" json:"must"`                     // 发布产品或商品时是否为必选属性
	Value         string `gorm:"column:value;type:varchar(50)" json:"value"`                           // 属性所匹配过的值
	IsFixedValue  int8   `gorm:"column:is_fixed_value;type:tinyint(1);not null" json:"is_fixed_value"` // 是否固定属性值
	IsFinished    int8   `gorm:"column:is_finished;type:tinyint(1);not null" json:"is_finished"`       // 是否完成属性匹配
}

type GbVirtual struct {
	ID            int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	ProductID     string `gorm:"column:product_id;type:varchar(60)" json:"product_id"`       // 商品编号的部分信息
	Name          string `gorm:"column:name;type:varchar(150)" json:"name"`                  // 产品名称
	Specification string `gorm:"column:specification;type:varchar(20)" json:"specification"` // 规格 1：周卡  2：月卡
	Cots          int    `gorm:"column:cots;type:int(10);not null" json:"cots"`              // 面值
	Type          int8   `gorm:"column:type;type:tinyint(1);not null" json:"type"`           // 类型 1:卡密 2:兑换码
	SupplierIDs   int    `gorm:"column:supplier_ids;type:int(11)" json:"supplier_ids"`       // 供货商id
	CreateTime    int    `gorm:"column:create_time;type:int(11)" json:"create_time"`         // 创建时间
	Number        string `gorm:"column:number;type:varchar(4)" json:"number"`                // 编号的后四位随机
	SupplierID    int    `gorm:"column:supplier_id;type:int(11)" json:"supplier_id"`         // 供应商账号ID
}

type GbPaymentTerminal struct {
	ID        int  `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Terminal  int8 `gorm:"column:terminal;type:tinyint(1);not null" json:"terminal"`           // 终端  1：H5端，2:小程序端，3：APP，4：PC
	PaymentID int  `gorm:"column:payment_id;type:int(2);not null" json:"payment_id"`           // 支付方式
	Sort      int  `gorm:"column:sort;type:int(11)" json:"sort"`                               // 排序，倒序
	Status    int8 `gorm:"column:status;type:tinyint(1)" json:"status"`                        // 是否启用 0为未启用 1为启用
	CreatedAt int  `gorm:"column:created_at;type:int(10) unsigned;not null" json:"created_at"` // 创建时间
	UpdatedAt int  `gorm:"column:updated_at;type:int(10) unsigned;not null" json:"updated_at"` // 更新时间
}

// GbSupProParam 供应商产品参数
type GbSupProParam struct {
	ID      int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	SupID   int    `gorm:"column:sup_id;type:int(11) unsigned;not null" json:"sup_id"`    // 供应商ID
	ProID   int    `gorm:"index;column:pro_id;type:int(11)" json:"pro_id"`                // 产品ID
	ProCode string `gorm:"column:pro_code;type:varchar(60);not null" json:"pro_code"`     // 产品编码 ENA13
	SupCode string `gorm:"column:sup_code;type:varchar(40);not null" json:"sup_code"`     // 供应商编号
	Name    string `gorm:"column:name;type:varchar(20);not null" json:"name"`             // 参数名称
	Value   string `gorm:"column:value;type:varchar(220);not null" json:"value"`          // 参数值
	Sort    int8   `gorm:"column:sort;type:smallint(4) unsigned;not null" json:"sort"`    // 排序
	IsDel   int8   `gorm:"column:is_del;type:tinyint(1) unsigned;not null" json:"is_del"` // 是否删除  1：是   0：否
}

type GbExchange struct {
	ID                  int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	Worksheet           string `gorm:"column:worksheet;type:char(32);not null" json:"worksheet"`                         // 服务工单号
	ChildrenIncrementID string `gorm:"column:children_increment_id;type:char(32);not null" json:"children_increment_id"` // 子订单流水号
	Cause               int8   `gorm:"column:cause;type:tinyint(2);not null" json:"cause"`                               // 原因
	Recipients          string `gorm:"column:recipients;type:varchar(10)" json:"recipients"`                             // 联系人
	Tel                 string `gorm:"column:tel;type:char(11)" json:"tel"`                                              // 电话
	Description         string `gorm:"column:description;type:varchar(255)" json:"description"`                          // 问题描述
	Voucher             string `gorm:"column:voucher;type:varchar(255)" json:"voucher"`                                  // 图片凭证，多图用，号隔开
	MerchantAcceptance  int8   `gorm:"column:merchant_acceptance;type:tinyint(2);not null" json:"merchant_acceptance"`   // 商家受理情况，1受理，2不受理
	MerchantReason      string `gorm:"column:merchant_reason;type:varchar(255)" json:"merchant_reason"`                  // 商家理由
	PlatformAcceptance  int8   `gorm:"column:platform_acceptance;type:tinyint(2);not null" json:"platform_acceptance"`   // 平台受理情况，1受理，2不受理
	PlatformReason      string `gorm:"column:platform_reason;type:varchar(255)" json:"platform_reason"`                  // 平台理由
	SupAcceptance       int8   `gorm:"column:sup_acceptance;type:tinyint(2);not null" json:"sup_acceptance"`             // 供应商受理情况，1受理，2不受理
	SupReason           string `gorm:"column:sup_reason;type:varchar(255)" json:"sup_reason"`                            // 供应商理由
	FreightNumber       string `gorm:"column:freight_number;type:char(30)" json:"freight_number"`                        // 用户发货物流单号
	FreightTime         int    `gorm:"column:freight_time;type:int(10) unsigned" json:"freight_time"`                    // 用户发货时间
	LogisticsID         int    `gorm:"column:logistics_id;type:int(10)" json:"logistics_id"`                             // 用户发货快递公司ID
	FreightTel          string `gorm:"column:freight_tel;type:char(11)" json:"freight_tel"`                              // 用户发货物流联系人电话
	SupFreightNumber    string `gorm:"column:sup_freight_number;type:char(30)" json:"sup_freight_number"`                // 供应商发货物流单号
	SupFreightTime      int    `gorm:"column:sup_freight_time;type:int(10) unsigned" json:"sup_freight_time"`            // 供应商发货时间
	SupLogisticsID      int    `gorm:"column:sup_logistics_id;type:int(10)" json:"sup_logistics_id"`                     // 供应商发货快递公司ID
	Status              int8   `gorm:"column:status;type:tinyint(2);not null" json:"status"`                             // 状态：0,不受理驳回，1发起申请（等待商家处理），2等待平台受理，3等待供应商受理，4等待用户填入快递单号，5等待供应商接收商品，6供应商发货，7用户收货流程完成
	Creattime           int    `gorm:"column:creattime;type:int(10) unsigned;not null" json:"creattime"`                 // 创建时间
	Updatatime          int    `gorm:"column:updatatime;type:int(10)" json:"updatatime"`                                 // 更新时间
	SupplierCode        string `gorm:"column:supplier_code;type:char(32)" json:"supplier_code"`                          // 供应商编号
	SupplierID          int    `gorm:"column:supplier_id;type:int(10)" json:"supplier_id"`                               // 供应商ID
	MerchantCode        string `gorm:"column:merchant_code;type:char(32)" json:"merchant_code"`                          // 商家编号
}

// GbMerchantBargainTextConfig 砍价活动文案
type GbMerchantBargainTextConfig struct {
	ID                       int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Name                     string `gorm:"column:name;type:varchar(50)" json:"name"`                                                // 文案名称
	FocusBtnText             string `gorm:"column:focus_btn_text;type:varchar(50)" json:"focus_btn_text"`                            // 关注按钮内容
	BargainBtnText           string `gorm:"column:bargain_btn_text;type:varchar(50)" json:"bargain_btn_text"`                        // 砍价按钮内容
	TransLinkText            string `gorm:"column:trans_link_text;type:varchar(100)" json:"trans_link_text"`                         // 转发链接语
	BargainSuccessToast      string `gorm:"column:bargain_success_toast;type:varchar(100)" json:"bargain_success_toast"`             // 砍价成功提示
	ShareSuccessText         string `gorm:"column:share_success_text;type:varchar(200)" json:"share_success_text"`                   // 分享完成提示
	BargainErrorToast        string `gorm:"column:bargain_error_toast;type:varchar(100)" json:"bargain_error_toast"`                 // 砍价出错提示
	ActivityEndText          string `gorm:"column:activity_end_text;type:varchar(200)" json:"activity_end_text"`                     // 活动结束通知
	ActivityTemporaryEndText string `gorm:"column:activity_temporary_end_text;type:varchar(200)" json:"activity_temporary_end_text"` // 临时下架通知
	ActivityDesc             string `gorm:"column:activity_desc;type:text" json:"activity_desc"`                                     // 活动说明
	Status                   int8   `gorm:"column:status;type:tinyint(1)" json:"status"`                                             // 启用状态
	IsDel                    int8   `gorm:"column:is_del;type:tinyint(1)" json:"is_del"`
	Ctime                    int    `gorm:"column:ctime;type:int(11)" json:"ctime"`
	SellerID                 int    `gorm:"column:seller_id;type:int(11)" json:"seller_id"`
	UserStatus               int8   `gorm:"column:user_status;type:tinyint(1)" json:"user_status"` // 使用状态 0 未使用，1 正在使用
	IsDefault                int8   `gorm:"column:is_default;type:tinyint(1)" json:"is_default"`   // 是否为默认模板 1是 0 否
}

// GbProBrandTop  品牌推广置顶
type GbProBrandTop struct {
	ID         int     `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"` // 品牌表
	StoreID    int     `gorm:"column:store_id;type:int(11);not null" json:"store_id"`         // 店铺ID
	BrandID    int     `gorm:"column:brand_id;type:int(11);not null" json:"brand_id"`         // 品牌ID
	Status     int8    `gorm:"column:status;type:tinyint(4);not null" json:"status"`          // 是否首页显示（1显示0不显示）
	RealPraise float32 `gorm:"column:real_praise;type:float" json:"real_praise"`              // 真实好评率
	Praise     float64 `gorm:"column:praise;type:float(11,2)" json:"praise"`                  // 虚拟好评率
	Sort       int     `gorm:"column:sort;type:int(11)" json:"sort"`                          // 店铺品牌排序
	IsDel      int8    `gorm:"column:is_del;type:tinyint(1);not null" json:"is_del"`          // 状态 1：正常  0：删除
}

type GbWhOrderBatchTrade struct {
	ID          int       `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	OrderNo     string    `gorm:"index;column:order_no;type:varchar(50);not null" json:"order_no"` // 订单流水号
	TradeNo     string    `gorm:"index;column:trade_no;type:varchar(50);not null" json:"trade_no"` // 批量付款流水号
	CreateAt    time.Time `gorm:"column:create_at;type:timestamp;not null" json:"create_at"`
	Status      int8      `gorm:"column:status;type:tinyint(1);not null" json:"status"`                // 状态 -1 付款失败 1 正在付款处理 2 付款完成
	TotalAmount float64   `gorm:"column:total_amount;type:decimal(10,2) unsigned" json:"total_amount"` // 总金额
}

type GbCrmWorkOrderDeal struct {
	ID          int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	WorkOrderID int    `gorm:"column:work_order_id;type:int(11);not null" json:"work_order_id"` // 工单ID
	UserID      int    `gorm:"column:user_id;type:int(11);not null" json:"user_id"`             // 处理人ID
	Content     string `gorm:"column:content;type:text;not null" json:"content"`                // 跟进结果
	CreateAt    int    `gorm:"column:create_at;type:int(11);not null" json:"create_at"`         // 处理时间
	Status      int8   `gorm:"column:status;type:tinyint(1);not null" json:"status"`            // 处理后的状态 0创建工单  1已处理  2已完结  3重新开始  4已转接
	Enclosure   string `gorm:"column:enclosure;type:text" json:"enclosure"`
}

type GbLogMerchantPurchaserChange struct {
	ID                  int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	UId                 int    `gorm:"column:uid;type:int(11) unsigned;not null" json:"uid"`                             // 操作者ID
	PurchaseIncrementID string `gorm:"column:purchase_increment_id;type:char(20);not null" json:"purchase_increment_id"` // 商家采购单流水号
	Remarks             string `gorm:"column:remarks;type:varchar(255)" json:"remarks"`                                  // 备注
	AfterStatus         int    `gorm:"column:after_status;type:int(11);not null" json:"after_status"`                    // 处理后状态
	Creattime           int    `gorm:"column:creattime;type:int(10)" json:"creattime"`                                   // 创建时间
}

type GbAdminModule struct {
	ID        int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`      // ID
	Name      string `gorm:"column:name;type:varchar(50);not null" json:"name"`                  // 名称
	CreatedAt int    `gorm:"column:created_at;type:int(10) unsigned;not null" json:"created_at"` // 创建时间
	UpdatedAt int    `gorm:"column:updated_at;type:int(10) unsigned;not null" json:"updated_at"` // 更新时间
}

type GbAppProfit struct {
	ID         int       `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	AppID      int       `gorm:"index:index_app_id_user_tag_spu_id;column:app_id;type:int(11);not null" json:"app_id"`                // 终端类型 0 h5 5 wxapp
	UserTag    int8      `gorm:"index:index_app_id_user_tag_spu_id;column:user_tag;type:tinyint(2);not null" json:"user_tag"`         // 用户类型 0 消费者 1BD 2代理商
	Commision  float64   `gorm:"column:commision;type:decimal(26,6) unsigned;not null" json:"commision"`                              // 佣金
	SkuID      int       `gorm:"index;column:sku_id;type:int(11) unsigned;not null" json:"sku_id"`                                    // 商家产品规格id
	SpuID      int       `gorm:"index;index:index_app_id_user_tag_spu_id;column:spu_id;type:int(11) unsigned;not null" json:"spu_id"` // 商家产品id spu
	SellerID   int       `gorm:"column:seller_id;type:int(11) unsigned;not null" json:"seller_id"`                                    // 商家ID
	CreatedAt  time.Time `gorm:"column:created_at;type:datetime;not null" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;type:datetime;not null" json:"updated_at"`
	SecretKey  string    `gorm:"unique;column:secret_key;type:char(16);not null" json:"secret_key"`    // 密钥
	IsCheapest int8      `gorm:"index;column:is_cheapest;type:tinyint(1);not null" json:"is_cheapest"` // 是否为商品到手最低价
}

type GbCrmWorkOrderDealCopy1 struct {
	ID          int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	WorkOrderID int    `gorm:"column:work_order_id;type:int(11);not null" json:"work_order_id"` // 工单ID
	UserID      int    `gorm:"column:user_id;type:int(11);not null" json:"user_id"`             // 处理人ID
	Content     string `gorm:"column:content;type:text;not null" json:"content"`                // 跟进结果
	CreateAt    int    `gorm:"column:create_at;type:int(11);not null" json:"create_at"`         // 处理时间
	Status      int8   `gorm:"column:status;type:tinyint(1);not null" json:"status"`            // 处理后的状态 0创建工单  1已处理  2已完结  3重新开始  4已转接
	Enclosure   string `gorm:"column:enclosure;type:text" json:"enclosure"`
}

// GbMerchantBargainTextMsg 砍价段子信息表
type GbMerchantBargainTextMsg struct {
	ID            int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	BargainTextID int    `gorm:"column:bargain_text_id;type:int(11)" json:"bargain_text_id"` // 所属文案ID
	Text          string `gorm:"column:text;type:varchar(255)" json:"text"`
}

type GbUserJifenSpuCont struct {
	ID         int  `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`      // 用户积分商品购买数量表
	StoreID    int  `gorm:"column:store_id;type:int(11) unsigned" json:"store_id"`              // 店铺ID
	UId        int  `gorm:"column:uid;type:int(11) unsigned" json:"uid"`                        // 用户ID
	SpuID      int  `gorm:"column:spu_id;type:int(11) unsigned" json:"spu_id"`                  // 商家商品ID
	Count      int  `gorm:"column:count;type:int(10) unsigned" json:"count"`                    // 已购总和
	IsDel      int8 `gorm:"column:is_del;type:tinyint(2)" json:"is_del"`                        // 是否删除0为未删除，1为已删除
	Creattime  int  `gorm:"column:creattime;type:int(10) unsigned;not null" json:"creattime"`   // 创建时间
	Updatatime int  `gorm:"column:updatatime;type:int(10) unsigned;not null" json:"updatatime"` // 更新时间
}

// GbSupProStandard 供应商产品规格
type GbSupProStandard struct {
	ID            int     `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	ProID         int     `gorm:"index;column:pro_id;type:int(11) unsigned;not null" json:"pro_id"`                 // 产品ID
	APIProCode    string  `gorm:"index;column:api_pro_code;type:varchar(100)" json:"api_pro_code"`                  // 接口商品编码
	Code          string  `gorm:"index;column:code;type:varchar(60);not null" json:"code"`                          // 编号
	ProCode       string  `gorm:"index;column:pro_code;type:varchar(40);not null" json:"pro_code"`                  // 产品编码 ENA13
	SupCode       string  `gorm:"index;column:sup_code;type:varchar(40);not null" json:"sup_code"`                  // 供应商编码
	SupID         int     `gorm:"column:sup_id;type:int(11)" json:"sup_id"`                                         // 供应商ID
	Property1     string  `gorm:"column:property_1;type:varchar(255)" json:"property_1"`                            // 属性1
	Property1Sort int8    `gorm:"column:property_1_sort;type:smallint(2) unsigned;not null" json:"property_1_sort"` // 规格固定排序
	Property2     string  `gorm:"column:property_2;type:varchar(255)" json:"property_2"`                            // 属性2
	Property2Sort int8    `gorm:"column:property_2_sort;type:smallint(2) unsigned;not null" json:"property_2_sort"` // 颜色固定排序
	MarketPrice   float64 `gorm:"column:market_price;type:decimal(10,2)" json:"market_price"`                       // 市场价
	MinPrice      float64 `gorm:"column:min_price;type:decimal(10,2)" json:"min_price"`                             // 最低价
	MinStatus     int8    `gorm:"column:min_status;type:tinyint(2) unsigned" json:"min_status"`                     // 是否可低限价出售 1代表是
	BatchType     int8    `gorm:"column:batch_type;type:smallint(2) unsigned" json:"batch_type"`                    // 含税类型 1未税 2含增值税 3含普票
	BatchPrice    float64 `gorm:"column:batch_price;type:decimal(10,2);not null" json:"batch_price"`                // 批发价(京东供货价)
	SellPrice     float64 `gorm:"column:sell_price;type:decimal(10,2)" json:"sell_price"`                           // 建议售价
	Qty           int     `gorm:"column:qty;type:int(11)" json:"qty"`                                               // 库存
	MinQty        int     `gorm:"column:min_qty;type:int(11)" json:"min_qty"`                                       // 预警库存
	Sku           string  `gorm:"column:sku;type:varchar(50)" json:"sku"`                                           // 供应商sku
	Pic           string  `gorm:"column:pic;type:varchar(255)" json:"pic"`                                          // 图片
	Status        int8    `gorm:"column:status;type:smallint(2) unsigned;not null" json:"status"`                   // 定价变更状态 1为定价已变更
	SetPrice      int8    `gorm:"column:set_price;type:smallint(2) unsigned;not null" json:"set_price"`             // 定价状态   默认为0未定价状态   1为已定价状态
	IsDel         int8    `gorm:"column:is_del;type:tinyint(1);not null" json:"is_del"`                             // 是否删除  1：是   0：否
	IsLeadingIn   int8    `gorm:"column:is_leading_in;type:tinyint(1)" json:"is_leading_in"`                        // 是否已经导入到平台， 0未导入 1已导入
	Ctime         int     `gorm:"column:ctime;type:int(10) unsigned;not null" json:"ctime"`                         // 插入时间
	JdPrice       float64 `gorm:"column:jd_price;type:decimal(10,2)" json:"jd_price"`                               // 京东挂价
}

// GbTaobaoCatesOriginal 淘宝类目表
type GbTaobaoCatesOriginal struct {
	ID              int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Level           int8   `gorm:"column:level;type:tinyint(1);not null" json:"level"`                           // 1一级分类  2二级分类 3三级分类 4三级分类
	Cid             int    `gorm:"index;column:cid;type:int(11);not null" json:"cid"`                            // 类目ID
	Name            string `gorm:"column:name;type:varchar(20);not null" json:"name"`                            // 类目名称
	ParentCid       int    `gorm:"column:parent_cid;type:int(11);not null" json:"parent_cid"`                    // 父类目ID
	IsParent        int8   `gorm:"column:is_parent;type:tinyint(1);not null" json:"is_parent"`                   // 是否为父级分类
	LocalCid        int    `gorm:"index;column:local_cid;type:int(11);not null" json:"local_cid"`                // 匹配的本地分类ID
	IsFixedLocalCid int8   `gorm:"column:is_fixed_local_cid;type:tinyint(1);not null" json:"is_fixed_local_cid"` // 是否固定本地分类
	IsFinished      int8   `gorm:"column:is_finished;type:tinyint(1);not null" json:"is_finished"`               // 是否完成属性匹配
}

type GbPaymentLog struct {
	ID            int     `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	PaymentID     int     `gorm:"column:payment_id;type:int(11);not null" json:"payment_id"`                  // 支付方式
	Type          int8    `gorm:"column:type;type:tinyint(1);not null" json:"type"`                           // 支付类型  1微信  2支付宝  3银行卡
	OrderNo       string  `gorm:"index;column:order_no;type:varchar(50);not null" json:"order_no"`            // 订单号
	PaymentNo     string  `gorm:"column:payment_no;type:varchar(50)" json:"payment_no"`                       // 支付订单号
	TransactionID string  `gorm:"column:transaction_id;type:varchar(50);not null" json:"transaction_id"`      // 第三方订单号，用支付系统后为支付系统订单号
	ThirdTradeNo  string  `gorm:"column:third_trade_no;type:varchar(50)" json:"third_trade_no"`               // 第三方订单号，用支付系统后
	TotalPrice    float64 `gorm:"column:total_price;type:decimal(10,2) unsigned;not null" json:"total_price"` // 支付金额
	Account       string  `gorm:"column:account;type:varchar(50)" json:"account"`                             // 支付账号
	Status        int8    `gorm:"column:status;type:tinyint(1) unsigned;not null" json:"status"`              // 订单状态 1已支付
	CreatedAt     int     `gorm:"column:created_at;type:int(10) unsigned;not null" json:"created_at"`         // 创建时间
	UpdatedAt     int     `gorm:"column:updated_at;type:int(10) unsigned;not null" json:"updated_at"`         // 更新时间
}

// GbPostage 邮费模板表
type GbPostage struct {
	ID             int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Name           string    `gorm:"column:name;type:varchar(20);not null" json:"name"`                   // 模板名称
	Area           string    `gorm:"column:area;type:varchar(100);not null" json:"area"`                  // 邮寄目的地
	PostTimeRange  int       `gorm:"column:post_time_range;type:int(11);not null" json:"post_time_range"` // 发货时间（1小时内或1天内）单位为小时
	IsFreeShipping int       `gorm:"column:is_free_shipping;type:int(1)" json:"is_free_shipping"`         // 是否包邮(0不包邮1包邮)
	BillingMethod  int       `gorm:"column:billing_method;type:int(11);not null" json:"billing_method"`   // 计费单位（1件数2体积3重量）
	TransportMode  int       `gorm:"column:transport_mode;type:int(11)" json:"transport_mode"`            // 运送方式（1快递、2物流暂时只有快递）
	IsCondition    int8      `gorm:"column:isCondition;type:tinyint(4)" json:"isCondition"`
	UserID         int       `gorm:"column:user_id;type:int(11);not null" json:"user_id"`  // 模板所属供应商
	Module         string    `gorm:"column:module;type:varchar(50)" json:"module"`         // 所属模块
	AddCount       float64   `gorm:"column:add_count;type:decimal(10,2)" json:"add_count"` // 递增的数量
	IsDel          int       `gorm:"column:is_del;type:int(1)" json:"is_del"`
	AddPostage     float64   `gorm:"column:add_postage;type:decimal(10,2)" json:"add_postage"`         // 所需增加的邮费
	DefaultCount   float64   `gorm:"column:default_count;type:decimal(10,2)" json:"default_count"`     // 首重单位数量
	DefaultPostage float64   `gorm:"column:default_postage;type:decimal(10,2)" json:"default_postage"` // 首重邮费
	Createtime     time.Time `gorm:"column:createtime;type:timestamp" json:"createtime"`               // 邮费创建时间
	ForbidArea     string    `gorm:"column:forbid_area;type:text" json:"forbid_area"`                  // 禁止配送地区
}

type GbProfitRuleRatio struct {
	ID     int     `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	RuleID int     `gorm:"column:rule_id;type:int(11) unsigned;not null" json:"rule_id"`  // 外部分配规则ID 对应 gb_profit_rule表
	Key    string  `gorm:"column:key;type:char(60);not null" json:"key"`                  // 参数名称
	Val    float64 `gorm:"column:val;type:float(6,2) unsigned;not null" json:"val"`       // 分配比例
	Status int8    `gorm:"column:status;type:tinyint(2) unsigned;not null" json:"status"` // 状态1打开 0关
	IsDel  int8    `gorm:"column:is_del;type:tinyint(2) unsigned;not null" json:"is_del"` // 删除状态 1已删 0未删
}

type GbStoreSpread struct {
	ID          int  `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	StoreID     int  `gorm:"column:store_id;type:int(11) unsigned;not null" json:"store_id"`         // 店铺ID
	Type        int8 `gorm:"column:type;type:tinyint(2) unsigned;not null" json:"type"`              // 营销类型  1网红（礼品精选）、3超值、4好物、5爆品专区、6京东直送
	CateID      int  `gorm:"column:cate_id;type:int(11)" json:"cate_id"`                             // 热卖榜单分类【其它类型没有】
	StoreProID  int  `gorm:"column:store_pro_id;type:int(11) unsigned;not null" json:"store_pro_id"` // 产品ID
	IsTop       int8 `gorm:"column:is_top;type:tinyint(2) unsigned;not null" json:"is_top"`          // 置顶在首页的状态  1置顶 0未置顶
	CTime       int  `gorm:"column:c_time;type:int(10) unsigned;not null" json:"c_time"`
	Display     int8 `gorm:"column:display;type:tinyint(2) unsigned;not null" json:"display"` // 展示状态  0不展示  1展示
	Sort        int  `gorm:"column:sort;type:int(11);not null" json:"sort"`                   // 排序
	YStoreProID int  `gorm:"column:y_store_pro_id;type:int(11)" json:"y_store_pro_id"`        // 导入之前此记录对应的(源）店铺商品ID
}

// GbStoreUserStoreProduct 用户店铺货品表
type GbStoreUserStoreProduct struct {
	ID          int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	StoreID     int    `gorm:"column:store_id;type:int(11) unsigned;not null" json:"store_id"`         // 店铺ID
	UserStoreID int    `gorm:"column:user_store_id;type:int(11);not null" json:"user_store_id"`        // 用户店铺ID
	UserID      int    `gorm:"column:user_id;type:int(11) unsigned;not null" json:"user_id"`           // 用户ID
	StoreProID  int    `gorm:"column:store_pro_id;type:int(11) unsigned;not null" json:"store_pro_id"` // 商铺产品ID
	PlatProCode string `gorm:"column:plat_pro_code;type:varchar(60);not null" json:"plat_pro_code"`    // 平台货品编码
	Ctime       int    `gorm:"column:ctime;type:int(11);not null" json:"ctime"`                        // 创建时间
	IsDel       int8   `gorm:"column:is_del;type:tinyint(2);not null" json:"is_del"`
}

type GbSupLogisticsFacilitator struct {
	ID                uint   `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	FacilitatorID     int    `gorm:"column:facilitator_id;type:int(10);not null" json:"facilitator_id"` // 服务商id
	BranchesCode      string `gorm:"column:branches_code;type:varchar(20)" json:"branches_code"`        // 网店编码
	BranchesName      string `gorm:"column:branches_name;type:varchar(100)" json:"branches_name"`       // 网店名称
	Accout            string `gorm:"column:accout;type:varchar(255)" json:"accout"`                     // 月结账户
	Time              string `gorm:"column:time;type:varchar(20);not null" json:"time"`                 // 创建时间
	CpType            int8   `gorm:"column:cp_type;type:tinyint(1);not null" json:"cp_type"`            // 1-直营 2-加盟
	Quantity          int    `gorm:"column:quantity;type:int(32)" json:"quantity"`                      // 面单余额
	AllocatedQuantity int    `gorm:"column:allocated_quantity;type:int(32)" json:"allocated_quantity"`  // 面单总数
	ServiceCode       string `gorm:"column:service_code;type:varchar(20)" json:"service_code"`          // 可用服务代号,如需要快递公司代收货款等
	ServiceName       string `gorm:"column:service_name;type:varchar(20)" json:"service_name"`          // 可用服务名称
	CancelQuantity    int    `gorm:"column:cancel_quantity;type:int(11)" json:"cancel_quantity"`        // 取消的面单数
	UpdateTime        string `gorm:"column:update_time;type:varchar(30)" json:"update_time"`            // 更新的时间
	SupID             int    `gorm:"column:sup_id;type:int(11);not null" json:"sup_id"`                 // 供应商id
	PrintQuantity     int    `gorm:"column:print_quantity;type:int(11)" json:"print_quantity"`          // 取号数量
}

// GbAPIUser 访问restful接口权限表
type GbAPIUser struct {
	ID                 int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	Username           string `gorm:"unique;column:username;type:varchar(20);not null" json:"username"`                       // 用户名
	Password           string `gorm:"column:password;type:varchar(100);not null" json:"password"`                             // 密码
	Status             int8   `gorm:"column:status;type:tinyint(3) unsigned;not null" json:"status"`                          // 状态 0关闭帐户权限  1为打开
	AccessToken        string `gorm:"unique;column:access_token;type:varchar(255);not null" json:"access_token"`              // restful请求token
	TokenTime          int    `gorm:"column:token_time;type:int(10) unsigned;not null" json:"token_time"`                     // token生成时间
	Allowance          int    `gorm:"column:allowance;type:int(10) unsigned;not null" json:"allowance"`                       // restful剩余的允许的请求数
	AllowanceUpdatedAt int    `gorm:"column:allowance_updated_at;type:int(10) unsigned;not null" json:"allowance_updated_at"` // restful请求的UNIX时间戳数
	RateLimit          int    `gorm:"column:rate_limit;type:int(11) unsigned;not null" json:"rate_limit"`                     // 一分钟限制请求的次数
	CreatedAt          int    `gorm:"column:created_at;type:int(10)" json:"created_at"`                                       // 创建时间
	UpdatedAt          int    `gorm:"column:updated_at;type:int(10)" json:"updated_at"`                                       // 更改时间
	LastAccessToken    string `gorm:"column:last_access_token;type:varchar(255)" json:"last_access_token"`                    // 上一个token
	LastTokenTime      int64  `gorm:"column:last_token_time;type:bigint(20)" json:"last_token_time"`                          // 上一个token的有效时间
}

type GbApplyBuy struct {
	ID           int     `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	PlatformType string  `gorm:"index;column:platform_type;type:varchar(12);not null" json:"platform_type"` // 所属平台编号：taobao淘宝 jingdong京东
	Tid          string  `gorm:"column:tid;type:varchar(50)" json:"tid"`                                    // 平台订购单号
	UserID       int     `gorm:"column:user_id;type:int(10) unsigned" json:"user_id"`                       // 订购用户ID
	UserName     string  `gorm:"index;column:user_name;type:varchar(255)" json:"user_name"`                 // 订购用户名
	UId          int     `gorm:"column:uid;type:int(10) unsigned" json:"uid"`                               // 本地用户ID
	Type         int8    `gorm:"column:type;type:tinyint(1) unsigned;not null" json:"type"`                 // 订购类型：1=新订 2=续订 3=升级 4=后台赠送 5=后台自动续订 6=订单审核后生成订购关系
	ArticleName  string  `gorm:"column:article_name;type:varchar(30)" json:"article_name"`                  // 订购应用名称
	Fee          float64 `gorm:"column:fee;type:decimal(10,2);not null" json:"fee"`                         // 订购应用原价
	PromFee      float64 `gorm:"column:prom_fee;type:decimal(10,2);not null" json:"prom_fee"`               // 订购应用优惠金额
	TotalPayFee  float64 `gorm:"column:total_pay_fee;type:decimal(10,2);not null" json:"total_pay_fee"`     // 订购应用实付金额
	ReturnFee    float64 `gorm:"column:return_fee;type:decimal(10,2);not null" json:"return_fee"`           // 应返还金额
	StartTime    int     `gorm:"column:start_time;type:int(10) unsigned" json:"start_time"`                 // 订购周期开始时间
	EndTime      int     `gorm:"column:end_time;type:int(10) unsigned" json:"end_time"`                     // 订购周期结束时间
}

// GbMerchantProStandard 商户产品规格
type GbMerchantProStandard struct {
	ID                int               `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	SellerID          int               `gorm:"index;column:seller_id;type:int(11);not null" json:"seller_id"`              // 商户ID
	ProID             int               `gorm:"index;column:pro_id;type:int(11) unsigned;not null" json:"pro_id"`           // 商户产品ID
	Sku               string            `gorm:"index;column:sku;type:varchar(60);not null" json:"sku"`                      // 产品SKU
	PlatProCode       string            `gorm:"index;column:plat_pro_code;type:varchar(100);not null" json:"plat_pro_code"` // 编号(产品平台编号)
	BarCode           string            `gorm:"index;column:bar_code;type:varchar(40);not null" json:"bar_code"`            // 产品编码 EAN13编码
	VirSales          int               `gorm:"column:vir_sales;type:int(11);not null" json:"vir_sales"`                    // 虚拟销售数
	MarketPrice       float64           `gorm:"column:market_price;type:decimal(20,6);not null" json:"market_price"`        // 市场价
	Price             float64           `gorm:"column:price;type:decimal(20,6);not null" json:"price"`                      // 售价
	IsPrice           int8              `gorm:"column:is_price;type:tinyint(2) unsigned;not null" json:"is_price"`          // 定价状态  当SKU填完售价时该字段要为1，默认为0
	IsDel             int8              `gorm:"column:is_del;type:tinyint(1);not null" json:"is_del"`                       // 是否删除  1：是   0：否
	SourceStandardID  int               `gorm:"column:source_standard_id;type:int(11)" json:"source_standard_id"`
	GbMerchantProduct GbMerchantProduct `gorm:"ForeignKey:ProID;"`
}

// GbCommentImage 商品评论图片
type GbCommentImage struct {
	ID        int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	CommentID int    `gorm:"column:comment_id;type:int(11)" json:"comment_id"`    // 评论ID
	ImageURL  string `gorm:"column:image_url;type:varchar(150)" json:"image_url"` // 评论图片URL
}

type GbMigration struct {
	Version   string `gorm:"primary_key;column:version;type:varchar(180);not null" json:"version"`
	ApplyTime int    `gorm:"column:apply_time;type:int(11)" json:"apply_time"`
}

type GbOrderStatusLog struct {
	ID          int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	OrderID     int    `gorm:"column:order_id;type:int(11) unsigned;not null" json:"order_id"`   // 订单ID
	OrderCode   string `gorm:"column:order_code;type:char(100);not null" json:"order_code"`      // 订单号
	IncrementID string `gorm:"column:increment_id;type:char(15)" json:"increment_id"`            // 订单流水号（YYMMDD+）
	UId         int    `gorm:"column:uid;type:int(11) unsigned;not null" json:"uid"`             // 用户ID
	Status      int8   `gorm:"column:status;type:tinyint(2);not null" json:"status"`             // 订单状态： -1：已取消；1:待付款；2:待商家采购； 3:待平台采购；4:商品发货中；5:已完成；6：待供应商退款；7：待平台退款；8：待商家退款；9：已关闭；10：待成,11:等待银行退款
	Creattime   int    `gorm:"column:creattime;type:int(10) unsigned;not null" json:"creattime"` // 创建时间
}

type GbWebsiteContentType struct {
	ID          int       `gorm:"primary_key;column:id;type:int(10);not null" json:"-"`
	WebsiteType int8      `gorm:"column:website_type;type:tinyint(1)" json:"website_type"` // 官网类型 1神雀 2鼎翰
	Type        int8      `gorm:"column:type;type:tinyint(1)" json:"type"`                 // 类型 1发布渠道 2新闻分类 3活动分类 4合作案例
	Name        string    `gorm:"column:name;type:varchar(255)" json:"name"`               // 名称
	Status      int8      `gorm:"column:status;type:tinyint(1)" json:"status"`             // 状态 0删除 1正常
	CreateTime  time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`     // 创建时间
	UpdateTime  time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`     // 更新时间
}

type GbWhSupplierWhPartnerRel struct {
	ID         int  `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	SupplierID int  `gorm:"column:supplier_id;type:int(11) unsigned;not null" json:"supplier_id"` // 供应商ID
	WhUserID   int  `gorm:"column:wh_user_id;type:int(11) unsigned;not null" json:"wh_user_id"`   // 网红用户ID
	CreatedAt  int  `gorm:"column:created_at;type:int(11)" json:"created_at"`                     // 创建时间
	IsInvite   int8 `gorm:"column:is_invite;type:tinyint(1) unsigned" json:"is_invite"`           // 是否为邀请人
	IsDel      int8 `gorm:"column:is_del;type:tinyint(2) unsigned" json:"is_del"`                 // 删除状态 0 1
}

type GbAPIUserExtend struct {
	ID       int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	UserID   int    `gorm:"unique_index:idx_0;column:user_id;type:int(11);not null" json:"user_id"`
	Name     string `gorm:"unique_index:idx_0;column:name;type:varchar(64);not null" json:"name"` // 名称
	Value    string `gorm:"column:value;type:text;not null" json:"value"`                         // 值
	Describe string `gorm:"column:describe;type:varchar(64)" json:"describe"`
}

// GbMerchantShareDimension 顶级分成维度信息
type GbMerchantShareDimension struct {
	ID    int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Name  string `gorm:"column:name;type:varchar(50)" json:"name"` // 顶级维度名称
	IsDel int8   `gorm:"column:is_del;type:tinyint(1)" json:"is_del"`
}

// GbStoreUserClosure 用户上下级从属关系表
type GbStoreUserClosure struct {
	ID         int     `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	UIdParent  int     `gorm:"index:union_index;column:uid_parent;type:int(11) unsigned;not null" json:"uid_parent"` // 父UID
	UIdChild   int     `gorm:"index:union_index;column:uid_child;type:int(11) unsigned;not null" json:"uid_child"`   // 子UID
	Level      int     `gorm:"index:union_index;column:level;type:int(11);not null" json:"level"`                    //  层级
	Commission float64 `gorm:"column:commission;type:decimal(20,4) unsigned;not null" json:"commission"`             // 佣金总额
}

type GbActivityStoreJoin struct {
	ID      int  `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	AlqID   int  `gorm:"column:alq_id;type:int(11) unsigned;not null" json:"alq_id"`     // 活动总表ID  gb_activity_lock_qty.id
	StoreID int  `gorm:"column:store_id;type:int(11) unsigned;not null" json:"store_id"` // 参与活动的店铺ID
	Status  int  `gorm:"column:status;type:int(255) unsigned;not null" json:"status"`    // 加入状态 1已加入  0已退出
	Ctime   int  `gorm:"column:ctime;type:int(10);not null" json:"ctime"`
	IsDel   int8 `gorm:"column:is_del;type:tinyint(1)" json:"is_del"`
}

type GbPaySuccessOrder struct {
	ID          int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	OrderID     int    `gorm:"column:order_id;type:int(10)" json:"order_id"`          // 订单流水号（YYMMDD+）
	IncrementID string `gorm:"column:increment_id;type:char(15)" json:"increment_id"` // 订单流水号（YYMMDD+）
	Status      int8   `gorm:"column:status;type:tinyint(2);not null" json:"status"`  // 1，接收到了数据；2处理数据完成。
	UpdateAt    int    `gorm:"column:update_at;type:int(10)" json:"update_at"`        // 修复时间
	CreateAt    int    `gorm:"column:create_at;type:int(10)" json:"create_at"`        // 创建时间
}

// GbStoreProductCateRela 店铺商品分类关联表
type GbStoreProductCateRela struct {
	ID      int `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Cid     int `gorm:"index:index_cid_sp_id;column:cid;type:int(11);not null" json:"cid"`     // 运营分类id
	SpID    int `gorm:"index:index_cid_sp_id;column:sp_id;type:int(11);not null" json:"sp_id"` // 店铺产品id
	YCid    int `gorm:"column:y_cid;type:int(11)" json:"y_cid"`                                // 导入之前的商品分类ID
	YSpID   int `gorm:"column:y_sp_id;type:int(11)" json:"y_sp_id"`                            // 导入之前的商品ID
	StoreID int `gorm:"column:store_id;type:int(11)" json:"store_id"`                          // 店铺ID
}

type GbWebsiteConfig struct {
	ID          int       `gorm:"primary_key;column:id;type:int(10);not null" json:"-"`    // 基本设置自增ID
	WebsiteType int8      `gorm:"column:website_type;type:tinyint(1)" json:"website_type"` // 官网类型 1神雀 2鼎翰
	HeadURL     string    `gorm:"column:head_url;type:varchar(255)" json:"head_url"`       // 页首图片
	BannerNum   int       `gorm:"column:banner_num;type:int(2)" json:"banner_num"`         // banner数量
	BannerURL   string    `gorm:"column:banner_url;type:text" json:"banner_url"`           // banner图片集（json格式）
	FootURL     string    `gorm:"column:foot_url;type:varchar(255)" json:"foot_url"`       // 页尾图片
	CreateTime  time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`     // 创建时间
	UpdateTime  time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`     // 修改时间
}

// GbShopingCart 购物车信息表
type GbShopingCart struct {
	ID           int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	Bcat         int    `gorm:"column:bcat;type:int(11);not null" json:"bcat"`                    // 一级分类
	UId          int    `gorm:"column:uid;type:int(11) unsigned;not null" json:"uid"`             // 用户ID
	MerchantCode string `gorm:"column:merchant_code;type:char(20);not null" json:"merchant_code"` // 商家Code
	StoreID      int    `gorm:"column:store_id;type:int(11);not null" json:"store_id"`            // 店铺ID
	ProCode      string `gorm:"column:pro_code;type:char(20);not null" json:"pro_code"`           // 商户编码
	Sku          string `gorm:"column:sku;type:char(20);not null" json:"sku"`                     // 产品规格编码
	Count        int    `gorm:"column:count;type:int(11) unsigned;not null" json:"count"`         // 商品数量，数量为0测视为从购物车中删除了该产品
	Status       int8   `gorm:"column:status;type:tinyint(2);not null" json:"status"`             // 购物车商品状态 -1，为删除商品（count 值为0）0 已经提交订单的商品 ，1有效，2无效
	CreateTime   int    `gorm:"column:create_time;type:int(10)" json:"create_time"`               // 创建时间
	UpdateTime   int    `gorm:"column:update_time;type:int(10)" json:"update_time"`               // 修改时间
}

type GbStoreTrend struct {
	ID           int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	StoreID      int    `gorm:"column:store_id;type:int(11) unsigned;not null" json:"store_id"`       // 店铺ID
	Title        string `gorm:"column:title;type:varchar(250);not null" json:"title"`                 // 标题
	BannerImg    string `gorm:"column:banner_img;type:varchar(250);not null" json:"banner_img"`       // banner图片地址
	Image        string `gorm:"column:image;type:varchar(250);not null" json:"image"`                 // 草文图片地址
	ContentImage string `gorm:"column:content_image;type:varchar(250);not null" json:"content_image"` // 内容图片
	Num          int    `gorm:"column:num;type:int(11) unsigned;not null" json:"num"`                 // 商品数量
	Hits         int    `gorm:"column:hits;type:int(11) unsigned;not null" json:"hits"`               // 阅读量
	Transmit     int    `gorm:"column:transmit;type:int(11) unsigned;not null" json:"transmit"`       // 转发量
	IsTop        int8   `gorm:"column:is_top;type:tinyint(2) unsigned;not null" json:"is_top"`        // 置顶在首页的状态  1置顶 0未置顶
	Display      int8   `gorm:"column:display;type:tinyint(2) unsigned;not null" json:"display"`      // 展示状态  0不展示  1展示
	CTime        int    `gorm:"column:c_time;type:int(10) unsigned;not null" json:"c_time"`
	Sort         int    `gorm:"column:sort;type:int(11);not null" json:"sort"` // 排序
}

type GbStoreUserLoginBindLog struct {
	ID       int64     `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"`
	Level    int       `gorm:"index;column:level;type:int(11)" json:"level"`
	Category string    `gorm:"index;column:category;type:varchar(255)" json:"category"`
	LogTime  time.Time `gorm:"column:log_time;type:datetime" json:"log_time"`
	Prefix   string    `gorm:"column:prefix;type:text" json:"prefix"`
	Message  string    `gorm:"column:message;type:text" json:"message"`
}

type GbSupAddress struct {
	ID              int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Name            string    `gorm:"column:name;type:varchar(255);not null" json:"name"` // 收货点名称
	SupID           int       `gorm:"column:sup_id;type:int(11);not null" json:"sup_id"`
	SupName         string    `gorm:"column:sup_name;type:varchar(255);not null" json:"sup_name"` // 供应商收货人
	SupTel          string    `gorm:"column:sup_tel;type:varchar(255);not null" json:"sup_tel"`   // 供应商收货人号码
	AreaAddressCode string    `gorm:"column:area_address_code;type:varchar(255)" json:"area_address_code"`
	AreaAddress     string    `gorm:"column:area_address;type:varchar(255);not null" json:"area_address"`     // 收货区域地址
	DetailAddress   string    `gorm:"column:detail_address;type:varchar(255);not null" json:"detail_address"` // 收货详细街道地址
	Sort            int       `gorm:"column:sort;type:int(11);not null" json:"sort"`                          // 排序
	Status          int       `gorm:"column:status;type:int(11);not null" json:"status"`                      // 状态：0停用1启用
	Createtime      time.Time `gorm:"column:createtime;type:datetime;not null" json:"createtime"`             // 创建时间
}

type GbAuthPriceLog struct {
	ID       int64     `gorm:"primary_key;column:id;type:bigint(20);not null" json:"-"`
	Level    int       `gorm:"index;column:level;type:int(11)" json:"level"`
	Category string    `gorm:"index;column:category;type:varchar(255)" json:"category"`
	LogTime  time.Time `gorm:"column:log_time;type:datetime" json:"log_time"`
	Prefix   string    `gorm:"column:prefix;type:text" json:"prefix"`
	Message  string    `gorm:"column:message;type:text" json:"message"`
}

type GbCardRechargeLog struct {
	ID          int     `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Orderid     string  `gorm:"column:orderid;type:varchar(55);not null" json:"orderid"`        // 欧飞生成的订单号
	Cardid      int     `gorm:"column:cardid;type:int(11)" json:"cardid"`                       // 卡编码
	Cardnum     int     `gorm:"column:cardnum;type:int(11)" json:"cardnum"`                     // 购买卡数量（如果是直充默认为1）
	Ordercash   float64 `gorm:"column:ordercash;type:float(25,0)" json:"ordercash"`             // 订单金额
	Cardname    string  `gorm:"column:cardname;type:varchar(255)" json:"cardname"`              // 卡名称
	IncrementID string  `gorm:"column:increment_id;type:char(25);not null" json:"increment_id"` // 订单流水号
	PostData    string  `gorm:"column:post_data;type:text;not null" json:"post_data"`           // 接收的虚拟商品返回的JSON数据
	CreatedAt   int     `gorm:"column:created_at;type:int(13);not null" json:"created_at"`      // 创建时间
	UpdatedAt   int     `gorm:"column:updated_at;type:int(13);not null" json:"updated_at"`      // 更新时间
	CostType    int8    `gorm:"column:cost_type;type:tinyint(2);not null" json:"cost_type"`     // 1为直充，2为卡密
	RetCode     int8    `gorm:"column:ret_code;type:tinyint(2)" json:"ret_code"`                // 1为成功，9位失败，0为充值中
}

type GbCrmKnowledgebase struct {
	ID         int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Title      string    `gorm:"column:title;type:varchar(200)" json:"title"`                // 问题
	Answer     string    `gorm:"column:answer;type:varchar(500);not null" json:"answer"`     // 答案
	Companyid  int       `gorm:"column:companyid;type:int(11)" json:"companyid"`             // 商户id
	Bcate      int       `gorm:"column:bcate;type:int(11);not null" json:"bcate"`            // 一级分类
	Mcate      int       `gorm:"column:mcate;type:int(11);not null" json:"mcate"`            // 二级分类
	Status     int8      `gorm:"column:status;type:smallint(6);not null" json:"status"`      // 状态：0已禁用1已启用
	Createtime time.Time `gorm:"column:createtime;type:datetime;not null" json:"createtime"` // 创建时间
}

// GbCrmQuickreply 快捷回复
type GbCrmQuickreply struct {
	ID         int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Keyword    string    `gorm:"column:keyword;type:varchar(50);not null" json:"keyword"`    // 回复关键字
	Companyid  int       `gorm:"column:companyid;type:int(11)" json:"companyid"`             // 商户id
	Content    string    `gorm:"column:content;type:varchar(500);not null" json:"content"`   // 回复内容
	Cateid     int       `gorm:"column:cateid;type:int(11);not null" json:"cateid"`          // 回复所属分类
	IsDel      int8      `gorm:"column:is_del;type:smallint(6);not null" json:"is_del"`      // 逻辑删除0未删除1已删除
	Status     int8      `gorm:"column:status;type:smallint(6);not null" json:"status"`      // 回复内容状态：0已禁用1已启用
	Createtime time.Time `gorm:"column:createtime;type:datetime;not null" json:"createtime"` // 快捷回复创建时间
}

type GbPaymentRefund struct {
	ID          int     `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	RequestID   string  `gorm:"column:request_id;type:varchar(64)" json:"request_id"`                         // 惟一请求标志
	PaymentID   int     `gorm:"column:payment_id;type:int(11);not null" json:"payment_id"`                    // 支付方式
	OrderNo     string  `gorm:"index;column:order_no;type:varchar(50);not null" json:"order_no"`              // 订单号
	PaymentNo   string  `gorm:"column:payment_no;type:varchar(50)" json:"payment_no"`                         // 支付订单号即支付系统内的支付订单号
	RefundNo    string  `gorm:"column:refund_no;type:varchar(50);not null" json:"refund_no"`                  // 退款订单号
	TotalPrice  float64 `gorm:"column:total_price;type:decimal(10,2) unsigned;not null" json:"total_price"`   // 订单总金额
	RefundPrice float64 `gorm:"column:refund_price;type:decimal(10,2) unsigned;not null" json:"refund_price"` // 退款金额
	Status      int8    `gorm:"index;column:status;type:tinyint(1) unsigned;not null" json:"status"`          // 退款订单状态：0待退款 1退款成功 2退款异常 3退款关闭
	Message     string  `gorm:"column:message;type:varchar(100);not null" json:"message"`                     // 备注信息
	RefundAt    int     `gorm:"column:refund_at;type:int(10)" json:"refund_at"`                               // 退款成功时间
	CreatedAt   int     `gorm:"column:created_at;type:int(10) unsigned;not null" json:"created_at"`           // 创建时间
	UpdatedAt   int     `gorm:"column:updated_at;type:int(10) unsigned;not null" json:"updated_at"`           // 更新时间
}

// GbProParam 平台产品参数
type GbProParam struct {
	ID      int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	ProID   int    `gorm:"column:pro_id;type:int(11)" json:"pro_id"`                  // 商品ID
	ProCode string `gorm:"column:pro_code;type:varchar(40);not null" json:"pro_code"` // 产品编号
	Name    string `gorm:"column:name;type:varchar(20);not null" json:"name"`         // 参数名称
	Value   string `gorm:"column:value;type:varchar(50);not null" json:"value"`       // 参数值
	Sort    int8   `gorm:"column:sort;type:tinyint(4);not null" json:"sort"`          // 排序
	IsDel   int8   `gorm:"column:is_del;type:tinyint(1);not null" json:"is_del"`      // 是否删除  1：是   0：否
}

type GbMerchantBargainTemplateItem struct {
	ID         int  `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	TemplateID int  `gorm:"column:template_id;type:int(11)" json:"template_id"` // 砍价模板ID
	Position   int  `gorm:"column:position;type:int(11)" json:"position"`       // 第几号砍价人
	Percentage int  `gorm:"column:percentage;type:int(2)" json:"percentage"`    // 百分比，总数100
	IsDel      int8 `gorm:"column:is_del;type:tinyint(1)" json:"is_del"`
}

type GbMerchantYigroupAd struct {
	ID            int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Type          int8   `gorm:"column:type;type:tinyint(2) unsigned;not null" json:"type"`                    // 广告类型  1：头部 2：中部
	Image         string `gorm:"column:image;type:varchar(250);not null" json:"image"`                         // 图片
	Title         string `gorm:"column:title;type:varchar(250);not null" json:"title"`                         // 广告名称
	MerchantSpuID int    `gorm:"column:merchant_spu_id;type:int(11) unsigned;not null" json:"merchant_spu_id"` // 商家产品ID
	MerchantSkuID int    `gorm:"column:merchant_sku_id;type:int(11)" json:"merchant_sku_id"`                   // 商家产品规格ID
	Desc          string `gorm:"column:desc;type:text" json:"desc"`                                            // 描述
	IsDel         int8   `gorm:"column:is_del;type:tinyint(1);not null" json:"is_del"`                         // 1删除，0未删除
	Sort          int    `gorm:"column:sort;type:int(11);not null" json:"sort"`                                // 排序
	CreatedAt     int    `gorm:"column:created_at;type:int(10) unsigned;not null" json:"created_at"`
	UpdatedAt     int    `gorm:"column:updated_at;type:int(10) unsigned;not null" json:"updated_at"`
	IsOnline      int8   `gorm:"column:is_online;type:tinyint(1)" json:"is_online"` // 启用状态 1 已启用 0 未启用
}

// GbOperateLog 操作日志
type GbOperateLog struct {
	ID          int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	UserID      int    `gorm:"index;column:user_id;type:int(10);not null" json:"user_id"`           // 操作人编号
	TableName   string `gorm:"index;column:table_name;type:varchar(35);not null" json:"table_name"` // 操作表
	ProEditTime int    `gorm:"column:pro_edit_time;type:int(11)" json:"pro_edit_time"`              // 当前审核的产品修改时间
	Uptime      int    `gorm:"column:uptime;type:int(11)" json:"uptime"`                            // 操作时间
	BeforeData  string `gorm:"column:before_data;type:text" json:"before_data"`                     // 操作之前的操作
	AfterData   string `gorm:"column:after_data;type:text;not null" json:"after_data"`              // 操作之后的数据
	SQL         string `gorm:"column:sql;type:text;not null" json:"sql"`                            // 操作语句
	UpdateID    string `gorm:"column:update_id;type:text;not null" json:"update_id"`                // 修改id
}

// GbProduct 平台产品库
type GbProduct struct {
	ID             int     `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	SerialNumber   string  `gorm:"column:serial_number;type:char(12)" json:"serial_number"`             // 商品流水号
	Code           string  `gorm:"index;column:code;type:varchar(100);not null" json:"code"`            // 编码
	ProName        string  `gorm:"column:pro_name;type:varchar(150);not null" json:"pro_name"`          // 产品名称
	ShortName      string  `gorm:"column:short_name;type:varchar(100);not null" json:"short_name"`      // 产品简称
	SupID          string  `gorm:"column:sup_id;type:varchar(20)" json:"sup_id"`                        // 供应商ID，多个供应商用，分割,  从小到大排序
	SupCode        string  `gorm:"index;column:sup_code;type:varchar(200);not null" json:"sup_code"`    // 供应商编码多个以“，”隔开
	BarCode        string  `gorm:"column:bar_code;type:varchar(40);not null" json:"bar_code"`           // 商品条码
	BarCodeImg     string  `gorm:"column:bar_code_img;type:varchar(200)" json:"bar_code_img"`           // 条码图片
	Daifa          int8    `gorm:"column:daifa;type:tinyint(2) unsigned;not null" json:"daifa"`         // 一件代发，默认否为0    是为1
	SellPoint      string  `gorm:"column:sell_point;type:varchar(200)" json:"sell_point"`               // 卖点
	BrandID        int     `gorm:"column:brand_id;type:int(11);not null" json:"brand_id"`               // 品牌ID
	OriginPlace    int     `gorm:"column:origin_place;type:int(20);not null" json:"origin_place"`       // 产地地区ID
	ShipTemplate   int     `gorm:"column:ship_template;type:varchar(20);not null" json:"ship_template"` // 运费模板编号
	Volume         float64 `gorm:"column:volume;type:decimal(10,2)" json:"volume"`                      // 体积
	Weight         float64 `gorm:"column:weight;type:float(10,3);not null" json:"weight"`               // 重量
	Keyword        string  `gorm:"column:keyword;type:varchar(150);not null" json:"keyword"`            // 关键字
	Service        string  `gorm:"column:service;type:varchar(100);not null" json:"service"`            // 服务保证
	Type           int8    `gorm:"column:type;type:tinyint(1);not null" json:"type"`                    // 类型  1：实物产品   0：虚拟产品  2服务商品
	GlobalSendType int8    `gorm:"column:global_send_type;type:tinyint(1)" json:"global_send_type"`     // 跨境商品发货类型（1为保税，2直邮，3为国内） @author by 万朝
	Pic            string  `gorm:"column:pic;type:longtext" json:"pic"`                                 // 商品图片
	Standard1      string  `gorm:"column:standard_1;type:varchar(150)" json:"standard_1"`               // 规格1
	Standard2      string  `gorm:"column:standard_2;type:varchar(150)" json:"standard_2"`               // 规格2
	Content        string  `gorm:"column:content;type:text" json:"content"`                             // 描述
	StatusRemark   string  `gorm:"column:status_remark;type:varchar(200)" json:"status_remark"`         // 审核失败原因记录
	Status         int8    `gorm:"column:status;type:tinyint(1);not null" json:"status"`                // 状态,1待审核,2审核失败,3审核通过,4已上架,5下架审核,6已下架
	IsDel          int8    `gorm:"column:is_del;type:tinyint(1);not null" json:"is_del"`                // 是否删除  1：是  0：否 2:回收站
	CreateTime     int     `gorm:"column:create_time;type:int(10);not null" json:"create_time"`         // 创建时间
	IsPrice        int8    `gorm:"column:is_price;type:tinyint(1) unsigned" json:"is_price"`            // 是否定价 0 未定价 ，1 已定价 , 2 部分定价 3待变更
	ModifyTime     int     `gorm:"column:modify_time;type:int(11)" json:"modify_time"`                  // 修改或者审核等的变更时间
	ModifyOperator string  `gorm:"column:modify_operator;type:varchar(20)" json:"modify_operator"`      // 变更操作者用户名
	ModifyOpID     int     `gorm:"column:modify_op_id;type:int(11)" json:"modify_op_id"`                // 变更操作人 用户ID
	UnExamNum      int8    `gorm:"column:un_exam_num;type:tinyint(4)" json:"un_exam_num"`               // 未审核的修改次数
	MinPrice       float64 `gorm:"column:min_price;type:decimal(10,2)" json:"min_price"`                // 商品中所有SKU中最低的售价
	PackingList    string  `gorm:"column:packing_list;type:varchar(255)" json:"packing_list"`           // 包装清单
	CostType       int8    `gorm:"column:cost_type;type:tinyint(1)" json:"cost_type"`                   // 充值方式（1直充，2卡密,3兑换码）
	ProductType    int8    `gorm:"column:product_type;type:tinyint(1)" json:"product_type"`             // 商品类型（1为视频）可扩展
	DeliveryType   int8    `gorm:"column:delivery_type;type:tinyint(1)" json:"delivery_type"`           // 发货方式（1为虚拟仓，2为api发送）
	IsSetZyPrice   int8    `gorm:"column:is_set_zy_price;type:tinyint(1)" json:"is_set_zy_price"`       // 是否有SKU设置了自营价格（使用于代运营商品导入）
}

// GbSupProduct 供应商产品
type GbSupProduct struct {
	ID                   int     `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	SupplierID           int     `gorm:"index;column:supplier_id;type:int(11)" json:"supplier_id"`                   // 供应商账号ID
	Code                 string  `gorm:"index;column:code;type:varchar(60);not null" json:"code"`                    // 编码
	SupCode              string  `gorm:"index;column:sup_code;type:varchar(40);not null" json:"sup_code"`            // 供应商编码
	ProName              string  `gorm:"column:pro_name;type:varchar(150)" json:"pro_name"`                          // 产品名称
	ShortName            string  `gorm:"column:short_name;type:varchar(100);not null" json:"short_name"`             // 产品简称
	Bcat                 int     `gorm:"column:bcat;type:int(11);not null" json:"bcat"`                              // 产品分类（一级分类）
	Mcat                 int     `gorm:"column:mcat;type:int(11);not null" json:"mcat"`                              // 产品分类（二级分类）
	Scat                 int     `gorm:"column:scat;type:int(11);not null" json:"scat"`                              // 产品分类（三级分类）
	SupSku               string  `gorm:"column:sup_sku;type:varchar(50);not null" json:"sup_sku"`                    // 供应商货号
	BarCode              string  `gorm:"index;column:bar_code;type:varchar(50);not null" json:"bar_code"`            // 商品条码 ENA-13编码值
	BarCodeImg           string  `gorm:"column:bar_code_img;type:varchar(200)" json:"bar_code_img"`                  // 条码图片
	SellPoint            string  `gorm:"column:sell_point;type:varchar(200)" json:"sell_point"`                      // 卖点
	BrandID              int     `gorm:"column:brand_id;type:int(11);not null" json:"brand_id"`                      // 品牌ID
	OriginPlace          int     `gorm:"column:origin_place;type:int(11);not null" json:"origin_place"`              // 产地 国家编码
	ShipID               int     `gorm:"column:ship_id;type:int(11)" json:"ship_id"`                                 // 运费模板ID
	Volume               float64 `gorm:"column:volume;type:decimal(5,2)" json:"volume"`                              // 体积
	Weight               float64 `gorm:"column:weight;type:float(10,3);not null" json:"weight"`                      // 重量
	Keyword              string  `gorm:"column:keyword;type:varchar(150)" json:"keyword"`                            // 关键字
	Service              string  `gorm:"column:service;type:varchar(100);not null" json:"service"`                   // 服务保证
	Type                 int8    `gorm:"column:type;type:tinyint(1);not null" json:"type"`                           // 类型 2服务商品 1：实物产品   0：虚拟产品
	Pic                  string  `gorm:"column:pic;type:longtext" json:"pic"`                                        // 商品图片
	Standard1            string  `gorm:"column:standard_1;type:varchar(150)" json:"standard_1"`                      // 规格1
	Standard2            string  `gorm:"column:standard_2;type:varchar(150)" json:"standard_2"`                      // 规格2
	Content              string  `gorm:"column:content;type:text" json:"content"`                                    // 描述
	StatusRemark         string  `gorm:"column:status_remark;type:varchar(200)" json:"status_remark"`                // 审核信息
	Status               int8    `gorm:"column:status;type:tinyint(1);not null" json:"status"`                       // 状态1待审核2审核失败3审核通过4已上架5下架待审核6已下架7上架待审核
	IsDel                int8    `gorm:"index;column:is_del;type:tinyint(1);not null" json:"is_del"`                 // 是否删除  1：是  0：否
	CreateTime           int     `gorm:"column:create_time;type:int(10);not null" json:"create_time"`                // 创建时间
	Daifa                int8    `gorm:"column:daifa;type:tinyint(1)" json:"daifa"`                                  // 一件代发 1是  0否
	SetPrice             int8    `gorm:"column:set_price;type:smallint(2) unsigned;not null" json:"set_price"`       // 定价状态
	LeadingInStatus      int8    `gorm:"column:leading_in_status;type:tinyint(1) unsigned" json:"leading_in_status"` // 导入状态： 0未导入 1已经导入  （是否从供应商表 复制到平台商品表中）
	PackingList          string  `gorm:"column:packing_list;type:varchar(255);not null" json:"packing_list"`         // 包装清单
	UpdateTime           int     `gorm:"column:update_time;type:int(10) unsigned" json:"update_time"`                // 更新时间
	OnlineTime           int     `gorm:"column:online_time;type:int(11)" json:"online_time"`                         // 上线时间
	ProductSup           string  `gorm:"column:product_sup;type:varchar(50)" json:"product_sup"`                     // 商品供货商
	SubPerson            string  `gorm:"column:sub_person;type:varchar(20)" json:"sub_person"`                       // 商品提报人
	SupplierManagementID int     `gorm:"column:supplier_management_id;type:int(11)" json:"supplier_management_id"`   // 商品供货商管理ID
	OfflineTime          int     `gorm:"column:offline_time;type:int(11)" json:"offline_time"`                       // 下架时间
	CostType             int8    `gorm:"column:cost_type;type:tinyint(1)" json:"cost_type"`                          // 充值方式（1直充，2卡密，3兑换码）
	ProductType          int8    `gorm:"column:product_type;type:tinyint(1) unsigned zerofill" json:"product_type"`  // 商品类型（1为视频）可扩展
	DeliveryType         int8    `gorm:"column:delivery_type;type:tinyint(1)" json:"delivery_type"`                  // 发货方式（1为虚拟仓，2为api发送）
	GlobalSendType       int8    `gorm:"column:global_send_type;type:tinyint(1)" json:"global_send_type"`            // 跨境商品发货类型（1为保税，2直邮，3为国内） @author by 万朝
	ShipModelID          int     `gorm:"column:ship_model_id;type:int(1) unsigned zerofill" json:"ship_model_id"`    // 运费模式ID   1单品运费   2运单运费
	SupplyType           int8    `gorm:"index;column:supply_type;type:tinyint(2);not null" json:"supply_type"`       // 产品供货类型 1公开
	AkcActivityCode      string  `gorm:"column:akc_activity_code;type:varchar(100)" json:"akc_activity_code"`        // 第三方活动ID(liveId)
	AkcProductID         string  `gorm:"column:akc_product_id;type:varchar(100)" json:"akc_product_id"`              // 第三方产品ID(spuId)
	TaxDate              int     `gorm:"column:tax_date;type:int(11)" json:"tax_date"`                               // 税率
	ActivityID           int     `gorm:"column:activity_id;type:int(11)" json:"activity_id"`                         // 第三方活动主键
	IsSpecial            int8    `gorm:"column:is_special;type:tinyint(1);not null" json:"is_special"`               // 是否设置过特供 0 未设置 1 设置过
}

type GbCrmChat struct {
	ID          int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	BearerLogID int    `gorm:"column:bearer_log_id;type:int(11);not null" json:"bearer_log_id"`
	BearerID    int    `gorm:"column:bearer_id;type:int(10);not null" json:"bearer_id"`
	SessionID   int    `gorm:"column:session_id;type:int(11);not null" json:"session_id"`
	UserID      int    `gorm:"column:user_id;type:int(11);not null" json:"user_id"`
	ToUserID    int    `gorm:"column:to_user_id;type:int(11);not null" json:"to_user_id"`
	Msg         string `gorm:"column:msg;type:text;not null" json:"msg"`
	Status      int8   `gorm:"column:status;type:tinyint(1);not null" json:"status"` // 0 未读  1已读
	CreateAt    int    `gorm:"column:create_at;type:int(11);not null" json:"create_at"`
	Type        int8   `gorm:"column:type;type:tinyint(2);not null" json:"type"`                      // 1普通  2商品  3订单
	Originator  int8   `gorm:"column:originator;type:tinyint(2) unsigned;not null" json:"originator"` // 1 用户  2客服
	UserAvatar  string `gorm:"column:user_avatar;type:varchar(255)" json:"user_avatar"`
}

type GbExchangeStatus struct {
	ID         int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	ExchangeID int    `gorm:"column:exchange_id;type:int(11)" json:"exchange_id"`                 // 退货ID
	Worksheet  string `gorm:"column:worksheet;type:char(32)" json:"worksheet"`                    // 退货工单
	Status     int8   `gorm:"column:status;type:tinyint(2);not null" json:"status"`               // 当前状态
	UId        int    `gorm:"column:uid;type:int(10)" json:"uid"`                                 // 操作人
	Creattime  int    `gorm:"column:creattime;type:int(10) unsigned;not null" json:"creattime"`   // 创建时间
	Updatetime int    `gorm:"column:updatetime;type:int(10) unsigned;not null" json:"updatetime"` // 修改时间
}

// GbMerchantBargainTemplate 砍价返佣比例信息表，如，五人砍价，每个人分别返佣的百分比
type GbMerchantBargainTemplate struct {
	ID         int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Name       string `gorm:"column:name;type:varchar(255)" json:"name"`
	BargainNum int    `gorm:"column:bargain_num;type:int(11)" json:"bargain_num"` // 砍价次数
	SellerID   int    `gorm:"column:seller_id;type:int(11)" json:"seller_id"`
	IsDel      int8   `gorm:"column:is_del;type:tinyint(1)" json:"is_del"`
	IsDefault  int8   `gorm:"column:is_default;type:tinyint(1)" json:"is_default"` // 是否为默认 1 是，0否
}

// GbMerchantSeckillStandard 秒杀活动SKU定价
type GbMerchantSeckillStandard struct {
	ID                int     `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	SeckillID         int     `gorm:"column:seckill_id;type:int(11)" json:"seckill_id"`                     // 秒杀ID
	SeckillProID      int     `gorm:"column:seckill_pro_id;type:int(11)" json:"seckill_pro_id"`             // 秒杀商品ID(gb_merchant_seckill_product.id)
	PlatProID         int     `gorm:"column:plat_pro_id;type:int(11)" json:"plat_pro_id"`                   // 平台商品ID（gb.product表的ID）
	PlatProStandardID int     `gorm:"column:plat_pro_standard_id;type:int(11)" json:"plat_pro_standard_id"` // 平台商品SKU的ID（gb_prd_standard表的ID）
	PlatProSku        string  `gorm:"column:plat_pro_sku;type:varchar(50)" json:"plat_pro_sku"`             // 平台商品SKU
	PlatProCode       string  `gorm:"column:plat_pro_code;type:varchar(40)" json:"plat_pro_code"`
	MarketPrice       float64 `gorm:"column:market_price;type:decimal(10,2)" json:"market_price"` // 市场价
	CostPrice         float64 `gorm:"column:cost_price;type:decimal(10,2)" json:"cost_price"`     // 成本价，对应 平台的批发价
	SellPrice         float64 `gorm:"column:sell_price;type:decimal(10,2)" json:"sell_price"`     // 秒杀售价
	PlanNum           int     `gorm:"column:plan_num;type:int(11)" json:"plan_num"`               // 计划销售数量
	BuyCount          int     `gorm:"column:buy_count;type:int(11) unsigned" json:"buy_count"`    // 购买数量
	IsTop             int8    `gorm:"column:is_top;type:tinyint(1)" json:"is_top"`                // 是佛置顶
	MerchantID        int     `gorm:"column:merchant_id;type:int(11)" json:"merchant_id"`         // 商家ID号
}

type GbPurchaseIndex struct {
	ID                  int     `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	PurchaseIncrementID string  `gorm:"column:purchase_increment_id;type:char(20)" json:"purchase_increment_id"`                // 平台采购流水号（YYMMDD+）
	Freight             float64 `gorm:"column:freight;type:decimal(10,2) unsigned;not null" json:"freight"`                     // 运费总价
	GoodsPrice          float64 `gorm:"column:goods_price;type:decimal(10,2) unsigned;not null" json:"goods_price"`             // 商品总价
	TotalPrice          float64 `gorm:"column:total_price;type:decimal(10,2) unsigned;not null" json:"total_price"`             // 订单总价（商品总价+运费总价）
	InputFreight        float64 `gorm:"column:input_freight;type:decimal(10,2) unsigned;not null" json:"input_freight"`         // 进账运费总价
	InputGoodsPrice     float64 `gorm:"column:input_goods_price;type:decimal(10,2) unsigned;not null" json:"input_goods_price"` // 进账商品总价
	InputTotalPrice     float64 `gorm:"column:input_total_price;type:decimal(10,2) unsigned;not null" json:"input_total_price"` // 进账订单总价（商品总价+运费总价）
	CountPurchases      int     `gorm:"column:count_purchases;type:int(10)" json:"count_purchases"`                             // 订单数量
	Status              int8    `gorm:"column:status;type:tinyint(2);not null" json:"status"`                                   // 0失效，1有效
	Creattime           int     `gorm:"column:creattime;type:int(10) unsigned;not null" json:"creattime"`                       // 创建时间
	Updatetime          int     `gorm:"column:updatetime;type:int(10) unsigned;not null" json:"updatetime"`                     // 修改时间
	AdminID             int     `gorm:"column:admin_id;type:int(10)" json:"admin_id"`                                           // 操作人（平台用户ID）
}

type GbMerchantApply struct {
	ID             int       `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	CheckUId       int       `gorm:"column:check_uid;type:int(11) unsigned;not null" json:"check_uid"`      // 审核人
	CheckType      int8      `gorm:"column:check_type;type:tinyint(2) unsigned;not null" json:"check_type"` // 0 入驻 1变更
	MerchantID     int       `gorm:"column:merchant_id;type:int(11) unsigned;not null" json:"merchant_id"`
	CheckNote      string    `gorm:"column:check_note;type:varchar(512);not null" json:"check_note"`            // 审核备注
	CheckStatus    int8      `gorm:"column:check_status;type:tinyint(2) unsigned;not null" json:"check_status"` // 1通过 2不通过
	ApplyContent   string    `gorm:"column:apply_content;type:text" json:"apply_content"`                       // 申请内容 serialize
	CreateTime     time.Time `gorm:"column:create_time;type:datetime;not null" json:"create_time"`
	UpdateTime     time.Time `gorm:"column:update_time;type:datetime;not null" json:"update_time"`
	IsDel          int8      `gorm:"column:is_del;type:tinyint(2) unsigned;not null" json:"is_del"` // 软删除
	ChangeUserInfo string    `gorm:"column:change_user_info;type:text" json:"change_user_info"`     // 申请变更人
}

// GbMerchantShareModel 分成模式
type GbMerchantShareModel struct {
	ID                int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	SellerID          int    `gorm:"column:seller_id;type:int(11)" json:"seller_id"` // 分成模式所属商家ID
	Name              string `gorm:"column:name;type:varchar(100)" json:"name"`
	Description       string `gorm:"column:description;type:varchar(200)" json:"description"`            // 模板描述
	RelateTemplateNum int    `gorm:"column:relate_template_num;type:int(11)" json:"relate_template_num"` // 管理模板数量
	Ctime             int    `gorm:"column:ctime;type:int(11)" json:"ctime"`                             // 创建时间
	IsDel             int8   `gorm:"column:is_del;type:tinyint(1)" json:"is_del"`                        // 删除状态 1已经删除，0未删除
}

// GbOrderItem 订单产品表
type GbOrderItem struct {
	ID                    int     `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	OrderCode             string  `gorm:"column:order_code;type:char(100);not null" json:"order_code"`                   // 订单编码
	ProCode               string  `gorm:"column:pro_code;type:varchar(60);not null" json:"pro_code"`                     // 平台产品编码
	Sku                   string  `gorm:"column:sku;type:char(20);not null" json:"sku"`                                  // 产品规格编码
	ShipID                int     `gorm:"column:ship_id;type:int(11) unsigned" json:"ship_id"`                           // 运费模板ID
	UnitPrice             float64 `gorm:"column:unit_price;type:decimal(10,4) unsigned;not null" json:"unit_price"`      // 商品单价
	Count                 string  `gorm:"column:count;type:int(11) unsigned;not null" json:"count"`                      // 商品数量
	Commision             float64 `gorm:"column:commision;type:decimal(10,2)" json:"commision"`                          // 优惠券金额
	CountPrice            float64 `gorm:"column:count_price;type:decimal(10,4) unsigned;not null" json:"count_price"`    // 商品总价
	PackageID             string  `gorm:"column:package_id;type:varchar(20);not null" json:"package_id"`                 // 包裹单号（废除）
	Status                int8    `gorm:"column:status;type:tinyint(2) unsigned;not null" json:"status"`                 // 0未发货 1已发货
	Creattime             int     `gorm:"column:creattime;type:int(10) unsigned;not null" json:"creattime"`              // 创建时间
	EvaluateID            int     `gorm:"column:evaluate_id;type:int(11);not null" json:"evaluate_id"`                   // 评价ID 0未评价
	LogisticsID           int     `gorm:"column:logistics_id;type:int(11)" json:"logistics_id"`                          // 快递公司ID（废除）
	IncrementID           string  `gorm:"index;column:increment_id;type:char(20);not null" json:"increment_id"`          // 总订单流水号:123456789451245
	ChildrenIncrementID   string  `gorm:"index;column:children_increment_id;type:char(20)" json:"children_increment_id"` // 子订单流水号：123456789451245-0001
	BackendStatus         int     `gorm:"column:backend_status;type:int(2)" json:"backend_status"`                       // 参考参数文件：20种状态
	PayStatus             int     `gorm:"column:pay_status;type:int(1)" json:"pay_status"`                               // 0未付款1已经付款
	OrderStatus           int8    `gorm:"column:order_status;type:tinyint(2);not null" json:"order_status"`              // 1有效，0无效（取消）
	UpdateAt              int     `gorm:"column:update_at;type:int(10)" json:"update_at"`                                // 更新时间
	MBatchPrice           float64 `gorm:"column:m_batch_price;type:decimal(10,4)" json:"m_batch_price"`                  // 平台批发价（附加字段）
	Freight               float64 `gorm:"column:freight;type:decimal(10,2)" json:"freight"`                              // 运费总数
	OriginalPrice         float64 `gorm:"column:original_price;type:decimal(10,2)" json:"original_price"`                // 原始售价
	OriginalFreight       float64 `gorm:"column:original_freight;type:decimal(10,2)" json:"original_freight"`            // 原始运费
	UnitFreight           float64 `gorm:"column:unit_freight;type:decimal(10,2)" json:"unit_freight"`                    // 实际运费单价,暂时保留两位小数，实际计算时请用运费除以数量，结果自定义。
	MerchantProCode       string  `gorm:"column:merchant_pro_code;type:char(50)" json:"merchant_pro_code"`               // 商家产品编号
	MerchantProStandardID int     `gorm:"column:merchant_pro_standard_id;type:int(11)" json:"merchant_pro_standard_id"`  // 商家产品规格ID
	Weight                float64 `gorm:"column:weight;type:int(10)" json:"weight"`                                      // 重量
	Volume                float64 `gorm:"column:volume;type:int(10)" json:"volume"`                                      // 体积
	Property1             string  `gorm:"column:property_1;type:varchar(30)" json:"property_1"`                          // 规格属性1
	Property2             string  `gorm:"column:property_2;type:varchar(30)" json:"property_2"`                          // 规格属性2
	Pic                   string  `gorm:"column:pic;type:varchar(256)" json:"pic"`                                       // 图片
	ProName               string  `gorm:"column:pro_name;type:varchar(100)" json:"pro_name"`                             // 商品名称
	Jifenbuynum           int     `gorm:"column:jifenbuynum;type:int(10)" json:"jifenbuynum"`                            // 积分购买数量
	GroupID               int     `gorm:"column:group_id;type:int(3)" json:"group_id"`                                   // 运费分组id
	IsSpecial             int     `gorm:"column:is_special;type:int(255)" json:"is_special"`                             // 网红订单是否为公开还是特供：1为公开，2为特供
	WhPrice               string  `gorm:"column:wh_price;type:decimal(10,4);not null" json:"wh_price"`                   // 网红商品单价格
	WhPriceCount          float64 `gorm:"column:wh_price_count;type:decimal(10,4);not null" json:"wh_price_count"`       // 网红商品总价格
	CrossBorderTax        float64 `gorm:"column:cross_border_tax;type:decimal(10,2);not null" json:"cross_border_tax"`   // 跨境综合税额度
}

type GbUserTag struct {
	ID         int  `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	UId        int  `gorm:"column:uid;type:int(10);not null" json:"uid"`                      // 用户id
	MerchantID int  `gorm:"column:merchant_id;type:int(11)" json:"merchant_id"`               // 商家ID
	StoreID    int  `gorm:"column:store_id;type:int(11)" json:"store_id"`                     // 店铺ID
	Tag        int8 `gorm:"column:tag;type:tinyint(2);not null" json:"tag"`                   // 用户标签 ，0：消费者；1：BD；2：代理商(团长)
	Status     int8 `gorm:"column:status;type:tinyint(2);not null" json:"status"`             // 状态： 0：不可用；1：可用
	Creattime  int  `gorm:"column:creattime;type:int(10) unsigned;not null" json:"creattime"` // 创建时间
	Updatetime int  `gorm:"column:updatetime;type:int(10)" json:"updatetime"`                 // 更新时间
}

type GbStoreProductQtyLimit struct {
	ID         int64     `gorm:"primary_key;column:id;type:bigint(20) unsigned;not null" json:"-"`
	StoreProID int64     `gorm:"unique;column:store_pro_id;type:bigint(20) unsigned;not null" json:"store_pro_id"` // 店铺产品ID
	QtyLimit   int       `gorm:"column:qty_limit;type:int(11);not null" json:"qty_limit"`                          // 库存限制
	QtyUsed    int       `gorm:"column:qty_used;type:int(11);not null" json:"qty_used"`                            // 已消耗库存
	CreatedAt  time.Time `gorm:"column:created_at;type:datetime;not null" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;type:datetime;not null" json:"updated_at"`
}

type GbSupLogisticsRelevance struct {
	ID            int `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	FacilitatorID int `gorm:"column:facilitator_id;type:int(11);not null" json:"facilitator_id"` // 网点id
	AddressID     int `gorm:"column:address_id;type:int(11);not null" json:"address_id"`         // 发货地址id
}

// GbUserVirtualAccount 用户平台账户
type GbUserVirtualAccount struct {
	ID               int     `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	UserID           int     `gorm:"column:user_id;type:int(11) unsigned;not null" json:"user_id"`                  // 用户ID
	Pwd              string  `gorm:"column:pwd;type:char(40)" json:"pwd"`                                           // 提现密码
	Balance          float64 `gorm:"column:balance;type:decimal(10,6) unsigned;not null" json:"balance"`            // 账户金额
	BalanceAvailable float64 `gorm:"column:balance_available;type:decimal(10,6);not null" json:"balance_available"` // 可用余额
	Frozen           float64 `gorm:"column:frozen;type:decimal(10,6) unsigned;not null" json:"frozen"`              // 冻结金额
	Status           int8    `gorm:"column:status;type:tinyint(2);not null" json:"status"`                          // 状态：1,正常；2，关闭
	Creattime        int     `gorm:"column:creattime;type:int(10) unsigned;not null" json:"creattime"`              // 创建时间
	Updatatime       int     `gorm:"column:updatatime;type:int(10)" json:"updatatime"`                              // 更新时间
	ThirdBookkeeping float64 `gorm:"column:third_bookkeeping;type:decimal(10,2)" json:"third_bookkeeping"`          // 第三方用户记账金额
	CashMoney        float64 `gorm:"column:cash_money;type:decimal(10,2) unsigned;not null" json:"cash_money"`      // 已提现金额
}

// GbAccount 平台账户
type GbAccount struct {
	ID            int     `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	Pwd           string  `gorm:"column:pwd;type:char(40)" json:"pwd"`                                              // 提现密码
	Balance       float64 `gorm:"column:balance;type:decimal(10,2) unsigned;not null" json:"balance"`               // 账户金额
	Frozen        float64 `gorm:"column:frozen;type:decimal(10,2) unsigned;not null" json:"frozen"`                 // 冻结金额
	RefundBalance float64 `gorm:"column:refund_balance;type:decimal(10,2) unsigned;not null" json:"refund_balance"` // 退款临时存储金额
	Status        int8    `gorm:"column:status;type:tinyint(2);not null" json:"status"`                             // 状态：1,正常；2，关闭
	Creattime     int     `gorm:"column:creattime;type:int(10) unsigned;not null" json:"creattime"`                 // 创建时间
	Updatatime    int     `gorm:"column:updatatime;type:int(10)" json:"updatatime"`                                 // 更新时间
}

type GbProfitType struct {
	ID        int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	CreateUId int    `gorm:"column:create_uid;type:int(11) unsigned;not null" json:"create_uid"` // 创建人
	AppID     int    `gorm:"unique;column:app_id;type:int(11) unsigned;not null" json:"app_id"`  // 应用ID
	Name      string `gorm:"column:name;type:char(60);not null" json:"name"`                     // 业务名称
	Status    int8   `gorm:"column:status;type:tinyint(2) unsigned;not null" json:"status"`      // 状态默认1打开 0关闭
	IsDel     int8   `gorm:"column:is_del;type:tinyint(2) unsigned;not null" json:"is_del"`      // 删除状态 0未删除 1已删除
	Ctime     int    `gorm:"column:ctime;type:int(10) unsigned;not null" json:"ctime"`           // 创建时间
}

// GbRetailBank 零售商-银行卡表
type GbRetailBank struct {
	ID       int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	RetailID int    `gorm:"column:retail_id;type:int(11) unsigned;not null" json:"retail_id"`   // 零售商ID
	BankID   int    `gorm:"index;column:bank_id;type:int(11) unsigned;not null" json:"bank_id"` // 银行卡ID
	Network  string `gorm:"column:network;type:varchar(100);not null" json:"network"`           // 网点名称
	Account  string `gorm:"column:account;type:varchar(50);not null" json:"account"`            // 开户人
	Card     string `gorm:"column:card;type:varchar(100);not null" json:"card"`                 // 银行卡号
	Ctime    int    `gorm:"column:ctime;type:int(11);not null" json:"ctime"`                    // 创建时间
}

// GbAccountRefundDetails 平台账户退款流水
type GbAccountRefundDetails struct {
	ID            int     `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	WorksheetCode string  `gorm:"column:worksheet_code;type:char(32);not null" json:"worksheet_code"`      // 工单编码
	BalanceAction float64 `gorm:"column:balance_action;type:decimal(10,2);not null" json:"balance_action"` // 退款金额
	ActionTepy    int8    `gorm:"column:action_tepy;type:tinyint(2);not null" json:"action_tepy"`          // 操作类型：1,订单取消，2，
	AdminID       int     `gorm:"column:admin_id;type:int(11) unsigned;not null" json:"admin_id"`          // 平台操作员ID
	Message       string  `gorm:"column:message;type:varchar(255)" json:"message"`
	Status        int8    `gorm:"column:status;type:tinyint(2);not null" json:"status"`             // 状态：1,正常；2，关闭
	Creattime     int     `gorm:"column:creattime;type:int(10) unsigned;not null" json:"creattime"` // 创建时间
	Updatatime    int     `gorm:"column:updatatime;type:int(10)" json:"updatatime"`                 // 更新时间
	UId           int     `gorm:"column:uid;type:int(11)" json:"uid"`                               // 用户ID
}

type GbAPIRequestLog struct {
	ID           int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	RequestID    string `gorm:"column:request_id;type:char(36);not null" json:"request_id"`
	Route        string `gorm:"column:route;type:varchar(64)" json:"route"`
	Requests     string `gorm:"column:requests;type:text" json:"requests"`
	Ua           string `gorm:"column:ua;type:text" json:"ua"`
	IP           string `gorm:"column:ip;type:varchar(128)" json:"ip"`
	RequestTime  int64  `gorm:"column:request_time;type:bigint(15)" json:"request_time"`
	Response     string `gorm:"column:response;type:text" json:"response"`
	ResponseTime int64  `gorm:"column:response_time;type:bigint(15)" json:"response_time"`
	AccessToken  string `gorm:"column:access_token;type:varchar(255)" json:"access_token"`
	Timestamp    int    `gorm:"column:timestamp;type:int(11)" json:"timestamp"`
}

// GbCrmRole CRM系统 用户所属单位信息
type GbCrmRole struct {
	ID       int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	RoleName string `gorm:"column:role_name;type:varchar(255);not null" json:"role_name"`     // 单位名称
	ParentID int    `gorm:"column:parent_id;type:int(11) unsigned;not null" json:"parent_id"` // 父级ID
	Content  string `gorm:"column:content;type:text;not null" json:"content"`                 // 单位描述
	Status   int8   `gorm:"column:status;type:tinyint(2) unsigned;not null" json:"status"`    // 0不启用 1启用
	IsDel    int8   `gorm:"column:is_del;type:tinyint(2) unsigned;not null" json:"is_del"`    // 0未删除 1已删除
	Ctime    int    `gorm:"column:ctime;type:int(10) unsigned;not null" json:"ctime"`         // 创建时间
}

type GbMerchantBargainLog struct {
	ID         int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	BargainID  int    `gorm:"column:bargain_id;type:int(11)" json:"bargain_id"`
	EditorID   int    `gorm:"column:editor_id;type:int(11)" json:"editor_id"`  // 操作人ID
	PreContent string `gorm:"column:pre_content;type:text" json:"pre_content"` // 编辑之前的内容
	Content    string `gorm:"column:content;type:text" json:"content"`         // 编辑之后的内容
	Ctime      int    `gorm:"column:ctime;type:int(11)" json:"ctime"`          // 编辑时间
	Type       int8   `gorm:"column:type;type:tinyint(1)" json:"type"`         // 操作类型  1 新增 2编辑
}

// GbCrmQuickreplyCate 快捷回复分类表
type GbCrmQuickreplyCate struct {
	ID         int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Name       string    `gorm:"column:name;type:varchar(100);not null" json:"name"`         // 快捷回复分类名称
	Companyid  int       `gorm:"column:companyid;type:int(11)" json:"companyid"`             // 商户id
	Desc       string    `gorm:"column:desc;type:varchar(200);not null" json:"desc"`         // 分类描述
	IsDel      int8      `gorm:"column:is_del;type:smallint(6);not null" json:"is_del"`      // 逻辑删除0未删除1已删除
	Status     int8      `gorm:"column:status;type:smallint(6);not null" json:"status"`      // 回复内容状态：0已禁用1已启用
	Createtime time.Time `gorm:"column:createtime;type:datetime;not null" json:"createtime"` // 分类创建时间
}

// GbMerchantVirtualAccount 商家平台账户
type GbMerchantVirtualAccount struct {
	ID           int     `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	MerchantCode string  `gorm:"column:merchant_code;type:char(30);not null" json:"merchant_code"`     // 商家编号
	MerchantID   int     `gorm:"column:merchant_id;type:int(11) unsigned;not null" json:"merchant_id"` // 商家ID
	Pwd          string  `gorm:"column:pwd;type:char(40)" json:"pwd"`                                  // 提现密码
	Balance      float64 `gorm:"column:balance;type:decimal(10,2) unsigned;not null" json:"balance"`   // 账户可用金额
	Frozen       float64 `gorm:"column:frozen;type:decimal(10,2) unsigned;not null" json:"frozen"`     // 冻结金额
	Status       int8    `gorm:"column:status;type:tinyint(2);not null" json:"status"`                 // 状态：1,正常；2，关闭
	Creattime    int     `gorm:"column:creattime;type:int(10) unsigned;not null" json:"creattime"`     // 创建时间
	Updatatime   int     `gorm:"column:updatatime;type:int(10)" json:"updatatime"`                     // 更新时间
}

type GbOrderPurchaseMergeInvoice struct {
	ID                  int     `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`                       // 主键ID
	PayJifen            float64 `gorm:"column:pay_jifen;type:decimal(10,2);not null" json:"pay_jifen"`                       // 用户花费积分
	PayMoney            float64 `gorm:"column:pay_money;type:decimal(10,2);not null" json:"pay_money"`                       // 用户实际支付金额
	Commision           float64 `gorm:"column:commision;type:decimal(10,2);not null" json:"commision"`                       // 用户优惠券金额
	Freight             float64 `gorm:"column:freight;type:decimal(10,2);not null" json:"freight"`                           // 用户订单商品运费总价
	Price               float64 `gorm:"column:price;type:decimal(10,2);not null" json:"price"`                               // 商品总价，不含邮费
	CountPrice          float64 `gorm:"column:count_price;type:decimal(10,2);not null" json:"count_price"`                   // 用户订单商品总价
	BatchValue          int8    `gorm:"column:batch_value;type:smallint(2) unsigned" json:"batch_value"`                     // 1表示可以开发票，0表示不可以开发票
	PurchaseIncrementID string  `gorm:"column:purchase_increment_id;type:char(20);not null" json:"purchase_increment_id"`    // 平台采购订单号
	PurchasePrice       float64 `gorm:"column:purchase_price;type:decimal(10,2);not null" json:"purchase_price"`             // 采购总价，不含邮费
	PurchaseFreight     float64 `gorm:"column:purchase_freight;type:decimal(10,2);not null" json:"purchase_freight"`         // 采购运费
	CountPurchasePrice  float64 `gorm:"column:count_purchase_price;type:decimal(10,2);not null" json:"count_purchase_price"` // 采购订单商品总价，包含邮费
	OrderQuantity       int     `gorm:"column:order_quantity;type:int(11) unsigned;not null" json:"order_quantity"`          // 订单数量
	GoodsQuantity       int     `gorm:"column:goods_quantity;type:int(11) unsigned;not null" json:"goods_quantity"`          // 商品数量
	PayStatus           int8    `gorm:"column:pay_status;type:tinyint(4) unsigned;not null" json:"pay_status"`               // 0表示未操作付款，1表示已经操作付款了
	Creattime           int     `gorm:"index;column:creattime;type:int(10) unsigned;not null" json:"creattime"`              // 创建时间
	Updatetime          int     `gorm:"column:updatetime;type:int(10) unsigned;not null" json:"updatetime"`                  // 更新时间
}

type GbPlatformPrice struct {
	ID               int     `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	SkuID            int     `gorm:"column:sku_id;type:int(11) unsigned;not null" json:"sku_id"`                               // 平台库产品SKUID
	SupplyGoodsPrice float64 `gorm:"column:supply_goods_price;type:decimal(11,2) unsigned;not null" json:"supply_goods_price"` // 平台供货价
	Ctime            int     `gorm:"column:ctime;type:int(10) unsigned;not null" json:"ctime"`                                 // 最新时间
}

// GbSupVirtualAccount 供应商平台账户
type GbSupVirtualAccount struct {
	ID             int     `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	SupCode        string  `gorm:"column:sup_code;type:char(30);not null" json:"sup_code"`             // 供应商编号
	SupID          int     `gorm:"column:sup_id;type:int(11) unsigned;not null" json:"sup_id"`         // 供应商ID
	CashWithDrawal float64 `gorm:"column:cash_with_drawal;type:decimal(10,4)" json:"cash_with_drawal"` // 提现中金额
	TotalPayment   float64 `gorm:"column:total_payment;type:decimal(10,4)" json:"total_payment"`       // 货款总额
	Withdraw       float64 `gorm:"column:Withdraw;type:decimal(10,4)" json:"Withdraw"`                 // 已提现金额
	Balance        float64 `gorm:"column:balance;type:decimal(10,4) unsigned;not null" json:"balance"` // 账户可用金额
	Frozen         float64 `gorm:"column:frozen;type:decimal(10,4) unsigned;not null" json:"frozen"`   // 账户冻结金额
	Status         int8    `gorm:"column:status;type:tinyint(2);not null" json:"status"`               // 状态：1,正常；2，关闭
	Creattime      int     `gorm:"column:creattime;type:int(10) unsigned;not null" json:"creattime"`   // 创建时间
	Updatatime     int     `gorm:"column:updatatime;type:int(10)" json:"updatatime"`                   // 更新时间
	Pwd            string  `gorm:"column:pwd;type:varchar(255)" json:"pwd"`                            // 密码
}

type GbAdministrator struct {
	ID                 int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`   // ID
	Point              int8   `gorm:"column:point;type:tinyint(4) unsigned;not null" json:"point"`     // 终端(1-雀信后端，2-商家后端，3-财务后端)
	Username           string `gorm:"index;column:username;type:varchar(20);not null" json:"username"` // 用户名
	Password           string `gorm:"column:password;type:varchar(255);not null" json:"password"`      // 密码
	Name               string `gorm:"column:name;type:varchar(255);not null" json:"name"`              // 姓名
	Mail               string `gorm:"column:mail;type:varchar(255)" json:"mail"`                       // 邮箱
	Salt               string `gorm:"column:salt;type:varchar(16);not null" json:"salt"`               // 混淆码
	Telephone          string `gorm:"column:telephone;type:varchar(20)" json:"telephone"`              // 联系电话
	Mobile             string `gorm:"column:mobile;type:char(11)" json:"mobile"`                       // 手机号
	Sex                int8   `gorm:"column:sex;type:tinyint(1)" json:"sex"`                           // 1、男性   2、女性
	LastLoginTime      int    `gorm:"column:last_login_time;type:int(11)" json:"last_login_time"`      // 上次登陆时间
	LastLoginIP        string `gorm:"column:last_login_ip;type:varchar(15)" json:"last_login_ip"`      // 上次登陆IP
	Status             int8   `gorm:"column:status;type:tinyint(1) unsigned" json:"status"`            // 状态  0：不可用  1：可用
	IsDel              int8   `gorm:"column:is_del;type:smallint(2) unsigned;not null" json:"is_del"`  // 删除状态   1表示已删除
	Remark             string `gorm:"column:remark;type:text" json:"remark"`                           // 备注信息
	CreateTime         int    `gorm:"column:create_time;type:int(11)" json:"create_time"`              // 创建时间
	AccessToken        string `gorm:"column:access_token;type:varchar(255)" json:"access_token"`
	AccessTokenExpire  int    `gorm:"column:access_token_expire;type:int(10) unsigned" json:"access_token_expire"`
	Allowance          int    `gorm:"column:allowance;type:int(10) unsigned" json:"allowance"`
	AllowanceUpdatedAt int    `gorm:"column:allowance_updated_at;type:int(10) unsigned" json:"allowance_updated_at"`
	RateLimit          int    `gorm:"column:rate_limit;type:int(10) unsigned" json:"rate_limit"`
}

// GbStoreZone 店铺-广告专区表
type GbStoreZone struct {
	ID      int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	StoreID int    `gorm:"column:store_id;type:int(11) unsigned;not null" json:"store_id"` // 店铺ID
	Name    string `gorm:"column:name;type:varchar(50);not null" json:"name"`              // 名称
	Pic     string `gorm:"column:pic;type:varchar(255);not null" json:"pic"`               // 广告图
	URL     string `gorm:"column:url;type:varchar(255);not null" json:"url"`               // 链接
	Sort    int    `gorm:"column:sort;type:int(11);not null" json:"sort"`                  // 排序
	IsDel   int8   `gorm:"column:is_del;type:tinyint(1);not null" json:"is_del"`           // 1删除，0未删除
	Status  int8   `gorm:"column:status;type:tinyint(1);not null" json:"status"`           // 状态 1 开启，0 关闭
	Ctime   int    `gorm:"column:ctime;type:int(11);not null" json:"ctime"`                // 创建时间
}

type GbCollectCategory struct {
	ID        int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	SiteID    int    `gorm:"column:site_id;type:int(11);not null" json:"site_id"`
	Pid       int    `gorm:"column:pid;type:int(11);not null" json:"pid"`        // 上级ID
	Fid       int    `gorm:"column:fid;type:int(11);not null" json:"fid"`        // 上上级ID
	Name      string `gorm:"column:name;type:varchar(255);not null" json:"name"` // 名称
	CreatedAt int    `gorm:"column:created_at;type:int(11);not null" json:"created_at"`
	UpdatedAt int    `gorm:"column:updated_at;type:int(11);not null" json:"updated_at"`
	IsDel     int8   `gorm:"column:is_del;type:tinyint(1);not null" json:"is_del"`
}

type GbMerchantProShareImage struct {
	ID            int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	StoreID       int    `gorm:"column:store_id;type:int(11) unsigned;not null" json:"store_id"`               // 店铺ID
	MerchantSpuID int    `gorm:"column:merchant_spu_id;type:int(11)" json:"merchant_spu_id"`                   // 商家SPU ID
	MerchantSkuID int    `gorm:"column:merchant_sku_id;type:int(11) unsigned;not null" json:"merchant_sku_id"` // 商家规格ID
	LockBaseImage string `gorm:"column:lock_base_image;type:varchar(250)" json:"lock_base_image"`              // 锁粉分享基图（包含商品信息，不包含团长信息）
	Image         string `gorm:"column:image;type:varchar(250);not null" json:"image"`                         // 分享图片
	CreatedAt     int    `gorm:"column:created_at;type:int(10) unsigned;not null" json:"created_at"`
	UpdatedAt     int    `gorm:"column:updated_at;type:int(10) unsigned;not null" json:"updated_at"`
}

// GbProEditLog 平台商品编辑日志记录
type GbProEditLog struct {
	ID         int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	OptUId     int    `gorm:"column:opt_uid;type:int(11)" json:"opt_uid"` // 操作人ID
	ProID      int    `gorm:"column:pro_id;type:int(11)" json:"pro_id"`
	ProCode    string `gorm:"column:pro_code;type:varchar(40)" json:"pro_code"`        // 商品编码
	PreContent string `gorm:"column:pre_content;type:varchar(200)" json:"pre_content"` // 编辑前的内容
	EndContent string `gorm:"column:end_content;type:varchar(255)" json:"end_content"` // 编辑后的内容
	Ctime      int    `gorm:"column:ctime;type:int(11)" json:"ctime"`                  // 创建时间
	IsDel      int8   `gorm:"column:is_del;type:tinyint(1)" json:"is_del"`             // 是否删除 0未删除 1已删除
}

type GbSequence struct {
	Name           string `gorm:"primary_key;column:name;type:varchar(50);not null" json:"name"`          // 自增序列名称
	StartValue     int64  `gorm:"column:start_value;type:bigint(12);not null" json:"start_value"`         // 起始值
	IncrementValue int64  `gorm:"column:increment_value;type:bigint(12);not null" json:"increment_value"` // 递增步长
}

type GbThirdOrder struct {
	ID            int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	RequestID     string `gorm:"column:request_id;type:varchar(36)" json:"request_id"`
	APIUserID     int    `gorm:"column:api_user_id;type:int(10) unsigned;not null" json:"api_user_id"`
	OutTradeNo    string `gorm:"column:out_trade_no;type:varchar(64);not null" json:"out_trade_no"` // 商城订单号
	Username      string `gorm:"column:username;type:varchar(64)" json:"username"`                  // 客户姓名
	Consignee     string `gorm:"column:consignee;type:varchar(64);not null" json:"consignee"`       // 收货人
	Mobile        string `gorm:"column:mobile;type:char(11);not null" json:"mobile"`                // 收货手机号
	Address       string `gorm:"column:address;type:text" json:"address"`
	PackageNo     string `gorm:"column:package_no;type:varchar(64)" json:"package_no"`         // 服务商系统的套餐编号
	TransactionID string `gorm:"column:transaction_id;type:varchar(64)" json:"transaction_id"` // 服务商 订单号
	CreatedAt     int    `gorm:"column:created_at;type:int(11)" json:"created_at"`
	UpdatedAt     int    `gorm:"column:updated_at;type:int(11)" json:"updated_at"`
	IsDel         int8   `gorm:"column:is_del;type:tinyint(1)" json:"is_del"` // 0已生效   1未生效
}

type GbWebsiteStructure struct {
	ID          int       `gorm:"primary_key;column:id;type:int(10);not null" json:"-"`      // 网站结构ID
	WebsiteType int8      `gorm:"column:website_type;type:tinyint(1)" json:"website_type"`   // 官网类型 1神雀 2鼎翰
	Pid         int       `gorm:"column:pid;type:int(10)" json:"pid"`                        // 父级ID 0表示一级 其他均对应其父级id
	Name        string    `gorm:"column:name;type:varchar(255)" json:"name"`                 // 名称
	LinkType    int8      `gorm:"column:link_type;type:tinyint(1)" json:"link_type"`         // 连接类型 1URL 2内容分类(发布渠道+新闻类型)
	LinkContent string    `gorm:"column:link_content;type:varchar(255)" json:"link_content"` // 连接类型对应内容值
	IconA       string    `gorm:"column:icon_a;type:varchar(255)" json:"icon_a"`             // 图标一
	IconB       string    `gorm:"column:icon_b;type:varchar(255)" json:"icon_b"`             // 图标二
	Sort        int       `gorm:"column:sort;type:int(10)" json:"sort"`                      // 排序
	HideType    int       `gorm:"column:hide_type;type:int(1)" json:"hide_type"`             // 隐藏状态 0隐藏 1显示
	Status      int8      `gorm:"column:status;type:tinyint(1)" json:"status"`               // 状态 0删除 1正常
	CreateTime  time.Time `gorm:"column:create_time;type:datetime" json:"create_time"`       // 创建时间
	UpdateTime  time.Time `gorm:"column:update_time;type:datetime" json:"update_time"`       // 修改时间
}

type GbCashCommissionRelation struct {
	ID     int  `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	CashID int  `gorm:"index;column:cash_id;type:int(11);not null" json:"cash_id"` // 提现ID
	FlowID int  `gorm:"column:flow_id;type:int(11);not null" json:"flow_id"`       // 用户流水ID
	Status int8 `gorm:"column:status;type:tinyint(1);not null" json:"status"`      // 0待发放 1已发放  2发放失败
	Type   int8 `gorm:"column:type;type:tinyint(1);not null" json:"type"`          // 1财务系统  2人工 3自动提现
}

type GbCollectArticleContent struct {
	ID        int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	ArticleID int    `gorm:"column:article_id;type:int(11);not null" json:"article_id"`
	Content   string `gorm:"column:content;type:longtext;not null" json:"content"`
}

// GbCrmSection CRM部门表
type GbCrmSection struct {
	ID        int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`      // 部门ID
	CompanyID int    `gorm:"column:company_id;type:int(11);not null" json:"company_id"` // 关联的公司ID
	Name      string `gorm:"column:name;type:varchar(255);not null" json:"name"`        // 部门名称
	Desc      string `gorm:"column:desc;type:text;not null" json:"desc"`                // 部门描述
	Enable    int8   `gorm:"column:enable;type:tinyint(2)" json:"enable"`               // 启用状态  1为启用
	IsDel     int8   `gorm:"column:is_del;type:tinyint(2)" json:"is_del"`               // 删除状态 1为删除
	Ctime     int    `gorm:"column:ctime;type:int(10);not null" json:"ctime"`           //  创建时间
}

// GbMerchant 平台商家表
type GbMerchant struct {
	ID                    int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Code                  string    `gorm:"column:code;type:varchar(30);not null" json:"code"` // 编号，管理后台给商家的编号
	Username              string    `gorm:"column:username;type:varchar(20)" json:"username"`  // 登录用户名
	Pwd                   string    `gorm:"column:pwd;type:char(40)" json:"pwd"`               // 登录密码
	Salt                  string    `gorm:"column:salt;type:char(4)" json:"salt"`
	RegPhone              string    `gorm:"column:reg_phone;type:varchar(30);not null" json:"reg_phone"`                               // 注册手机号
	CompanyName           string    `gorm:"column:company_name;type:varchar(50)" json:"company_name"`                                  // 公司名称
	CompanyTaxType        int8      `gorm:"column:company_tax_type;type:tinyint(1) unsigned;not null" json:"company_tax_type"`         // 1一般纳税人 2小规模纳税人
	CompanyArea           int8      `gorm:"column:company_area;type:tinyint(1) unsigned;not null" json:"company_area"`                 // 0 大陆 1港澳台 2境外
	IsMergeID             int8      `gorm:"column:is_merge_id;type:tinyint(1) unsigned;not null" json:"is_merge_id"`                   // 是否多证合一 0 否 1是
	CompanyType           int8      `gorm:"column:company_type;type:tinyint(1)" json:"company_type"`                                   // 企业类型 1普通企业 2 一般纳税人企业，3 其他
	CompanyAddrProvinceID int       `gorm:"column:company_addr_province_id;type:int(10)" json:"company_addr_province_id"`              // 公司所在地址省ID
	CompanyAddrCityID     int       `gorm:"column:company_addr_city_id;type:int(10)" json:"company_addr_city_id"`                      // 公司所在地址市ID
	CompanyAddrCountyID   int       `gorm:"column:company_addr_county_id;type:int(10)" json:"company_addr_county_id"`                  // 公司所在地址区县ID
	CompanyAddressPre     string    `gorm:"column:company_address_pre;type:varchar(50)" json:"company_address_pre"`                    // 公司地址前缀，由省、市、区县组成
	CompanyAddress        string    `gorm:"column:company_address;type:varchar(100)" json:"company_address"`                           // 办公地址
	ServiceLine           string    `gorm:"column:service_line;type:varchar(30)" json:"service_line"`                                  // 业务线
	CreditCode            string    `gorm:"column:credit_code;type:varchar(30)" json:"credit_code"`                                    // 组织机构代码/信用代码
	BusinessLicense       string    `gorm:"column:business_license;type:varchar(500);not null" json:"business_license"`                // 营业执照副本上传地址
	PlatDockingPeople     string    `gorm:"column:plat_docking_people;type:varchar(255);not null" json:"plat_docking_people"`          // 平台对接人
	BusinessLicence       string    `gorm:"column:business_licence;type:varchar(100)" json:"business_licence"`                         // 营业执照图片地址
	ContactsName          string    `gorm:"column:contacts_name;type:varchar(30)" json:"contacts_name"`                                // 联系人姓名
	ContactsMobile        string    `gorm:"column:contacts_mobile;type:varchar(30)" json:"contacts_mobile"`                            // 联系人手机号码
	ContactsPhone         string    `gorm:"column:contacts_phone;type:varchar(30)" json:"contacts_phone"`                              // 联系人固定电话
	ContactsEmail         string    `gorm:"column:contacts_email;type:varchar(20)" json:"contacts_email"`                              // 联系人邮箱地址
	BusinessTimeLimit     time.Time `gorm:"column:business_time_limit;type:datetime;not null" json:"business_time_limit"`              // 营业期限
	ContactsFax           string    `gorm:"column:contacts_fax;type:varchar(30)" json:"contacts_fax"`                                  // 联系人传真
	LeaderName            string    `gorm:"column:leader_name;type:varchar(30)" json:"leader_name"`                                    // 负责人姓名
	LeaderMobile          string    `gorm:"column:leader_mobile;type:varchar(30)" json:"leader_mobile"`                                // 负责人手机号码
	BusinessScope         string    `gorm:"column:business_scope;type:varchar(250)" json:"business_scope"`                             // 经营范围
	Ctime                 int       `gorm:"column:ctime;type:int(10)" json:"ctime"`                                                    // 加入时间
	Mtime                 int       `gorm:"column:mtime;type:int(11)" json:"mtime"`                                                    // 后台修改时间
	SellProductNum        int       `gorm:"column:sell_product_num;type:int(11)" json:"sell_product_num"`                              // 在售产品种类数量
	Status                int8      `gorm:"column:status;type:tinyint(1)" json:"status"`                                               // 启用状态 0 未启用 ，1 已启用
	IsDel                 int8      `gorm:"column:is_del;type:tinyint(1);not null" json:"is_del"`                                      // 是否删除  1：是   0：否
	LastLoginTime         int       `gorm:"column:last_login_time;type:int(11)" json:"last_login_time"`                                // 最后登录时间
	LastLoginIP           int64     `gorm:"column:last_login_ip;type:bigint(15)" json:"last_login_ip"`                                 // 最后登录的IP
	ManagementType        int8      `gorm:"column:management_type;type:tinyint(1)" json:"management_type"`                             // 经营类型 1 自营，2 第三方
	IsCheck               int8      `gorm:"column:is_check;type:tinyint(1) unsigned;not null" json:"is_check"`                         // 入驻审核状态 0 未审核 1 通过 2不通过
	AccessToken           string    `gorm:"column:access_token;type:varchar(255);not null" json:"access_token"`                        // 身份凭证
	AccessTokenExpire     int       `gorm:"column:access_token_expire;type:int(11) unsigned;not null" json:"access_token_expire"`      // 存活时间
	SystemType            int8      `gorm:"column:system_type;type:tinyint(1);not null" json:"system_type"`                            // 系统类型  1商城系统 2团购平台
	AkcAvailableStatus    int8      `gorm:"column:akc_available_status;type:tinyint(1) unsigned zerofill" json:"akc_available_status"` // 0为没选中，1为选中
	AutoPurchase          int8      `gorm:"column:auto_purchase;type:tinyint(1)" json:"auto_purchase"`                                 // 0为手动采购，1为自动采购；陈伟强加字段
}

type GbStoreCardCouponBuyQuota struct {
	ID         int64     `gorm:"primary_key;column:id;type:bigint(64) unsigned;not null" json:"-"`
	StoreSpuID int       `gorm:"column:store_spu_id;type:int(11) unsigned;not null" json:"store_spu_id"` // 店铺SPU_ID
	StoreID    int       `gorm:"column:store_id;type:int(11) unsigned;not null" json:"store_id"`         // 店铺ID
	Num        int       `gorm:"column:num;type:int(11);not null" json:"num"`                            // 限购数量
	IsDel      int8      `gorm:"column:is_del;type:tinyint(2);not null" json:"is_del"`                   // 软删除  0 否
	CreatedAt  time.Time `gorm:"column:created_at;type:datetime;not null" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;type:datetime;not null" json:"updated_at"`
}

type GbOrderIDentity struct {
	ID              int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`        // 主键ID
	UId             int    `gorm:"column:uid;type:int(11);not null" json:"uid"`                          // 用户id
	IncrementID     string `gorm:"index;column:increment_id;type:char(20);not null" json:"increment_id"` // 总订单流水号:123456789451245
	Name            string `gorm:"column:name;type:varchar(36);not null" json:"name"`                    // 真实姓名
	IDcard          string `gorm:"column:idcard;type:varchar(36);not null" json:"idcard"`                // 身份证号码
	FrontendCardPic string `gorm:"column:frontend_card_pic;type:varchar(256)" json:"frontend_card_pic"`  // 身份证正面图片地址
	BackendCardPic  string `gorm:"column:backend_card_pic;type:varchar(256)" json:"backend_card_pic"`    // 身份证反面图片地址
	Creattime       int    `gorm:"column:creattime;type:int(10) unsigned;not null" json:"creattime"`     // 创建时间
}

type GbSupWithdrawLog struct {
	ID               int     `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`                            // 编号
	SupUId           int     `gorm:"column:sup_uid;type:int(10) unsigned;not null" json:"sup_uid"`                             // 提现用户编号
	FlowNumber       string  `gorm:"column:flow_number;type:varchar(64);not null" json:"flow_number"`                          // 提现流水号
	Type             string  `gorm:"column:type;type:varchar(255);not null" json:"type"`                                       // 提现方式（0 支付宝）
	Money            float64 `gorm:"column:money;type:decimal(19,4) unsigned;not null" json:"money"`                           // 可提现余额
	CanWithdrawMoney float64 `gorm:"column:can_withdraw_money;type:decimal(19,4) unsigned;not null" json:"can_withdraw_money"` // 可提现金额
	CreateAt         int     `gorm:"column:create_at;type:int(10) unsigned;not null" json:"create_at"`                         // 创建时间
	UpdateAt         int     `gorm:"column:update_at;type:int(10) unsigned;not null" json:"update_at"`                         // 更新时间
	Msg              string  `gorm:"column:msg;type:varchar(2000);not null" json:"msg"`                                        // 反馈信息
	Status           int8    `gorm:"column:status;type:tinyint(1) unsigned;not null" json:"status"`                            // 提现状态 （0 待处理 1 处理中  2  审核通过 3 审核不通过）
	OpID             int     `gorm:"column:op_id;type:int(10)" json:"op_id"`                                                   // 操作人员编号
	OpName           string  `gorm:"column:op_name;type:varchar(50)" json:"op_name"`                                           // 操作人员名称
	Account          string  `gorm:"column:account;type:varchar(50);not null" json:"account"`                                  // 收款账号
	RealName         string  `gorm:"column:real_name;type:varchar(255);not null" json:"real_name"`                             // 真实姓名
	NickName         string  `gorm:"column:nick_name;type:varchar(50);not null" json:"nick_name"`                              // 昵称
	Cktime           int     `gorm:"column:cktime;type:int(10)" json:"cktime"`                                                 // 审核时间
	IsPay            int8    `gorm:"column:is_pay;type:tinyint(2)" json:"is_pay"`                                              // 是否支付 1支付 0未支付
}

// GbTaobaoPropValuesOriginal 淘宝类目属性表
type GbTaobaoPropValuesOriginal struct {
	ID   int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Cid  int    `gorm:"index:ids_cid_pid;column:cid;type:int(11);not null" json:"cid"` // 类目ID
	Pid  int    `gorm:"index:ids_cid_pid;column:pid;type:int(11);not null" json:"pid"` // 属性ID
	Vid  int    `gorm:"index;column:vid;type:int(11);not null" json:"vid"`             // 属性值ID
	Name string `gorm:"column:name;type:varchar(50);not null" json:"name"`             // 属性值名称
}

type GbUserCardCouponOptionLog struct {
	ID               int       `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	CardCouponID     int64     `gorm:"column:card_coupon_id;type:bigint(64) unsigned" json:"card_coupon_id"`       // 卡券ID 充值会要求填写卡券ID
	StoreID          int       `gorm:"column:store_id;type:int(11) unsigned;not null" json:"store_id"`             // 店铺ID
	UserID           int       `gorm:"column:user_id;type:int(11) unsigned;not null" json:"user_id"`               // 充值用户ID
	ActionType       int8      `gorm:"column:action_type;type:tinyint(2);not null" json:"action_type"`             // 操作类型 0 消费  1 充值
	ActionDetail     string    `gorm:"column:action_detail;type:varchar(255);not null" json:"action_detail"`       // 操作详细
	ActionNote       string    `gorm:"column:action_note;type:varchar(512);not null" json:"action_note"`           // 操作备注
	OrderIncrementID string    `gorm:"column:order_increment_id;type:char(20);not null" json:"order_increment_id"` // 订单流水号
	RechargeIP       string    `gorm:"column:recharge_ip;type:varchar(30)" json:"recharge_ip"`                     // 操作用户IP
	CreatedAt        time.Time `gorm:"column:created_at;type:datetime;not null" json:"created_at"`                 // 创建时间
}

type GbUserToken struct {
	ID                 int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	Point              int8   `gorm:"column:point;type:tinyint(3) unsigned;not null" json:"point"`                            // 业务端（1-雀信h5）
	Mobile             string `gorm:"unique;column:mobile;type:varchar(40);not null" json:"mobile"`                           // 手机号
	Password           string `gorm:"column:password;type:varchar(100);not null" json:"password"`                             // 密码
	Status             int8   `gorm:"column:status;type:tinyint(3) unsigned;not null" json:"status"`                          // 状态 0关闭帐户  1为打开
	AccessToken        string `gorm:"unique;column:access_token;type:varchar(50);not null" json:"access_token"`               // 用户token
	TokenTime          int    `gorm:"column:token_time;type:int(10) unsigned;not null" json:"token_time"`                     // token生成时间
	AccessTokenExpire  int    `gorm:"column:access_token_expire;type:int(10) unsigned;not null" json:"access_token_expire"`   // token过期时间
	Allowance          int    `gorm:"column:allowance;type:int(10) unsigned;not null" json:"allowance"`                       // 允许的请求数
	AllowanceUpdatedAt int    `gorm:"column:allowance_updated_at;type:int(10) unsigned;not null" json:"allowance_updated_at"` // restful请求的UNIX时间戳
	RateLimit          int    `gorm:"column:rate_limit;type:int(11) unsigned;not null" json:"rate_limit"`                     // 允许的最大请求数
}

type GbWhUserRegisterRelation struct {
	ID              int  `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	UserID          int  `gorm:"column:user_id;type:int(11) unsigned;not null" json:"user_id"`                       // 注册用户ID (可能是供应商/网红)
	BeInvitedUserID int  `gorm:"column:be_invited_user_id;type:int(11) unsigned;not null" json:"be_invited_user_id"` // 被邀请的用户ID(供应商/网红)
	Type            int8 `gorm:"column:type;type:tinyint(2) unsigned;not null" json:"type"`                          // 关系链类型 1 网红+网红 2 网红+供应商 3 供应商+网红 4 供应商+供应商
	CreatedAt       int  `gorm:"column:created_at;type:int(11) unsigned;not null" json:"created_at"`                 // 创建时间
}

// GbAreaCountry 国家代码对应表，用于订单生成,用于商品编码生成
type GbAreaCountry struct {
	ID    int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Code  string `gorm:"column:code;type:char(3)" json:"code"`        // 国家代码值
	Name  string `gorm:"column:name;type:varchar(20)" json:"name"`    // 国家名称
	IsDel int8   `gorm:"column:is_del;type:tinyint(1)" json:"is_del"` // 删除状态，1 已删除，0未删除
}

// GbCrmCompany CRM公司表
type GbCrmCompany struct {
	ID     int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`          // 公司表
	Name   string `gorm:"column:name;type:varchar(255);not null" json:"name"`            // 公司名称
	Desc   string `gorm:"column:desc;type:text;not null" json:"desc"`                    // 公司描述
	Enable int8   `gorm:"column:enable;type:tinyint(2) unsigned;not null" json:"enable"` // 启用状态 0为未启用1为启用
	IsDel  int8   `gorm:"column:is_del;type:tinyint(2) unsigned;not null" json:"is_del"` // 删除状态0为未删除 1为已删除
	Ctime  int    `gorm:"column:ctime;type:int(10);not null" json:"ctime"`               // 添加时间
}

type GbCrmWorkOrder struct {
	ID                   int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	JobNo                string `gorm:"column:job_no;type:varchar(64);not null" json:"job_no"`                             // 工单号
	Title                string `gorm:"column:title;type:varchar(255);not null" json:"title"`                              // 标题
	Content              string `gorm:"column:content;type:text;not null" json:"content"`                                  // 内容
	Enclosure            string `gorm:"column:enclosure;type:text" json:"enclosure"`                                       // 附件
	Level                int8   `gorm:"column:level;type:tinyint(1);not null" json:"level"`                                // 1非常紧急  2紧急  3一般  4低
	AcceptanceSectionID  int    `gorm:"column:acceptance_section_id;type:int(11);not null" json:"acceptance_section_id"`   // 受理部门
	AcceptancePassportID int    `gorm:"column:acceptance_passport_id;type:int(11);not null" json:"acceptance_passport_id"` // 受理人
	CateID               int8   `gorm:"column:cate_id;type:tinyint(2);not null" json:"cate_id"`                            // 工单分类
	Username             string `gorm:"column:username;type:varchar(64);not null" json:"username"`                         // 用户姓名
	Mobile               string `gorm:"column:mobile;type:char(11);not null" json:"mobile"`                                // 用户手机
	Status               int8   `gorm:"column:status;type:tinyint(1);not null" json:"status"`                              // 0 受理中  1已处理  2已完结
	CreateAt             int    `gorm:"column:create_at;type:int(11);not null" json:"create_at"`
	IsDel                int8   `gorm:"column:is_del;type:tinyint(1);not null" json:"is_del"`
	CustomerID           int    `gorm:"column:customer_id;type:int(11);not null" json:"customer_id"` // 发起人ID
	UserID               int    `gorm:"column:user_id;type:int(11);not null" json:"user_id"`         // 用户ID
}

// GbCtCode 编码生成记录
type GbCtCode struct {
	ID     int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Nums   int    `gorm:"column:nums;type:int(11)" json:"nums"`
	Module string `gorm:"column:module;type:varchar(20);not null" json:"module"`
}

// GbMerchantAccountRefundDetails 商家账户退款流水
type GbMerchantAccountRefundDetails struct {
	ID               int     `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	WorksheetCode    string  `gorm:"column:worksheet_code;type:char(32);not null" json:"worksheet_code"`       // 工单编码
	BalanceAction    float64 `gorm:"column:balance_action;type:decimal(10,2);not null" json:"balance_action"`  // 退款金额
	ActionTepy       int8    `gorm:"column:action_tepy;type:tinyint(2);not null" json:"action_tepy"`           // 操作类型：1,订单取消，2，
	OrganizationCode string  `gorm:"column:organization_code;type:char(32);not null" json:"organization_code"` // 商家编码
	UId              int     `gorm:"column:uid;type:int(11) unsigned;not null" json:"uid"`                     // 用户ID
	AdminID          int     `gorm:"column:admin_id;type:int(11) unsigned;not null" json:"admin_id"`           // 平台、商家操作员ID
	Message          string  `gorm:"column:message;type:varchar(255)" json:"message"`
	Status           int8    `gorm:"column:status;type:tinyint(2);not null" json:"status"`             // 状态：1,正常；2，关闭
	Creattime        int     `gorm:"column:creattime;type:int(10) unsigned;not null" json:"creattime"` // 创建时间
	Updatatime       int     `gorm:"column:updatatime;type:int(10)" json:"updatatime"`                 // 更新时间
}

// GbOrderAddressToorders 订单地址表（存储用户下单时选择的地址详情）
type GbOrderAddressToorders struct {
	ID         int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	UId        int    `gorm:"column:uid;type:int(11) unsigned;not null" json:"uid"`  // 购买者ID
	Province   string `gorm:"column:province;type:varchar(18)" json:"province"`      // 省
	City       string `gorm:"column:city;type:varchar(18)" json:"city"`              // 市
	Area       string `gorm:"column:area;type:varchar(18)" json:"area"`              // 区县
	Address    string `gorm:"column:address;type:varchar(150)" json:"address"`       // 街道详细地址
	Recipients string `gorm:"column:recipients;type:varchar(255)" json:"recipients"` // 收件人
	Tel        string `gorm:"column:tel;type:varchar(255)" json:"tel"`               // 收件人电话
	Postcodes  string `gorm:"column:postcodes;type:varchar(6)" json:"postcodes"`     // 邮编
	Creattime  int    `gorm:"column:creattime;type:int(10)" json:"creattime"`        // 创建时间
}

// GbBrankRefundDetails 平台退费记录表
type GbBrankRefundDetails struct {
	ID              int     `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	WorksheetCode   string  `gorm:"column:worksheet_code;type:char(32);not null" json:"worksheet_code"`       // 工单编码
	IncrementID     string  `gorm:"column:increment_id;type:char(32)" json:"increment_id"`                    // 订单流水号
	RequestID       string  `gorm:"column:request_id;type:char(32)" json:"request_id"`                        // 退款记录查询编码
	Balance         float64 `gorm:"column:balance;type:decimal(10,2);not null" json:"balance"`                // 操作金额动作
	UId             int     `gorm:"column:uid;type:int(10)" json:"uid"`                                       // 用户ID
	UserAccountInfo string  `gorm:"column:user_account_info;type:char(32);not null" json:"user_account_info"` // 用户账户信息
	Params          string  `gorm:"column:params;type:varchar(100);not null" json:"params"`                   // 请求参数（josn）
	BrankCallbackID string  `gorm:"column:brank_callback_id;type:char(64)" json:"brank_callback_id"`          // 银行回调查询编号
	Status          int8    `gorm:"column:status;type:tinyint(2);not null" json:"status"`                     // 状态：1：退款新记录，2：退款记录已经向支付系统提交了请求；3：成功退款；4：退款错误需要人工介入
	Creattime       int     `gorm:"column:creattime;type:int(10) unsigned;not null" json:"creattime"`         // 创建时间
	Updatatime      int     `gorm:"column:updatatime;type:int(10)" json:"updatatime"`                         // 更新时间
	Countsub        int     `gorm:"column:countsub;type:int(10);not null" json:"countsub"`                    // 提交次数
	CallbackMsg     string  `gorm:"column:callbackMsg;type:varchar(100)" json:"callbackMsg"`                  // 支付接口回复的错误信息
	Jifen           float64 `gorm:"column:jifen;type:decimal(10,2)" json:"jifen"`                             // 退积分
}

// GbCrmPassport CRM系统用户表
type GbCrmPassport struct {
	ID            int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	CompanyID     int    `gorm:"column:company_id;type:int(11) unsigned" json:"company_id"`   // 所属公司 关联gb_crm_role表
	MerchantID    int    `gorm:"column:merchant_id;type:int(11)" json:"merchant_id"`          // 商户ID
	SectionID     int    `gorm:"column:section_id;type:int(11) unsigned" json:"section_id"`   // 所属部门 关联gb_crm_role表
	PositionID    int    `gorm:"column:position_id;type:int(11) unsigned" json:"position_id"` // 所属职位 关联gb_crm_role表
	Username      string `gorm:"column:username;type:char(60);not null" json:"username"`      // 登录用户名
	Password      string `gorm:"column:password;type:char(60);not null" json:"password"`      // 用户密码
	HeadImg       string `gorm:"column:head_img;type:varchar(255)" json:"head_img"`           // 用户头像
	Mail          string `gorm:"column:mail;type:char(60)" json:"mail"`                       // 邮箱地址
	Role          int8   `gorm:"column:role;type:tinyint(4)" json:"role"`
	Status        int8   `gorm:"column:status;type:tinyint(2) unsigned;not null" json:"status"`       // 0禁用 1开启
	IsDel         int8   `gorm:"column:is_del;type:tinyint(2) unsigned;not null" json:"is_del"`       // 1已删除 0未删除
	CreateUId     int    `gorm:"column:create_uid;type:int(11) unsigned" json:"create_uid"`           // 创建人
	LastLoginTime int    `gorm:"column:last_login_time;type:int(10) unsigned" json:"last_login_time"` // 最后登陆时间
	Ctime         int    `gorm:"column:ctime;type:int(10) unsigned" json:"ctime"`                     // 创建时间
	Menu          string `gorm:"column:menu;type:text" json:"menu"`                                   // 能见菜单
	Remark        string `gorm:"column:remark;type:varchar(255)" json:"remark"`                       // 备注信息
	Nickname      string `gorm:"column:nickname;type:varchar(60);not null" json:"nickname"`           // 姓名
	Mobile        string `gorm:"column:mobile;type:char(11);not null" json:"mobile"`
	IsCustomer    int8   `gorm:"column:is_customer;type:tinyint(2)" json:"is_customer"` // 0不是   1售前  2售后
}

// GbMerchantPurchase 商家采购订单信息表
type GbMerchantPurchase struct {
	ID                      int     `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	PurchaseIncrementID     string  `gorm:"column:purchase_increment_id;type:char(20)" json:"purchase_increment_id"`    // 商家采购流水号（YYMMDD+）
	Code                    string  `gorm:"column:code;type:char(100)" json:"code"`                                     // 商家采购订单号
	MerchantCode            string  `gorm:"column:merchant_code;type:char(20);not null" json:"merchant_code"`           // 商户编码
	Freight                 float64 `gorm:"column:freight;type:decimal(10,2) unsigned;not null" json:"freight"`         // 运费总价
	GoodsPrice              float64 `gorm:"column:goods_price;type:decimal(10,2) unsigned;not null" json:"goods_price"` // 运费总价
	TotalPrice              float64 `gorm:"column:total_price;type:decimal(10,2) unsigned;not null" json:"total_price"` // 订单总价（商品总价+运费总价）
	Status                  int8    `gorm:"column:status;type:tinyint(2);not null" json:"status"`                       // 订单状态： 待付款 1； 已付款  2；平台已采购 3
	IsConfirm               int8    `gorm:"column:is_confirm;type:tinyint(2) unsigned" json:"is_confirm"`               // 客服确认 0未确认 1确认
	Creattime               int     `gorm:"column:creattime;type:int(10) unsigned;not null" json:"creattime"`           // 创建时间
	UpdateAt                int     `gorm:"column:update_at;type:int(10)" json:"update_at"`
	IsDel                   int8    `gorm:"column:is_del;type:tinyint(1)" json:"is_del"`
	PlatPurchaseIncrementID string  `gorm:"column:plat_purchase_increment_id;type:char(20)" json:"plat_purchase_increment_id"` // 平台采购流水
	PlatPurchaseID          int     `gorm:"column:plat_purchase_id;type:int(10)" json:"plat_purchase_id"`                      // 平台采购ID
	UId                     int     `gorm:"column:uid;type:int(10)" json:"uid"`                                                // 操作人ID
}

// GbProBrand 产品品牌表
type GbProBrand struct {
	ID      int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"` // 品牌表
	Name    string `gorm:"column:name;type:varchar(60);not null" json:"name"`             // 品牌名称
	Char    string `gorm:"column:char;type:varchar(20)" json:"char"`                      // 品牌首字母
	Number  int    `gorm:"column:number;type:int(11);not null" json:"number"`             // 编号
	Tag     string `gorm:"column:tag;type:varchar(255)" json:"tag"`                       // 标签
	Country string `gorm:"column:country;type:char(40)" json:"country"`                   // 国家
	Logo    string `gorm:"column:logo;type:varchar(255)" json:"logo"`                     // 品牌LOGO
	Pic     string `gorm:"column:pic;type:varchar(255)" json:"pic"`                       // 品牌专区大图
	Story   string `gorm:"column:story;type:text" json:"story"`                           // 品牌故事
	Status  int8   `gorm:"column:status;type:tinyint(1);not null" json:"status"`          // 状态 1：正常  0：删除
	IsDel   int8   `gorm:"column:is_del;type:tinyint(1);not null" json:"is_del"`          // 是否删除 (0:未删 1:已删)
	Sort    int    `gorm:"column:sort;type:int(11) unsigned;not null" json:"sort"`        // 排序
	Ctime   int64  `gorm:"column:ctime;type:bigint(20) unsigned;not null" json:"ctime"`   // 创建时间
}

// GbRefundGoods 商品退款表
type GbRefundGoods struct {
	ID                   int       `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	Worksheet            string    `gorm:"column:worksheet;type:char(32);not null" json:"worksheet"`                               // 服务工单号
	IncrementID          string    `gorm:"column:increment_id;type:char(40)" json:"increment_id"`                                  // 订单流水号
	ChildrenIncrementID  string    `gorm:"index;column:children_increment_id;type:char(32);not null" json:"children_increment_id"` // 子订单流水号
	Cause                int8      `gorm:"column:cause;type:tinyint(2);not null" json:"cause"`                                     // 原因
	Recipients           string    `gorm:"column:recipients;type:varchar(10)" json:"recipients"`                                   // 联系人
	Tel                  string    `gorm:"column:tel;type:varchar(11)" json:"tel"`                                                 // 电话
	Description          string    `gorm:"column:description;type:varchar(255)" json:"description"`                                // 问题描述
	Voucher              string    `gorm:"column:voucher;type:varchar(255)" json:"voucher"`                                        // 图片凭证，多图用，号隔开
	MerchantAcceptance   int8      `gorm:"column:merchant_acceptance;type:tinyint(2);not null" json:"merchant_acceptance"`         // 商家受理情况，1受理，2不受理
	MerchantReason       string    `gorm:"column:merchant_reason;type:varchar(255)" json:"merchant_reason"`                        // 商家理由
	PlatformAcceptance   int8      `gorm:"column:platform_acceptance;type:tinyint(2);not null" json:"platform_acceptance"`         // 平台受理情况，1受理，2不受理
	PlatformReason       string    `gorm:"column:platform_reason;type:varchar(255)" json:"platform_reason"`                        // 平台理由
	SupAcceptance        int8      `gorm:"column:sup_acceptance;type:tinyint(2);not null" json:"sup_acceptance"`                   // 供应商受理情况，1受理，2不受理
	SupReason            string    `gorm:"column:sup_reason;type:varchar(255)" json:"sup_reason"`                                  // 供应商理由
	FreightNumber        string    `gorm:"column:freight_number;type:char(30)" json:"freight_number"`                              // 物流单号
	FreightTime          int       `gorm:"column:freight_time;type:int(10) unsigned" json:"freight_time"`                          // 发货时间
	LogisticsID          int       `gorm:"column:logistics_id;type:int(10)" json:"logistics_id"`                                   // 快递公司ID
	FreightTel           string    `gorm:"column:freight_tel;type:varchar(11)" json:"freight_tel"`                                 // 物流联系人电话
	Status               int8      `gorm:"column:status;type:tinyint(2);not null" json:"status"`                                   // 状态：-1,商家端取消 -2 平台端取消，-3供应商取消，-4用户取消；0,不受理驳回，1发起申请（等待商家处理），2等待平台受理，3等待供应商受理，4等待用户填入快递单号，5等待供应商接收商品，6待拆损定价，7'待供应商打款，8待平台打款，9待商家打款，10退款（用户收货）完成,13待供应商重新发货 ,11待用户收货
	Creattime            int       `gorm:"column:creattime;type:int(10) unsigned;not null" json:"creattime"`                       // 创建时间
	Updatatime           int       `gorm:"column:updatatime;type:int(10)" json:"updatatime"`                                       // 更新时间
	Zhejia               float64   `gorm:"column:zhejia;type:decimal(10,3)" json:"zhejia"`                                         // 折价比 80  代表80%
	MerchantReasonZhejia string    `gorm:"column:merchant_reason_zhejia;type:varchar(100)" json:"merchant_reason_zhejia"`          // 折价留言
	SupplierCode         string    `gorm:"column:supplier_code;type:char(32)" json:"supplier_code"`                                // 供应商编号
	SupplierID           int       `gorm:"column:supplier_id;type:int(10)" json:"supplier_id"`                                     // 供应商ID
	SupShitui            float64   `gorm:"column:sup_shitui;type:decimal(10,2)" json:"sup_shitui"`                                 // 供应商实际退款
	PlatformShitui       float64   `gorm:"column:platform_shitui;type:decimal(10,2)" json:"platform_shitui"`                       // 平台实际退款
	MerchantCode         string    `gorm:"column:merchant_code;type:char(32)" json:"merchant_code"`                                // 商家编号
	MerchantShitui       float64   `gorm:"column:merchant_shitui;type:decimal(10,0)" json:"merchant_shitui"`                       // 商家实际退款
	SupFreightNumber     []byte    `gorm:"column:sup_freight_number;type:varbinary(32)" json:"sup_freight_number"`                 // 供应商货运单
	SupFreightTime       int       `gorm:"column:sup_freight_time;type:int(10)" json:"sup_freight_time"`                           // 供货商发货时间
	SupLogisticsID       int       `gorm:"column:sup_logistics_id;type:int(10)" json:"sup_logistics_id"`                           // 供货商发货快点公司ID
	Tepy                 int8      `gorm:"column:tepy;type:tinyint(2)" json:"tepy"`                                                // 售后类别 1、退款  2、换货
	UId                  int       `gorm:"column:uid;type:int(10)" json:"uid"`                                                     // 用户ID
	StoreCode            string    `gorm:"column:store_code;type:char(32)" json:"store_code"`                                      // 店铺code
	ThirdWorkSheet       string    `gorm:"column:third_work_sheet;type:char(255)" json:"third_work_sheet"`                         // 第三方服务单号（京东）
	NewThirdOrderID      string    `gorm:"column:new_third_order_id;type:char(100)" json:"new_third_order_id"`                     // 京东退货生成新的订单号
	RefundAddress        string    `gorm:"column:refund_address;type:varchar(255)" json:"refund_address"`                          // 第三方平台供应商收货地址
	Activation           int       `gorm:"column:activation;type:int(2)" json:"activation"`                                        // 是否激活
	PickwareType         int       `gorm:"column:pickwareType;type:int(2)" json:"pickwareType"`                                    // （商品返回京东方式 ）4上门取件  40客户发货 7客户送货
	UserAddress          string    `gorm:"column:user_address;type:varchar(255);not null" json:"user_address"`                     // 用户上门取件地址
	UserAddresDetail     string    `gorm:"column:user_addres_detail;type:varchar(255);not null" json:"user_addres_detail"`         // 用户上门取件详细地址
	SuAddressID          int       `gorm:"column:su_address_id;type:int(11)" json:"su_address_id"`                                 // 用户地址id
	Date                 time.Time `gorm:"column:date;type:datetime" json:"date"`                                                  // 用户取件时间
}

type GbWhProduct struct {
	ID           int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	SupProID     int    `gorm:"index;column:sup_pro_id;type:int(11);not null" json:"sup_pro_id"`     // 供应商商品ID
	StoProID     int    `gorm:"index;column:sto_pro_id;type:int(11);not null" json:"sto_pro_id"`     // 店铺商品ID
	ProName      string `gorm:"column:pro_name;type:varchar(150)" json:"pro_name"`                   // 网红自定义商品名称
	Pic          string `gorm:"column:pic;type:longtext" json:"pic"`                                 // 网红自定义商品图片
	Content      string `gorm:"column:content;type:longtext" json:"content"`                         // 网红自定义商品描述图
	UserID       int    `gorm:"column:user_id;type:int(11);not null" json:"user_id"`                 // 网红用户id
	Status       int8   `gorm:"column:status;type:tinyint(1);not null" json:"status"`                // 选品状态 -2 移出选品 -1 已下架， 0 已选品， 1 申请铺货， 2已铺货
	PlatProCode  string `gorm:"column:plat_pro_code;type:varchar(40);not null" json:"plat_pro_code"` // 编号(产品平台编号)
	PlatformCids string `gorm:"column:platform_cids;type:text" json:"platform_cids"`                 // 商品第三方类目ID集合
	Props        string `gorm:"column:props;type:text" json:"props"`                                 // 商品属性信息
	IsUpdate     int8   `gorm:"column:is_update;type:tinyint(4);not null" json:"is_update"`          // 是否修改了类目属性数据 0否 1是
	ErrorLog     string `gorm:"column:error_log;type:text" json:"error_log"`                         // 最近一次的错误日志
}

// GbLogOrderChildrenChange 子订单状态流水
type GbLogOrderChildrenChange struct {
	ID                  int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	UId                 int    `gorm:"column:uid;type:int(11) unsigned;not null" json:"uid"`                             // 操作者ID
	ChildrenIncrementID string `gorm:"column:children_increment_id;type:char(30);not null" json:"children_increment_id"` // 订单流水号
	Remarks             string `gorm:"column:remarks;type:varchar(255)" json:"remarks"`                                  // 备注
	AfterStatus         int    `gorm:"column:after_status;type:int(11);not null" json:"after_status"`                    // 处理后状态
	Creattime           int    `gorm:"column:creattime;type:int(10)" json:"creattime"`                                   // 创建时间
}

// GbOrderGroupRela 订单拼团关联表
type GbOrderGroupRela struct {
	ID                 int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`                 // id
	GroupRelaID        int       `gorm:"column:group_rela_id;type:int(11);not null" json:"group_rela_id"`      // 拼团产品关联id
	ParentGroup        int       `gorm:"column:parent_group;type:int(11)" json:"parent_group"`                 // 所属拼团id 0为团长订单
	OrderCode          string    `gorm:"column:order_code;type:varchar(255);not null" json:"order_code"`       // 订单流水号
	StoreID            int       `gorm:"column:store_id;type:int(11)" json:"store_id"`                         // 拼团所属店铺
	Status             int8      `gorm:"column:status;type:tinyint(1) unsigned zerofill" json:"status"`        // 0.待成团未付款1.待成团已付款2.已成团 3.已失败
	IsDel              int8      `gorm:"column:is_del;type:tinyint(1)" json:"is_del"`                          // 逻辑删除字段
	UserID             int       `gorm:"column:user_id;type:int(11);not null" json:"user_id"`                  // 下单用户id
	Createtime         time.Time `gorm:"column:createtime;type:datetime" json:"createtime"`                    // 订单创建时间
	MerchantStandardID int       `gorm:"column:merchant_standard_id;type:int(11)" json:"merchant_standard_id"` // 拼团产品供应商规格ID
}

// GbProCatesRela 平台商品关联的分类信息表,一个商品可关联到多个分类
type GbProCatesRela struct {
	ID      int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	ProID   int    `gorm:"index;column:pro_id;type:int(11)" json:"pro_id"`      // 平台商品ID 对应 gb_product.id
	BarCode string `gorm:"index;column:bar_code;type:char(15)" json:"bar_code"` // 商品ENA13编码
	Bcat    int    `gorm:"column:bcat;type:int(11)" json:"bcat"`                // 产品分类（三级分类）
	Mcat    int    `gorm:"column:mcat;type:int(11)" json:"mcat"`                // 产品分类（三级分类）
	Scat    int    `gorm:"column:scat;type:int(11)" json:"scat"`                // 产品分类（三级分类）
	IsDel   int8   `gorm:"column:is_del;type:tinyint(4)" json:"is_del"`         // 删除状态 0未删除 1已经删除
}

// GbSuproRealAPISupro 平台供应商产品关联第三方(接口)供应商产品
type GbSuproRealAPISupro struct {
	ID                 int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	SupID              int    `gorm:"column:sup_id;type:int(11) unsigned;not null" json:"sup_id"`                       // 平台供应商ID
	APISupID           int    `gorm:"column:api_sup_id;type:int(11) unsigned;not null" json:"api_sup_id"`               // 接口供应商类型 1京东
	SupProductCode     string `gorm:"column:sup_product_code;type:char(60);not null" json:"sup_product_code"`           // 供应商主产品编码
	SupProStandardCode string `gorm:"column:sup_pro_standard_code;type:char(60);not null" json:"sup_pro_standard_code"` // 供应商SKU子产品编码
	APISupProCode      string `gorm:"column:api_sup_pro_code;type:char(60);not null" json:"api_sup_pro_code"`           // 接口供应商产品维一编码
}

// GbUser 会员表
type GbUser struct {
	ID            int       `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"` // 会员表
	Openid        string    `gorm:"column:openid;type:varchar(255)" json:"openid"`
	Mobile        string    `gorm:"column:mobile;type:char(11);not null" json:"mobile"`          // 手机号码
	StoreID       int       `gorm:"column:store_id;type:int(11) unsigned" json:"store_id"`       // 店铺ID
	MerchantID    int       `gorm:"column:merchant_id;type:int(11) unsigned" json:"merchant_id"` // 商家ID
	Nickname      string    `gorm:"column:nickname;type:varchar(60)" json:"nickname"`            // 昵称
	Address       string    `gorm:"column:address;type:varchar(100)" json:"address"`             // 所在城市
	UUID          int       `gorm:"column:uuid;type:int(11) unsigned" json:"uuid"`               // 父级ID
	Username      string    `gorm:"column:username;type:varchar(30)" json:"username"`            // 姓名
	Realname      string    `gorm:"column:realname;type:varchar(50)" json:"realname"`            // 真实姓名 用于提现
	Headimg       string    `gorm:"column:headimg;type:varchar(255)" json:"headimg"`             // 头像
	AvatarURL     string    `gorm:"column:avatarUrl;type:varchar(255)" json:"avatarUrl"`         // 微信头像
	Gender        int8      `gorm:"column:gender;type:tinyint(1)" json:"gender"`                 // 性别 1男 2女
	Country       string    `gorm:"column:country;type:varchar(60)" json:"country"`
	Source        int8      `gorm:"column:source;type:tinyint(1)" json:"source"`                               // 注册来源 1为H5端 2为程序
	Birthday      time.Time `gorm:"column:birthday;type:date" json:"birthday"`                                 // 生日
	Ctime         int64     `gorm:"column:ctime;type:bigint(20)" json:"ctime"`                                 // 注册时间
	Level         int8      `gorm:"column:level;type:tinyint(1)" json:"level"`                                 // 等级(1:普通会员)
	Status        int8      `gorm:"column:status;type:smallint(1)" json:"status"`                              // 用户状态   1：正常   0：删除
	LastLoginIP   string    `gorm:"column:last_login_ip;type:varchar(20)" json:"last_login_ip"`                // 最后登录IP
	LastLoginTime int64     `gorm:"column:last_login_time;type:bigint(20)" json:"last_login_time"`             // 最后登录时间
	IsCertified   int8      `gorm:"column:is_certified;type:tinyint(1) unsigned;not null" json:"is_certified"` // 是否实名认证
}

type GbWhGoodsSkuMatching struct {
	ID           int     `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	Sid          int     `gorm:"unique_index:idx_uid_sid_oid_osid;column:sid;type:int(20) unsigned;not null" json:"sid"`               // 店铺ID
	UId          int64   `gorm:"unique_index:idx_uid_sid_oid_osid;column:uid;type:bigint(20) unsigned;not null" json:"uid"`            // 用户ID
	Gid          int     `gorm:"column:gid;type:int(11) unsigned;not null" json:"gid"`                                                 // 平台商品匹配表主键ID
	ProID        int     `gorm:"column:pro_id;type:int(11) unsigned;not null" json:"pro_id"`                                           // 本地匹配的商品ID
	OpNumIid     string  `gorm:"column:op_num_iid;type:varchar(50);not null" json:"op_num_iid"`                                        // 平台商品ID
	OuterSkuID   string  `gorm:"column:outer_sku_id;type:varchar(100)" json:"outer_sku_id"`                                            // 本地商品sku编码
	ProSkuID     int     `gorm:"unique_index:idx_uid_sid_oid_osid;column:pro_sku_id;type:int(11) unsigned;not null" json:"pro_sku_id"` // 本地匹配的商品规格ID
	OpSkuIid     string  `gorm:"unique_index:idx_uid_sid_oid_osid;column:op_sku_iid;type:varchar(40);not null" json:"op_sku_iid"`      // 平台商品规格ID
	SkuPrice     float64 `gorm:"column:sku_price;type:decimal(10,2)" json:"sku_price"`                                                 // 平台商品sku售价
	SkuStatus    int8    `gorm:"column:sku_status;type:tinyint(2);not null" json:"sku_status"`                                         // 平台sku状态 0-可用 1-上架 2-下架 3-删除
	Properties   string  `gorm:"column:properties;type:varchar(600)" json:"properties"`                                                // 商品sku属性，用于sku后续操作(平台的sku属性都会在这里记录)
	PlatformType string  `gorm:"column:platform_type;type:varchar(12);not null" json:"platform_type"`                                  // 平台类型 淘宝taobao 京东 jingdong
	LocalPic     string  `gorm:"column:local_pic;type:varchar(255)" json:"local_pic"`                                                  // 本地sku图片
	PlatformPic  string  `gorm:"column:platform_pic;type:varchar(255)" json:"platform_pic"`                                            // 平台商品图片
}

// GbBanks 银行卡表
type GbBanks struct {
	ID     int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	Name   string `gorm:"column:name;type:varchar(100);not null" json:"name"`   // 银行名称
	Logo   string `gorm:"column:logo;type:varchar(100);not null" json:"logo"`   // 银行LOGO
	Status int8   `gorm:"column:status;type:tinyint(1);not null" json:"status"` // 状态(0:未启用   1:启用)
	IsDel  int8   `gorm:"column:is_del;type:tinyint(1)" json:"is_del"`          // 删除状态(0:未删除  1:已删除)
	Ctime  int    `gorm:"column:ctime;type:int(11);not null" json:"ctime"`      // 创建时间
}

// GbMerchantAccountDetails 商家账户流水
type GbMerchantAccountDetails struct {
	ID            int     `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	AccountID     int     `gorm:"column:account_id;type:int(11) unsigned;not null" json:"account_id"`      // 账户ID
	WorksheetCode string  `gorm:"column:worksheet_code;type:char(32);not null" json:"worksheet_code"`      // 工单编码
	BalanceAction float64 `gorm:"column:balance_action;type:decimal(10,2);not null" json:"balance_action"` // 操作金额动作
	ActionTepy    int8    `gorm:"column:action_tepy;type:tinyint(2);not null" json:"action_tepy"`          // 操作类型：1,采购；2，提现，3，退款，4，冻结,5，解冻，6，采购，
	Status        int8    `gorm:"column:status;type:tinyint(2);not null" json:"status"`                    // 状态：1,正常；2，关闭
	Creattime     int     `gorm:"column:creattime;type:int(10) unsigned;not null" json:"creattime"`        // 创建时间
	Updatatime    int     `gorm:"column:updatatime;type:int(10)" json:"updatatime"`                        // 更新时间
}

// GbMerchantSeckillProduct 秒杀商品表（SPU）
type GbMerchantSeckillProduct struct {
	ID         int     `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	SeckillID  int     `gorm:"column:seckill_id;type:int(11)" json:"seckill_id"`         // 秒杀活动ID
	ProID      int     `gorm:"column:pro_id;type:int(11)" json:"pro_id"`                 // 产品ID(商家产品ID) gb_merchant_product.id
	ProBarCode string  `gorm:"column:pro_bar_code;type:varchar(15)" json:"pro_bar_code"` // 商品条码 ENA-13编码
	MinPrice   float64 `gorm:"column:min_price;type:decimal(10,2)" json:"min_price"`     // 最低售价
}

// GbOrderAddress 用户地址表
type GbOrderAddress struct {
	ID         int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	UId        int    `gorm:"column:uid;type:int(11) unsigned;not null" json:"uid"`   // 购买者ID
	Province   string `gorm:"column:province;type:varchar(18)" json:"province"`       // 省
	City       string `gorm:"column:city;type:varchar(18)" json:"city"`               // 市
	Area       string `gorm:"column:area;type:varchar(18)" json:"area"`               // 区县
	Address    string `gorm:"column:address;type:varchar(150)" json:"address"`        // 街道详细地址
	Recipients string `gorm:"column:recipients;type:varchar(255)" json:"recipients"`  // 收件人
	Tel        string `gorm:"column:tel;type:varchar(255)" json:"tel"`                // 收件人电话
	Postcodes  string `gorm:"column:postcodes;type:varchar(6)" json:"postcodes"`      // 邮编
	Creattime  int    `gorm:"column:creattime;type:int(10)" json:"creattime"`         // 创建时间
	Default    int8   `gorm:"column:default;type:tinyint(1);not null" json:"default"` // 1为默认地址
	Isdel      int8   `gorm:"column:isdel;type:tinyint(1)" json:"isdel"`              // 1为删除
}

// GbProCates 平台产品分类
type GbProCates struct {
	ID      int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Token   string `gorm:"index;column:token;type:varchar(20);not null" json:"token"`       // 公众号标识
	Pid     int    `gorm:"index;column:pid;type:int(11);not null" json:"pid"`               // 上级ID
	Name    string `gorm:"column:name;type:varchar(30);not null" json:"name"`               // 类别名称
	Level   int8   `gorm:"column:level;type:tinyint(1);not null" json:"level"`              // 0:一级分类  1：二级分类 2：三级分类
	Logo    string `gorm:"column:logo;type:varchar(200)" json:"logo"`                       // 图片
	Sort    int    `gorm:"column:sort;type:int(11);not null" json:"sort"`                   // 排序
	Display int8   `gorm:"column:display;type:tinyint(1) unsigned;not null" json:"display"` // 是否默认展示：0否1是
	Status  int8   `gorm:"column:status;type:tinyint(1);not null" json:"status"`            // 状态  1：可用   0：不可用  2：特价
	Showpic string `gorm:"column:showpic;type:varchar(200);not null" json:"showpic"`        // 展示图
	IsDel   int8   `gorm:"column:is_del;type:tinyint(1);not null" json:"is_del"`            // 是否删除  1：是   0：否
}

// GbLogistics 物流快递公司表
type GbLogistics struct {
	ID         int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Name       string `gorm:"column:name;type:varchar(50);not null" json:"name"`           // 物流公司名称
	API        string `gorm:"column:api;type:varchar(150)" json:"api"`                     // 物流查询API地址
	Status     int8   `gorm:"column:status;type:tinyint(1);not null" json:"status"`        // 状态1正常，0暂停使用
	IsDel      int8   `gorm:"column:is_del;type:tinyint(1);not null" json:"is_del"`        // 是否删除  1：是  0：否
	CreateTime int    `gorm:"column:create_time;type:int(10);not null" json:"create_time"` // 创建时间
	UpdateTime int    `gorm:"column:update_time;type:int(10) unsigned" json:"update_time"` // 更新时间
}

// GbOrder 订单信息表
type GbOrder struct {
	ID                 int     `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	Code               string  `gorm:"column:code;type:char(100);not null" json:"code"`                            // 订单号
	UId                int     `gorm:"index;column:uid;type:int(11) unsigned;not null" json:"uid"`                 // 用户ID
	AddressID          int     `gorm:"column:address_id;type:int(11) unsigned;not null" json:"address_id"`         // 收货信息
	MerchantCode       string  `gorm:"column:merchant_code;type:char(20);not null" json:"merchant_code"`           // 商户编码
	StoreCode          string  `gorm:"column:store_code;type:char(20);not null" json:"store_code"`                 // 店铺CODE
	Freight            float64 `gorm:"column:freight;type:decimal(10,2) unsigned;not null" json:"freight"`         // 运费总价
	GoodsPrice         float64 `gorm:"column:goods_price;type:decimal(10,2) unsigned;not null" json:"goods_price"` // 运费总价
	TotalPrice         float64 `gorm:"column:total_price;type:decimal(10,2) unsigned;not null" json:"total_price"` // 订单总价（商品总价+运费总价）
	Buytype            int     `gorm:"column:buytype;type:int(2);not null" json:"buytype"`                         // 0普通订单，1秒杀，2团购，3子店铺订单,5壹团，6国寿
	Message            string  `gorm:"column:message;type:varchar(255)" json:"message"`
	Status             int8    `gorm:"column:status;type:tinyint(2);not null" json:"status"`             // 订单状态： -1：已取消；1:待付款；2:待商家采购； 3:待平台采购；4:待发货；12：已发货；5:已完成；6：待供应商退款；7：待平台退款；8：待商家退款；9：已关闭；10：待成,11:等待银行退款; 12:已发货
	IsConfirm          int8    `gorm:"column:is_confirm;type:tinyint(2) unsigned" json:"is_confirm"`     // 客服确认 0未确认 1确认
	Creattime          int     `gorm:"column:creattime;type:int(10) unsigned;not null" json:"creattime"` // 创建时间
	Upstatustime       int     `gorm:"column:upstatustime;type:int(10)" json:"upstatustime"`             // 订单状态更改时间
	IncrementID        string  `gorm:"index;column:increment_id;type:char(48)" json:"increment_id"`      // 订单流水号（YYMMDD+）
	UpdateAt           int     `gorm:"column:update_at;type:int(10)" json:"update_at"`
	PurchaserID        string  `gorm:"column:purchaser_id;type:char(20)" json:"purchaser_id"`                      // 商家采购流水号
	CountOriginalPrice float64 `gorm:"column:count_original_price;type:decimal(10,2)" json:"count_original_price"` // 原始售价总额
	CountMBatchPrice   float64 `gorm:"column:count_m_batch_price;type:decimal(10,2)" json:"count_m_batch_price"`   // 平台供货总额
	OriginalFreight    float64 `gorm:"column:original_freight;type:decimal(10,2)" json:"original_freight"`         // 原始运费总额
	Updatatime         int     `gorm:"column:updatatime;type:int(10)" json:"updatatime"`                           // 更新时间
	UserStoreCode      string  `gorm:"column:user_store_code;type:char(20)" json:"user_store_code"`                // 导购小店CODE
	PtRemark           string  `gorm:"column:pt_remark;type:varchar(255)" json:"pt_remark"`                        // 平台备注
	MsgID              int8    `gorm:"column:msg_id;type:tinyint(2)" json:"msg_id"`                                // 用户取消的消息ID
	ActivityID         int     `gorm:"column:activity_id;type:int(11) unsigned;not null" json:"activity_id"`       // 活动关联id
	IsDel              int8    `gorm:"column:is_del;type:tinyint(2)" json:"is_del"`                                // 是否已经删除0未删除，1已删除
	YigroupProID       int     `gorm:"column:yigroup_pro_id;type:int(11)" json:"yigroup_pro_id"`                   // 团购产品ID 用于小程序再次购买或通过订单查看团品
	Allocation         int8    `gorm:"column:allocation;type:tinyint(4);not null" json:"allocation"`               // 0为默认，1表示可分配，
	Type               int8    `gorm:"column:type;type:tinyint(1);not null" json:"type"`                           // 订单类型  1：实物订单   0：虚拟订单
	CostType           int8    `gorm:"column:cost_type;type:tinyint(1)" json:"cost_type"`                          // 充值方式（1直充，2卡密,3兑换码）
	AccountMessage     string  `gorm:"column:account_message;type:varchar(50)" json:"account_message"`             // 充值账号
	Ratio              int     `gorm:"column:ratio;type:int(10)" json:"ratio"`                                     // 积分与人民币的比值，大于0的正整数才有效
	IsPay              int8    `gorm:"column:is_pay;type:tinyint(3);not null" json:"is_pay"`                       // 判断用户是否付过款，1表示付过，0表示没有
	ThirdParty         int8    `gorm:"column:third_party;type:tinyint(3)" json:"third_party"`                      // 0代表自营订单，1代表是国寿订单
	WhTradeCode        string  `gorm:"column:wh_trade_code;type:varchar(50)" json:"wh_trade_code"`                 // 网红第三方平方订单号(陈磊添加)
	WhTradeType        int8    `gorm:"column:wh_trade_type;type:tinyint(3)" json:"wh_trade_type"`                  // 网红第三方平台类型(陈磊添加) 1 淘宝
	GoodsSource        int8    `gorm:"column:goods_source;type:tinyint(2)" json:"goods_source"`                    // 订单商品来源，0代表境内，1代表跨境
}

// GbStoreProduct 商铺-产品表
type GbStoreProduct struct {
	ID           int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	SellerID     int       `gorm:"index;column:seller_id;type:int(11);not null" json:"seller_id"`                         // 商家ID
	StoreID      int       `gorm:"index;index:index_supplier_type;column:store_id;type:int(11);not null" json:"store_id"` // 商铺id
	MpCode       string    `gorm:"index;column:mp_code;type:varchar(15);not null" json:"mp_code"`                         // 商家产品编号
	BrandID      int       `gorm:"column:brand_id;type:int(11)" json:"brand_id"`                                          // 品牌id
	BrandName    string    `gorm:"column:brand_name;type:varchar(255)" json:"brand_name"`                                 // 产品品牌名称
	PlatProCode  string    `gorm:"index;column:plat_pro_code;type:varchar(100);not null" json:"plat_pro_code"`            // 平台产品编号
	PlatProID    int       `gorm:"column:plat_pro_id;type:int(11)" json:"plat_pro_id"`                                    // 平台产品库id
	BarCode      string    `gorm:"index;column:bar_code;type:varchar(40);not null" json:"bar_code"`                       // 商品条码
	Cateids      string    `gorm:"column:cateids;type:varchar(255)" json:"cateids"`                                       // 店铺商品分类
	Bcat         int       `gorm:"column:bcat;type:int(10)" json:"bcat"`                                                  // 一级分类
	Mcat         int       `gorm:"column:mcat;type:int(10)" json:"mcat"`                                                  // 二级分类
	Scat         int       `gorm:"column:scat;type:int(10)" json:"scat"`                                                  // 三级分类
	Sign         string    `gorm:"column:sign;type:varchar(50)" json:"sign"`                                              // 商品标签
	ProName      string    `gorm:"column:pro_name;type:varchar(200)" json:"pro_name"`
	IsDel        int8      `gorm:"index:index_supplier_type;column:is_del;type:tinyint(1)" json:"is_del"`                        // 删除状态 1删除，0未删除
	Sort         int       `gorm:"column:sort;type:int(11);not null" json:"sort"`                                                // 商品排序 越小越靠前
	Status       int8      `gorm:"index;index:index_supplier_type;column:status;type:tinyint(1);not null" json:"status"`         // 店铺状态 1 上架，0 下架
	Lftime       int       `gorm:"column:lftime;type:int(11);not null" json:"lftime"`                                            // 下架时间
	Ctime        int       `gorm:"column:ctime;type:int(11);not null" json:"ctime"`                                              // 创建时间
	YID          int       `gorm:"column:y_id;type:int(11)" json:"y_id"`                                                         // 导入的记录的之前的店铺商品ID
	StoreBrandID int       `gorm:"column:store_brand_id;type:int(11) unsigned" json:"store_brand_id"`                            // 店铺品牌ID
	GoodsSource  int8      `gorm:"column:goods_source;type:tinyint(1) unsigned;not null" json:"goods_source"`                    // 商品来源[0 境内 1 境外]
	SupplierType int8      `gorm:"index:index_supplier_type;column:supplier_type;type:tinyint(2);not null" json:"supplier_type"` // 0自营1代表京东 2代表爱库存  3全球仓  4网红（10是没有供应商）
	GbProduct    GbProduct `gorm:"ForeignKey:Code;AssociationForeignKey:PlaProCode"`
}

// GbRetail 零售店(分销店   导购店)  注册登录即成为分销者
type GbRetail struct {
	ID              int     `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	StoreID         int     `gorm:"index;column:store_id;type:int(11);not null" json:"store_id"`                 // 店铺ID
	UserID          int     `gorm:"column:user_id;type:int(11);not null" json:"user_id"`                         // 用户ID
	Name            string  `gorm:"column:name;type:varchar(60);not null" json:"name"`                           // 店铺名称
	Logo            string  `gorm:"column:logo;type:varchar(100);not null" json:"logo"`                          // 店铺头像
	Cover           string  `gorm:"column:cover;type:varchar(100)" json:"cover"`                                 // 店招图片
	Desc            string  `gorm:"column:desc;type:varchar(255);not null" json:"desc"`                          // 店铺描述
	Color           string  `gorm:"column:color;type:varchar(100);not null" json:"color"`                        // 店铺配色
	Realname        string  `gorm:"column:realname;type:varchar(50)" json:"realname"`                            // 真实姓名
	IDentityCard    string  `gorm:"column:identity_card;type:varchar(100)" json:"identity_card"`                 // 身份证号码
	NegativePic     string  `gorm:"column:negative_pic;type:varchar(100)" json:"negative_pic"`                   // 身份证反面照
	PositivePic     string  `gorm:"column:positive_pic;type:varchar(100)" json:"positive_pic"`                   // 身份证正面照
	LockingIncome   float64 `gorm:"column:locking_income;type:decimal(10,2);not null" json:"locking_income"`     // 待到账金额(累计已申请提现-待审核)
	FinishedIncome  float64 `gorm:"column:finished_income;type:decimal(10,2);not null" json:"finished_income"`   // 累计提现金额(累计已到账金额)
	AvailableIncome float64 `gorm:"column:available_income;type:decimal(10,2);not null" json:"available_income"` // 可提现佣金(累计可提现金额)
	Ctime           int64   `gorm:"column:ctime;type:bigint(20);not null" json:"ctime"`                          // 开店时间
	SetMode         int8    `gorm:"column:set_mode;type:smallint(1)" json:"set_mode"`                            // 默认结算方式(1:银行卡   2:支付宝  3:微信)
	SetAccount      string  `gorm:"column:set_account;type:varchar(60)" json:"set_account"`                      // 默认结算账户(银行卡号/支付宝账号/微信账号)
}

// GbStoreAdvert 店铺-轮播图表
type GbStoreAdvert struct {
	ID       int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	StoreID  int    `gorm:"column:store_id;type:int(11) unsigned;not null" json:"store_id"` // 店铺ID
	Pic      string `gorm:"column:pic;type:varchar(100);not null" json:"pic"`               // 广告图
	URL      string `gorm:"column:url;type:varchar(100);not null" json:"url"`               // 链接
	Sort     int    `gorm:"column:sort;type:int(11);not null" json:"sort"`                  // 排序
	Position string `gorm:"column:position;type:varchar(255);not null" json:"position"`     // 轮播位置
	IsDel    int8   `gorm:"column:is_del;type:tinyint(1);not null" json:"is_del"`           // 1删除，0未删除
	Status   int8   `gorm:"column:status;type:tinyint(1);not null" json:"status"`           // 状态 1 开启，0 关闭
	Ctime    int    `gorm:"column:ctime;type:int(11);not null" json:"ctime"`                // 创建时间
}

// GbAdmin 管理员
type GbAdmin struct {
	ID            int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	SectionID     int    `gorm:"column:section_id;type:int(11) unsigned;not null" json:"section_id"` // 角色ID
	Username      string `gorm:"index;column:username;type:varchar(20);not null" json:"username"`    // 用户名
	Password      string `gorm:"column:password;type:char(40);not null" json:"password"`             // 密码
	Name          string `gorm:"column:name;type:varchar(255);not null" json:"name"`                 // 姓名
	Mail          string `gorm:"column:mail;type:varchar(255)" json:"mail"`                          // 邮箱
	Salt          string `gorm:"column:salt;type:varchar(10);not null" json:"salt"`                  // 混淆码
	Telphone      string `gorm:"column:telphone;type:varchar(20)" json:"telphone"`                   // 联系电话
	Mobile        string `gorm:"column:mobile;type:char(11)" json:"mobile"`                          // 手机号
	Sex           int8   `gorm:"column:sex;type:tinyint(1)" json:"sex"`                              // 1、男性   2、女性
	LastLoginTime int    `gorm:"column:last_login_time;type:int(11)" json:"last_login_time"`         // 上次登陆时间
	LastLoginIP   string `gorm:"column:last_login_ip;type:varchar(15)" json:"last_login_ip"`         // 上次登陆IP
	CreateTime    int    `gorm:"column:create_time;type:int(11)" json:"create_time"`                 // 创建时间
	UpdatedAt     int    `gorm:"column:updated_at;type:int(11)" json:"updated_at"`                   // 更改时间
	Status        int8   `gorm:"column:status;type:tinyint(1) unsigned" json:"status"`               // 状态  0：不可用  1：可用
	IsDel         int8   `gorm:"column:is_del;type:smallint(2) unsigned;not null" json:"is_del"`     // 删除状态   1表示已删除
	Remark        string `gorm:"column:remark;type:text" json:"remark"`                              // 备注信息
}

type GbAdminRoleUser struct {
	ID        int  `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`      // ID
	AdminID   int  `gorm:"column:admin_id;type:int(11);not null" json:"admin_id"`              // 用户ID
	RoleID    int  `gorm:"column:role_id;type:int(11);not null" json:"role_id"`                // 角色ID
	Module    int8 `gorm:"column:module;type:tinyint(1);not null" json:"module"`               // 后台模块 1平台端 2商家端 3店铺端 4供应商
	Status    int8 `gorm:"column:status;type:tinyint(1);not null" json:"status"`               // 状态：1有效、0无效
	CreatedAt int  `gorm:"column:created_at;type:int(10) unsigned;not null" json:"created_at"` // 创建时间
	UpdatedAt int  `gorm:"column:updated_at;type:int(10) unsigned;not null" json:"updated_at"` // 更新时间
}

type GbAPIInsideLog struct {
	ID        int    `gorm:"primary_key;column:id;type:int(10) unsigned;not null" json:"-"`
	RequestID string `gorm:"column:request_id;type:char(36)" json:"request_id"`
	Requests  string `gorm:"column:requests;type:text;not null" json:"requests"`
	Response  string `gorm:"column:response;type:text;not null" json:"response"`
	Timestamp int    `gorm:"column:timestamp;type:int(11);not null" json:"timestamp"`
	Type      int8   `gorm:"column:type;type:tinyint(3);not null" json:"type"` // 1 下单接口  2取消订单  3搜索订单
	CreatedAt int    `gorm:"column:created_at;type:int(11);not null" json:"created_at"`
}

// GbMerchantPercentConfigure 爱库存售价配置百分比
type GbMerchantPercentConfigure struct {
	ID    int    `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Value int    `gorm:"column:value;type:int(11)" json:"value"`    // 商家售价百分比
	Name  string `gorm:"column:name;type:varchar(100)" json:"name"` // 商家默认配置名称
}

// GbMerchantProduct 商家-产品表
type GbMerchantProduct struct {
	ID              int     `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	SellerID        int     `gorm:"column:seller_id;type:int(11);not null" json:"seller_id"`                                  // 商家ID
	Code            string  `gorm:"index;index:plat_pro_code;column:code;type:varchar(50);not null" json:"code"`              // 商家产品编号
	PlatProCode     string  `gorm:"index:plat_pro_code;column:plat_pro_code;type:varchar(100);not null" json:"plat_pro_code"` // 平台产品编号
	BarCode         string  `gorm:"index;column:bar_code;type:varchar(40);not null" json:"bar_code"`                          // 商品条码
	ProName         string  `gorm:"index;column:pro_name;type:varchar(150);not null" json:"pro_name"`                         // 产品名称
	SellPoint       string  `gorm:"column:sell_point;type:varchar(200);not null" json:"sell_point"`                           // 商品卖点
	Type            int8    `gorm:"column:type;type:tinyint(1);not null" json:"type"`                                         // 类型  1：使用商户库存   0：使用平台库存
	Tag             string  `gorm:"column:tag;type:varchar(100);not null" json:"tag"`                                         // 商品标签
	Pic             string  `gorm:"column:pic;type:longtext" json:"pic"`                                                      // 商品图片
	Content         string  `gorm:"column:content;type:text" json:"content"`                                                  // 描述
	Price           float64 `gorm:"column:price;type:decimal(20,2);not null" json:"price"`                                    // 售价，只作显示使用，不具有实际意义
	IsDel           int8    `gorm:"column:is_del;type:tinyint(1);not null" json:"is_del"`                                     // 删除状态 1删除，0未删除
	Status          int8    `gorm:"column:status;type:tinyint(1);not null" json:"status"`                                     // 商铺产品状态：1 上架，0 下架
	ProChangeStatus int8    `gorm:"column:pro_change_status;type:tinyint(2) unsigned;not null" json:"pro_change_status"`      // 产品变更状态，当供应商有新SKU入到平台库时，需要设置该字段为1，默认为0
	Ctime           int     `gorm:"column:ctime;type:int(11);not null" json:"ctime"`                                          // 创建时间
	JdChangeStatus  int8    `gorm:"column:jd_change_status;type:tinyint(1)" json:"jd_change_status"`                          // 京东价格是否发生变化 1没有变化 2变化
	SourceProID     int     `gorm:"column:source_pro_id;type:int(11)" json:"source_pro_id"`                                   // 源商品ID
	GoodsSource     int8    `gorm:"column:goods_source;type:tinyint(1) unsigned;not null" json:"goods_source"`                // 商品来源[0 境内 1 境外]
	GlobalSendType  int8    `gorm:"column:global_send_type;type:tinyint(1) unsigned;not null" json:"global_send_type"`        // 发货类型[1为保税，2直邮，3为国内]
}

type GbMerchantYigroupPromotion struct {
	ID            int  `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Type          int8 `gorm:"column:type;type:tinyint(2) unsigned;not null" json:"type"`                    // 促销类型  1：9.9元包邮、2团长热卖、3网红爆品、4京东好货
	MerchantSpuID int  `gorm:"column:merchant_spu_id;type:int(11) unsigned;not null" json:"merchant_spu_id"` // 商家产品ID
	MerchantSkuID int  `gorm:"column:merchant_sku_id;type:int(11) unsigned;not null" json:"merchant_sku_id"` // 商家规格ID
	IsTop         int8 `gorm:"column:is_top;type:tinyint(2) unsigned;not null" json:"is_top"`                // 置顶在首页的状态  1置顶 0未置顶
	IsDel         int8 `gorm:"column:is_del;type:tinyint(1);not null" json:"is_del"`                         // 1删除，0未删除
	Sort          int  `gorm:"column:sort;type:int(11);not null" json:"sort"`                                // 排序
	CreatedAt     int  `gorm:"column:created_at;type:int(10) unsigned;not null" json:"created_at"`
	UpdatedAt     int  `gorm:"column:updated_at;type:int(10) unsigned;not null" json:"updated_at"`
}

// GbStoreUserStore 用户店铺
type GbStoreUserStore struct {
	ID        int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	UId       int    `gorm:"column:uid;type:int(11) unsigned;not null" json:"uid"`                  // UID
	StoreID   int    `gorm:"column:store_id;type:int(11) unsigned;not null" json:"store_id"`        // 店铺ID   gb_store.id
	Code      string `gorm:"column:code;type:char(80);not null" json:"code"`                        // 店铺code 本表唯一  store_id+uid
	Title     string `gorm:"column:title;type:varchar(255);not null" json:"title"`                  // 店铺名称
	Logo      string `gorm:"column:logo;type:varchar(255);not null" json:"logo"`                    // 店铺LOGO
	Banner    string `gorm:"column:banner;type:varchar(255);not null" json:"banner"`                // 店铺店招
	BannerDef int8   `gorm:"column:banner_def;type:tinyint(2) unsigned;not null" json:"banner_def"` // 店铺默认店招ID
	Info      string `gorm:"column:info;type:varchar(255);not null" json:"info"`                    // 店铺描述
	Ctime     int    `gorm:"column:ctime;type:int(11) unsigned;not null" json:"ctime"`              // 创建时间
}

type GbCrmSessions struct {
	ID               int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	BearerID         int    `gorm:"column:bearer_id;type:int(11);not null" json:"bearer_id"`
	BearerLogID      int    `gorm:"column:bearer_log_id;type:int(11);not null" json:"bearer_log_id"`
	CreateAt         int    `gorm:"column:create_at;type:int(11);not null" json:"create_at"`                             // 进线时间(创建时间)
	CustomerID       int    `gorm:"column:customer_id;type:int(10) unsigned;not null" json:"customer_id"`                // 客服ID
	Username         string `gorm:"column:username;type:varchar(255);not null" json:"username"`                          // 客服昵称
	Nickname         string `gorm:"column:nickname;type:varchar(255)" json:"nickname"`                                   // 用户昵称
	SessionStartAt   int    `gorm:"column:session_start_at;type:int(11) unsigned;not null" json:"session_start_at"`      // 会话开始时间
	SessionInitiator int8   `gorm:"column:session_initiator;type:tinyint(2) unsigned;not null" json:"session_initiator"` // 会话发起方  1用户 2客服
	SessionClose     int8   `gorm:"column:session_close;type:tinyint(2) unsigned;not null" json:"session_close"`         // 会话结束方  1用户  2客服  3系统
	SessionSource    int8   `gorm:"column:session_source;type:tinyint(2);not null" json:"session_source"`                // 会话来源  1店铺H5  2供应商  3商家  4客服  5客服转接
	UserIP           string `gorm:"column:user_ip;type:varchar(20);not null" json:"user_ip"`
	Status           int8   `gorm:"column:status;type:tinyint(1);not null" json:"status"` // 1 进行中  2已结束
	CustomerFd       int    `gorm:"column:customer_fd;type:int(10)" json:"customer_fd"`
	UserFd           int    `gorm:"column:user_fd;type:int(10)" json:"user_fd"`
	UserID           int    `gorm:"column:user_id;type:int(11);not null" json:"user_id"`
	EndAt            int    `gorm:"column:end_at;type:int(11) unsigned" json:"end_at"` // 结束时间
}

// GbPostageSpecial 邮费模板特殊区域邮费设置
type GbPostageSpecial struct {
	ID             int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Area           string    `gorm:"column:area;type:varchar(100);not null" json:"area"`
	AddPostage     float64   `gorm:"column:add_postage;type:decimal(10,2);not null" json:"add_postage"`
	DefaultCount   float64   `gorm:"column:default_count;type:decimal(10,2);not null" json:"default_count"`     // 首重单位数量
	DefaultPostage float64   `gorm:"column:default_postage;type:decimal(10,2);not null" json:"default_postage"` // 首重邮费
	IsDel          int       `gorm:"column:is_del;type:int(1)" json:"is_del"`
	PostageID      int       `gorm:"index;column:postage_id;type:int(11);not null" json:"postage_id"`        // 对应的邮费模板
	GbPostage      GbPostage `gorm:"association_foreignkey:postage_id;foreignkey:id" json:"gb_postage_list"` // 邮费模板表
	AddCount       float64   `gorm:"column:add_count;type:decimal(10,2);not null" json:"add_count"`
}

// GbMerchantFrontendBargain H5端砍价专区发起砍价
type GbMerchantFrontendBargain struct {
	ID              int    `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	UId             int    `gorm:"column:uid;type:int(11) unsigned;not null" json:"uid"`                             // 用户ID
	ActID           int    `gorm:"column:act_id;type:int(11) unsigned;not null" json:"act_id"`                       // 活动ID
	StoreID         int    `gorm:"column:store_id;type:int(11) unsigned;not null" json:"store_id"`                   // 店铺id
	MerchantStandID int    `gorm:"column:merchant_stand_id;type:int(11) unsigned;not null" json:"merchant_stand_id"` // 商家产品属性ID
	ProductStandID  int    `gorm:"column:product_stand_id;type:int(11) unsigned;not null" json:"product_stand_id"`   // 平台产品属性ID
	Status          int8   `gorm:"column:status;type:tinyint(2) unsigned;not null" json:"status"`                    // 0关闭 1开启 2完成
	Ctime           int    `gorm:"column:ctime;type:int(10) unsigned;not null" json:"ctime"`                         // 砍价发起时间
	IsDel           int8   `gorm:"column:is_del;type:tinyint(2) unsigned;not null" json:"is_del"`                    // 0已删除 1未删除
	IsOrdered       int8   `gorm:"column:is_ordered;type:tinyint(2) unsigned;not null" json:"is_ordered"`            // 是否已经下单 不包含已付款
	IncrementID     string `gorm:"column:increment_id;type:char(15);not null" json:"increment_id"`                   // 下单的主订单表流水号 用来做定时红包发放
}

// GbProStandard 平台产品规格
type GbProStandard struct {
	ID                 int     `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	IsDel              int8    `gorm:"column:is_del;type:tinyint(1);not null" json:"is_del"`                         // 是否删除  1：是   0：否
	ProID              int     `gorm:"index;column:pro_id;type:int(11) unsigned;not null" json:"pro_id"`             // 产品ID
	Code               string  `gorm:"index;column:code;type:varchar(100);not null" json:"code"`                     // 编号(产品平台编号)
	ProCode            string  `gorm:"index;column:pro_code;type:varchar(40);not null" json:"pro_code"`              // 产品编码 EAN13编码
	Sku                string  `gorm:"column:sku;type:varchar(60)" json:"sku"`                                       // 平台商品SKU
	Property1          string  `gorm:"column:property_1;type:varchar(255)" json:"property_1"`                        // 属性1
	Property1Sort      int8    `gorm:"column:property_1_sort;type:smallint(2)" json:"property_1_sort"`               // 颜色固定排序
	Property2          string  `gorm:"column:property_2;type:varchar(255)" json:"property_2"`                        // 属性2
	Property2Sort      int8    `gorm:"column:property_2_sort;type:smallint(2)" json:"property_2_sort"`               // 规格固定排序
	MarketPrice        float64 `gorm:"column:market_price;type:decimal(10,2);not null" json:"market_price"`          // 市场价
	MinPrice           float64 `gorm:"column:min_price;type:decimal(10,2);not null" json:"min_price"`                // 最低价
	BatchPrice         float64 `gorm:"column:batch_price;type:decimal(10,2);not null" json:"batch_price"`            // 批发价(供应商给平台的进货价)
	SellPrice          float64 `gorm:"column:sell_price;type:decimal(10,2);not null" json:"sell_price"`              // 建议售价
	MCostPrice         float64 `gorm:"column:m_cost_price;type:decimal(10,2)" json:"m_cost_price"`                   // 平台成本价
	MBatchPrice        float64 `gorm:"column:m_batch_price;type:decimal(10,2)" json:"m_batch_price"`                 // 平台批发价(平台卖给商家的价格)
	MBatchUntaxedPrice float64 `gorm:"column:m_batch_untaxed_price;type:decimal(10,2)" json:"m_batch_untaxed_price"` // 未税价格
	TaxRate            float64 `gorm:"column:tax_rate;type:decimal(10,4)" json:"tax_rate"`                           // 税率
	Qty                int     `gorm:"column:qty;type:int(11);not null" json:"qty"`                                  // 库存
	MinQty             int     `gorm:"column:min_qty;type:int(11);not null" json:"min_qty"`                          // 预警库存
	SupSku             string  `gorm:"column:sup_sku;type:varchar(60)" json:"sup_sku"`                               // 供应商sku
	Pic                string  `gorm:"column:pic;type:varchar(255)" json:"pic"`                                      // 图片
	OccupyQty          int     `gorm:"column:occupy_qty;type:int(11)" json:"occupy_qty"`                             // 订单占有库存数量
	ZiyingPrice        float64 `gorm:"column:ziying_price;type:decimal(10,2)" json:"ziying_price"`                   // 自营商品定价（定售价，用于代运营，导入商品）
}

type GbStoreCardCoupon struct {
	ID             int       `gorm:"primary_key;column:id;type:int(11) unsigned;not null" json:"-"`
	TaskRecordID   int       `gorm:"index;column:task_record_id;type:int(11);not null" json:"task_record_id"`                         // 生成批次ID
	StoreID        int       `gorm:"unique_index:search_and_validate;column:store_id;type:int(11) unsigned;not null" json:"store_id"` // 店铺ID
	CardNo         string    `gorm:"unique;column:card_no;type:char(12);not null" json:"card_no"`                                     // 卡号
	CardSecret     string    `gorm:"unique_index:search_and_validate;column:card_secret;type:char(6);not null" json:"card_secret"`    // 卡密
	CardScore      int       `gorm:"column:card_score;type:int(11) unsigned;not null" json:"card_score"`                              // 卡券积分
	ActiveStatus   int8      `gorm:"column:active_status;type:tinyint(2) unsigned;not null" json:"active_status"`                     // 激活状态 0未激活 1激活
	RechargeStatus int8      `gorm:"column:recharge_status;type:tinyint(2);not null" json:"recharge_status"`                          // 充值状态 0 未充值 1已充值
	CreatedAt      time.Time `gorm:"column:created_at;type:datetime;not null" json:"created_at"`                                      // 创建时间
	UpdatedAt      time.Time `gorm:"column:updated_at;type:datetime;not null" json:"updated_at"`                                      // 更新时间
}

type GbSupProductExamineLog struct {
	ID            int       `gorm:"primary_key;column:id;type:int(11);not null" json:"-"`
	Changetype    int8      `gorm:"column:changetype;type:tinyint(4);not null" json:"changetype"`    // 变更类型：1.新增2.编辑3.上架4.下架5撤销上架6撤销下架
	BeforeData    string    `gorm:"column:before_data;type:text" json:"before_data"`                 // 变更前json数据
	SupProid      int       `gorm:"column:sup_proid;type:int(11);not null" json:"sup_proid"`         // 变更供应商产品id
	AfterData     string    `gorm:"column:after_data;type:text;not null" json:"after_data"`          // 变更后json数据
	SubmitUserid  int       `gorm:"column:submit_userid;type:int(11);not null" json:"submit_userid"` // 提交变更审核供应商用户id
	SubmitTime    time.Time `gorm:"column:submit_time;type:datetime;not null" json:"submit_time"`    // 变更审核提交时间
	ExamineTime   time.Time `gorm:"column:examine_time;type:datetime" json:"examine_time"`           // 审核时间
	ExamineUserid int       `gorm:"column:examine_userid;type:int(11)" json:"examine_userid"`        // 变更审核平台用户id
	ExamineStatus int8      `gorm:"column:examine_status;type:tinyint(11)" json:"examine_status"`    // 变更审核状态1.审核通过2.审核不通过
	ExamineRemark string    `gorm:"column:examine_remark;type:varchar(255)" json:"examine_remark"`   // 变更审核备注
}
