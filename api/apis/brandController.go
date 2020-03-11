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

func BrandAll(c *gin.Context) {
	var Brand []models.BrandInfo
	var page, _ = strconv.Atoi(c.DefaultQuery("page", "1"))
	var pageSize, _ = strconv.Atoi(c.DefaultQuery("pageSize", "1"))
	var total int = 0
	db.Eloquent.Find(&Brand).Count(&total)
	db.Eloquent.Limit(pageSize).Offset((page - 1) * pageSize).Find(&Brand)
	c.JSON(http.StatusOK, gin.H{
		"code":     200,
		"data":     Brand,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}
func BrandtOne(c *gin.Context) {
	var Brand []models.BrandInfo
	id := c.Query("id")
	db.Eloquent.Find(&Brand, id)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": Brand,
	})
}
func BrandSave(c *gin.Context) {
	form := struct {
		BrandName    string `form:"brand_name" binding:"required"`
		Telephone    string `form:"telephone" binding:"required"`    // 联系电话
		BrandWeb     string `form:"brand_web" `                      // 品牌网络
		BrandLogo    string `form:"brand_logo" `                     // 品牌logo URL
		BrandDesc    string `form:"brand_desc" `                     // 品牌描述
		BrandStatus  int8   `form:"brand_status" binding:"required"` // 品牌状态,0禁用,1启用
		BrandOrder   int8   `form:"brand_order" binding:"required"`  // 排序
		ModifiedTime string `form:"modified_time" binding:"required"`
	}{
		BrandWeb:  "null",
		BrandLogo: "null",
		BrandDesc: "0",
	}
	if err := c.BindJSON(&form); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	//	fmt.Println(form)
	Brand := models.BrandInfo{
		BrandName:    form.BrandName,
		Telephone:    form.Telephone,
		BrandWeb:     form.BrandWeb,
		BrandLogo:    form.BrandLogo,
		BrandDesc:    form.BrandDesc,
		BrandStatus:  form.BrandStatus,
		BrandOrder:   form.BrandOrder,
		ModifiedTime: form.ModifiedTime,
	}
	ok := db.Eloquent.NewRecord(Brand) // => 主键为空返回`true`
	if !ok {
		panic("保存失败")
	}
	db.Eloquent.Create(&Brand)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": form,
	})
}
func BrandDel(c *gin.Context) {

	var Brand []models.BrandInfo
	db.Eloquent.Delete(&Brand, c.Query("id"))
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": Brand,
	})
}
func BrandUpdate(c *gin.Context) {
	var user map[string]interface{}
	body, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(body, &user)
	var Brand []models.BrandInfo
	id := user["BrandId"]
	db.Eloquent.Model(&Brand).Where("brand_id = ?", id).
		Updates(user)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": user,
	})
}
