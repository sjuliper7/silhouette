package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

//InitDatabase is function to create connection to Database
func (cf *Config) initDatabase() {
	db, err := sql.Open("mysql", populateStringConnection())

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Database connected successfully..")
	}

	cf.DB = db

}

func populateStringConnection() string {
	stringConnection := ""

	stringConnection += os.Getenv("DATABASE_USER") + ":" + os.Getenv("DATABASE_PASSWORD") +
		"@tcp(" + os.Getenv("DATABASE_HOST") + ":" + os.Getenv("DATABASE_PORT") + ")/" +
		os.Getenv("DATABASE_NAME")

	return stringConnection
}
