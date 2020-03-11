package apis

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	. "zhiji/api/components"
)

func Uploadfiles(c *gin.Context) {
	var Up map[string]interface{}
	body, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(body, &Up)

	co := new(Cosupload)
	result, err := co.Uploadfile(Up["serverPath"].(string), Up["localPath"].(string))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "失败",
			"result":  result,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "成功",
		"result":  result,
	})
}
func DelFile(c *gin.Context) {
	serverPath := c.Query("serverPath")
	co := new(Cosupload)
	result, err := co.Delfile(serverPath)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "失败",
			"result":  result,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "成功",
		"result":  result,
	})
}
