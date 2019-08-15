package users

import (
	"time"

	"github.com/yogaagungk/assets-management/services/roles"
)

/*
User is representation object of users table
*/
type User struct {
	ID          uint64
	Name        string
	Username    string
	Password    string `json:",omitempty"`
	Role        roles.Role
	CreatedTime time.Time `db:"created_time" json:"created_time"`
	UpdatedTime time.Time `db:"updated_time" json:"updated_time"`
}
