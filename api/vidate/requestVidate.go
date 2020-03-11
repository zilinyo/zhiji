package vidate

type ProducVidate struct {
	ProductCore     string  `form:"product_core" binding:"required"`      // 商品编码
	ProductName     string  `form:"product_name" binding:"required"`      // 商品名称
	BarCode         string  `form:"bar_code" binding:"required"`          // 国条码
	BrandID         int     `form:"brand_id" binding:"required"`          // 品牌表的ID
	OneCategoryID   int16   `form:"one_category_id" binding:"required"`   // 一级分类ID
	TwoCategoryID   int16   `form:"two_category_id" binding:"required"`   // 二级分类ID
	ThreeCategoryID int16   `form:"three_category_id" binding:"required"` // 三级分类ID
	Price           float64 `form:"price" binding:"required"`             // 商品销售价格
	AverageCost     float64 `form:"average_cost" binding:"required"`      // 商品加权平均成本
	PublishStatus   int8    `form:"publish_status" binding:"required"`    // 上下架状态：0下架1上架
	AuditStatus     int8    `form:"audit_status" binding:"required"`      // 审核状态：0未审核，1已审核
	Weight          float32 `form:"weight"`                               // 排序
	Length          float32 `form:"length"`                               // 排序
	Height          float32 `form:"height"`
	Width           float32 `form:"width"`
	ColorType       string  `form:"color_type"`
	ProductionDate  string  `form:"production_date"`
	ShelfLife       int     `form:"shelf_life"`
	Descript        string  `form:"descript"`
	Indate          string  `form:"indate" `
	ModifiedTime    string  `form:"modified_time"`
	PicInFo         []*PPCvidate
}

type PPCvidate struct {
	ProductID    int    `form:"product_core" binding:"required"` // 商品图片ID
	PicDesc      string `form:"product_core" binding:"required"` // 商品ID
	PicURL       string `form:"product_core" binding:"required"` // 图片描述
	IsMaster     int8   `form:"product_core" binding:"required"` // 图片URL
	PicOrder     int8   `form:"product_core" binding:"required"` // 是否主图：0.非主图1.主图
	PicStatus    int8   `form:"product_core" binding:"required"` // 图片是否有效：0无效 1有效
	ModifiedTime string `form:"product_core" binding:"required"` // 最后修改时间
}
