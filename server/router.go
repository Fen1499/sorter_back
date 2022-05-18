package server

import (
	"sorter_gin/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	gin.SetMode("")
	router := gin.Default()
	router.Use(cors.Default())

	sorted := new(controllers.SortedController)
	health := new(controllers.HealthController)
	router.GET("/health", health.HealthCheck)

	v1 := router.Group("v1")
	{
		v1.POST("/sort", sorted.MergeSort)
	}
	return router
}
