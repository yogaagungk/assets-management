package category_assets

import (
	"database/sql"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

// InjectDep is a function for inject db to Repository object
func ProvideRepo(db *sqlx.DB) *Repository {
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

	result := repo.db.Select(&categoryAssets, sql)

	return categoryAssets, result == nil
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

	repo.db.Get(&result, sql)

	return result
}

// FindByID is function to find specific object by id as a param
// visit http://gorm.io/docs/query.html for reference
func (repo *Repository) FindByID(id uint64) (CategoryAsset, bool) {
	var categoryAsset CategoryAsset

	result := repo.db.Get(&categoryAsset, "SELECT * FROM category_assets WHERE id = ? ", id)

	return categoryAsset, result == nil
}

// Save is function to save data to table
// visit http://gorm.io/docs/create.html
func (repo *Repository) Save(tx *sql.Tx, entity CategoryAsset) (CategoryAsset, error) {
	query := "INSERT INTO category_assets(name, created_time, updated_time) VALUES(?,?,?)"

	result, err := tx.Exec(query, entity.Name, time.Now(), time.Now())

	if err != nil {
		id, _ := result.LastInsertId()

		entity.ID = uint64(id)
	}

	return entity, err
}

// Update is function to update data those changed fields
// visit http://gorm.io/docs/update.html
func (repo *Repository) Update(tx *sql.Tx, entity CategoryAsset) (CategoryAsset, int64) {
	query := "UPDATE category_assets SET name = ? WHERE id = ?"

	result, err := tx.Exec(query, entity.Name, entity.ID)

	rowAffected, _ := result.RowsAffected()

	if err != nil {
		log.Println(err.Error())
	} else {
		id, _ := result.LastInsertId()

		entity.ID = uint64(id)
	}

	return entity, rowAffected
}

// Delete is function to delete data (flagged)
// visit http://gorm.io/docs/delete.html
// using approach to not permanently delete data, just update on deleteAt column
// to delete permanently use db.Unscoped().Delete(&entity)
func (repo *Repository) Delete(tx *sql.Tx, entity CategoryAsset) (CategoryAsset, int64) {
	result, err := tx.Exec("DELETE FROM category_assets WHERE id = ?", entity.ID)

	rowAffected, _ := result.RowsAffected()

	if err != nil {
		log.Println(err.Error())
	} else {
		id, _ := result.LastInsertId()

		entity.ID = uint64(id)
	}

	return entity, rowAffected
}
