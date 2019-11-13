package main

import (
	"log"
	"path"

	"github.com/jessevdk/go-flags"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/krini-project/ProjectHandler/docs"
	"github.com/krini-project/ProjectHandler/persistence"
	"github.com/krini-project/ProjectHandler/server"
)

var opts struct {
	Port   int    `short:"p" long:"port" description:"Server port" default:"8080"`
	SQLite string `short:"d" long:"dbPath" description:"Path to the sqlite database file" default:"/tmp/krinidatabase"`
}

func main() {
	databaseHandler := persistence.DatabaseHandler{}

	_, err := flags.Parse(&opts)
	if err != nil {
		log.Println(err.Error())
		return
	}

	//sqliteTmpFileName := uuid.New().String()
	databaseHandler.InitSQLite(path.Join("/tmp", "sqliteTmpFileName"))

	server := server.Server{}
	server.Init(&databaseHandler, opts.Port)

}
