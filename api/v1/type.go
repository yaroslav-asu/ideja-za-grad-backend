package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/yaroslav-asu/urban-map/internal/utils/db"
	"github.com/yaroslav-asu/urban-map/models/gorm/marker"
	"github.com/yaroslav-asu/urban-map/pkg/app"
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
