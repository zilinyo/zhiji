package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	db "zhiji/api/database"
	"zhiji/api/models"
)

//列表
func GetAllProduct(c *gin.Context) {

	var productInfo []models.ProductInfo

	Db := db.Eloquent

	if Id, isExist := c.GetQuery("ProductID"); isExist == true && Id != "0" { //商品id
		if IdInt, _ := strconv.Atoi(Id); IdInt != 0 {
			Db = Db.Where("product_id = ?", IdInt)
		}

	}

	if ProductCore, isExist := c.GetQuery("ProductCore"); isExist == true { //商品编号
		if ProductCore, _ := strconv.Atoi(ProductCore); ProductCore != 0 {
			Db = Db.Where("product_id = ?", ProductCore)
		}
	}

	if ProductName, isExist := c.GetQuery("ProductName"); isExist == true { //商品编号
		if ProductName, _ := strconv.Atoi(ProductName); ProductName != 0 {
			Db = Db.Where("product_id = ?", ProductName)
		}
	}

	if IsNew, isExist := c.GetQuery("IsNew"); isExist == true { //新品
		if IsNew, _ := strconv.Atoi(IsNew); IsNew != 0 {
			Db = Db.Where("is_new = ?", IsNew)
		}
	}

	if IsHot, isExist := c.GetQuery("IsHot"); isExist == true { //热卖品
		if IsHot, _ := strconv.Atoi(IsHot); IsHot != 0 {
			Db = Db.Where("is_hot = ?", IsHot)
		}
	}

	if IsBanner, isExist := c.GetQuery("IsBanner"); isExist == true { //banner图
		if IsBanner, _ := strconv.Atoi(IsBanner); IsBanner != 0 {
			Db = Db.Where("is_banner = ?", IsBanner)
		}
	}

	err := Db.Order("product_id desc").Find(&productInfo).Error

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errNotExist.Error()})

		return
	}

	c.JSON(http.StatusOK, gin.H{"code": "200", "data": &productInfo})
}

//单条数据
func GetOneProduct(c *gin.Context) {

	var productInfo models.ProductInfo

	if err := db.Eloquent.Where("product_id = ?", c.Param("id")).Preload("PicInFo").
		Preload("ProductParam").Find(&productInfo).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"code": "200", "err": err})

	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": productInfo})

}

/**
 * @author lin
 * 保存
 */
func FetchSingleProduct(c *gin.Context) {

	var product = new(models.ProductInfo)

	if err := c.BindJSON(product); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errNotExist.Error(), "err": err})

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
func DestroyProduct(c *gin.Context) {

	if err := db.Eloquent.Where("product_id = ?", c.Param("id")).Delete(&models.ProductInfo{}).Error; err != nil {

		c.JSON(http.StatusOK, gin.H{"code": "200", "err": err})

	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": "success"})

}
