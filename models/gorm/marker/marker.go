package marker

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"os"
	"strings"
)

type Marker struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	CoordsID    uint    `json:"coords_id"`
	Coords      Coords  `json:"coords"`
	TypeID      uint    `json:"type_id"`
	Type        Type    `json:"type"`
	Description string  `gorm:"size:512" json:"description"`
	Approved    bool    `gorm:"default:false" json:"approved"`
	Images      []Image `gorm:"many2many:marker_images;" json:"images"`
}

func (m *Marker) Save(db *gorm.DB) error {
	err := m.Type.Save(db)
	if err != nil {
		return err
	}
	err = m.Coords.Save(db)
	if err != nil {
		return err
	}
	return db.Save(m).Error
}

func (m *Marker) Delete(db *gorm.DB) error {
	var t Type
	err := db.Model(&Type{}).Where("title = ?", m.Type.Title).First(&t).Error
	if err != nil {
		return err
	}
	var c Coords
	err = db.Model(&Coords{}).Where("lat = ? AND lng = ?", m.Coords.Lat, m.Coords.Lng).First(&c).Error
	if err != nil {
		return err
	}
	db.Model(Marker{}).Preload("Images").Where("type_id = ? AND coords_id = ? AND description = ?", t.ID, c.ID, m.Description).First(&m)
	db.Select(clause.Associations).Delete(&m)
	db.Model(Coords{}).Where("id = ?", c.ID).Delete(&Coords{})
	titles := make([]string, len(m.Images))
	for i, image := range m.Images {
		titles[i] = fmt.Sprintf("'%s'", image.Title)
		fmt.Println(os.Remove(fmt.Sprintf("static/images/%s", image.Title)))
	}
	db.Model(Image{}).Where(fmt.Sprintf("title in (%s)", strings.Join(titles, ","))).Delete(&Image{})
	return nil
}

func (m *Marker) Approve(db *gorm.DB) error {
	var t Type
	err := db.Model(&Type{}).Where("title = ?", m.Type.Title).First(&t).Error
	if err != nil {
		return err
	}
	var c Coords
	err = db.Model(&Coords{}).Where("lat = ? AND lng = ?", m.Coords.Lat, m.Coords.Lng).First(&c).Error
	if err != nil {
		return err
	}
	return db.Model(&Marker{}).Where("type_id = ? AND coords_id = ? AND description = ?", t.ID, c.ID, m.Description).Update("approved", true).Error
}

func Get(db *gorm.DB, id uint) (Marker, error) {
	var marker Marker
	err := db.Preload("Type").Preload("Coords").First(&marker, id).Error
	if err != nil {
		return Marker{}, err
	}
	return marker, nil
}

func GetAllApproved(db *gorm.DB) ([]Marker, error) {
	var markers []Marker
	err := db.Preload("Type").Preload("Coords").Preload("Images").Where("approved = true").Find(&markers).Error
	if err != nil {
		return nil, err
	}
	return markers, nil
}

func GetImages(db *gorm.DB, id uint) ([]Image, error) {
	var images []Image
	err := db.Model(&Marker{ID: id}).Association("Images").Find(&images)
	if err != nil {
		return nil, err
	}
	return images, nil
}
