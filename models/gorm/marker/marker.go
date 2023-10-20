package marker

import "gorm.io/gorm"

type Marker struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	CoordsID    uint   `json:"coords_id"`
	Coords      Coords `json:"coords"`
	TypeID      uint   `json:"type_id"`
	Type        Type   `json:"type"`
	Description string `gorm:"size:512" json:"description"`
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

func (m *Marker) Delete(db *gorm.DB) {
	db.Delete(m)
}

func Get(db *gorm.DB, id uint) Marker {
	var marker Marker
	db.Preload("Type").Preload("Coords").First(&marker, id)
	return marker
}

func GetAll(db *gorm.DB) []Marker {
	var markers []Marker
	db.Preload("Type").Preload("Coords").Find(&markers)
	return markers
}
