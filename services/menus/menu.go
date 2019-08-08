package menus

import (
	"github.com/jinzhu/gorm"
)

/*
Menu is representation of menu table
*/
type Menu struct {
	gorm.Model
	Name string `gorm:"type:varchar(60)"`
}
