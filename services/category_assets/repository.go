package category_assets

import "github.com/jinzhu/gorm"

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
func (repo *Repository) Find(param CategoryAsset, offset string, limit string) ([]CategoryAsset, bool) {
	var categoryAssets []CategoryAsset

	var sql = "SELECT * FROM category_assets WHERE 1=1"

	if param.Name != "" {
		sql += " AND name LIKE '%" + param.Name + "%'"
	}

	sql += " LIMIT " + offset + "," + limit

	result := repo.db.Raw(sql).Scan(&categoryAssets).RecordNotFound()

	return categoryAssets, result
}

// Count is a function to count length list of object with parameter
// Using native query for get dynamic parameter
// visit http://gorm.io/docs/sql_builder.html for reference
func (repo *Repository) Count(param CategoryAsset) uint {
	var result uint

	var sql = "SELECT * FROM category_assets WHERE 1=1"

	if param.Name != "" {
		sql += " AND name LIKE '%" + param.Name + "%'"
	}

	repo.db.Raw(sql).Scan(&result)

	return result
}

// FindByID is function to find specific object by id as a param
// visit http://gorm.io/docs/query.html for reference
func (repo *Repository) FindByID(id uint64) (CategoryAsset, bool) {
	var categoryAsset CategoryAsset

	result := repo.db.Where("id = ?", id).First(&categoryAsset).RecordNotFound()

	return categoryAsset, result
}

// Save is function to save data to table
// visit http://gorm.io/docs/create.html
func (repo *Repository) Save(entity CategoryAsset) (CategoryAsset, error) {
	err := repo.db.Create(&entity)

	return entity, err.Error
}

// Update is function to update data those changed fields
// visit http://gorm.io/docs/update.html
func (repo *Repository) Update(entity CategoryAsset) (CategoryAsset, int64) {
	result := repo.db.Model(&entity).Updates(CategoryAsset{Name: entity.Name})

	return entity, result.RowsAffected
}

// Delete is function to delete data (flagged)
// visit http://gorm.io/docs/delete.html
// using approach to not permanently delete data, just update on deleteAt column
// to delete permanently use db.Unscoped().Delete(&entity)
func (repo *Repository) Delete(entity CategoryAsset) (CategoryAsset, int64) {
	result := repo.db.Delete(&entity)

	return entity, result.RowsAffected
}
