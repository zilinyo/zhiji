package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	. "zhiji/api/components"
)

//上传文件
func Uploadfiles(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		return
	}
	filename := file.Filename
	fmt.Println("........", filename)
	errs := c.SaveUploadedFile(file, "upload/"+filename)
	if err != nil {
		fmt.Println(errs)
	}
	co := new(Cosupload)
	//上传到cos
	result, err := co.Uploadfile("images/"+filename, "upload/"+filename)
	os.Remove("upload" + filename)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": "失败",
			"result":  result,
			"err":     err,
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
	fmt.Println(err)
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
