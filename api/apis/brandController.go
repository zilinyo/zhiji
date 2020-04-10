package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	db "zhiji/api/database"
	"zhiji/api/models"
)

/*
品牌数据
*/
func GetAllBrand(c *gin.Context) {

	var brandInfo []models.BrandInfo

	err := db.Eloquent.Find(&brandInfo).Error

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errNotExist.Error()})

		return
	}

	c.JSON(http.StatusOK, gin.H{"code": "200", "data": &brandInfo})
}

/*
单条记录
*/
func GetOneBrand(c *gin.Context) {

	var brandInfo models.BrandInfo

	id := c.Query("id")

	err := db.Eloquent.Find(&brandInfo, id).Error

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errNotExist.Error()})

		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": &brandInfo})
}

/*
添加
*/
func FetchSingleBrand(c *gin.Context) {
	var brandInfo models.BrandInfo

	if err := c.ShouldBind(&brandInfo); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidBody.Error(), "err": err.Error()})

		return
	}

	err := db.Eloquent.Create(&brandInfo).Error

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInsertionFailed.Error(), "err": err.Error()})

		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": brandInfo.BrandID})
}

/*
删除
*/
func DestroyBrand(c *gin.Context) {

	var brandInfo models.BrandInfo

	err := db.Eloquent.Delete(&brandInfo, "brand_id ="+c.Param("id")).Error

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errDeletionFailed.Error(), "err": err.Error()})

		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": "success"})
}

/*
更新
*/
func UpdateBrand(c *gin.Context) {

	var brandInfo models.BrandInfo

	if err := c.ShouldBind(&brandInfo); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidBody.Error(), "err": err.Error()})

		return
	}

	err := db.Eloquent.Model(&brandInfo).Where("brand_id = ?", c.Param("id")).Updates(&brandInfo).Error

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInsertionFailed.Error(), "err": err.Error()})

		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": brandInfo})
}
