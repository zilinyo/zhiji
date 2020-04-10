package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	db "zhiji/api/database"
	"zhiji/api/models"
)

//列表
func GetAllProduct(c *gin.Context) {

	var productInfo []models.ProductInfo

	err := db.Eloquent.Find(&productInfo).Error

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errNotExist.Error()})

		return
	}

	c.JSON(http.StatusOK, gin.H{"code": "200", "data": &productInfo})
}

/**
 * @author lin
 * 保存
 */
func FetchSingleProduct(c *gin.Context) {
	var product = new(models.ProductInfo)

	if err := c.BindJSON(product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errNotExist.Error()})

		return
	}

	ProductID, err := product.Insert()

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errNotExist.Error()})

		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "200", "data": ProductID})

}

//修改
func UpdateProduct(c *gin.Context) {

	var product = new(models.ProductInfo)

	if err := c.BindJSON(product); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errNotExist.Error()})

		return
	}
	db.Eloquent.Where("product_id in (?)", c.Query("product_id")).Delete(&models.ProductPicInfo{})

	db.Eloquent.Model(product).Updates(&product)

	c.JSON(http.StatusOK, gin.H{"code": "200", "data": product})

}

//删除
func ProductyDel(c *gin.Context) {

	db.Eloquent.Where("product_id in (?)", c.Query("product_id")).Delete(&models.ProductPicInfo{})

	db.Eloquent.Where("product_id = ?", c.Query("product_id")).Delete(&models.ProductInfo{})

	c.JSON(http.StatusOK, gin.H{"code": "200", "data": "sucess"})

}
