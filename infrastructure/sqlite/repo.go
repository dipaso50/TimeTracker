package sqlite

import (
	"fmt"
	"tracker/domain"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	db     *gorm.DB
	dbType = "sqlite3"
	dbName string
)

//InitialMigration configure our database
func InitialMigration(dbNa string) {

	dbName = dbNa

	db, err := gorm.Open(dbType, dbName)

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}

	defer db.Close()

	db.AutoMigrate(&domain.Project{}, &domain.Task{}, &domain.Heads{}, &domain.TaskTime{})
}
