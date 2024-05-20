package database

import (
	"database/sql"
	"strconv"

	"github.com/mass584/twitter-clone-app/backend/config"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	config := config.New()

	database_url :=
		config.DatabaseUser + ":" + config.DatabasePass +
			"@tcp(" + config.DatabaseHost + ":" + strconv.Itoa(config.DatabasePort) + ")" +
			"/" + config.DatabaseName +
			"?multiStatements=true"

	var error error

	DB, error = sql.Open("mysql", database_url)
	if error != nil {
		panic(error)
	}
}
