package api

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {
	router.GET("/get_users", GetUsers)
	router.GET("/get_user/:id/", GetUser)
	router.POST("/create_user", CreateUser)
	router.GET("/ingest/:id/", GetIngest)
	router.GET("/trimmer/:id/", GetTrimmer)
	router.PUT("/trimmer/:id/", PutTrimmer)
}
