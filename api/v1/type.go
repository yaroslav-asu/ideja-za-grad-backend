package v1

import (
	"github.com/gin-gonic/gin"
	"urban-map/internal/utils/db"
	"urban-map/models/gorm/marker"
	"urban-map/pkg/app"
)

func GetTypes(c *gin.Context) {
	appG := app.Gin{C: c}
	types := marker.GetAllTypes(db.GetDB())
	if len(types) > 0 {
		appG.Response(200, types)
		return
	}
	appG.Response(200, []marker.Type{})
}
