package roles

import (
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
func (repo *Repository) Find(param Role, offset string, limit string) ([]Role, bool) {
	var roles []Role

	var sql = "SELECT * FROM roles WHERE 1=1"

	if param.Name != "" {
		sql += " AND name LIKE '%" + param.Name + "%'"
	}

	sql += " LIMIT " + offset + "," + limit

	result := repo.db.Select(&roles, "SELECT * FROM roles")

	return roles, result == nil
}

// FindByID is function to find specific object by id as a param
// visit http://gorm.io/docs/query.html for reference
func (repo *Repository) FindByID(id uint64) (Role, bool) {
	var role Role

	result := repo.db.Get(&role, "SELECT * FROM roles WHERE id = ? ", id)

	return role, result == nil
}

// FindByName is function to find specific object by id as a param
// visit http://gorm.io/docs/query.html for reference
func (repo *Repository) FindByName(name string) (Role, bool) {
	var role Role

	result := repo.db.Get(&role, "SELECT * FROM roles WHERE name = ? ", name)

	return role, result == nil
}
