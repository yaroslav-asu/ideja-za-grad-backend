package v1

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"urban-map/internal/utils/db"
	"urban-map/models/gorm/marker"
	"urban-map/pkg/app"
)

func GetMarker(c *gin.Context) {
	appG := app.Gin{C: c}
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		appG.Response(400, err.Error())
		return
	}
	appG.Response(200, marker.Get(db.GetDB(), uint(id)))
}

func GetMarkers(c *gin.Context) {
	appG := app.Gin{C: c}
	appG.Response(200, marker.GetAll(db.GetDB()))
}
