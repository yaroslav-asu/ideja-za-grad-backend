package marker

import (
	"gorm.io/gorm"
)

type Marker struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	CoordsID    uint   `json:"coords_id"`
	Coords      Coords `json:"coords"`
	TypeID      uint   `json:"type_id"`
	Type        Type   `json:"type"`
	Description string `gorm:"size:512" json:"description"`
	Approved    bool   `gorm:"default:false" json:"approved"`
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
	m.CoordsID = c.ID
	err = db.Model(&Marker{}).Where("type_id = ? AND coords_id = ? AND description = ?", t.ID, c.ID, m.Description).Delete(&m).Error
	if err != nil {
		return err
	}
	return c.Delete(db)
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

func Get(db *gorm.DB, id uint) Marker {
	var marker Marker
	db.Preload("Type").Preload("Coords").First(&marker, id)
	return marker
}

func GetAllApproved(db *gorm.DB) []Marker {
	var markers []Marker
	db.Preload("Type").Preload("Coords").Where("approved = true").Find(&markers)
	return markers
}
