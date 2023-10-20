package marker

import "gorm.io/gorm"

type Marker struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	CoordsID    uint   `json:"coords_id"`
	Coords      Coords `json:"coords"`
	TypeID      uint   `json:"type_id"`
	Type        Type   `json:"type"`
	ImageID     uint   `json:"image_id"`
	Image       Image  `json:"image"`
	Description string `gorm:"size:512" json:"description"`
}

func (m *Marker) Save(db *gorm.DB) {
	db.Save(m)
}

func (m *Marker) Delete(db *gorm.DB) {
	db.Delete(m)
}
func Get(db *gorm.DB, id uint) Marker {
	var marker Marker
	db.First(&marker, id)
	return marker
}

func GetAll(db *gorm.DB) []Marker {
	var markers []Marker
	db.Find(&markers)
	return markers
}
