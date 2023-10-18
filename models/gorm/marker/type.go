package marker

import "gorm.io/gorm"

type Type struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Title string `gorm:"size:64" json:"title"`
}

func (t *Type) Save(db *gorm.DB) {
	db.FirstOrCreate(t)
}

func (t *Type) Delete(db *gorm.DB) {
	db.Delete(t)
}
