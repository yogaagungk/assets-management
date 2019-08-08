package category_assets

import (
	"github.com/jinzhu/gorm"
)

type CategoryAsset struct {
	gorm.Model
	Name string `gorm:"type:varchar(60)"`
}
