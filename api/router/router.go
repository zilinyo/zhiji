package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
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
	//login register

	router.POST("/login", Login)
	router.POST("/register", Register)

	router.Use(jwt.JWTAuth())
	{

		router.GET("/api/user/list", List)

		router.GET("/tree", Tree)

		//分类
		category := router.Group("/category")
		{
			category.GET("", GetAllCategory)
			category.GET("/:id", GetOneCategory)
			category.POST("", FetchSingleCategory)
			category.DELETE("/:id", DestroyCategory)
			category.PUT("/:id", UpdateCategory)

		}

		//品牌
		brand := router.Group("/brand")
		{
			brand.GET("", GetAllBrand)
			brand.GET("/:id", GetOneBrand)
			brand.POST("", FetchSingleBrand)
			brand.DELETE("/:id", DestroyBrand)
			brand.PUT("/:id", UpdateBrand)

		}

		//品牌
		product := router.Group("/product")
		{
			product.GET("", GetAllProduct)
			//product.GET("/:id", GetOneProduct)
			product.POST("", FetchSingleProduct)
			//product.DELETE("/:id", DestroyProduct)
			product.PUT("/:id", UpdateProduct)
		}
		//router.GET("/brand/all", BrandAll)
		//router.GET("/brand/one", BrandtOne)
		//router.POST("/brand/save", BrandSave)
		//router.DELETE("/brand/del", BrandDel)
		//router.PUT("/brand/update", BrandUpdate)

		//商品
		//router.POST("/product/save", ProductSave)
		//router.GET("/product/all", ProductAll)
		//router.PUT("/product/update", ProductUpdate)
		//router.DELETE("/product/del", ProductyDel)

		//图片
		router.POST("/upload", Uploadfiles)
		router.DELETE("/deleteFile", DelFile)

	}

	return router
}
