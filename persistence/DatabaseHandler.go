package persistence

import (
	"log"

	"github.com/jinzhu/gorm"
)

// DatabaseHandler Used to hold and initialize the database connection via GORM
type DatabaseHandler struct {
	Database *gorm.DB
}

// InitSQLite Initiates a database handler using SQLite
// path Path to store the database
func (handler *DatabaseHandler) InitSQLite(path string) error {
	db, err := gorm.Open("sqlite3", path)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	handler.Database = db
	handler.handleMigrations()

	return nil
}

func (handler *DatabaseHandler) handleMigrations() {
	if err := handler.Database.AutoMigrate(&User{}, &Project{}, &WorkflowRun{}, &Right{}, &Role{}, &ProjectUser{}).Error; err != nil {
		log.Fatalln(err)
	}
}
