package apis

import (
	"github.com/gin-gonic/gin"
	db "zhiji/api/database"
	"zhiji/api/models"
)

//sup_pro_standard   供应商产品规格

// 平台产品规格  pro_standard_rela_sup

//$proStandardRelaSupModel

//stand-id   和  count

func OrderSave(c *gin.Context) {
	//var Pid string //产品id
	//
	//var StantId string   //  产品规格id
	//
	var gbproduct models.GbProduct

	db.Eloquent.Find(&gbproduct)

	//判断库存  gb_product gb_sup_product    gb_pro_standard    gb_sup_pro_standard

	//下单入队列 gb_order gb_order_children  gb_order_item

	//运费

	//积分
}
