package menus

import (
	"database/sql"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
)

// Repository is a object for bind sqlx.DB instrance
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
func (repo *Repository) Find(param Menu, offset string, limit string) ([]Menu, bool) {
	var menus []Menu

	var sql = "SELECT * FROM menus WHERE 1=1"

	if param.Name != "" {
		sql += " AND name LIKE '%" + param.Name + "%'"
	}

	sql += " LIMIT " + offset + "," + limit

	result := repo.db.Select(&menus, sql)

	return menus, result == nil
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

	repo.db.Get(&result, sql)

	return result
}

// FindByID is function to find specific object by id as a param
// visit http://gorm.io/docs/query.html for reference
func (repo *Repository) FindByID(id uint64) (Menu, bool) {
	var menu Menu

	result := repo.db.Get(&menu, "SELECT * FROM menus WHERE id = ? ", id)

	return menu, result == nil
}

// Save is function to save data to table
// visit http://gorm.io/docs/create.html
func (repo *Repository) Save(tx *sql.Tx, entity Menu) (Menu, error) {
	query := "INSERT INTO menus(name, created_time, updated_time) VALUES(?,?,?)"

	result, err := tx.Exec(query, entity.Name, time.Now(), time.Now())

	if err != nil {
		id, _ := result.LastInsertId()

		entity.ID = uint64(id)
	}

	return entity, err
}

// Update is function to update data those changed fields
// visit http://gorm.io/docs/update.html
func (repo *Repository) Update(tx *sql.Tx, entity Menu) (Menu, int64) {
	query := "UPDATE menus SET name = ? WHERE id = ?"

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
func (repo *Repository) Delete(tx *sql.Tx, entity Menu) (Menu, int64) {
	result, err := tx.Exec("DELETE FROM menus WHERE id = ?", entity.ID)

	rowAffected, _ := result.RowsAffected()

	if err != nil {
		log.Println(err.Error())
	} else {
		id, _ := result.LastInsertId()

		entity.ID = uint64(id)
	}

	return entity, rowAffected
}
