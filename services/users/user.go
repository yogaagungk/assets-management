package users

import (
	"github.com/yogaagungk/assets-management/services/roles"
)

/*
User is representation object of users table
*/
type User struct {
	ID       uint64
	Name     string
	Username string
	Password string
	Role     roles.Role
}
