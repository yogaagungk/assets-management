package assets

import (
	"time"

	"github.com/yogaagungk/assets-management/services/category_assets"
	"github.com/yogaagungk/assets-management/services/users"
)

type Asset struct {
	ID          uint64
	Name        string
	Price       float32
	Category    category_assets.CategoryAsset
	User        users.User
	CreatedTime time.Time `db:"created_time" json:"created_time"`
	UpdatedTime time.Time `db:"updated_time" json:"updated_time"`
}
