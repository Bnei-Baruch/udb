package api

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {
	router.GET("/get_users", GetUsers)
	router.GET("/get_user/:id", GetUser)
	router.POST("/create_user", CreateUser)
	router.GET("/ingest/:id", GetIngest)

	router.GET("/trimmer/find", FindTrimmer)
	router.GET("/trimmer/:id", GetTrimmer)
	router.GET("/trim", GetTrimmed)
	router.PUT("/trimmer/:id", PutTrimmer)
	router.POST("/trimmer/:id/wfstatus/:key", TrimmerStatusValue)
	router.DELETE("/trimmer/:id", RemoveTrimmer)
}
