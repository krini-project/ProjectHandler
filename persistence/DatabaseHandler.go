package persistence

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/volatiletech/sqlboiler/boil"
)

// DatabaseHandler Used to hold and initialize the database connection via GORM
type DatabaseHandler struct {
	DB *sql.DB
}

//InitSQL Initiliazes an SQL session
func (handler *DatabaseHandler) InitSQL() error {
	dbUsername := "root"
	dbPassword := "test123"
	dbHost := "localhost"

	if username := os.Getenv("dbUsername"); username != "" {
		dbUsername = username
	}

	if password := os.Getenv("dbPassword"); password != "" {
		dbPassword = password
	}

	if host := os.Getenv("dbHost"); host != "" {
		dbHost = host
	}

	dbConnectionString := fmt.Sprintf("%s:%s@tcp(%s:3306)/cwlab?charset=utf8&parseTime=True&loc=Local",
		dbUsername, dbPassword, dbHost)

	db, err := sql.Open("mysql", dbConnectionString)
	if err != nil {
		log.Println(err.Error())
		return fmt.Errorf("Error while connecting to database")
	}

	handler.DB = db
	boil.SetDB(db)
	return nil
}
