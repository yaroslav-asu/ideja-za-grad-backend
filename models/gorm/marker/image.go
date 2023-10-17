package marker

import "gorm.io/gorm"

type Image struct {
	ID       uint `gorm:"primaryKey"`
	FileName string
}

func (i *Image) Save(db *gorm.DB) {
	db.FirstOrCreate(i)
}

func (i *Image) Delete(db *gorm.DB) {
	db.Delete(i)
}
