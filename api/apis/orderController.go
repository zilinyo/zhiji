package apis

import (
	"github.com/gin-gonic/gin"
	"zhiji/api/models"
)

func OrderSave(c *gin.Context) {
	var product = new(models.OrderMaster)
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
	}

	c.JSON(200, gin.H{
		"data": ProductID,
	})
}
