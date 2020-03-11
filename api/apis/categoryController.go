package apis

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
	db "zhiji/api/database"
	"zhiji/api/models"
)

//列表
func GetAll(c *gin.Context) {
	var prodict []models.ProductCategory
	var page, _ = strconv.Atoi(c.DefaultQuery("page", "1"))
	var pageSize, _ = strconv.Atoi(c.DefaultQuery("pageSize", "1"))
	var total int = 0
	db.Eloquent.Find(&prodict).Count(&total)
	db.Eloquent.Limit(pageSize).Offset((page - 1) * pageSize).Find(&prodict)
	c.JSON(http.StatusOK, gin.H{
		"code":     200,
		"data":     prodict,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}

//一条
func GetOne(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": 11111,
	})
	var prodict models.ProductCategory
	id := c.Query("id")
	db.Eloquent.Find(&prodict, id)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": prodict,
	})
}

type CaregoryForm struct {
	CategoryID   string `form:"category_id"`
	CategoryName string `form:"category_name" binding:"required"`
	CategoryCode string `form:"CategoryCode" binding:"required"`
	ParentID     int16  `jsonn:"parent_id" binding:"required"`
	//CategoryLevel int8 `form:"category_level" binding:"required"`
	CategoryStatus int8   `form:"category_status"binding:"required"`
	ModifiedTime   string `form:"modified_time"binding:"required"`
	CategoryLevel  int8   `form:"category_level,default=1"`
}

func CategorySave(c *gin.Context) {
	form := struct {
		CategoryID   string `form:"category_id"`
		CategoryName string `form:"category_name" binding:"required"`
		CategoryCode string `form:"CategoryCode" binding:"required"`
		ParentID     int16  `jsonn:"parent_id" binding:"required"`
		//CategoryLevel int8 `form:"category_level" binding:"required"`
		CategoryStatus int8   `form:"category_status"binding:"required"`
		ModifiedTime   string `form:"modified_time"binding:"required"`
		CategoryLevel  int8   `form:"category_level,default=1"`
	}{
		CategoryLevel: 1,
	}
	if err := c.BindJSON(&form); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	//	fmt.Println(form)
	Product := models.ProductCategory{
		CategoryName:   form.CategoryName,
		CategoryCode:   form.CategoryCode,
		ParentID:       form.ParentID,
		CategoryLevel:  form.CategoryLevel,
		CategoryStatus: form.CategoryStatus,
		ModifiedTime:   form.ModifiedTime,
	}
	ok := db.Eloquent.NewRecord(Product) // => 主键为空返回`true`
	if !ok {
		panic("保存失败")
	}
	db.Eloquent.Create(&Product)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": form,
	})
}
func CategoryDel(c *gin.Context) {
	var prodict models.ProductCategory
	db.Eloquent.Delete(&prodict, c.Query("id"))
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": prodict,
	})
}
func CategoryUpdate(c *gin.Context) {
	var user map[string]interface{}
	body, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(body, &user)
	var prodict models.ProductCategory
	id := user["CategoryID"]
	db.Eloquent.Model(&prodict).Where("brand_id = ?", id).
		Updates(user)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": prodict,
	})
}
