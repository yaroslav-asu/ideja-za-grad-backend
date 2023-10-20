package marker

import "gorm.io/gorm"

type Coords struct {
	ID  uint `gorm:"primaryKey" json:"id"`
	Lat uint `json:"lat"`
	Lng uint `json:"lng"`
}

func (c *Coords) Save(db *gorm.DB) error {
	return db.Model(Coords{}).Where("lat = ? and lng = ?", c.Lat, c.Lng).FirstOrCreate(&c).Error
}

func (c *Coords) Delete(db *gorm.DB) error {
	return db.Delete(c).Error
}
