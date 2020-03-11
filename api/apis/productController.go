package apis

import (
	"github.com/gin-gonic/gin"

	"net/http"
	"strconv"
	db "zhiji/api/database"
	"zhiji/api/models"
)

//列表
func ProductAll(c *gin.Context) {
	var product []models.ProductInfo
	var page, _ = strconv.Atoi(c.DefaultQuery("page", "1"))
	var pageSize, _ = strconv.Atoi(c.DefaultQuery("pageSize", "1"))
	var total int = 0
	db.Eloquent.Find(&product).Count(&total)
	db.Eloquent.Limit(pageSize).Offset((page - 1) * pageSize).Preload("PicInFo").Find(&product)
	c.JSON(http.StatusOK, gin.H{
		"code":     200,
		"data":     product,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}

/**
 * @author lin
 * 保存
 */
func ProductSave(c *gin.Context) {
	var product = new(models.ProductInfo)
	if err := c.BindJSON(product); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	ProductID, err := product.Insert()
	if err != nil {
		c.JSON(500, gin.H{
			"err": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"data": ProductID,
	})

}
func ProductUpdate(c *gin.Context) {

	var product = new(models.ProductInfo)
	if err := c.BindJSON(product); err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	db.Eloquent.Where("product_id in (?)", c.Query("product_id")).Delete(&models.ProductPicInfo{})
	db.Eloquent.Model(product).Updates(&product)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": product,
	})
}
func ProductyDel(c *gin.Context) {
	db.Eloquent.Where("product_id in (?)", c.Query("product_id")).Delete(&models.ProductPicInfo{})
	db.Eloquent.Where("product_id = ?", c.Query("product_id")).Delete(&models.ProductInfo{})
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": "success",
	})
}
