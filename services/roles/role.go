package roles

import "time"

/*
Role is representation of role table
*/
type Role struct {
	ID          uint64
	Name        string
	CreatedTime time.Time `db:"created_time" json:"created_time"`
	UpdatedTime time.Time `db:"updated_time" json:"updated_time"`
}
