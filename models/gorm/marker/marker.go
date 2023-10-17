package marker

import "gorm.io/gorm"

type Marker struct {
	ID          uint `gorm:"primaryKey"`
	TypeID      uint
	Type        Type
	ImageID     uint
	Image       Image
	Description string `gorm:"size:512"`
}

func (m *Marker) Save(db *gorm.DB) {
	db.FirstOrCreate(m)
}

func (m *Marker) Delete(db *gorm.DB) {
	db.Delete(m)
}
