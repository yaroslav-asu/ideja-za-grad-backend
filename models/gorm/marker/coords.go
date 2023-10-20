package marker

import "gorm.io/gorm"

type Coords struct {
	ID  uint `gorm:"primaryKey" json:"id"`
	Lat uint `json:"lat"`
	Lng uint `json:"lng"`
}

func (c *Coords) Save(db *gorm.DB) {
	db.FirstOrCreate(c)
}

func (c *Coords) Delete(db *gorm.DB) {
	db.Delete(c)
}
