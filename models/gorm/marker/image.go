package marker

import (
	"gorm.io/gorm"
	"os"
)

type Image struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Title string `json:"title"`
}

func (i *Image) Save(db *gorm.DB) {
	db.Create(i)
}

func (i *Image) Delete(db *gorm.DB) error {
	err := os.Remove(i.Title)
	if err != nil {
		return err
	}
	return db.Delete(i).Error
}
