package menus

import "time"

/*
Menu is representation of menu table
*/
type Menu struct {
	ID          uint64
	Name        string
	CreatedTime time.Time `db:"created_time" json:"created_time"`
	UpdatedTime time.Time `db:"updated_time" json:"updated_time"`
}
