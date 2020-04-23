package MiniProgram

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	face "zhiji/api/interface"
	"zhiji/api/models"
)

/*
@Time : 2020/4/14 1:04 PM
@Author : apple
@File : IndexController
@Software: GoLand
*/

type data struct {
	banner []models.ProductInfo
}

/**
*首页数据
 */
func Index(c *gin.Context) {

	m1 := make(map[string]interface{})

	var indexing = []face.MiniIndex{models.ProductInfo{}}

	for _, animal := range indexing {

		m1["banner"] = animal.Banner()

		m1["channel"] = animal.Channel()

		m1["newGoodsList"] = animal.NewGoodsList()

	}
	//for i := 1; i < len(m1["banner"]); i++ {
	//	fmt.Printf("i=%d\n", i)
	//}
	c.JSON(http.StatusOK, gin.H{"code": "200", "data": m1})

}
