package marker

import "gorm.io/gorm"

type Coords struct {
	ID  uint    `gorm:"primaryKey" json:"id"`
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func (c *Coords) Save(db *gorm.DB) error {
	return db.Model(Coords{}).Where("lat = ? and lng = ?", c.Lat, c.Lng).FirstOrCreate(&c).Error
}

func (c *Coords) Delete(db *gorm.DB) error {
	return db.Where("lat = ? and lng = ?", c.Lat, c.Lng).Delete(c).Error
}
