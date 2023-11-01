package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"strconv"
	"strings"
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
	m, err := marker.Get(db.GetDB(), uint(id))
	if err != nil {
		appG.Response(500, err.Error())
		return
	}
	appG.Response(200, m)
}

func GetMarkers(c *gin.Context) {
	appG := app.Gin{C: c}
	markers, err := marker.GetAllApproved(db.GetDB())
	if err != nil {
		appG.Response(500, err.Error())
		return
	}
	appG.Response(200, markers)
}
func GetImages(c *gin.Context) {
	appG := app.Gin{C: c}
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		appG.Response(400, err.Error())
		return
	}
	images, err := marker.GetImages(db.GetDB(), uint(id))
	if err != nil {
		appG.Response(500, err.Error())
		return
	}
	appG.Response(200, images)
}

func CreateMarker(c *gin.Context) {
	appG := app.Gin{C: c}
	var m marker.Marker

	form, err := c.MultipartForm()
	if err != nil {
		zap.L().Error(err.Error())
		appG.Response(400, err.Error())
		return
	}
	typeId, err := strconv.ParseUint(form.Value["type"][0], 10, 32)
	if err != nil {
		zap.L().Error(err.Error())
		appG.Response(400, err.Error())
		return
	}
	m.Type.ID = uint(typeId)
	err = m.Type.IsExist(db.GetDB())
	if err != nil {
		zap.L().Error("marker type not found: " + err.Error())
		appG.Response(422, "type doesn't exist")
		return
	}
	m.Description = form.Value["description"][0]
	m.Coords.Lat, _ = strconv.ParseFloat(form.Value["lat"][0], 64)
	m.Coords.Lng, _ = strconv.ParseFloat(form.Value["lng"][0], 64)
	files := form.File["images"]
	for _, file := range files {
		uniqueId := uuid.New()
		fileExt := strings.Split(file.Filename, ".")[1]
		imageTitle := fmt.Sprintf("%s.%s", uniqueId, fileExt)
		path := fmt.Sprintf("static/images/%s", imageTitle)
		err = c.SaveUploadedFile(file, path)
		if err != nil {
			zap.L().Error(err.Error())
		}
		image := marker.Image{Title: imageTitle}
		image.Save(db.GetDB())
		m.Images = append(m.Images, image)
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
