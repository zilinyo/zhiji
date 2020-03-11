package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
	"zhiji/api/apis"
	. "zhiji/api/apis"
	jwt "zhiji/api/middleware"
)

func InitRouter() *gin.Engine {

	router := gin.Default()
	// 跨域
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "X-Requested-With", "X-CSRF-TOKEN", "TOKEN"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))
	router.POST("/login", apis.Login)
	router.POST("/register", apis.Register)
	router.Use(jwt.JWTAuth())

	{
		router.GET("/api/user/list", List)
		router.GET("/routinne", Routine)
		router.GET("/lock", Lock)
		router.POST("/save-order", Saveorder)
		router.POST("/setcount", Setcount)
		router.POST("/getcount", Getcount)
		//分类
		router.GET("/category/all", GetAll)
		router.GET("/category/one", GetOne)
		router.POST("/category/save", CategorySave)
		router.DELETE("/category/del", CategoryDel)
		router.PUT("/category/update", CategoryUpdate)
		//品牌
		router.GET("/brand/all", BrandAll)
		router.GET("/brand/one", BrandtOne)
		router.POST("/brand/save", BrandSave)
		router.DELETE("/brand/del", BrandDel)
		router.PUT("/brand/update", BrandUpdate)

		//添加商品
		router.POST("/product/save", ProductSave)
		router.GET("/product/all", ProductAll)
		router.PUT("/product/update", ProductUpdate)
		router.DELETE("/product/del", ProductyDel)
		router.POST("/upload", Uploadfiles)
		router.DELETE("/deleteFile", DelFile)

		//router.GET("/SetClient", Uploadfile)
		//router.GET("/DelClient", Delfile)

	}

	//router.GET("/api/user/list", List)
	//router.GET("/api/statistics", Statistics)

	return router
}
