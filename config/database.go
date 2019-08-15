package config

import (
	"log"

	"github.com/yogaagungk/assets-management/database"

	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

/*
OpenDatabaseConnection : Open connection to database
*/
func OpenDatabaseConnection() *sqlx.DB {
	DB, errConnect := sqlx.Connect(
		"mysql",
		"root:root@tcp(localhost:3306)/assets_management?charset=utf8&parseTime=True&loc=Local",
	)

	if errConnect != nil {
		log.Fatal(errConnect.Error())
	}

	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(10)

	for _, sql := range database.SqlScheme {
		_, errExec := DB.Exec(sql)

		if errExec != nil {
			log.Panic(errExec.Error())
		}
	}

	return DB
}

func ProvideDatabase() *sqlx.DB {
	if DB == nil {
		return OpenDatabaseConnection()
	}

	return DB
}
