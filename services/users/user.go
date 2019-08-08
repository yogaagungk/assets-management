package users

import (
	"github.com/jinzhu/gorm"
	"github.com/yogaagungk/assets-management/services/roles"
)

/*
User is representation object of users table
*/
type User struct {
	gorm.Model
	Name     string     `gorm:"type:varchar(100)" json:"name"`
	Username string     `gorm:"type:varchar(60);unique_index" json:"username"`
	Password string     `gorm:"type:varchar(255)" json:"password"`
	Role     roles.Role `gorm:"foreignkey:role_id"`
	RoleId   uint64     `gorm:"column:role_id" json:"role_id"`
}
