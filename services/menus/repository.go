package menus

import (
	"github.com/jinzhu/gorm"
)

// Repository is a object for bind gorm.DB instrance
type Repository struct {
	db *gorm.DB
}

// InjectDep is a function for inject db to Repository object
func ProvideRepo(db *gorm.DB) *Repository {
	return &Repository{db}
}

// Find is a function to find list of object with parameter, offset and limit
// Using native query for get dynamic parameter
// visit http://gorm.io/docs/sql_builder.html for reference
func (repo *Repository) Find(param Menu, offset string, limit string) ([]Menu, bool) {
	var menus []Menu

	var sql = "SELECT * FROM menus WHERE 1=1"

	if param.Name != "" {
		sql += " AND name LIKE '%" + param.Name + "%'"
	}

	sql += " LIMIT " + offset + "," + limit

	result := repo.db.Raw(sql).Scan(&menus).RecordNotFound()

	return menus, result
}

// Count is a function to count length list of object with parameter
// Using native query for get dynamic parameter
// visit http://gorm.io/docs/sql_builder.html for reference
func (repo *Repository) Count(param Menu) uint {
	var result uint

	var sql = "SELECT * FROM menu WHERE 1=1"

	if param.Name != "" {
		sql += " AND name LIKE '%" + param.Name + "%'"
	}

	repo.db.Raw(sql).Scan(&result)

	return result
}

// FindByID is function to find specific object by id as a param
// visit http://gorm.io/docs/query.html for reference
func (repo *Repository) FindByID(id uint64) (Menu, bool) {
	var menu Menu

	result := repo.db.Where("id = ?", id).First(&menu).RecordNotFound()

	return menu, result
}

// Save is function to save data to table
// visit http://gorm.io/docs/create.html
func (repo *Repository) Save(entity Menu) (Menu, error) {
	err := repo.db.Create(&entity)

	return entity, err.Error
}

// Update is function to update data those changed fields
// visit http://gorm.io/docs/update.html
func (repo *Repository) Update(entity Menu) (Menu, int64) {
	result := repo.db.Model(&entity).Updates(Menu{Name: entity.Name})

	return entity, result.RowsAffected
}

// Delete is function to delete data (flagged)
// visit http://gorm.io/docs/delete.html
// using approach to not permanently delete data, just update on deleteAt column
// to delete permanently use db.Unscoped().Delete(&entity)
func (repo *Repository) Delete(entity Menu) (Menu, int64) {
	result := repo.db.Unscoped().Delete(&entity)

	return entity, result.RowsAffected
}
