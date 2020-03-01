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
	router.POST("/login", apis.Login)
	router.POST("/register", apis.Register)
	router.Use(jwt.JWTAuth())

	{
		router.GET("/api/user/list", List)
		router.GET("/routinne", Routine)
		router.GET("/lock", Lock)
		router.GET("/channels", Getchannel)
		router.POST("/save-order", Saveorder)
		router.POST("/setcount", Setcount)
		router.POST("/getcount", Getcount)
	}
	// 跨域
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "X-Requested-With", "X-CSRF-TOKEN"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))

	//router.GET("/api/user/list", List)
	//router.GET("/api/statistics", Statistics)

	return router
}
