package marker

import "gorm.io/gorm"

type Type struct {
	ID    uint   `gorm:"primaryKey"`
	Title string `gorm:"size:64"`
}

func (t *Type) Save(db *gorm.DB) {
	db.FirstOrCreate(t)
}

func (t *Type) Delete(db *gorm.DB) {
	db.Delete(t)
}
