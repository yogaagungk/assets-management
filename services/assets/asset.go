package assets

// import (
// 	"github.com/yogaagungk/assets-management/services/category_assets"
// 	"github.com/yogaagungk/assets-management/services/users"
// )

// type Asset struct {
// 	ID         uint64                        `gorm:"column:id;primary_key"`
// 	Name       string                        `gorm:"type:varchar(60)"`
// 	Price      float32                       `gorm:"type:decimal(20)"`
// 	Category   category_assets.CategoryAsset `gorm:"foreignkey:category_asset_id"`
// 	CategoryID uint64                        `gorm:"column:category_asset_id" json:"category_asset_id"`
// 	User       users.User                    `gorm:"foreignkey:user_id"`
// 	UserID     uint64                        `gorm:"column:user_id" json:"user_id"`
// }
