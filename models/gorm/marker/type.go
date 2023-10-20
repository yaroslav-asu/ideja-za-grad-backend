package marker

import "gorm.io/gorm"

type Type struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Title string `gorm:"size:64" json:"title"`
}

func (t *Type) Save(db *gorm.DB) error {
	return db.Model(Type{}).Where("title = ? or id = ?", t.Title, t.ID).FirstOrCreate(&t).Error
}

func (t *Type) Delete(db *gorm.DB) error {
	return db.Delete(t).Error
}
func (t *Type) IsExist(db *gorm.DB) error {
	return db.Model(Type{}).Where("title = ? or id = ?", t.Title, t.ID).First(&Type{}).Error
}

func GetAllTypes(db *gorm.DB) []Type {
	var types []Type
	db.Find(&types)
	return types
}
