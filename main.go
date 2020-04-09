package main

import (
	"log"

	"github.com/jessevdk/go-flags"
	"github.com/krini-project/ProjectHandler/api/server"
	_ "github.com/krini-project/ProjectHandler/docs"
	"github.com/krini-project/ProjectHandler/persistence"

	_ "github.com/go-sql-driver/mysql"
)

var opts struct {
	Port   int    `short:"p" long:"port" description:"Server port" default:"8080"`
	SQLite string `short:"d" long:"dbPath" description:"Path to the sqlite database file" default:"/tmp/krinidatabase"`
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	databaseHandler := persistence.DatabaseHandler{}

	_, err := flags.Parse(&opts)
	if err != nil {
		log.Println(err.Error())
		return
	}

	//sqliteTmpFileName := uuid.New().String()
	databaseHandler.InitSQL()

	databaseHandler.CreateDefaultRoles()

	server := server.Server{}
	server.Init(&databaseHandler, opts.Port)

}
