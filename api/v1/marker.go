package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"urban-map/internal/utils/db"
	"urban-map/models/gorm/marker"
	"urban-map/pkg/app"
	"urban-map/telegram_bot"
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
	appG.Response(200, marker.GetAllApproved(db.GetDB()))
}

func CreateMarker(c *gin.Context) {
	appG := app.Gin{C: c}
	var m marker.Marker

	err := c.BindJSON(&m)
	if err != nil {
		zap.L().Error(err.Error())
		appG.Response(400, err.Error())
		return
	}
	err = m.Type.IsExist(db.GetDB())
	if err != nil {
		zap.L().Error("marker type not found: " + err.Error())
		appG.Response(422, "type doesn't exist")
		return
	}
	err = m.Save(db.GetDB())
	if err != nil {
		zap.L().Error("failed to create marker: " + err.Error())
		appG.Response(500, err.Error())
		return
	}
	appG.Response(200, m)
	zap.L().Info("marker created")
	telegram_bot.SendNotification(&m)
}
