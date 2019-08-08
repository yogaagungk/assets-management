package config

import (
	"log"

	"github.com/yogaagungk/assets-management/services/roles"
	"github.com/yogaagungk/assets-management/services/users"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/yogaagungk/assets-management/services/menus"
)

var DB *gorm.DB = nil

/*
OpenDatabaseConnection : Open connection to database
*/
func OpenDatabaseConnection() *gorm.DB {
	DB, err := gorm.Open(
		"mysql",
		"root:root@tcp(localhost:3306)/assets_management?charset=utf8&parseTime=True&loc=Local",
	)

	if err != nil {
		log.Fatal(err.Error())
	}

	DB.DB().SetMaxIdleConns(5)
	DB.AutoMigrate(&menus.Menu{}, &roles.Role{}, &users.User{})

	return DB
}

func ProvideDatabase() *gorm.DB {
	if DB == nil {
		return OpenDatabaseConnection()
	}

	return DB
}
