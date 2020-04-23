package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	db "zhiji/api/database"
	m "zhiji/api/models"
)

/**
*列表
 */
func GetAllCategory(c *gin.Context) {

	var productCategory []m.ProductCategory

	err := db.Eloquent.Find(&productCategory).Error

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errNotExist.Error()})

		return
	}

	c.JSON(http.StatusOK, gin.H{"code": "200", "data": &productCategory})
}

/**
*单条数据
 */
func GetOneCategory(c *gin.Context) {

	var productCategory m.ProductCategory

	id := c.Query("id")

	err := db.Eloquent.Find(&productCategory, id).Error

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errNotExist.Error()})

		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": &productCategory})

}

/**
*保存
 */
func FetchSingleCategory(c *gin.Context) {

	var productCategory m.ProductCategory

	if err := c.ShouldBind(&productCategory); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidBody.Error(), "err": err.Error()})

		return
	}

	err := db.Eloquent.Create(&productCategory).Error

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInsertionFailed.Error(), "err": err.Error()})

		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": productCategory})
}

/**
*删除分类
 */
func DestroyCategory(c *gin.Context) {

	var productCategory m.ProductCategory

	err := db.Eloquent.Delete(&productCategory, "category_id ="+c.Param("id")).Error

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errDeletionFailed.Error(), "err": err.Error()})

		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": productCategory})
}

/**
*更新分类
 */
func UpdateCategory(c *gin.Context) {

	var productCategory m.ProductCategory

	if err := c.ShouldBind(&productCategory); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidBody.Error(), "err": err.Error()})

		return
	}

	err := db.Eloquent.Model(&productCategory).Where("category_id = ?", c.Param("id")).Updates(&productCategory).Error

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInsertionFailed.Error(), "err": err.Error()})

		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": productCategory})
}

/**
无限极分类
*/
func Tree(c *gin.Context) {
	var productCategory []*m.ProductCategory

	db.Eloquent.Find(&productCategory)

	data := buildData(productCategory)

	result := makeTreeCore(0, data)

	c.JSON(http.StatusOK, gin.H{"status": "success", "result": result})
}

/**
数据组装
*/
func buildData(list []*m.ProductCategory) map[int]map[int]*m.ProductCategory {

	var data = make(map[int]map[int]*m.ProductCategory)

	for _, v := range list {

		CategoryID := int(v.CategoryID)

		ParentID := int(v.ParentID)

		if _, ok := data[ParentID]; !ok {

			data[ParentID] = make(map[int]*m.ProductCategory)

		}

		data[ParentID][CategoryID] = v
	}

	return data
}

/**
数据排序
*/
func makeTreeCore(index int, data map[int]map[int]*m.ProductCategory) []*m.ProductCategory {

	tmp := make([]*m.ProductCategory, 0)

	for id, item := range data[index] {

		if data[id] != nil {

			item.Children = makeTreeCore(id, data)

		}

		tmp = append(tmp, item)
	}

	return tmp
}

/**
查处上级分类
*/
func GetParentCategory(c *gin.Context) {

}
