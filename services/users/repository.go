package users

import (
	"database/sql"
	"fmt"
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
func (repo *Repository) Find(param User, offset string, limit string) ([]User, bool) {
	var users []User

	var sql = `SELECT 
	users.id, 
	users.name,
	users.username, 
	roles.id, 
	roles.name FROM users INNER JOIN roles ON users.role_id = roles.id WHERE 1=1`

	if param.Username != "" {
		sql += " AND username = '" + param.Username + "'"
	}

	if param.Name != "" {
		sql += " AND users.name LIKE '%" + param.Name + "%'"
	}

	if param.Role.ID != 0 {
		sql += " AND role_id = " + fmt.Sprint(param.Role.ID)
	}

	sql += " LIMIT " + offset + "," + limit

	rows, result := repo.db.Query(sql)

	for rows.Next() {
		var user User

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Username,
			&user.Role.ID,
			&user.Role.Name)

		if err != nil {
			log.Panic(err.Error())
		} else {
			users = append(users, user)
		}
	}

	return users, result == nil
}

// Count is a function to count length list of object with parameter
// Using native query for get dynamic parameter
// visit http://gorm.io/docs/sql_builder.html for reference
func (repo *Repository) Count(param User) uint {
	var result uint

	var sql = "SELECT * FROM user WHERE 1=1"

	if param.Name != "" {
		sql += " AND name LIKE '%" + param.Name + "%'"
	}

	if param.Role.ID != 0 {
		sql += " AND role_id = " + fmt.Sprint(param.Role.ID)
	}

	repo.db.Get(&result, sql)

	return result
}

// FindByID is function to find specific object by id as a param
// visit http://gorm.io/docs/query.html for reference
func (repo *Repository) FindByID(id uint64) (User, bool) {
	var user User

	sql := `SELECT 
	users.id, 
	users.name,
	users.username, 
	roles.id, 
	roles.name FROM users INNER JOIN roles ON users.role_id = roles.id WHERE users.id = ?`

	row := repo.db.QueryRow(sql, id)

	result := row.Scan(
		&user.ID,
		&user.Name,
		&user.Username,
		&user.Role.ID,
		&user.Role.Name)

	if result != nil {
		log.Panic(result.Error())
	}

	return user, result == nil
}

// FindByUsername is function to find specific object by username as a param
// visit http://gorm.io/docs/query.html for reference
func (repo *Repository) FindByUsername(username string) (User, bool) {
	var user User

	sql := `SELECT 
	users.id, 
	users.name,
	users.username, 
	users.password, 
	roles.id, 
	roles.name FROM users INNER JOIN roles ON users.role_id = roles.id WHERE users.username = ?`

	row := repo.db.QueryRow(sql, user)

	result := row.Scan(
		&user.ID,
		&user.Name,
		&user.Username,
		&user.Password,
		&user.Role.ID,
		&user.Role.Name)

	if result != nil {
		log.Panic(result.Error())
	}

	return user, result == nil
}

// Save is function to save data to table
// visit http://gorm.io/docs/create.html
func (repo *Repository) Save(tx *sql.Tx, entity User) (User, error) {
	query := `INSERT INTO users(
		name, 
		username, 
		password, 
		role_id, 
		created_time, 
		updated_time) VALUES(?,?,?,?,?,?)`

	result, err := tx.Exec(query,
		entity.Name,
		entity.Username,
		entity.Password,
		entity.Role.ID,
		time.Now(),
		time.Now())

	if err != nil {
		id, _ := result.LastInsertId()

		entity.ID = uint64(id)
	}

	return entity, err
}

// Update is function to update data those changed fields
// visit http://gorm.io/docs/update.html
func (repo *Repository) Update(tx *sql.Tx, entity User) (User, int64) {
	query := `UPDATE users SET 
	name = ?, 
	username = ?, 
	password = ?,
	role_id = ? WHERE id = ?`

	result, err := tx.Exec(query,
		entity.Name,
		entity.Username,
		entity.Password,
		entity.ID)

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
func (repo *Repository) Delete(tx *sql.Tx, entity User) (User, int64) {
	result, err := tx.Exec("DELETE FROM users WHERE id = ?", entity.ID)

	rowAffected, _ := result.RowsAffected()

	if err != nil {
		log.Println(err.Error())
	} else {
		id, _ := result.LastInsertId()

		entity.ID = uint64(id)
	}

	return entity, rowAffected
}
