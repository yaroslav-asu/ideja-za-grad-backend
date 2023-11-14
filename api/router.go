package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	v1 "github.com/yaroslav-asu/urban-map/api/v1"
	"log"
	"sync"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"*"},
	}))
	groupV1 := r.Group("/api/v1")
	groupV1.Static("/static", "./static/images")
	groupV1.GET("/markers", v1.GetMarkers)
	groupV1.GET("/markers/:id", v1.GetMarker)
	groupV1.GET("/markers/:id/images", v1.GetImages)
	groupV1.POST("/markers", v1.CreateMarker)
	groupV1.GET("/types", v1.GetTypes)
	return r
}
func Run(wg *sync.WaitGroup) {
	defer wg.Done()
	r := InitRouter()
	err := r.Run()
	if err != nil {
		log.Fatal(err.Error())
	}
}
