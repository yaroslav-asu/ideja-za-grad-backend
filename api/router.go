package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	v1 "urban-map/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"*"},
	}))
	groupV1 := r.Group("/api/v1")
	groupV1.GET("/markers/:id", v1.GetMarker)
	groupV1.GET("/markers", v1.GetMarkers)
	groupV1.POST("/markers", v1.CreateMarker)
	groupV1.GET("/types", v1.GetTypes)
	return r
}
