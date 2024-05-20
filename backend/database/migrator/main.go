package main

import (
	"database/sql"
	"log"
	"os"
	"strconv"

	"github.com/mass584/twitter-clone-app/backend/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	config := config.New()

	database_url :=
		config.DatabaseUser + ":" + config.DatabasePass +
			"@tcp(" + config.DatabaseHost + ":" + strconv.Itoa(config.DatabasePort) + ")" +
			"/" + config.DatabaseName +
			"?multiStatements=true"

	db, error := sql.Open("mysql", database_url)

	if error != nil {
		log.Fatal(error)
		panic(error)
	}

	driver, error := mysql.WithInstance(db, &mysql.Config{})

	if error != nil {
		log.Fatal(error)
		panic(error)
	}

	exec_path, error := os.Getwd()

	if error != nil {
		log.Fatal(error)
		panic(error)
	}

	source_url :=
		"file://" +
			exec_path + "/database/migrations"

	migrator, error := migrate.NewWithDatabaseInstance(
		source_url,
		"mysql",
		driver,
	)

	if error != nil {
		log.Fatal(error)
		panic(error)
	}

	error = migrator.Up()

	if error != nil {
		log.Fatal(error)
		panic(error)
	}
}
