package roles

import (
	"github.com/jinzhu/gorm"
)

/*
Role is representation of role table
*/
type Role struct {
	gorm.Model
	Name string `gorm:"type:varchar(60)"`
}
