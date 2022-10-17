package database

import (
	"github.com/sepehrmohseni/go-web-boilerplate/config"
	"github.com/sepehrmohseni/go-web-boilerplate/entities"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func ConnectDB() error {
	conf := config.AppConfig.GetDatabaseURI()
	db, connectionErr := gorm.Open(mysql.Open(conf), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if connectionErr != nil {
		panic(connectionErr)
	}
	Database = db
	return db.AutoMigrate(
		&entities.Test{},
	)
}
